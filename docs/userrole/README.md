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
