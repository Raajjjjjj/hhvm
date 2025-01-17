// Autogenerated by Thrift for thrift/compiler/test/fixtures/interactions/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module

import (
    "maps"

    shared "shared"
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

var _ = shared.GoUnusedProtection__
// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = maps.Copy[map[int]int, map[int]int]
var _ = metadata.GoUnusedProtection__

// Premade Thrift types
var (
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
            )
    premadeThriftType_void = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_VOID_TYPE.Ptr(),
            )
    premadeThriftType_shared_DoSomethingResult = metadata.NewThriftType().SetTStruct(
        metadata.NewThriftStructType().
            SetName("shared.DoSomethingResult"),
            )
)

var structMetadatas = []*metadata.ThriftStruct{
}

var exceptionMetadatas = []*metadata.ThriftException{
    metadata.NewThriftException().
    SetName("module.CustomException").
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("message").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
}

var enumMetadatas = []*metadata.ThriftEnum{
}

var serviceMetadatas = []*metadata.ThriftService{
    metadata.NewThriftService().
    SetName("module.MyService").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("foo").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void),
        },
    ),
    metadata.NewThriftService().
    SetName("module.Factories").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("foo").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void),
        },
    ),
    metadata.NewThriftService().
    SetName("module.Perform").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("foo").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_void),
        },
    ),
    metadata.NewThriftService().
    SetName("module.InteractWithShared").
    SetFunctions(
        []*metadata.ThriftFunction{
            metadata.NewThriftFunction().
    SetName("do_some_similar_things").
    SetIsOneway(false).
    SetReturnType(premadeThriftType_shared_DoSomethingResult),
        },
    ),
}

// GetThriftMetadata returns complete Thrift metadata for current and imported packages.
func GetThriftMetadata() *metadata.ThriftMetadata {
    allEnums := GetEnumsMetadata()
    allStructs := GetStructsMetadata()
    allExceptions := GetExceptionsMetadata()
    allServices := GetServicesMetadata()

    return metadata.NewThriftMetadata().
        SetEnums(allEnums).
        SetStructs(allStructs).
        SetExceptions(allExceptions).
        SetServices(allServices)
}

// GetEnumsMetadata returns Thrift metadata for enums in the current and recursively included packages.
func GetEnumsMetadata() map[string]*metadata.ThriftEnum {
    allEnumsMap := make(map[string]*metadata.ThriftEnum)

    // Add enum metadatas from the current program...
    for _, enumMetadata := range enumMetadatas {
        allEnumsMap[enumMetadata.GetName()] = enumMetadata
    }

    // ...now add enum metadatas from recursively included programs.
    maps.Copy(allEnumsMap, shared.GetEnumsMetadata())

    return allEnumsMap
}

// GetStructsMetadata returns Thrift metadata for structs in the current and recursively included packages.
func GetStructsMetadata() map[string]*metadata.ThriftStruct {
    allStructsMap := make(map[string]*metadata.ThriftStruct)

    // Add struct metadatas from the current program...
    for _, structMetadata := range structMetadatas {
        allStructsMap[structMetadata.GetName()] = structMetadata
    }

    // ...now add struct metadatas from recursively included programs.
    maps.Copy(allStructsMap, shared.GetStructsMetadata())

    return allStructsMap
}

// GetExceptionsMetadata returns Thrift metadata for exceptions in the current and recursively included packages.
func GetExceptionsMetadata() map[string]*metadata.ThriftException {
    allExceptionsMap := make(map[string]*metadata.ThriftException)

    // Add exception metadatas from the current program...
    for _, exceptionMetadata := range exceptionMetadatas {
        allExceptionsMap[exceptionMetadata.GetName()] = exceptionMetadata
    }

    // ...now add exception metadatas from recursively included programs.
    maps.Copy(allExceptionsMap, shared.GetExceptionsMetadata())

    return allExceptionsMap
}

// GetServicesMetadata returns Thrift metadata for services in the current and recursively included packages.
func GetServicesMetadata() map[string]*metadata.ThriftService {
    allServicesMap := make(map[string]*metadata.ThriftService)

    // Add service metadatas from the current program...
    for _, serviceMetadata := range serviceMetadatas {
        allServicesMap[serviceMetadata.GetName()] = serviceMetadata
    }

    // ...now add service metadatas from recursively included programs.
    maps.Copy(allServicesMap, shared.GetServicesMetadata())

    return allServicesMap
}

// GetThriftMetadataForService returns Thrift metadata for the given service.
func GetThriftMetadataForService(scopedServiceName string) *metadata.ThriftMetadata {
    thriftMetadata := GetThriftMetadata()

    allServicesMap := thriftMetadata.GetServices()
    relevantServicesMap := make(map[string]*metadata.ThriftService)

    serviceMetadata := allServicesMap[scopedServiceName]
    // Visit and record all recursive parents of the target service.
    for serviceMetadata != nil {
        relevantServicesMap[serviceMetadata.GetName()] = serviceMetadata
        if serviceMetadata.IsSetParent() {
            serviceMetadata = allServicesMap[serviceMetadata.GetParent()]
        } else {
            serviceMetadata = nil
        }
    }

    thriftMetadata.SetServices(relevantServicesMap)

    return thriftMetadata
}
