// Autogenerated by Thrift for thrift/compiler/test/fixtures/interactions/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module

import (
    "fmt"
    "strings"

    shared "shared"
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

var _ = shared.GoUnusedProtection__
// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = strings.Split
var _ = thrift.ZERO


type CustomException struct {
    Message string `thrift:"message,1" json:"message" db:"message"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*CustomException)(nil)

func NewCustomException() *CustomException {
    return (&CustomException{}).
        SetMessageNonCompat("")
}

func (x *CustomException) GetMessage() string {
    return x.Message
}

func (x *CustomException) SetMessageNonCompat(value string) *CustomException {
    x.Message = value
    return x
}

func (x *CustomException) SetMessage(value string) *CustomException {
    x.Message = value
    return x
}

func (x *CustomException) writeField1(p thrift.Encoder) error {  // Message
    if err := p.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Message
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *CustomException) readField1(p thrift.Decoder) error {  // Message
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Message = result
    return nil
}

func (x *CustomException) toString1() string {  // Message
    return fmt.Sprintf("%v", x.Message)
}



func (x *CustomException) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("CustomException"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *CustomException) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        var fieldReadErr error
        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRING)):  // message
            fieldReadErr = x.readField1(p)
        default:
            fieldReadErr = p.Skip(wireType)
        }

        if fieldReadErr != nil {
            return fieldReadErr
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *CustomException) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("CustomException({")
    sb.WriteString(fmt.Sprintf("Message:%s", x.toString1()))
    sb.WriteString("})")

    return sb.String()
}
func (x *CustomException) Error() string {
    return x.String()
}

// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {

}
