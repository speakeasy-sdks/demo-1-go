# ResourceAccount

## Overview

ResourceAccount represents the account being used to access a resource.

Resource Accounts hold credentials that are required to provision and manage resources.
<SchemaDefinition schemaRef="#/components/schemas/ResourceAccountRequest" />


### Available Operations

* [GetOrgsOrgIDResourcesAccounts](#getorgsorgidresourcesaccounts) - List Resource Accounts in the organization.
* [GetOrgsOrgIDResourcesAccountsAccID](#getorgsorgidresourcesaccountsaccid) - Get a Resource Account.
* [PatchOrgsOrgIDResourcesAccountsAccID](#patchorgsorgidresourcesaccountsaccid) - Update a Resource Account.
* [PostOrgsOrgIDResourcesAccounts](#postorgsorgidresourcesaccounts) - Create a new Resource Account in the organization.

## GetOrgsOrgIDResourcesAccounts

List Resource Accounts in the organization.

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
    res, err := s.ResourceAccount.GetOrgsOrgIDResourcesAccounts(ctx, operations.GetOrgsOrgIDResourcesAccountsRequest{
        OrgID: "repudiandae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceAccountResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetOrgsOrgIDResourcesAccountsRequest](../../models/operations/getorgsorgidresourcesaccountsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetOrgsOrgIDResourcesAccountsResponse](../../models/operations/getorgsorgidresourcesaccountsresponse.md), error**


## GetOrgsOrgIDResourcesAccountsAccID

Get a Resource Account.

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
    res, err := s.ResourceAccount.GetOrgsOrgIDResourcesAccountsAccID(ctx, operations.GetOrgsOrgIDResourcesAccountsAccIDRequest{
        AccID: "cum",
        OrgID: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetOrgsOrgIDResourcesAccountsAccIDRequest](../../models/operations/getorgsorgidresourcesaccountsaccidrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetOrgsOrgIDResourcesAccountsAccIDResponse](../../models/operations/getorgsorgidresourcesaccountsaccidresponse.md), error**


## PatchOrgsOrgIDResourcesAccountsAccID

Update a Resource Account.

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
    res, err := s.ResourceAccount.PatchOrgsOrgIDResourcesAccountsAccID(ctx, operations.PatchOrgsOrgIDResourcesAccountsAccIDRequest{
        UpdateResourceAccountRequestRequest: shared.UpdateResourceAccountRequestRequest{
            Credentials: map[string]interface{}{
                "veniam": "animi",
                "dolores": "nam",
                "dicta": "consequuntur",
                "necessitatibus": "nobis",
            },
            Name: test1.String("Mr. Dora Wuckert"),
        },
        AccID: "pariatur",
        OrgID: "libero",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PatchOrgsOrgIDResourcesAccountsAccIDRequest](../../models/operations/patchorgsorgidresourcesaccountsaccidrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PatchOrgsOrgIDResourcesAccountsAccIDResponse](../../models/operations/patchorgsorgidresourcesaccountsaccidresponse.md), error**


## PostOrgsOrgIDResourcesAccounts

Create a new Resource Account in the organization.

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
    res, err := s.ResourceAccount.PostOrgsOrgIDResourcesAccounts(ctx, operations.PostOrgsOrgIDResourcesAccountsRequest{
        CreateResourceAccountRequestRequest: shared.CreateResourceAccountRequestRequest{
            Credentials: map[string]interface{}{
                "occaecati": "nemo",
                "aliquam": "nostrum",
                "doloribus": "eligendi",
            },
            ID: test1.String("95fa8897-0e18-49db-b30f-cb33ea055b19"),
            Name: test1.String("Sheri Schuppe"),
            Type: test1.String("itaque"),
        },
        OrgID: "sed",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostOrgsOrgIDResourcesAccountsRequest](../../models/operations/postorgsorgidresourcesaccountsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostOrgsOrgIDResourcesAccountsResponse](../../models/operations/postorgsorgidresourcesaccountsresponse.md), error**

