# AccountType

## Overview

Resource Account Types define cloud providers or protocols to which a resource account can belong.
<SchemaDefinition schemaRef="#/components/schemas/AccountTypeRequest" />


### Available Operations

* [GetOrgsOrgIDResourcesAccountTypes](#getorgsorgidresourcesaccounttypes) - List Resource Account Types available to the organization.

## GetOrgsOrgIDResourcesAccountTypes

List Resource Account Types available to the organization.

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
    res, err := s.AccountType.GetOrgsOrgIDResourcesAccountTypes(ctx, operations.GetOrgsOrgIDResourcesAccountTypesRequest{
        OrgID: "provident",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountTypeResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GetOrgsOrgIDResourcesAccountTypesRequest](../../models/operations/getorgsorgidresourcesaccounttypesrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.GetOrgsOrgIDResourcesAccountTypesResponse](../../models/operations/getorgsorgidresourcesaccounttypesresponse.md), error**

