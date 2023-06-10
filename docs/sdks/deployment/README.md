# Deployment

## Overview

Deployments represent updates to the running state of an Environment.

Deployments are made by applying _Deltas_ to a state defined by an existing Deployment. The Environment’s from_deploy property defines the Deployment. This Deployment is usually but not always in the current Environment. If the Deployment is from another Environment, the state of that Environment will be "cloned" into the current Environment with the option to apply a Delta.
<SchemaDefinition schemaRef="#/components/schemas/DeploymentRequest" />


### Available Operations

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](#getorgsorgidappsappidenvsenviddeploys) - List Deployments in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployID](#getorgsorgidappsappidenvsenviddeploysdeployid) - Get a specific Deployment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrors](#getorgsorgidappsappidenvsenviddeploysdeployiderrors) - List errors that occurred in a Deployment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](#postorgsorgidappsappidenvsenviddeploys) - Start a new Deployment.

## GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploys

List all of the Deployments that have been carried out in the current Environment. Deployments are returned with the newest first.

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
    res, err := s.Deployment.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploys(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysRequest{
        AppID: "ullam",
        EnvID: "perferendis",
        OrgID: "illum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeploymentResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysRequest](../../models/operations/getorgsorgidappsappidenvsenviddeploysrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysResponse](../../models/operations/getorgsorgidappsappidenvsenviddeploysresponse.md), error**


## GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployID

Gets a specific Deployment in an Application and an Environment.

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
    res, err := s.Deployment.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployID(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDRequest{
        AppID: "totam",
        DeployID: "impedit",
        EnvID: "quibusdam",
        OrgID: "nam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeploymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDRequest](../../models/operations/getorgsorgidappsappidenvsenviddeploysdeployidrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDResponse](../../models/operations/getorgsorgidappsappidenvsenviddeploysdeployidresponse.md), error**


## GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrors

List errors that occurred in a Deployment.

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
    res, err := s.Deployment.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrors(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrorsRequest{
        AppID: "ipsam",
        DeployID: "culpa",
        EnvID: "dolor",
        OrgID: "aliquam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeploymentErrorResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrorsRequest](../../models/operations/getorgsorgidappsappidenvsenviddeploysdeployiderrorsrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrorsResponse](../../models/operations/getorgsorgidappsappidenvsenviddeploysdeployiderrorsresponse.md), error**


## PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploys

At Humanitec, Deployments are defined as changes to the state of the Environment. The state can be changed by defining a set of desired changes to the current state via a Deployment Delta or by resetting the current state after a previous Deployment. (See Environment Rebase.) Both types of changes can be combined into a single Deployment during which the Delta is applied to the Rebased state.

When specifying a Delta, a Delta ID must be used. That Delta must have been committed to the Delta store prior to the Deployment.

A Set ID can also be defined in the deployment to force the state of the environment to a particular state. This will be ignored if the Delta is specified.

**NOTE:**

Directly setting a `set_id` in a deployment is not recommended as it will not record history of where the set came from. If the intention is to replicate an existing environment, use the environment rebasing approach described above.

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
    res, err := s.Deployment.PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploys(ctx, operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploysRequest{
        DeploymentRequest: shared.DeploymentRequest{
            Comment: test1.String("inventore"),
            DeltaID: test1.String("deleniti"),
            ValueSetVersionID: test1.String("veritatis"),
        },
        AppID: "tempora",
        EnvID: "dolor",
        OrgID: "consequatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeploymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploysRequest](../../models/operations/postorgsorgidappsappidenvsenviddeploysrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploysResponse](../../models/operations/postorgsorgidappsappidenvsenviddeploysresponse.md), error**

