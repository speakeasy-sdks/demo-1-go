# DriverDefinition

## Overview

DriverDefinition describes the resource driver.

Resource Drivers are code that fulfils the Humanitec Resource Driver Interface. This interface allows for certain actions to be performed on resources such as creation and destruction. Humanitec provides numerous Resource Drivers “out of the box”. It is also possible to use 3rd party Resource Drivers or write your own.
<SchemaDefinition schemaRef="#/components/schemas/DriverDefinitionRequest" />


### Available Operations

* [DeleteOrgsOrgIDResourcesDriversDriverID](#deleteorgsorgidresourcesdriversdriverid) - Delete a Resources Driver.
* [GetOrgsOrgIDResourcesDrivers](#getorgsorgidresourcesdrivers) - List Resource Drivers.
* [GetOrgsOrgIDResourcesDriversDriverID](#getorgsorgidresourcesdriversdriverid) - Get a Resource Driver.
* [PostOrgsOrgIDResourcesDrivers](#postorgsorgidresourcesdrivers) - Register a new Resource Driver.
* [PutOrgsOrgIDResourcesDriversDriverID](#putorgsorgidresourcesdriversdriverid) - Update a Resource Driver.

## DeleteOrgsOrgIDResourcesDriversDriverID

Delete a Resources Driver.

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
    res, err := s.DriverDefinition.DeleteOrgsOrgIDResourcesDriversDriverID(ctx, operations.DeleteOrgsOrgIDResourcesDriversDriverIDRequest{
        DriverID: "architecto",
        OrgID: "sit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.DeleteOrgsOrgIDResourcesDriversDriverIDRequest](../../models/operations/deleteorgsorgidresourcesdriversdriveridrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.DeleteOrgsOrgIDResourcesDriversDriverIDResponse](../../models/operations/deleteorgsorgidresourcesdriversdriveridresponse.md), error**


## GetOrgsOrgIDResourcesDrivers

List Resource Drivers.

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
    res, err := s.DriverDefinition.GetOrgsOrgIDResourcesDrivers(ctx, operations.GetOrgsOrgIDResourcesDriversRequest{
        OrgID: "modi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DriverDefinitionResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetOrgsOrgIDResourcesDriversRequest](../../models/operations/getorgsorgidresourcesdriversrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetOrgsOrgIDResourcesDriversResponse](../../models/operations/getorgsorgidresourcesdriversresponse.md), error**


## GetOrgsOrgIDResourcesDriversDriverID

# Only drivers that belongs to the given organization or registered as `public` are accessible through this endpoint

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
    res, err := s.DriverDefinition.GetOrgsOrgIDResourcesDriversDriverID(ctx, operations.GetOrgsOrgIDResourcesDriversDriverIDRequest{
        DriverID: "fugit",
        OrgID: "ab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DriverDefinitionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.GetOrgsOrgIDResourcesDriversDriverIDRequest](../../models/operations/getorgsorgidresourcesdriversdriveridrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.GetOrgsOrgIDResourcesDriversDriverIDResponse](../../models/operations/getorgsorgidresourcesdriversdriveridresponse.md), error**


## PostOrgsOrgIDResourcesDrivers

Register a new Resource Driver.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.DriverDefinition.PostOrgsOrgIDResourcesDrivers(ctx, operations.PostOrgsOrgIDResourcesDriversRequest{
        CreateDriverRequestRequest: shared.CreateDriverRequestRequest{
            AccountTypes: []string{
                "quae",
                "dolor",
                "fugiat",
            },
            ID: test1.String("5208ece7-e253-4b66-8451-c6c6e205e16d"),
            InputsSchema: map[string]interface{}{
                "est": "harum",
                "sequi": "doloribus",
                "repudiandae": "optio",
                "occaecati": "nemo",
            },
            IsPublic: test1.Bool(false),
            Target: "voluptate",
            Template: test1.String("blanditiis"),
            Type: "officia",
        },
        OrgID: "voluptas",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DriverDefinitionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostOrgsOrgIDResourcesDriversRequest](../../models/operations/postorgsorgidresourcesdriversrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PostOrgsOrgIDResourcesDriversResponse](../../models/operations/postorgsorgidresourcesdriversresponse.md), error**


## PutOrgsOrgIDResourcesDriversDriverID

Update a Resource Driver.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.DriverDefinition.PutOrgsOrgIDResourcesDriversDriverID(ctx, operations.PutOrgsOrgIDResourcesDriversDriverIDRequest{
        UpdateDriverRequestRequest: shared.UpdateDriverRequestRequest{
            AccountTypes: []string{
                "nemo",
                "quos",
            },
            InputsSchema: map[string]interface{}{
                "aspernatur": "ducimus",
                "nesciunt": "fuga",
            },
            IsPublic: test1.Bool(false),
            Target: "laudantium",
            Template: test1.String("incidunt"),
            Type: "quasi",
        },
        DriverID: "rem",
        OrgID: "fugiat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DriverDefinitionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PutOrgsOrgIDResourcesDriversDriverIDRequest](../../models/operations/putorgsorgidresourcesdriversdriveridrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PutOrgsOrgIDResourcesDriversDriverIDResponse](../../models/operations/putorgsorgidresourcesdriversdriveridresponse.md), error**

