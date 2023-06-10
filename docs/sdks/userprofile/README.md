# UserProfile

## Overview

UserProfile holds the profile information of a user
<SchemaDefinition schemaRef="#/components/schemas/UserProfileRequest" />


### Available Operations

* [DeleteTokensTokenID](#deletetokenstokenid) - DEPRECATED
* [GetCurrentUser](#getcurrentuser) - Gets the extended profile of the current user
* [GetTokens](#gettokens) - DEPRECATED
* [GetUsersMe](#getusersme) - DEPRECATED
* [PatchCurrentUser](#patchcurrentuser) - Updates the extended profile of the current user.
* [PostOrgsOrgIDUsers](#postorgsorgidusers) - Creates a new service user.

## DeleteTokensTokenID

DEPRECATED

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
    res, err := s.UserProfile.DeleteTokensTokenID(ctx, operations.DeleteTokensTokenIDRequest{
        TokenID: "eligendi",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteTokensTokenIDRequest](../../models/operations/deletetokenstokenidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteTokensTokenIDResponse](../../models/operations/deletetokenstokenidresponse.md), error**


## GetCurrentUser

Gets the extended profile of the current user

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
    res, err := s.UserProfile.GetCurrentUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UserProfileExtendedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCurrentUserResponse](../../models/operations/getcurrentuserresponse.md), error**


## GetTokens

DEPRECATED

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
    res, err := s.UserProfile.GetTokens(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTokens200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTokensResponse](../../models/operations/gettokensresponse.md), error**


## GetUsersMe

DEPRECATED

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
    res, err := s.UserProfile.GetUsersMe(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUsersMe200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUsersMeResponse](../../models/operations/getusersmeresponse.md), error**


## PatchCurrentUser

Updates the extended profile of the current user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/shared"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.UserProfile.PatchCurrentUser(ctx, shared.UserProfileExtendedRequest{
        CreatedAt: test1.String("2020-06-22T09:37:23.523Z"),
        Email: test1.String("Gunnar98@hotmail.com"),
        ID: test1.String("eeb1c7cb-db6e-4ec7-8378-ba25317747dc"),
        Name: test1.String("Gregory Heidenreich"),
        Properties: map[string]interface{}{
            "maxime": "dolorum",
        },
        Roles: map[string]string{
            "nostrum": "illum",
            "quibusdam": "commodi",
            "esse": "explicabo",
            "consectetur": "temporibus",
        },
        Type: test1.String("optio"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserProfileExtendedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.UserProfileExtendedRequest](../../models/shared/userprofileextendedrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PatchCurrentUserResponse](../../models/operations/patchcurrentuserresponse.md), error**


## PostOrgsOrgIDUsers

Creates a new service user.

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
    res, err := s.UserProfile.PostOrgsOrgIDUsers(ctx, operations.PostOrgsOrgIDUsersRequest{
        NewServiceUserRequest: shared.NewServiceUserRequest{
            Email: test1.String("Wilford_Heller@gmail.com"),
            Name: test1.String("Dixie Doyle"),
            Role: test1.String("harum"),
        },
        OrgID: "ducimus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserProfileExtendedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PostOrgsOrgIDUsersRequest](../../models/operations/postorgsorgidusersrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PostOrgsOrgIDUsersResponse](../../models/operations/postorgsorgidusersresponse.md), error**

