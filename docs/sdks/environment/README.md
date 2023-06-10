# Environment

## Overview

Environments are independent spaces where Applications can run. An Application is always deployed into an Environment.
<SchemaDefinition schemaRef="#/components/schemas/EnvironmentRequest" />


### Available Operations

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvID](#deleteorgsorgidappsappidenvsenvid) - Delete a specific Environment.
* [GetOrgsOrgIDAppsAppIDEnvs](#getorgsorgidappsappidenvs) - List all Environments.
* [GetOrgsOrgIDAppsAppIDEnvsEnvID](#getorgsorgidappsappidenvsenvid) - Get a specific Environment.
* [PostOrgsOrgIDAppsAppIDEnvs](#postorgsorgidappsappidenvs) - Add a new Environment to an Application.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployID](#putorgsorgidappsappidenvsenvidfromdeployid) - Rebase to a different Deployment.

## DeleteOrgsOrgIDAppsAppIDEnvsEnvID

Deletes a specific Environment in an Application.

Deleting an Environment will also delete the Deployment history of the Environment.

_Deletions are currently irreversible._

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
    res, err := s.Environment.DeleteOrgsOrgIDAppsAppIDEnvsEnvID(ctx, operations.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRequest{
        AppID: "dicta",
        EnvID: "nisi",
        OrgID: "consequuntur",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRequest](../../models/operations/deleteorgsorgidappsappidenvsenvidrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResponse](../../models/operations/deleteorgsorgidappsappidenvsenvidresponse.md), error**


## GetOrgsOrgIDAppsAppIDEnvs

Lists all of the Environments in the Application.

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
    res, err := s.Environment.GetOrgsOrgIDAppsAppIDEnvs(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsRequest{
        AppID: "consectetur",
        OrgID: "aperiam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetOrgsOrgIDAppsAppIDEnvsRequest](../../models/operations/getorgsorgidappsappidenvsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDEnvsResponse](../../models/operations/getorgsorgidappsappidenvsresponse.md), error**


## GetOrgsOrgIDAppsAppIDEnvsEnvID

Gets a specific Environment in an Application.

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
    res, err := s.Environment.GetOrgsOrgIDAppsAppIDEnvsEnvID(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDRequest{
        AppID: "cupiditate",
        EnvID: "reiciendis",
        OrgID: "soluta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDRequest](../../models/operations/getorgsorgidappsappidenvsenvidrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDResponse](../../models/operations/getorgsorgidappsappidenvsenvidresponse.md), error**


## PostOrgsOrgIDAppsAppIDEnvs

Creates a new Environment of the specified Type and associates it with the Application specified by `appId`.

The Environment is also initialized to the **current or past state of Deployment in another Environment**. This ensures that every Environment is derived from a previously known state. This means it is not possible to create a new Environment for an Application until at least one Deployment has occurred. (The Deployment does not have to be successful.)

The Type of the Environment must be already defined in the Organization.

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
    res, err := s.Environment.PostOrgsOrgIDAppsAppIDEnvs(ctx, operations.PostOrgsOrgIDAppsAppIDEnvsRequest{
        EnvironmentDefinitionRequest: shared.EnvironmentDefinitionRequest{
            FromDeployID: "alias",
            ID: "929921ae-fb9f-458c-8d86-e68e4be05601",
            Name: "Shawna Hamill",
            Type: "deserunt",
        },
        AppID: "esse",
        OrgID: "nemo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostOrgsOrgIDAppsAppIDEnvsRequest](../../models/operations/postorgsorgidappsappidenvsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostOrgsOrgIDAppsAppIDEnvsResponse](../../models/operations/postorgsorgidappsappidenvsresponse.md), error**


## PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployID

Rebasing an Environment means that the next Deployment to the Environment will be based on the Deployment specified in the rebase rather than the last one in the Environment. The Deployment to rebase to can either be current or a previous Deployment. The Deployment can be from any Environment of the same Application.

_Running code will only be affected on the next Deployment to the Environment._

Common use cases for rebasing an Environment:

* _Rollback_: Rebasing to a previous Deployment in the current Environment and then Deploying without additional changes will execute a rollback to the previous Deployment state.

* _Clone_: Rebasing to the current Deployment in a different Environment and then deploying without additional changes will clone all of the configuration of the other Environment into the current one. (NOTE: External Resources will not be cloned in the process - the current External Resources of the Environment will remain unchanged and will be used by the deployed Application in the Environment.

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
    res, err := s.Environment.PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployID(ctx, operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployIDRequest{
        RequestBody: "reprehenderit",
        AppID: "est",
        EnvID: "quis",
        OrgID: "sint",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployIDRequest](../../models/operations/putorgsorgidappsappidenvsenvidfromdeployidrequest.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |


### Response

**[*operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployIDResponse](../../models/operations/putorgsorgidappsappidenvsenvidfromdeployidresponse.md), error**

