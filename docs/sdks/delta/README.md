# Delta

## Overview

A Deployment Delta (or just "Delta") describes the changes that must be applied to one Deployment Set to generate another Deployment Set. Deployment Deltas are the only way to create new Deployment Sets.

Deployment Deltas can be created fully formed or combined together via PATCHing. They can also be generated from the difference between two Deployment Sets.

**Basic Structure**

```
 {
   "id": <ID of the Deployment Delta.>,
   "metadata": {
     <Properties such as who created the Delta, which Environment it is associated to and a Human-friendly name>
   }
   "modules" : {
     "add" : {
       <ID of Module to add to the Deployment Set> : {
         <An entire Modules object>
       }
     },
     "remove": [
       <An array of Module IDs that should be removed from the Deployment Set>
     ],
    "update": {
       <ID of Module already in the Set to be updated> : [
         <An array of JSON Patch (Search Results (RFC 6902) objects scoped to the module>
       ]
     }
   }
 }
```
<SchemaDefinition schemaRef="#/components/schemas/DeltaRequest" />


### Available Operations

* [GetDelta](#getdelta) - Fetch an existing Delta
* [GetOrgsOrgIDAppsAppIDDeltas](#getorgsorgidappsappiddeltas) - List Deltas in an Application
* [PatchOrgsOrgIDAppsAppIDDeltasDeltaID](#patchorgsorgidappsappiddeltasdeltaid) - Update an existing Delta
* [PostOrgsOrgIDAppsAppIDDeltas](#postorgsorgidappsappiddeltas) - Create a new Delta
* [PutDelta](#putdelta) - Update an existing Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived](#putorgsorgidappsappiddeltasdeltaidmetadataarchived) - Mark a Delta as "archived"
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID](#putorgsorgidappsappiddeltasdeltaidmetadataenvid) - Change the Environment of a Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName](#putorgsorgidappsappiddeltasdeltaidmetadataname) - Change the name of a Delta

## GetDelta

Fetch an existing Delta

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
    res, err := s.Delta.GetDelta(ctx, operations.GetDeltaRequest{
        AppID: "praesentium",
        DeltaID: "rem",
        OrgID: "voluptates",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeltaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetDeltaRequest](../../models/operations/getdeltarequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetDeltaResponse](../../models/operations/getdeltaresponse.md), error**


## GetOrgsOrgIDAppsAppIDDeltas

List Deltas in an Application

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
    res, err := s.Delta.GetOrgsOrgIDAppsAppIDDeltas(ctx, operations.GetOrgsOrgIDAppsAppIDDeltasRequest{
        AppID: "quasi",
        Archived: test1.Bool(false),
        Env: test1.String("repudiandae"),
        OrgID: "sint",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeltaResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetOrgsOrgIDAppsAppIDDeltasRequest](../../models/operations/getorgsorgidappsappiddeltasrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDDeltasResponse](../../models/operations/getorgsorgidappsappiddeltasresponse.md), error**


## PatchOrgsOrgIDAppsAppIDDeltasDeltaID

Update an existing Delta

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
	"test-1/pkg/types"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.Delta.PatchOrgsOrgIDAppsAppIDDeltasDeltaID(ctx, operations.PatchOrgsOrgIDAppsAppIDDeltasDeltaIDRequest{
        RequestBody: []shared.DeltaRequest{
            shared.DeltaRequest{
                ID: test1.String("e450ad2a-bd44-4269-802d-502a94bb4f63"),
                Metadata: &shared.DeltaMetadataRequest{
                    Archived: test1.Bool(false),
                    Contributers: []string{
                        "sint",
                        "aliquid",
                        "provident",
                        "necessitatibus",
                    },
                    CreatedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                    CreatedBy: test1.String("sint"),
                    EnvID: test1.String("officia"),
                    LastModifiedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                    Name: test1.String("Raquel Wilderman"),
                    Shared: test1.Bool(false),
                },
                Modules: &shared.ModuleDeltasRequest{
                    Add: map[string]map[string]shared.ControllerRequest{
                        "illum": map[string]shared.ControllerRequest{
                            "rerum": shared.ControllerRequest{
                                Kind: test1.String("dicta"),
                                Message: test1.String("magnam"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "aliquid": "laborum",
                                                "accusamus": "non",
                                            },
                                            map[string]interface{}{
                                                "enim": "accusamus",
                                                "delectus": "quidem",
                                                "provident": "nam",
                                            },
                                            map[string]interface{}{
                                                "blanditiis": "deleniti",
                                                "sapiente": "amet",
                                                "deserunt": "nisi",
                                            },
                                            map[string]interface{}{
                                                "natus": "omnis",
                                                "molestiae": "perferendis",
                                            },
                                        },
                                        Phase: test1.String("nihil"),
                                        PodName: test1.String("magnam"),
                                        Revision: test1.Int64(716075),
                                        Status: test1.String("id"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "suscipit": "natus",
                                                "nobis": "eum",
                                            },
                                            map[string]interface{}{
                                                "aspernatur": "architecto",
                                                "magnam": "et",
                                                "excepturi": "ullam",
                                                "provident": "quos",
                                            },
                                        },
                                        Phase: test1.String("sint"),
                                        PodName: test1.String("accusantium"),
                                        Revision: test1.Int64(653201),
                                        Status: test1.String("reiciendis"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "eum": "dolor",
                                                "necessitatibus": "odit",
                                            },
                                            map[string]interface{}{
                                                "quasi": "iure",
                                                "doloribus": "debitis",
                                            },
                                            map[string]interface{}{
                                                "maxime": "deleniti",
                                                "facilis": "in",
                                            },
                                        },
                                        Phase: test1.String("architecto"),
                                        PodName: test1.String("architecto"),
                                        Revision: test1.Int64(919483),
                                        Status: test1.String("ullam"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "repellat": "quibusdam",
                                                "sed": "saepe",
                                            },
                                            map[string]interface{}{
                                                "accusantium": "consequuntur",
                                                "praesentium": "natus",
                                                "magni": "sunt",
                                                "quo": "illum",
                                            },
                                            map[string]interface{}{
                                                "maxime": "ea",
                                                "excepturi": "odit",
                                                "ea": "accusantium",
                                                "ab": "maiores",
                                            },
                                        },
                                        Phase: test1.String("quidem"),
                                        PodName: test1.String("ipsam"),
                                        Revision: test1.Int64(453543),
                                        Status: test1.String("autem"),
                                    },
                                },
                                Replicas: test1.Int64(722056),
                                Revision: test1.Int64(50588),
                                Status: test1.String("pariatur"),
                            },
                            "nemo": shared.ControllerRequest{
                                Kind: test1.String("voluptatibus"),
                                Message: test1.String("perferendis"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "cumque": "corporis",
                                            },
                                        },
                                        Phase: test1.String("hic"),
                                        PodName: test1.String("libero"),
                                        Revision: test1.Int64(749999),
                                        Status: test1.String("dolores"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "dignissimos": "eaque",
                                                "quis": "nesciunt",
                                                "eos": "perferendis",
                                            },
                                            map[string]interface{}{
                                                "minus": "quam",
                                            },
                                        },
                                        Phase: test1.String("dolor"),
                                        PodName: test1.String("vero"),
                                        Revision: test1.Int64(345352),
                                        Status: test1.String("hic"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "facilis": "perspiciatis",
                                                "voluptatem": "porro",
                                                "consequuntur": "blanditiis",
                                            },
                                            map[string]interface{}{
                                                "eaque": "occaecati",
                                                "rerum": "adipisci",
                                                "asperiores": "earum",
                                            },
                                            map[string]interface{}{
                                                "iste": "dolorum",
                                                "deleniti": "pariatur",
                                            },
                                            map[string]interface{}{
                                                "nobis": "libero",
                                                "delectus": "quaerat",
                                                "quos": "aliquid",
                                            },
                                        },
                                        Phase: test1.String("dolorem"),
                                        PodName: test1.String("dolorem"),
                                        Revision: test1.Int64(222443),
                                        Status: test1.String("qui"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "excepturi": "cum",
                                                "voluptate": "dignissimos",
                                                "reiciendis": "amet",
                                                "dolorum": "numquam",
                                            },
                                        },
                                        Phase: test1.String("veritatis"),
                                        PodName: test1.String("ipsa"),
                                        Revision: test1.Int64(56418),
                                        Status: test1.String("iure"),
                                    },
                                },
                                Replicas: test1.Int64(487838),
                                Revision: test1.Int64(311796),
                                Status: test1.String("accusamus"),
                            },
                            "quidem": shared.ControllerRequest{
                                Kind: test1.String("voluptatibus"),
                                Message: test1.String("voluptas"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "sit": "fugiat",
                                                "ab": "soluta",
                                                "dolorum": "iusto",
                                            },
                                        },
                                        Phase: test1.String("voluptate"),
                                        PodName: test1.String("dolorum"),
                                        Revision: test1.Int64(536579),
                                        Status: test1.String("omnis"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "asperiores": "nihil",
                                                "ipsum": "voluptate",
                                                "id": "saepe",
                                            },
                                            map[string]interface{}{
                                                "aspernatur": "perferendis",
                                                "amet": "optio",
                                            },
                                            map[string]interface{}{
                                                "ad": "saepe",
                                                "suscipit": "deserunt",
                                                "provident": "minima",
                                                "repellendus": "totam",
                                            },
                                            map[string]interface{}{
                                                "alias": "at",
                                                "quaerat": "tempora",
                                                "vel": "quod",
                                            },
                                        },
                                        Phase: test1.String("officiis"),
                                        PodName: test1.String("qui"),
                                        Revision: test1.Int64(679880),
                                        Status: test1.String("a"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "iusto": "ipsum",
                                                "quisquam": "tenetur",
                                                "amet": "tempore",
                                            },
                                            map[string]interface{}{
                                                "numquam": "enim",
                                                "dolorem": "sapiente",
                                                "totam": "nihil",
                                                "sit": "expedita",
                                            },
                                        },
                                        Phase: test1.String("neque"),
                                        PodName: test1.String("sed"),
                                        Revision: test1.Int64(424685),
                                        Status: test1.String("libero"),
                                    },
                                },
                                Replicas: test1.Int64(374170),
                                Revision: test1.Int64(646265),
                                Status: test1.String("quam"),
                            },
                            "ipsum": shared.ControllerRequest{
                                Kind: test1.String("incidunt"),
                                Message: test1.String("qui"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "soluta": "dicta",
                                                "laborum": "totam",
                                                "incidunt": "aspernatur",
                                                "dolores": "distinctio",
                                            },
                                            map[string]interface{}{
                                                "aliquid": "quam",
                                                "molestias": "temporibus",
                                                "qui": "neque",
                                            },
                                            map[string]interface{}{
                                                "magni": "odio",
                                            },
                                            map[string]interface{}{
                                                "ullam": "nam",
                                            },
                                        },
                                        Phase: test1.String("hic"),
                                        PodName: test1.String("voluptatem"),
                                        Revision: test1.Int64(765326),
                                        Status: test1.String("soluta"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "saepe": "ipsum",
                                            },
                                            map[string]interface{}{
                                                "nobis": "quos",
                                            },
                                            map[string]interface{}{
                                                "cupiditate": "aperiam",
                                                "delectus": "dolorem",
                                                "dolore": "labore",
                                            },
                                        },
                                        Phase: test1.String("adipisci"),
                                        PodName: test1.String("dolorum"),
                                        Revision: test1.Int64(100294),
                                        Status: test1.String("quae"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "itaque": "consequatur",
                                                "est": "repellendus",
                                                "porro": "doloribus",
                                            },
                                        },
                                        Phase: test1.String("ut"),
                                        PodName: test1.String("facilis"),
                                        Revision: test1.Int64(586410),
                                        Status: test1.String("qui"),
                                    },
                                },
                                Replicas: test1.Int64(63955),
                                Revision: test1.Int64(512393),
                                Status: test1.String("odio"),
                            },
                        },
                        "occaecati": map[string]shared.ControllerRequest{
                            "quisquam": shared.ControllerRequest{
                                Kind: test1.String("vero"),
                                Message: test1.String("omnis"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "voluptate": "consectetur",
                                                "vero": "tenetur",
                                                "dignissimos": "hic",
                                                "distinctio": "quod",
                                            },
                                        },
                                        Phase: test1.String("odio"),
                                        PodName: test1.String("similique"),
                                        Revision: test1.Int64(708548),
                                        Status: test1.String("vero"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "quibusdam": "illum",
                                                "sequi": "natus",
                                            },
                                            map[string]interface{}{
                                                "aut": "voluptatibus",
                                                "exercitationem": "nulla",
                                                "fugit": "porro",
                                                "maiores": "doloribus",
                                            },
                                        },
                                        Phase: test1.String("iusto"),
                                        PodName: test1.String("eligendi"),
                                        Revision: test1.Int64(497391),
                                        Status: test1.String("alias"),
                                    },
                                },
                                Replicas: test1.Int64(639473),
                                Revision: test1.Int64(269479),
                                Status: test1.String("ipsam"),
                            },
                            "ea": shared.ControllerRequest{
                                Kind: test1.String("aspernatur"),
                                Message: test1.String("vel"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "ex": "laudantium",
                                            },
                                            map[string]interface{}{
                                                "dolor": "maiores",
                                            },
                                        },
                                        Phase: test1.String("quasi"),
                                        PodName: test1.String("ex"),
                                        Revision: test1.Int64(862192),
                                        Status: test1.String("excepturi"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "sapiente": "quisquam",
                                                "saepe": "ea",
                                            },
                                            map[string]interface{}{
                                                "corporis": "veniam",
                                                "aliquid": "inventore",
                                                "magnam": "ea",
                                                "quo": "consectetur",
                                            },
                                            map[string]interface{}{
                                                "aspernatur": "minima",
                                                "eaque": "a",
                                                "libero": "aut",
                                                "aut": "deleniti",
                                            },
                                            map[string]interface{}{
                                                "aliquam": "fugit",
                                                "accusamus": "inventore",
                                                "non": "et",
                                                "dolorum": "laborum",
                                            },
                                        },
                                        Phase: test1.String("placeat"),
                                        PodName: test1.String("velit"),
                                        Revision: test1.Int64(432148),
                                        Status: test1.String("autem"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "assumenda": "nulla",
                                                "voluptas": "libero",
                                                "quasi": "tempora",
                                            },
                                            map[string]interface{}{
                                                "explicabo": "provident",
                                                "ipsa": "molestiae",
                                            },
                                            map[string]interface{}{
                                                "odio": "eius",
                                                "esse": "esse",
                                            },
                                            map[string]interface{}{
                                                "fuga": "reprehenderit",
                                                "quidem": "fugiat",
                                                "ut": "eum",
                                            },
                                        },
                                        Phase: test1.String("suscipit"),
                                        PodName: test1.String("assumenda"),
                                        Revision: test1.Int64(181151),
                                        Status: test1.String("praesentium"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "ipsa": "id",
                                            },
                                            map[string]interface{}{
                                                "neque": "quo",
                                                "illum": "quo",
                                                "fuga": "eius",
                                            },
                                            map[string]interface{}{
                                                "voluptas": "ab",
                                            },
                                            map[string]interface{}{
                                                "consequatur": "tempora",
                                                "debitis": "ipsam",
                                                "aspernatur": "sequi",
                                            },
                                        },
                                        Phase: test1.String("quo"),
                                        PodName: test1.String("esse"),
                                        Revision: test1.Int64(925164),
                                        Status: test1.String("aperiam"),
                                    },
                                },
                                Replicas: test1.Int64(715179),
                                Revision: test1.Int64(799796),
                                Status: test1.String("dignissimos"),
                            },
                            "inventore": shared.ControllerRequest{
                                Kind: test1.String("nihil"),
                                Message: test1.String("totam"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "occaecati": "commodi",
                                                "sapiente": "dolores",
                                            },
                                            map[string]interface{}{
                                                "molestiae": "accusantium",
                                                "porro": "eum",
                                                "quas": "praesentium",
                                            },
                                        },
                                        Phase: test1.String("consequuntur"),
                                        PodName: test1.String("deleniti"),
                                        Revision: test1.Int64(143829),
                                        Status: test1.String("fuga"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "atque": "explicabo",
                                                "minima": "nisi",
                                            },
                                            map[string]interface{}{
                                                "sapiente": "consequuntur",
                                            },
                                            map[string]interface{}{
                                                "explicabo": "saepe",
                                            },
                                        },
                                        Phase: test1.String("occaecati"),
                                        PodName: test1.String("atque"),
                                        Revision: test1.Int64(92260),
                                        Status: test1.String("esse"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "veritatis": "esse",
                                                "quod": "nam",
                                                "vero": "aliquid",
                                                "quasi": "saepe",
                                            },
                                            map[string]interface{}{
                                                "harum": "molestiae",
                                                "rerum": "occaecati",
                                            },
                                            map[string]interface{}{
                                                "distinctio": "eligendi",
                                                "sit": "culpa",
                                            },
                                            map[string]interface{}{
                                                "adipisci": "cumque",
                                                "consequuntur": "consequatur",
                                                "minus": "quaerat",
                                            },
                                        },
                                        Phase: test1.String("sapiente"),
                                        PodName: test1.String("consectetur"),
                                        Revision: test1.Int64(458139),
                                        Status: test1.String("blanditiis"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "nulla": "quas",
                                                "esse": "quasi",
                                                "a": "error",
                                                "sint": "pariatur",
                                            },
                                            map[string]interface{}{
                                                "quia": "eveniet",
                                                "asperiores": "facere",
                                                "veritatis": "consequuntur",
                                                "quasi": "similique",
                                            },
                                            map[string]interface{}{
                                                "aliquid": "tenetur",
                                                "quae": "earum",
                                                "vel": "in",
                                            },
                                        },
                                        Phase: test1.String("eius"),
                                        PodName: test1.String("libero"),
                                        Revision: test1.Int64(849039),
                                        Status: test1.String("soluta"),
                                    },
                                },
                                Replicas: test1.Int64(33304),
                                Revision: test1.Int64(306986),
                                Status: test1.String("sapiente"),
                            },
                            "dicta": shared.ControllerRequest{
                                Kind: test1.String("ullam"),
                                Message: test1.String("reprehenderit"),
                                Pods: []shared.PodStateRequest{
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "voluptatum": "qui",
                                            },
                                            map[string]interface{}{
                                                "ex": "deleniti",
                                                "itaque": "dolorum",
                                                "architecto": "omnis",
                                                "tenetur": "quasi",
                                            },
                                        },
                                        Phase: test1.String("at"),
                                        PodName: test1.String("et"),
                                        Revision: test1.Int64(454162),
                                        Status: test1.String("ipsa"),
                                    },
                                    shared.PodStateRequest{
                                        ContainerStatuses: []map[string]interface{}{
                                            map[string]interface{}{
                                                "consectetur": "adipisci",
                                            },
                                            map[string]interface{}{
                                                "temporibus": "accusantium",
                                                "rem": "aut",
                                                "laudantium": "eum",
                                            },
                                        },
                                        Phase: test1.String("mollitia"),
                                        PodName: test1.String("ab"),
                                        Revision: test1.Int64(544591),
                                        Status: test1.String("non"),
                                    },
                                },
                                Replicas: test1.Int64(32465),
                                Revision: test1.Int64(221161),
                                Status: test1.String("occaecati"),
                            },
                        },
                    },
                    Remove: []string{
                        "impedit",
                        "explicabo",
                    },
                    Update: map[string][]shared.UpdateActionRequest{
                        "aut": []shared.UpdateActionRequest{
                            shared.UpdateActionRequest{
                                From: test1.String("dicta"),
                                Op: test1.String("maiores"),
                                Path: test1.String("natus"),
                                Value: test1.String("velit"),
                            },
                            shared.UpdateActionRequest{
                                From: test1.String("voluptatibus"),
                                Op: test1.String("voluptas"),
                                Path: test1.String("asperiores"),
                                Value: test1.String("aperiam"),
                            },
                        },
                        "ea": []shared.UpdateActionRequest{
                            shared.UpdateActionRequest{
                                From: test1.String("consequuntur"),
                                Op: test1.String("repellendus"),
                                Path: test1.String("officia"),
                                Value: test1.String("maxime"),
                            },
                            shared.UpdateActionRequest{
                                From: test1.String("dignissimos"),
                                Op: test1.String("officia"),
                                Path: test1.String("asperiores"),
                                Value: test1.String("nemo"),
                            },
                        },
                    },
                },
                Shared: []shared.UpdateActionRequest{
                    shared.UpdateActionRequest{
                        From: test1.String("quaerat"),
                        Op: test1.String("porro"),
                        Path: test1.String("quod"),
                        Value: test1.String("labore"),
                    },
                },
            },
        },
        AppID: "ab",
        DeltaID: "adipisci",
        OrgID: "fuga",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeltaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.PatchOrgsOrgIDAppsAppIDDeltasDeltaIDRequest](../../models/operations/patchorgsorgidappsappiddeltasdeltaidrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.PatchOrgsOrgIDAppsAppIDDeltasDeltaIDResponse](../../models/operations/patchorgsorgidappsappiddeltasdeltaidresponse.md), error**


