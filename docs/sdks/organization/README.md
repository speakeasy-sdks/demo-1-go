# Organization

## Overview

An Organization is the top level object in Humanitec. All other objects belong to an Organization.
<SchemaDefinition schemaRef="#/components/schemas/OrganizationRequest" />


### Available Operations

* [GetOrgs](#getorgs) - List active organizations the user has access to.
* [GetOrgsOrgID](#getorgsorgid) - Get the specified Organization.

## GetOrgs

List active organizations the user has access to.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.Organization.GetOrgs(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.OrganizationResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetOrgsResponse](../../models/operations/getorgsresponse.md), error**


## GetOrgsOrgID

Get the specified Organization.

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
    res, err := s.Organization.GetOrgsOrgID(ctx, operations.GetOrgsOrgIDRequest{
        OrgID: "cum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrganizationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetOrgsOrgIDRequest](../../models/operations/getorgsorgidrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetOrgsOrgIDResponse](../../models/operations/getorgsorgidresponse.md), error**

