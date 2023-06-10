# RuntimeInfo

## Overview

RuntimeInfo object returned by the runtime endpoint. Represents a list post statuses grouped by modules and controllers (deployments and stateful sets).
<SchemaDefinition schemaRef="#/components/schemas/RuntimeInfoRequest" />


### Available Operations

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntime](#getorgsorgidappsappidenvsenvidruntime) - Get Runtime information about the environment.
* [GetOrgsOrgIDAppsAppIDRuntime](#getorgsorgidappsappidruntime) - Get Runtime information about specific environments.
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicas](#patchorgsorgidappsappidenvsenvidruntimereplicas) - Set number of replicas for an environment's modules.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePaused](#putorgsorgidappsappidenvsenvidruntimepaused) - Pause / Resume an environment.

## GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntime

Get Runtime information about the environment.

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
    res, err := s.RuntimeInfo.GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntime(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeRequest{
        AppID: "aut",
        EnvID: "aut",
        OrgID: "eveniet",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RuntimeInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeRequest](../../models/operations/getorgsorgidappsappidenvsenvidruntimerequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeResponse](../../models/operations/getorgsorgidappsappidenvsenvidruntimeresponse.md), error**


## GetOrgsOrgIDAppsAppIDRuntime

Get Runtime information about specific environments.

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
    res, err := s.RuntimeInfo.GetOrgsOrgIDAppsAppIDRuntime(ctx, operations.GetOrgsOrgIDAppsAppIDRuntimeRequest{
        AppID: "odio",
        ID: test1.String("64ad7334-ec1b-4781-b36a-08088d100efa"),
        OrgID: "nulla",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentRuntimeInfoResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetOrgsOrgIDAppsAppIDRuntimeRequest](../../models/operations/getorgsorgidappsappidruntimerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDRuntimeResponse](../../models/operations/getorgsorgidappsappidruntimeresponse.md), error**


## PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicas

Set number of replicas for an environment's modules.

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
    res, err := s.RuntimeInfo.PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicas(ctx, operations.PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicasRequest{
        RequestBody: map[string]int64{
            "sed": 30208,
            "alias": 910073,
            "hic": 27982,
        },
        AppID: "incidunt",
        EnvID: "qui",
        OrgID: "qui",
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

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicasRequest](../../models/operations/patchorgsorgidappsappidenvsenvidruntimereplicasrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicasResponse](../../models/operations/patchorgsorgidappsappidenvsenvidruntimereplicasresponse.md), error**


## PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePaused

On pause requests, all the Kubernetes Deployment resources are scaled down to 0 replicas.

On resume requests, all the Kubernetes Deployment resources are scaled up to the number of replicas running before the environment was paused.

When an environment is paused, it is not possible to:

```
  - Deploy the environment within Humanitec.
  - Scale the number of replicas running of any workload.
```

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
    res, err := s.RuntimeInfo.PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePaused(ctx, operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePausedRequest{
        RequestBody: false,
        AppID: "necessitatibus",
        EnvID: "harum",
        OrgID: "explicabo",
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePausedRequest](../../models/operations/putorgsorgidappsappidenvsenvidruntimepausedrequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |


### Response

**[*operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePausedResponse](../../models/operations/putorgsorgidappsappidenvsenvidruntimepausedresponse.md), error**