## PostOrgsOrgIDAppsAppIDDeltas

Create a new Delta

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
	"test-1/pkg/types"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.Delta.PostOrgsOrgIDAppsAppIDDeltas(ctx, operations.PostOrgsOrgIDAppsAppIDDeltasRequest{
        DeltaRequest: shared.DeltaRequest{
            ID: test1.String("a63aae8d-6786-44db-b675-fd5e60b375ed"),
            Metadata: &shared.DeltaMetadataRequest{
                Archived: test1.Bool(false),
                Contributers: []string{
                    "doloribus",
                    "suscipit",
                },
                CreatedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                CreatedBy: test1.String("reiciendis"),
                EnvID: test1.String("quidem"),
                LastModifiedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                Name: test1.String("Dr. Terence Gulgowski"),
                Shared: test1.Bool(false),
            },
            Modules: &shared.ModuleDeltasRequest{
                Add: map[string]map[string]shared.ControllerRequest{
                    "amet": map[string]shared.ControllerRequest{
                        "dignissimos": shared.ControllerRequest{
                            Kind: test1.String("a"),
                            Message: test1.String("debitis"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "laboriosam": "ipsa",
                                            "voluptates": "libero",
                                            "vitae": "accusamus",
                                        },
                                        map[string]interface{}{
                                            "tempora": "aspernatur",
                                            "voluptas": "voluptas",
                                            "voluptas": "minima",
                                        },
                                    },
                                    Phase: test1.String("nobis"),
                                    PodName: test1.String("dolorum"),
                                    Revision: test1.Int64(237807),
                                    Status: test1.String("minus"),
                                },
                            },
                            Replicas: test1.Int64(171853),
                            Revision: test1.Int64(503934),
                            Status: test1.String("in"),
                        },
                    },
                },
                Remove: []string{
                    "aliquam",
                    "officiis",
                },
                Update: map[string][]shared.UpdateActionRequest{
                    "ullam": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("cum"),
                            Op: test1.String("blanditiis"),
                            Path: test1.String("quas"),
                            Value: test1.String("hic"),
                        },
                    },
                    "nesciunt": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("corrupti"),
                            Op: test1.String("pariatur"),
                            Path: test1.String("totam"),
                            Value: test1.String("hic"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("exercitationem"),
                            Op: test1.String("nobis"),
                            Path: test1.String("sit"),
                            Value: test1.String("rerum"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("sed"),
                            Op: test1.String("reiciendis"),
                            Path: test1.String("explicabo"),
                            Value: test1.String("asperiores"),
                        },
                    },
                    "facilis": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("expedita"),
                            Op: test1.String("ab"),
                            Path: test1.String("iste"),
                            Value: test1.String("dolore"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("laborum"),
                            Op: test1.String("sed"),
                            Path: test1.String("in"),
                            Value: test1.String("commodi"),
                        },
                    },
                    "quidem": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("voluptas"),
                            Op: test1.String("unde"),
                            Path: test1.String("architecto"),
                            Value: test1.String("suscipit"),
                        },
                    },
                },
            },
            Shared: []shared.UpdateActionRequest{
                shared.UpdateActionRequest{
                    From: test1.String("debitis"),
                    Op: test1.String("illo"),
                    Path: test1.String("reiciendis"),
                    Value: test1.String("perferendis"),
                },
                shared.UpdateActionRequest{
                    From: test1.String("corrupti"),
                    Op: test1.String("maiores"),
                    Path: test1.String("incidunt"),
                    Value: test1.String("sed"),
                },
                shared.UpdateActionRequest{
                    From: test1.String("provident"),
                    Op: test1.String("eius"),
                    Path: test1.String("necessitatibus"),
                    Value: test1.String("ipsum"),
                },
                shared.UpdateActionRequest{
                    From: test1.String("ea"),
                    Op: test1.String("occaecati"),
                    Path: test1.String("quos"),
                    Value: test1.String("voluptatibus"),
                },
            },
        },
        AppID: "tempora",
        OrgID: "tempora",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostOrgsOrgIDAppsAppIDDeltas200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PostOrgsOrgIDAppsAppIDDeltasRequest](../../models/operations/postorgsorgidappsappiddeltasrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PostOrgsOrgIDAppsAppIDDeltasResponse](../../models/operations/postorgsorgidappsappiddeltasresponse.md), error**


