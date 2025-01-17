/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package thrift

import (
	"encoding/base64"
	"fmt"
	"io"

	"github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

const (
	THRIFT_JSON_PROTOCOL_VERSION = 1
)

// for references to _ParseContext see simple_json_protocol.go

// JSONProtocol is the Compact JSON protocol implementation for thrift.
//
// This protocol produces/consumes a compact JSON output with field numbers as
// object keys and field values lightly encoded.
//
// Example: With the Message definition
//
//	struct Message {
//	  1: bool aBool
//	  2: map<string, bool> aBoolStringMap
//	},
//
//	Message(aBool=True, aBoolStringMap={"key1": True, "key2": False})
//
// will be encoded as:
//
//	{"1":{"tf":1},"2":{"map":["str","tf",2,{"key1": 1,"key2":0}]}}'
type jsonProtocol struct {
	*simpleJSONProtocol
}

var _ types.Format = (*jsonProtocol)(nil)

// Constructor
func NewJSONProtocol(buffer io.ReadWriteCloser) types.Format {
	v := &jsonProtocol{simpleJSONProtocol: newSimpleJSONProtocol(buffer)}
	v.parseContextStack = append(v.parseContextStack, int(_CONTEXT_IN_TOPLEVEL))
	v.dumpContext = append(v.dumpContext, int(_CONTEXT_IN_TOPLEVEL))
	return v
}

func (p *jsonProtocol) WriteMessageBegin(name string, typeID types.MessageType, seqID int32) error {
	p.resetContextStack() // THRIFT-3735
	if e := p.OutputListBegin(); e != nil {
		return e
	}
	if e := p.WriteI32(THRIFT_JSON_PROTOCOL_VERSION); e != nil {
		return e
	}
	if e := p.WriteString(name); e != nil {
		return e
	}
	if e := p.WriteByte(byte(typeID)); e != nil {
		return e
	}
	if e := p.WriteI32(seqID); e != nil {
		return e
	}
	return nil
}

func (p *jsonProtocol) WriteMessageEnd() error {
	return p.OutputListEnd()
}

func (p *jsonProtocol) WriteStructBegin(name string) error {
	if e := p.OutputObjectBegin(); e != nil {
		return e
	}
	return nil
}

func (p *jsonProtocol) WriteStructEnd() error {
	return p.OutputObjectEnd()
}

func (p *jsonProtocol) WriteFieldBegin(name string, typeID types.Type, id int16) error {
	if e := p.WriteI16(id); e != nil {
		return e
	}
	if e := p.OutputObjectBegin(); e != nil {
		return e
	}
	s, e1 := p.TypeIdToString(typeID)
	if e1 != nil {
		return e1
	}
	if e := p.WriteString(s); e != nil {
		return e
	}
	return nil
}

func (p *jsonProtocol) WriteFieldEnd() error {
	return p.OutputObjectEnd()
}

func (p *jsonProtocol) WriteFieldStop() error { return nil }

func (p *jsonProtocol) WriteMapBegin(keyType types.Type, valueType types.Type, size int) error {
	if e := p.OutputListBegin(); e != nil {
		return e
	}
	s, e1 := p.TypeIdToString(keyType)
	if e1 != nil {
		return e1
	}
	if e := p.WriteString(s); e != nil {
		return e
	}
	s, e1 = p.TypeIdToString(valueType)
	if e1 != nil {
		return e1
	}
	if e := p.WriteString(s); e != nil {
		return e
	}
	if e := p.WriteI64(int64(size)); e != nil {
		return e
	}
	return p.OutputObjectBegin()
}

func (p *jsonProtocol) WriteMapEnd() error {
	if e := p.OutputObjectEnd(); e != nil {
		return e
	}
	return p.OutputListEnd()
}

func (p *jsonProtocol) WriteListBegin(elemType types.Type, size int) error {
	return p.OutputElemListBegin(elemType, size)
}

func (p *jsonProtocol) WriteListEnd() error {
	return p.OutputListEnd()
}

func (p *jsonProtocol) WriteSetBegin(elemType types.Type, size int) error {
	return p.OutputElemListBegin(elemType, size)
}

func (p *jsonProtocol) WriteSetEnd() error {
	return p.OutputListEnd()
}

func (p *jsonProtocol) WriteBool(b bool) error {
	if b {
		return p.WriteI32(1)
	}
	return p.WriteI32(0)
}

func (p *jsonProtocol) WriteByte(b byte) error {
	return p.WriteI32(int32(b))
}

func (p *jsonProtocol) WriteI16(v int16) error {
	return p.WriteI32(int32(v))
}

func (p *jsonProtocol) WriteI32(v int32) error {
	return p.OutputI64(int64(v))
}

func (p *jsonProtocol) WriteI64(v int64) error {
	return p.OutputI64(int64(v))
}

func (p *jsonProtocol) WriteDouble(v float64) error {
	return p.OutputF64(v)
}

func (p *jsonProtocol) WriteFloat(v float32) error {
	return p.OutputF32(v)
}

func (p *jsonProtocol) WriteString(v string) error {
	return p.OutputString(v)
}

func (p *jsonProtocol) WriteBinary(v []byte) error {
	// JSON library only takes in a string,
	// not an arbitrary byte array, to ensure bytes are transmitted
	// efficiently we must convert this into a valid JSON string
	// therefore we use base64 encoding to avoid excessive escaping/quoting
	if e := p.OutputPreValue(); e != nil {
		return e
	}
	if _, e := p.write(types.JSON_QUOTE_BYTES); e != nil {
		return types.NewProtocolException(e)
	}
	if len(v) > 0 {
		writer := base64.NewEncoder(base64.StdEncoding, p.writer)
		if _, e := writer.Write(v); e != nil {
			p.writer.Reset(p.buffer) // THRIFT-3735
			return types.NewProtocolException(e)
		}
		if e := writer.Close(); e != nil {
			return types.NewProtocolException(e)
		}
	}
	if _, e := p.write(types.JSON_QUOTE_BYTES); e != nil {
		return types.NewProtocolException(e)
	}
	return p.OutputPostValue()
}

// Reading methods.
func (p *jsonProtocol) ReadMessageBegin() (name string, typeID types.MessageType, seqID int32, err error) {
	p.resetContextStack() // THRIFT-3735
	if isNull, err := p.ParseListBegin(); isNull || err != nil {
		return name, typeID, seqID, err
	}
	version, err := p.ReadI32()
	if err != nil {
		return name, typeID, seqID, err
	}
	if version != THRIFT_JSON_PROTOCOL_VERSION {
		e := fmt.Errorf("Unknown Protocol version %d, expected version %d", version, THRIFT_JSON_PROTOCOL_VERSION)
		return name, typeID, seqID, types.NewProtocolExceptionWithType(types.INVALID_DATA, e)

	}
	if name, err = p.ReadString(); err != nil {
		return name, typeID, seqID, err
	}
	bTypeID, err := p.ReadByte()
	typeID = types.MessageType(bTypeID)
	if err != nil {
		return name, typeID, seqID, err
	}
	if seqID, err = p.ReadI32(); err != nil {
		return name, typeID, seqID, err
	}
	return name, typeID, seqID, nil
}

func (p *jsonProtocol) ReadMessageEnd() error {
	err := p.ParseListEnd()
	return err
}

func (p *jsonProtocol) ReadStructBegin() (name string, err error) {
	_, err = p.ParseObjectStart()
	return "", err
}

func (p *jsonProtocol) ReadStructEnd() error {
	return p.ParseObjectEnd()
}

func (p *jsonProtocol) ReadFieldBegin() (string, types.Type, int16, error) {
	b, _ := p.reader.Peek(1)
	if len(b) < 1 || b[0] == types.JSON_RBRACE[0] || b[0] == types.JSON_RBRACKET[0] {
		return "", types.STOP, -1, nil
	}
	fieldID, err := p.ReadI16()
	if err != nil {
		return "", types.STOP, fieldID, err
	}
	if _, err = p.ParseObjectStart(); err != nil {
		return "", types.STOP, fieldID, err
	}
	sType, err := p.ReadString()
	if err != nil {
		return "", types.STOP, fieldID, err
	}
	fType, err := p.StringToTypeId(sType)
	return "", fType, fieldID, err
}

func (p *jsonProtocol) ReadFieldEnd() error {
	return p.ParseObjectEnd()
}

func (p *jsonProtocol) ReadMapBegin() (keyType types.Type, valueType types.Type, size int, e error) {
	if isNull, e := p.ParseListBegin(); isNull || e != nil {
		return types.VOID, types.VOID, 0, e
	}

	// read keyType
	sKeyType, e := p.ReadString()
	if e != nil {
		return keyType, valueType, size, e
	}
	keyType, e = p.StringToTypeId(sKeyType)
	if e != nil {
		return keyType, valueType, size, e
	}

	// read valueType
	sValueType, e := p.ReadString()
	if e != nil {
		return keyType, valueType, size, e
	}
	valueType, e = p.StringToTypeId(sValueType)
	if e != nil {
		return keyType, valueType, size, e
	}

	// read size
	iSize, e := p.ReadI64()
	if e != nil {
		return keyType, valueType, size, e
	}
	size = int(iSize)

	_, e = p.ParseObjectStart()
	return keyType, valueType, size, e
}

func (p *jsonProtocol) ReadMapEnd() error {
	e := p.ParseObjectEnd()
	if e != nil {
		return e
	}
	return p.ParseListEnd()
}

func (p *jsonProtocol) ReadListBegin() (elemType types.Type, size int, e error) {
	return p.ParseElemListBegin()
}

func (p *jsonProtocol) ReadListEnd() error {
	return p.ParseListEnd()
}

func (p *jsonProtocol) ReadSetBegin() (elemType types.Type, size int, e error) {
	return p.ParseElemListBegin()
}

func (p *jsonProtocol) ReadSetEnd() error {
	return p.ParseListEnd()
}

func (p *jsonProtocol) ReadBool() (bool, error) {
	value, err := p.ReadI32()
	return (value != 0), err
}

func (p *jsonProtocol) ReadByte() (byte, error) {
	v, err := p.ReadI64()
	return byte(v), err
}

func (p *jsonProtocol) ReadI16() (int16, error) {
	v, err := p.ReadI64()
	return int16(v), err
}

func (p *jsonProtocol) ReadI32() (int32, error) {
	v, err := p.ReadI64()
	return int32(v), err
}

func (p *jsonProtocol) ReadI64() (int64, error) {
	v, _, err := p.ParseI64()
	return v, err
}

func (p *jsonProtocol) ReadDouble() (float64, error) {
	v, _, err := p.ParseF64()
	return v, err
}

func (p *jsonProtocol) ReadFloat() (float32, error) {
	v, _, err := p.ParseF32()
	return v, err
}

func (p *jsonProtocol) ReadString() (string, error) {
	var v string
	if err := p.ParsePreValue(); err != nil {
		return v, err
	}
	f, _ := p.reader.Peek(1)
	if len(f) > 0 && f[0] == types.JSON_QUOTE {
		p.reader.ReadByte()
		value, err := p.ParseStringBody()
		v = value
		if err != nil {
			return v, err
		}
	} else if len(f) > 0 && f[0] == types.JSON_NULL[0] {
		b := make([]byte, len(types.JSON_NULL))
		_, err := p.reader.Read(b)
		if err != nil {
			return v, types.NewProtocolException(err)
		}
		if string(b) != string(types.JSON_NULL) {
			e := fmt.Errorf("Expected a JSON string, found unquoted data started with %s", string(b))
			return v, types.NewProtocolExceptionWithType(types.INVALID_DATA, e)
		}
	} else {
		e := fmt.Errorf("Expected a JSON string, found unquoted data started with %s", string(f))
		return v, types.NewProtocolExceptionWithType(types.INVALID_DATA, e)
	}
	return v, p.ParsePostValue()
}

func (p *jsonProtocol) ReadBinary() ([]byte, error) {
	var v []byte
	if err := p.ParsePreValue(); err != nil {
		return nil, err
	}
	f, _ := p.reader.Peek(1)
	if len(f) > 0 && f[0] == types.JSON_QUOTE {
		p.reader.ReadByte()
		value, err := p.ParseBase64EncodedBody()
		v = value
		if err != nil {
			return v, err
		}
	} else if len(f) > 0 && f[0] == types.JSON_NULL[0] {
		b := make([]byte, len(types.JSON_NULL))
		_, err := p.reader.Read(b)
		if err != nil {
			return v, types.NewProtocolException(err)
		}
		if string(b) != string(types.JSON_NULL) {
			e := fmt.Errorf("Expected a JSON string, found unquoted data started with %s", string(b))
			return v, types.NewProtocolExceptionWithType(types.INVALID_DATA, e)
		}
	} else {
		e := fmt.Errorf("Expected a JSON string, found unquoted data started with %s", string(f))
		return v, types.NewProtocolExceptionWithType(types.INVALID_DATA, e)
	}

	return v, p.ParsePostValue()
}

func (p *jsonProtocol) Flush() (err error) {
	err = p.writer.Flush()
	if err == nil {
		return flush(p.buffer)
	}
	return types.NewProtocolException(err)
}

func (p *jsonProtocol) Skip(fieldType types.Type) (err error) {
	return types.SkipDefaultDepth(p, fieldType)
}

func (p *jsonProtocol) Close() error {
	return p.buffer.Close()
}

func (p *jsonProtocol) OutputElemListBegin(elemType types.Type, size int) error {
	if e := p.OutputListBegin(); e != nil {
		return e
	}
	s, e1 := p.TypeIdToString(elemType)
	if e1 != nil {
		return e1
	}
	if e := p.WriteString(s); e != nil {
		return e
	}
	if e := p.WriteI64(int64(size)); e != nil {
		return e
	}
	return nil
}

func (p *jsonProtocol) ParseElemListBegin() (elemType types.Type, size int, e error) {
	if isNull, e := p.ParseListBegin(); isNull || e != nil {
		return types.VOID, 0, e
	}
	sElemType, err := p.ReadString()
	if err != nil {
		return types.VOID, size, err
	}
	elemType, err = p.StringToTypeId(sElemType)
	if err != nil {
		return elemType, size, err
	}
	nSize, err2 := p.ReadI64()
	size = int(nSize)
	return elemType, size, err2
}

func (p *jsonProtocol) readElemListBegin() (elemType types.Type, size int, e error) {
	if isNull, e := p.ParseListBegin(); isNull || e != nil {
		return types.VOID, 0, e
	}
	sElemType, err := p.ReadString()
	if err != nil {
		return types.VOID, size, err
	}
	elemType, err = p.StringToTypeId(sElemType)
	if err != nil {
		return elemType, size, err
	}
	nSize, err2 := p.ReadI64()
	size = int(nSize)
	return elemType, size, err2
}

func (p *jsonProtocol) writeElemListBegin(elemType types.Type, size int) error {
	if e := p.OutputListBegin(); e != nil {
		return e
	}
	s, e1 := p.TypeIdToString(elemType)
	if e1 != nil {
		return e1
	}
	if e := p.OutputString(s); e != nil {
		return e
	}
	if e := p.OutputI64(int64(size)); e != nil {
		return e
	}
	return nil
}

func (p *jsonProtocol) TypeIdToString(fieldType types.Type) (string, error) {
	switch byte(fieldType) {
	case types.BOOL:
		return "tf", nil
	case types.BYTE:
		return "i8", nil
	case types.I16:
		return "i16", nil
	case types.I32:
		return "i32", nil
	case types.I64:
		return "i64", nil
	case types.DOUBLE:
		return "dbl", nil
	case types.FLOAT:
		return "flt", nil
	case types.STRING:
		return "str", nil
	case types.STRUCT:
		return "rec", nil
	case types.MAP:
		return "map", nil
	case types.SET:
		return "set", nil
	case types.LIST:
		return "lst", nil
	}

	e := fmt.Errorf("Unknown fieldType: %d", int(fieldType))
	return "", types.NewProtocolExceptionWithType(types.INVALID_DATA, e)
}

func (p *jsonProtocol) StringToTypeId(fieldType string) (types.Type, error) {
	switch fieldType {
	case "tf":
		return types.Type(types.BOOL), nil
	case "i8":
		return types.Type(types.BYTE), nil
	case "i16":
		return types.Type(types.I16), nil
	case "i32":
		return types.Type(types.I32), nil
	case "i64":
		return types.Type(types.I64), nil
	case "dbl":
		return types.Type(types.DOUBLE), nil
	case "flt":
		return types.Type(types.FLOAT), nil
	case "str":
		return types.Type(types.STRING), nil
	case "rec":
		return types.Type(types.STRUCT), nil
	case "map":
		return types.Type(types.MAP), nil
	case "set":
		return types.Type(types.SET), nil
	case "lst":
		return types.Type(types.LIST), nil
	}

	e := fmt.Errorf("Unknown type identifier: %s", fieldType)
	return types.Type(types.STOP), types.NewProtocolExceptionWithType(types.INVALID_DATA, e)
}
