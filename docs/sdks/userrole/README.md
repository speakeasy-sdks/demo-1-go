# UserRole

## Overview

UserRole holds the mapping of role to user for a particular object.
<SchemaDefinition schemaRef="#/components/schemas/UserRoleRequest" />


### Available Operations

* [DeleteOrgsOrgIDAppsAppIDUsersUserID](#deleteorgsorgidappsappidusersuserid) - Remove the role of a User on an Application
* [DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserID](#deleteorgsorgidenvtypesenvtypeusersuserid) - Remove the role of a User on an Environment Type
* [DeleteOrgsOrgIDUsersUserID](#deleteorgsorgidusersuserid) - Remove the role of a User on an Organization
* [GetOrgsOrgIDAppsAppIDUsers](#getorgsorgidappsappidusers) - List Users with roles in an App
* [GetOrgsOrgIDAppsAppIDUsersUserID](#getorgsorgidappsappidusersuserid) - Get the role of a User on an Application
* [GetOrgsOrgIDEnvTypesEnvTypeUsersUserID](#getorgsorgidenvtypesenvtypeusersuserid) - Get the role of a User on an Environment Type
* [GetOrgsOrgIDUsers](#getorgsorgidusers) - List Users with roles in an Organization
* [GetOrgsOrgIDUsersUserID](#getorgsorgidusersuserid) - Get the role of a User on an Organization
* [PatchOrgsOrgIDAppsAppIDUsersUserID](#patchorgsorgidappsappidusersuserid) - Update the role of a User on an Application
* [PatchOrgsOrgIDEnvTypesEnvTypeUsersUserID](#patchorgsorgidenvtypesenvtypeusersuserid) - Update the role of a User on an Environment Type
* [PatchOrgsOrgIDUsersUserID](#patchorgsorgidusersuserid) - Update the role of a User on an Organization
* [PostOrgsOrgIDAppsAppIDUsers](#postorgsorgidappsappidusers) - Adds a User to an Application with a Role
* [PostOrgsOrgIDEnvTypesEnvTypeUsers](#postorgsorgidenvtypesenvtypeusers) - Adds a User to an Environment Type with a Role
* [PostOrgsOrgIDInvitations](#postorgsorgidinvitations) - Invites a user to an Organization with a specified role.

## DeleteOrgsOrgIDAppsAppIDUsersUserID

Remove the role of a User on an Application

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
    res, err := s.UserRole.DeleteOrgsOrgIDAppsAppIDUsersUserID(ctx, operations.DeleteOrgsOrgIDAppsAppIDUsersUserIDRequest{
        AppID: "doloremque",
        OrgID: "perferendis",
        UserID: "laudantium",
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.DeleteOrgsOrgIDAppsAppIDUsersUserIDRequest](../../models/operations/deleteorgsorgidappsappidusersuseridrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.DeleteOrgsOrgIDAppsAppIDUsersUserIDResponse](../../models/operations/deleteorgsorgidappsappidusersuseridresponse.md), error**


## DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserID

Remove the role of a User on an Environment Type

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
    res, err := s.UserRole.DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserID(ctx, operations.DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserIDRequest{
        EnvType: "iusto",
        OrgID: "corrupti",
        UserID: "molestiae",
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserIDRequest](../../models/operations/deleteorgsorgidenvtypesenvtypeusersuseridrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserIDResponse](../../models/operations/deleteorgsorgidenvtypesenvtypeusersuseridresponse.md), error**


## DeleteOrgsOrgIDUsersUserID

Remove the role of a User on an Organization

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
    res, err := s.UserRole.DeleteOrgsOrgIDUsersUserID(ctx, operations.DeleteOrgsOrgIDUsersUserIDRequest{
        OrgID: "quis",
        UserID: "iure",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.DeleteOrgsOrgIDUsersUserIDRequest](../../models/operations/deleteorgsorgidusersuseridrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.DeleteOrgsOrgIDUsersUserIDResponse](../../models/operations/deleteorgsorgidusersuseridresponse.md), error**


## GetOrgsOrgIDAppsAppIDUsers

List Users with roles in an App

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
    res, err := s.UserRole.GetOrgsOrgIDAppsAppIDUsers(ctx, operations.GetOrgsOrgIDAppsAppIDUsersRequest{
        AppID: "ab",
        OrgID: "quaerat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetOrgsOrgIDAppsAppIDUsersRequest](../../models/operations/getorgsorgidappsappidusersrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDUsersResponse](../../models/operations/getorgsorgidappsappidusersresponse.md), error**


## GetOrgsOrgIDAppsAppIDUsersUserID

Get the role of a User on an Application

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
    res, err := s.UserRole.GetOrgsOrgIDAppsAppIDUsersUserID(ctx, operations.GetOrgsOrgIDAppsAppIDUsersUserIDRequest{
        AppID: "amet",
        OrgID: "sapiente",
        UserID: "corporis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.GetOrgsOrgIDAppsAppIDUsersUserIDRequest](../../models/operations/getorgsorgidappsappidusersuseridrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDUsersUserIDResponse](../../models/operations/getorgsorgidappsappidusersuseridresponse.md), error**


## GetOrgsOrgIDEnvTypesEnvTypeUsersUserID

Get the role of a User on an Environment Type

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
    res, err := s.UserRole.GetOrgsOrgIDEnvTypesEnvTypeUsersUserID(ctx, operations.GetOrgsOrgIDEnvTypesEnvTypeUsersUserIDRequest{
        EnvType: "est",
        OrgID: "iure",
        UserID: "quisquam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetOrgsOrgIDEnvTypesEnvTypeUsersUserIDRequest](../../models/operations/getorgsorgidenvtypesenvtypeusersuseridrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.GetOrgsOrgIDEnvTypesEnvTypeUsersUserIDResponse](../../models/operations/getorgsorgidenvtypesenvtypeusersuseridresponse.md), error**


## GetOrgsOrgIDUsers

List Users with roles in an Organization

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
    res, err := s.UserRole.GetOrgsOrgIDUsers(ctx, operations.GetOrgsOrgIDUsersRequest{
        OrgID: "provident",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetOrgsOrgIDUsersRequest](../../models/operations/getorgsorgidusersrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetOrgsOrgIDUsersResponse](../../models/operations/getorgsorgidusersresponse.md), error**


## GetOrgsOrgIDUsersUserID

Get the role of a User on an Organization

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
    res, err := s.UserRole.GetOrgsOrgIDUsersUserID(ctx, operations.GetOrgsOrgIDUsersUserIDRequest{
        OrgID: "laudantium",
        UserID: "nam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetOrgsOrgIDUsersUserIDRequest](../../models/operations/getorgsorgidusersuseridrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetOrgsOrgIDUsersUserIDResponse](../../models/operations/getorgsorgidusersuseridresponse.md), error**


## PatchOrgsOrgIDAppsAppIDUsersUserID

Update the role of a User on an Application

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
    res, err := s.UserRole.PatchOrgsOrgIDAppsAppIDUsersUserID(ctx, operations.PatchOrgsOrgIDAppsAppIDUsersUserIDRequest{
        RoleRequest: shared.RoleRequest{
            Role: test1.String("nemo"),
        },
        AppID: "enim",
        OrgID: "ipsam",
        UserID: "minima",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.PatchOrgsOrgIDAppsAppIDUsersUserIDRequest](../../models/operations/patchorgsorgidappsappidusersuseridrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.PatchOrgsOrgIDAppsAppIDUsersUserIDResponse](../../models/operations/patchorgsorgidappsappidusersuseridresponse.md), error**


## PatchOrgsOrgIDEnvTypesEnvTypeUsersUserID

Update the role of a User on an Environment Type

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
    res, err := s.UserRole.PatchOrgsOrgIDEnvTypesEnvTypeUsersUserID(ctx, operations.PatchOrgsOrgIDEnvTypesEnvTypeUsersUserIDRequest{
        RoleRequest: shared.RoleRequest{
            Role: test1.String("tempora"),
        },
        EnvType: "perferendis",
        OrgID: "corrupti",
        UserID: "doloremque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.PatchOrgsOrgIDEnvTypesEnvTypeUsersUserIDRequest](../../models/operations/patchorgsorgidenvtypesenvtypeusersuseridrequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |


### Response

**[*operations.PatchOrgsOrgIDEnvTypesEnvTypeUsersUserIDResponse](../../models/operations/patchorgsorgidenvtypesenvtypeusersuseridresponse.md), error**


## PatchOrgsOrgIDUsersUserID

Update the role of a User on an Organization

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
    res, err := s.UserRole.PatchOrgsOrgIDUsersUserID(ctx, operations.PatchOrgsOrgIDUsersUserIDRequest{
        RoleRequest: shared.RoleRequest{
            Role: test1.String("fugiat"),
        },
        OrgID: "numquam",
        UserID: "doloremque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PatchOrgsOrgIDUsersUserIDRequest](../../models/operations/patchorgsorgidusersuseridrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PatchOrgsOrgIDUsersUserIDResponse](../../models/operations/patchorgsorgidusersuseridresponse.md), error**


## PostOrgsOrgIDAppsAppIDUsers

Adds a User to an Application with a Role

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
    res, err := s.UserRole.PostOrgsOrgIDAppsAppIDUsers(ctx, operations.PostOrgsOrgIDAppsAppIDUsersRequest{
        UserRoleRequest: shared.UserRoleRequest{
            CreatedAt: test1.String("2020-06-22T09:37:23.523Z"),
            Email: test1.String("Neva_Murphy39@hotmail.com"),
            ID: test1.String("cbd6b5f3-ec90-4930-8f92-6bad2553819b"),
            Invite: test1.String("tempora"),
            Name: test1.String("Dr. Monica Ratke"),
            Role: test1.String("fugit"),
            Type: test1.String("voluptatem"),
            User: test1.String("repudiandae"),
        },
        AppID: "corporis",
        OrgID: "ea",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PostOrgsOrgIDAppsAppIDUsersRequest](../../models/operations/postorgsorgidappsappidusersrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PostOrgsOrgIDAppsAppIDUsersResponse](../../models/operations/postorgsorgidappsappidusersresponse.md), error**


## PostOrgsOrgIDEnvTypesEnvTypeUsers

Adds a User to an Environment Type with a Role

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
    res, err := s.UserRole.PostOrgsOrgIDEnvTypesEnvTypeUsers(ctx, operations.PostOrgsOrgIDEnvTypesEnvTypeUsersRequest{
        UserRoleRequest: shared.UserRoleRequest{
            CreatedAt: test1.String("2020-06-22T09:37:23.523Z"),
            Email: test1.String("Elmer_Kuvalis99@gmail.com"),
            ID: test1.String("639a910a-bdca-4b62-a766-96e1ec00221b"),
            Invite: test1.String("sequi"),
            Name: test1.String("Yvonne Stamm"),
            Role: test1.String("similique"),
            Type: test1.String("eligendi"),
            User: test1.String("tempore"),
        },
        EnvType: "amet",
        OrgID: "debitis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PostOrgsOrgIDEnvTypesEnvTypeUsersRequest](../../models/operations/postorgsorgidenvtypesenvtypeusersrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.PostOrgsOrgIDEnvTypesEnvTypeUsersResponse](../../models/operations/postorgsorgidenvtypesenvtypeusersresponse.md), error**


## PostOrgsOrgIDInvitations

Invites a user to an Organization with a specified role.

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
    res, err := s.UserRole.PostOrgsOrgIDInvitations(ctx, operations.PostOrgsOrgIDInvitationsRequest{
        UserInviteRequestRequest: shared.UserInviteRequestRequest{
            Email: "Yasmine_Smith@hotmail.com",
            Role: "quibusdam",
        },
        OrgID: "sit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserRoleResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostOrgsOrgIDInvitationsRequest](../../models/operations/postorgsorgidinvitationsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PostOrgsOrgIDInvitationsResponse](../../models/operations/postorgsorgidinvitationsresponse.md), error**