## PutDelta

Update an existing Delta

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
	"test-1/pkg/models/shared"
	"test-1/pkg/types"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.Delta.PutDelta(ctx, operations.PutDeltaRequest{
        DeltaRequest: shared.DeltaRequest{
            ID: test1.String("7f603e8b-445e-480c-a55e-fd20e457e185"),
            Metadata: &shared.DeltaMetadataRequest{
                Archived: test1.Bool(false),
                Contributers: []string{
                    "cum",
                    "laboriosam",
                    "dolorum",
                },
                CreatedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                CreatedBy: test1.String("voluptatum"),
                EnvID: test1.String("error"),
                LastModifiedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                Name: test1.String("Rudolph Trantow"),
                Shared: test1.Bool(false),
            },
            Modules: &shared.ModuleDeltasRequest{
                Add: map[string]map[string]shared.ControllerRequest{
                    "officia": map[string]shared.ControllerRequest{
                        "corrupti": shared.ControllerRequest{
                            Kind: test1.String("accusamus"),
                            Message: test1.String("tempora"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "fugiat": "voluptatem",
                                            "culpa": "expedita",
                                        },
                                    },
                                    Phase: test1.String("magnam"),
                                    PodName: test1.String("consequatur"),
                                    Revision: test1.Int64(460220),
                                    Status: test1.String("ipsam"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "quas": "repudiandae",
                                            "corporis": "et",
                                            "blanditiis": "ex",
                                        },
                                    },
                                    Phase: test1.String("sed"),
                                    PodName: test1.String("sit"),
                                    Revision: test1.Int64(425508),
                                    Status: test1.String("nostrum"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "consequatur": "incidunt",
                                            "reiciendis": "dolorem",
                                            "harum": "dicta",
                                        },
                                        map[string]interface{}{
                                            "occaecati": "labore",
                                        },
                                        map[string]interface{}{
                                            "atque": "laborum",
                                            "nam": "tenetur",
                                            "laboriosam": "alias",
                                        },
                                        map[string]interface{}{
                                            "deserunt": "voluptate",
                                        },
                                    },
                                    Phase: test1.String("unde"),
                                    PodName: test1.String("reiciendis"),
                                    Revision: test1.Int64(588740),
                                    Status: test1.String("repellendus"),
                                },
                            },
                            Replicas: test1.Int64(962771),
                            Revision: test1.Int64(914791),
                            Status: test1.String("perferendis"),
                        },
                        "est": shared.ControllerRequest{
                            Kind: test1.String("quidem"),
                            Message: test1.String("reprehenderit"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "mollitia": "veniam",
                                            "voluptatem": "quisquam",
                                            "repudiandae": "quasi",
                                        },
                                        map[string]interface{}{
                                            "reprehenderit": "asperiores",
                                            "totam": "suscipit",
                                            "quidem": "maxime",
                                        },
                                        map[string]interface{}{
                                            "esse": "amet",
                                        },
                                    },
                                    Phase: test1.String("assumenda"),
                                    PodName: test1.String("ea"),
                                    Revision: test1.Int64(539118),
                                    Status: test1.String("error"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "accusamus": "natus",
                                            "minima": "aspernatur",
                                            "ex": "maiores",
                                            "corrupti": "at",
                                        },
                                        map[string]interface{}{
                                            "blanditiis": "suscipit",
                                            "repudiandae": "atque",
                                            "atque": "sunt",
                                        },
                                        map[string]interface{}{
                                            "dolorum": "repellendus",
                                            "labore": "reiciendis",
                                            "doloremque": "repudiandae",
                                            "dicta": "accusantium",
                                        },
                                        map[string]interface{}{
                                            "dolores": "enim",
                                        },
                                    },
                                    Phase: test1.String("laboriosam"),
                                    PodName: test1.String("velit"),
                                    Revision: test1.Int64(952143),
                                    Status: test1.String("molestias"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "consequuntur": "occaecati",
                                            "officiis": "perspiciatis",
                                            "in": "adipisci",
                                            "eveniet": "occaecati",
                                        },
                                        map[string]interface{}{
                                            "fugit": "id",
                                        },
                                    },
                                    Phase: test1.String("quis"),
                                    PodName: test1.String("reprehenderit"),
                                    Revision: test1.Int64(625528),
                                    Status: test1.String("illo"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "eveniet": "non",
                                            "vero": "doloremque",
                                            "iure": "ipsa",
                                        },
                                        map[string]interface{}{
                                            "quae": "molestiae",
                                            "eveniet": "qui",
                                            "cum": "iure",
                                        },
                                    },
                                    Phase: test1.String("necessitatibus"),
                                    PodName: test1.String("ratione"),
                                    Revision: test1.Int64(672582),
                                    Status: test1.String("distinctio"),
                                },
                            },
                            Replicas: test1.Int64(528940),
                            Revision: test1.Int64(523006),
                            Status: test1.String("aliquam"),
                        },
                        "ad": shared.ControllerRequest{
                            Kind: test1.String("repellat"),
                            Message: test1.String("alias"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "mollitia": "voluptas",
                                            "alias": "maiores",
                                        },
                                        map[string]interface{}{
                                            "dolores": "id",
                                            "minima": "dolore",
                                            "dolorum": "nesciunt",
                                            "quae": "recusandae",
                                        },
                                        map[string]interface{}{
                                            "quaerat": "molestiae",
                                            "ex": "ut",
                                            "culpa": "adipisci",
                                        },
                                    },
                                    Phase: test1.String("debitis"),
                                    PodName: test1.String("laudantium"),
                                    Revision: test1.Int64(432606),
                                    Status: test1.String("nemo"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "provident": "quis",
                                            "eum": "reiciendis",
                                        },
                                        map[string]interface{}{
                                            "aspernatur": "ullam",
                                            "quasi": "animi",
                                            "nostrum": "mollitia",
                                        },
                                        map[string]interface{}{
                                            "possimus": "animi",
                                            "ex": "aliquid",
                                            "accusantium": "repellat",
                                        },
                                        map[string]interface{}{
                                            "ullam": "in",
                                            "nam": "earum",
                                            "officia": "laborum",
                                            "placeat": "modi",
                                        },
                                    },
                                    Phase: test1.String("voluptatibus"),
                                    PodName: test1.String("molestias"),
                                    Revision: test1.Int64(889794),
                                    Status: test1.String("sapiente"),
                                },
                            },
                            Replicas: test1.Int64(764562),
                            Revision: test1.Int64(113486),
                            Status: test1.String("rerum"),
                        },
                    },
                    "tempora": map[string]shared.ControllerRequest{
                        "inventore": shared.ControllerRequest{
                            Kind: test1.String("fugit"),
                            Message: test1.String("cumque"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "aspernatur": "eum",
                                        },
                                    },
                                    Phase: test1.String("eius"),
                                    PodName: test1.String("rem"),
                                    Revision: test1.Int64(871083),
                                    Status: test1.String("impedit"),
                                },
                            },
                            Replicas: test1.Int64(179410),
                            Revision: test1.Int64(958741),
                            Status: test1.String("eum"),
                        },
                        "dicta": shared.ControllerRequest{
                            Kind: test1.String("minima"),
                            Message: test1.String("beatae"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "soluta": "hic",
                                            "illum": "eaque",
                                            "earum": "perspiciatis",
                                            "maiores": "debitis",
                                        },
                                        map[string]interface{}{
                                            "porro": "suscipit",
                                            "dolorem": "fugit",
                                        },
                                        map[string]interface{}{
                                            "fuga": "ratione",
                                            "animi": "necessitatibus",
                                            "nulla": "consequatur",
                                            "quasi": "et",
                                        },
                                    },
                                    Phase: test1.String("ducimus"),
                                    PodName: test1.String("natus"),
                                    Revision: test1.Int64(581082),
                                    Status: test1.String("suscipit"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "magni": "doloribus",
                                        },
                                    },
                                    Phase: test1.String("nulla"),
                                    PodName: test1.String("necessitatibus"),
                                    Revision: test1.Int64(58534),
                                    Status: test1.String("tempora"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "dicta": "iusto",
                                            "esse": "praesentium",
                                        },
                                        map[string]interface{}{
                                            "reiciendis": "vel",
                                            "architecto": "fugiat",
                                            "doloremque": "dicta",
                                            "odio": "tempora",
                                        },
                                    },
                                    Phase: test1.String("esse"),
                                    PodName: test1.String("ex"),
                                    Revision: test1.Int64(235263),
                                    Status: test1.String("aliquid"),
                                },
                            },
                            Replicas: test1.Int64(58870),
                            Revision: test1.Int64(671384),
                            Status: test1.String("sunt"),
                        },
                    },
                },
                Remove: []string{
                    "fugiat",
                    "expedita",
                },
                Update: map[string][]shared.UpdateActionRequest{
                    "officia": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("aliquid"),
                            Op: test1.String("perferendis"),
                            Path: test1.String("eum"),
                            Value: test1.String("voluptas"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("iste"),
                            Op: test1.String("id"),
                            Path: test1.String("ab"),
                            Value: test1.String("error"),
                        },
                    },
                    "possimus": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("mollitia"),
                            Op: test1.String("laborum"),
                            Path: test1.String("libero"),
                            Value: test1.String("ad"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("deleniti"),
                            Op: test1.String("enim"),
                            Path: test1.String("vitae"),
                            Value: test1.String("repellendus"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("ex"),
                            Op: test1.String("quo"),
                            Path: test1.String("ex"),
                            Value: test1.String("ut"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("ad"),
                            Op: test1.String("expedita"),
                            Path: test1.String("voluptatem"),
                            Value: test1.String("molestias"),
                        },
                    },
                },
            },
            Shared: []shared.UpdateActionRequest{
                shared.UpdateActionRequest{
                    From: test1.String("aliquid"),
                    Op: test1.String("beatae"),
                    Path: test1.String("voluptatum"),
                    Value: test1.String("omnis"),
                },
                shared.UpdateActionRequest{
                    From: test1.String("veritatis"),
                    Op: test1.String("rerum"),
                    Path: test1.String("est"),
                    Value: test1.String("culpa"),
                },
                shared.UpdateActionRequest{
                    From: test1.String("voluptatem"),
                    Op: test1.String("sapiente"),
                    Path: test1.String("officiis"),
                    Value: test1.String("architecto"),
                },
            },
        },
        AppID: "fuga",
        DeltaID: "pariatur",
        OrgID: "debitis",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.PutDeltaRequest](../../models/operations/putdeltarequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PutDeltaResponse](../../models/operations/putdeltaresponse.md), error**


## PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived

Archived Deltas are still accessible but can no longer be updated.

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
    res, err := s.Delta.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived(ctx, operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchivedRequest{
        RequestBody: false,
        AppID: "voluptatem",
        DeltaID: "alias",
        OrgID: "deleniti",
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

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchivedRequest](../../models/operations/putorgsorgidappsappiddeltasdeltaidmetadataarchivedrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchivedResponse](../../models/operations/putorgsorgidappsappiddeltasdeltaidmetadataarchivedresponse.md), error**


## PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID

Change the Environment of a Delta

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
    res, err := s.Delta.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID(ctx, operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvIDRequest{
        RequestBody: "earum",
        AppID: "ex",
        DeltaID: "sapiente",
        OrgID: "rem",
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
| `request`                                                                                                                                              | [operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvIDRequest](../../models/operations/putorgsorgidappsappiddeltasdeltaidmetadataenvidrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvIDResponse](../../models/operations/putorgsorgidappsappiddeltasdeltaidmetadataenvidresponse.md), error**


## PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName

Change the name of a Delta

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
    res, err := s.Delta.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName(ctx, operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameRequest{
        RequestBody: "minus",
        AppID: "nemo",
        DeltaID: "asperiores",
        OrgID: "ratione",
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

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameRequest](../../models/operations/putorgsorgidappsappiddeltasdeltaidmetadatanamerequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataNameResponse](../../models/operations/putorgsorgidappsappiddeltasdeltaidmetadatanameresponse.md), error**

