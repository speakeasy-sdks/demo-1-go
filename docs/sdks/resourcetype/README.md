# ResourceType

## Overview

Resources Types define the technology that Applications can have dependencies on.

Each Resource Type also defines a set of input parameters (`inputs_schema`), and a set of output data (`outputs_schema`). When provisioning a resource, these are treated as inputs and outputs respectively.
<SchemaDefinition schemaRef="#/components/schemas/ResourceTypeRequest" />


### Available Operations

* [GetOrgsOrgIDResourcesTypes](#getorgsorgidresourcestypes) - List Resource Types.

## GetOrgsOrgIDResourcesTypes

List Resource Types.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.ResourceType.GetOrgsOrgIDResourcesTypes(ctx, operations.GetOrgsOrgIDResourcesTypesRequest{
        OrgID: "modi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceTypeResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetOrgsOrgIDResourcesTypesRequest](../../models/operations/getorgsorgidresourcestypesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetOrgsOrgIDResourcesTypesResponse](../../models/operations/getorgsorgidresourcestypesresponse.md), error**

