# Set

## Overview

A Deployment Set (or just "Set") defines all of the non-Environment specific configuration for Modules and External Resources. Each of these Modules or External Resources has a unique name.

Deployment Sets are immutable and their ID is a cryptographic hash of their content. This means that a Deployment Set can be globally identified based on its ID. It also means that referencing a Deployment Set by ID will always return the same Deployment Set.

Deployment Sets cannot be created directly, instead they are created by applying a Deployment Delta to an existing Deployment Set.

**Basic Structure**

```
 {
   "id": <ID of the Deployment Set>,
   "modules" : {
     <ID of Module> : {
       "profile": <Defines how the optional "spec" property is interpreted>
       "spec": {
         <Properties that depend on the "profile" property.>
       }
       "externals": {
         <External Resource Name> : {
           "type": <Resource Type>,
           "params": {
             <Properties which parametrize the resource depending on the Resource Type.>
           }
         }
       }
     }
   }
 }
```

For details about how the Humanitec provided profiles work, see (Deployment Set Profiles)[].
<SchemaDefinition schemaRef="#/components/schemas/SetRequest" />


### Available Operations

* [GetSets](#getsets) - Get all Deployment Sets
* [GetOrgsOrgIDAppsAppIDSetsSetID](#getorgsorgidappsappidsetssetid) - Get a Deployment Set
* [GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID](#getorgsorgidappsappidsetssetiddiffsourcesetid) - Get the difference between 2 Deployment Sets
* [PostOrgsOrgIDAppsAppIDSetsSetID](#postorgsorgidappsappidsetssetid) - Apply a Deployment Delta to a Deployment Set

## GetSets

Get all Deployment Sets

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
    res, err := s.Set.GetSets(ctx, operations.GetSetsRequest{
        AppID: "beatae",
        OrgID: "aliquid",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SetResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDSetsSetID

Get a Deployment Set

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
    res, err := s.Set.GetOrgsOrgIDAppsAppIDSetsSetID(ctx, operations.GetOrgsOrgIDAppsAppIDSetsSetIDRequest{
        AppID: "modi",
        Diff: test1.String("optio"),
        OrgID: "voluptatibus",
        SetID: "molestias",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetOrgsOrgIDAppsAppIDSetsSetID200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID

Get the difference between 2 Deployment Sets

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
    res, err := s.Set.GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID(ctx, operations.GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetIDRequest{
        AppID: "officia",
        OrgID: "libero",
        SetID: "totam",
        SourceSetID: "sequi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PlainDeltaResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDSetsSetID

Apply a Deployment Delta to a Deployment Set

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
    res, err := s.Set.PostOrgsOrgIDAppsAppIDSetsSetID(ctx, operations.PostOrgsOrgIDAppsAppIDSetsSetIDRequest{
        DeltaRequest: shared.DeltaRequest{
            ID: test1.String("66c723ff-da9e-406b-ae48-25c1fc0e115c"),
            Metadata: &shared.DeltaMetadataRequest{
                Archived: test1.Bool(false),
                Contributers: []string{
                    "perferendis",
                    "facilis",
                    "reiciendis",
                },
                CreatedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                CreatedBy: test1.String("a"),
                EnvID: test1.String("iste"),
                LastModifiedAt: types.MustTimeFromString("2020-06-22T09:37:23.523Z"),
                Name: test1.String("Olga Hermiston"),
                Shared: test1.Bool(false),
            },
            Modules: &shared.ModuleDeltasRequest{
                Add: map[string]map[string]shared.ControllerRequest{
                    "maxime": map[string]shared.ControllerRequest{
                        "consequuntur": shared.ControllerRequest{
                            Kind: test1.String("assumenda"),
                            Message: test1.String("vero"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "accusamus": "totam",
                                            "reiciendis": "ab",
                                            "sint": "nihil",
                                            "esse": "iure",
                                        },
                                        map[string]interface{}{
                                            "nesciunt": "debitis",
                                            "vel": "neque",
                                        },
                                        map[string]interface{}{
                                            "voluptas": "consequuntur",
                                            "officia": "reprehenderit",
                                        },
                                        map[string]interface{}{
                                            "eius": "ipsa",
                                            "rem": "maiores",
                                            "accusantium": "veniam",
                                        },
                                    },
                                    Phase: test1.String("saepe"),
                                    PodName: test1.String("neque"),
                                    Revision: test1.Int64(816365),
                                    Status: test1.String("aliquam"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "fugiat": "est",
                                            "delectus": "velit",
                                            "vitae": "nesciunt",
                                            "similique": "illo",
                                        },
                                        map[string]interface{}{
                                            "nemo": "doloribus",
                                            "possimus": "unde",
                                            "incidunt": "explicabo",
                                            "ipsam": "cupiditate",
                                        },
                                        map[string]interface{}{
                                            "alias": "quidem",
                                            "nesciunt": "commodi",
                                            "sapiente": "consequuntur",
                                            "veniam": "debitis",
                                        },
                                    },
                                    Phase: test1.String("officia"),
                                    PodName: test1.String("sint"),
                                    Revision: test1.Int64(280859),
                                    Status: test1.String("numquam"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "libero": "in",
                                        },
                                        map[string]interface{}{
                                            "ex": "minus",
                                            "ab": "beatae",
                                        },
                                        map[string]interface{}{
                                            "nisi": "quisquam",
                                            "dolor": "ducimus",
                                            "fuga": "minima",
                                            "architecto": "qui",
                                        },
                                        map[string]interface{}{
                                            "magni": "incidunt",
                                            "adipisci": "praesentium",
                                        },
                                    },
                                    Phase: test1.String("dolor"),
                                    PodName: test1.String("exercitationem"),
                                    Revision: test1.Int64(709701),
                                    Status: test1.String("facilis"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "nemo": "culpa",
                                        },
                                        map[string]interface{}{
                                            "amet": "deserunt",
                                        },
                                        map[string]interface{}{
                                            "veniam": "quod",
                                            "itaque": "a",
                                        },
                                        map[string]interface{}{
                                            "enim": "doloribus",
                                            "assumenda": "officiis",
                                            "architecto": "alias",
                                            "culpa": "ipsa",
                                        },
                                    },
                                    Phase: test1.String("nobis"),
                                    PodName: test1.String("necessitatibus"),
                                    Revision: test1.Int64(155978),
                                    Status: test1.String("dicta"),
                                },
                            },
                            Replicas: test1.Int64(426002),
                            Revision: test1.Int64(595584),
                            Status: test1.String("debitis"),
                        },
                        "ullam": shared.ControllerRequest{
                            Kind: test1.String("architecto"),
                            Message: test1.String("accusantium"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "cumque": "iure",
                                            "quibusdam": "quod",
                                            "nemo": "recusandae",
                                        },
                                    },
                                    Phase: test1.String("velit"),
                                    PodName: test1.String("magnam"),
                                    Revision: test1.Int64(493591),
                                    Status: test1.String("laboriosam"),
                                },
                            },
                            Replicas: test1.Int64(152283),
                            Revision: test1.Int64(486272),
                            Status: test1.String("natus"),
                        },
                    },
                    "provident": map[string]shared.ControllerRequest{
                        "doloribus": shared.ControllerRequest{
                            Kind: test1.String("facilis"),
                            Message: test1.String("quidem"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "modi": "perspiciatis",
                                            "hic": "cum",
                                            "aspernatur": "libero",
                                        },
                                        map[string]interface{}{
                                            "incidunt": "recusandae",
                                            "quod": "id",
                                            "saepe": "autem",
                                        },
                                    },
                                    Phase: test1.String("quo"),
                                    PodName: test1.String("nesciunt"),
                                    Revision: test1.Int64(849383),
                                    Status: test1.String("nemo"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "non": "mollitia",
                                            "assumenda": "recusandae",
                                            "distinctio": "pariatur",
                                        },
                                        map[string]interface{}{
                                            "facere": "laborum",
                                            "eveniet": "laborum",
                                        },
                                        map[string]interface{}{
                                            "maxime": "ipsam",
                                            "alias": "suscipit",
                                        },
                                        map[string]interface{}{
                                            "molestias": "laborum",
                                            "est": "occaecati",
                                            "labore": "quo",
                                        },
                                    },
                                    Phase: test1.String("perferendis"),
                                    PodName: test1.String("fugit"),
                                    Revision: test1.Int64(399222),
                                    Status: test1.String("magnam"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "hic": "nostrum",
                                            "officiis": "unde",
                                            "nulla": "error",
                                            "mollitia": "magnam",
                                        },
                                        map[string]interface{}{
                                            "esse": "corrupti",
                                            "fuga": "facere",
                                        },
                                    },
                                    Phase: test1.String("impedit"),
                                    PodName: test1.String("quasi"),
                                    Revision: test1.Int64(647218),
                                    Status: test1.String("quod"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "voluptatem": "facere",
                                        },
                                        map[string]interface{}{
                                            "maxime": "consequatur",
                                            "eaque": "architecto",
                                            "similique": "porro",
                                            "blanditiis": "quae",
                                        },
                                    },
                                    Phase: test1.String("magni"),
                                    PodName: test1.String("officiis"),
                                    Revision: test1.Int64(148379),
                                    Status: test1.String("necessitatibus"),
                                },
                            },
                            Replicas: test1.Int64(773259),
                            Revision: test1.Int64(55981),
                            Status: test1.String("excepturi"),
                        },
                        "a": shared.ControllerRequest{
                            Kind: test1.String("maiores"),
                            Message: test1.String("laudantium"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "rem": "dicta",
                                            "suscipit": "earum",
                                            "doloribus": "velit",
                                            "eius": "esse",
                                        },
                                    },
                                    Phase: test1.String("in"),
                                    PodName: test1.String("eligendi"),
                                    Revision: test1.Int64(94697),
                                    Status: test1.String("neque"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "accusantium": "qui",
                                            "impedit": "beatae",
                                            "incidunt": "dicta",
                                        },
                                        map[string]interface{}{
                                            "corporis": "rerum",
                                        },
                                        map[string]interface{}{
                                            "error": "vel",
                                        },
                                        map[string]interface{}{
                                            "id": "laboriosam",
                                        },
                                    },
                                    Phase: test1.String("ex"),
                                    PodName: test1.String("quas"),
                                    Revision: test1.Int64(85794),
                                    Status: test1.String("ullam"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "incidunt": "quam",
                                            "magni": "deserunt",
                                            "delectus": "omnis",
                                        },
                                    },
                                    Phase: test1.String("sed"),
                                    PodName: test1.String("nesciunt"),
                                    Revision: test1.Int64(805463),
                                    Status: test1.String("quis"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "excepturi": "maiores",
                                            "laudantium": "velit",
                                        },
                                        map[string]interface{}{
                                            "amet": "nemo",
                                            "ipsa": "quisquam",
                                            "tenetur": "quas",
                                            "molestiae": "aliquid",
                                        },
                                        map[string]interface{}{
                                            "a": "nobis",
                                            "perspiciatis": "accusantium",
                                            "dicta": "minus",
                                            "commodi": "eveniet",
                                        },
                                    },
                                    Phase: test1.String("porro"),
                                    PodName: test1.String("tempore"),
                                    Revision: test1.Int64(693747),
                                    Status: test1.String("modi"),
                                },
                            },
                            Replicas: test1.Int64(916341),
                            Revision: test1.Int64(145435),
                            Status: test1.String("eius"),
                        },
                        "sequi": shared.ControllerRequest{
                            Kind: test1.String("eligendi"),
                            Message: test1.String("asperiores"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "repellat": "a",
                                            "animi": "maiores",
                                            "itaque": "nulla",
                                        },
                                        map[string]interface{}{
                                            "corporis": "velit",
                                            "officiis": "enim",
                                            "officia": "saepe",
                                        },
                                        map[string]interface{}{
                                            "repudiandae": "accusantium",
                                            "officia": "impedit",
                                        },
                                    },
                                    Phase: test1.String("quasi"),
                                    PodName: test1.String("blanditiis"),
                                    Revision: test1.Int64(260911),
                                    Status: test1.String("quisquam"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "natus": "minus",
                                            "quia": "magnam",
                                            "reprehenderit": "quod",
                                        },
                                    },
                                    Phase: test1.String("quos"),
                                    PodName: test1.String("corrupti"),
                                    Revision: test1.Int64(227870),
                                    Status: test1.String("molestiae"),
                                },
                            },
                            Replicas: test1.Int64(227156),
                            Revision: test1.Int64(675126),
                            Status: test1.String("modi"),
                        },
                    },
                    "perferendis": map[string]shared.ControllerRequest{
                        "architecto": shared.ControllerRequest{
                            Kind: test1.String("molestias"),
                            Message: test1.String("dolore"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "odit": "earum",
                                        },
                                        map[string]interface{}{
                                            "ipsam": "eaque",
                                            "exercitationem": "veniam",
                                        },
                                        map[string]interface{}{
                                            "ad": "nisi",
                                            "tenetur": "quis",
                                        },
                                        map[string]interface{}{
                                            "nemo": "suscipit",
                                            "pariatur": "sit",
                                            "quidem": "repellendus",
                                            "perferendis": "id",
                                        },
                                    },
                                    Phase: test1.String("sapiente"),
                                    PodName: test1.String("sed"),
                                    Revision: test1.Int64(823572),
                                    Status: test1.String("repellat"),
                                },
                            },
                            Replicas: test1.Int64(921060),
                            Revision: test1.Int64(100233),
                            Status: test1.String("adipisci"),
                        },
                        "pariatur": shared.ControllerRequest{
                            Kind: test1.String("harum"),
                            Message: test1.String("dolore"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "minus": "soluta",
                                        },
                                        map[string]interface{}{
                                            "velit": "earum",
                                            "praesentium": "error",
                                            "non": "quasi",
                                        },
                                    },
                                    Phase: test1.String("mollitia"),
                                    PodName: test1.String("accusamus"),
                                    Revision: test1.Int64(687589),
                                    Status: test1.String("cumque"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "corrupti": "eaque",
                                            "deserunt": "aliquid",
                                            "excepturi": "magni",
                                        },
                                    },
                                    Phase: test1.String("tempora"),
                                    PodName: test1.String("possimus"),
                                    Revision: test1.Int64(220824),
                                    Status: test1.String("rerum"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "optio": "delectus",
                                            "minus": "quo",
                                            "quos": "asperiores",
                                            "voluptatum": "iste",
                                        },
                                    },
                                    Phase: test1.String("corporis"),
                                    PodName: test1.String("accusantium"),
                                    Revision: test1.Int64(75850),
                                    Status: test1.String("aut"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "at": "possimus",
                                            "neque": "pariatur",
                                        },
                                        map[string]interface{}{
                                            "sapiente": "mollitia",
                                            "quae": "quos",
                                        },
                                        map[string]interface{}{
                                            "non": "voluptates",
                                        },
                                        map[string]interface{}{
                                            "aliquam": "quisquam",
                                            "quas": "consequuntur",
                                        },
                                    },
                                    Phase: test1.String("maiores"),
                                    PodName: test1.String("inventore"),
                                    Revision: test1.Int64(400448),
                                    Status: test1.String("laudantium"),
                                },
                            },
                            Replicas: test1.Int64(665872),
                            Revision: test1.Int64(221329),
                            Status: test1.String("aliquid"),
                        },
                        "consectetur": shared.ControllerRequest{
                            Kind: test1.String("cumque"),
                            Message: test1.String("rem"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "recusandae": "tempora",
                                        },
                                        map[string]interface{}{
                                            "numquam": "sequi",
                                            "voluptatum": "sit",
                                            "rerum": "veritatis",
                                        },
                                    },
                                    Phase: test1.String("tenetur"),
                                    PodName: test1.String("autem"),
                                    Revision: test1.Int64(694088),
                                    Status: test1.String("totam"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "magni": "nihil",
                                            "voluptas": "animi",
                                            "commodi": "alias",
                                        },
                                        map[string]interface{}{
                                            "aut": "dolore",
                                            "maxime": "aliquam",
                                            "iste": "ullam",
                                        },
                                        map[string]interface{}{
                                            "placeat": "voluptas",
                                            "occaecati": "unde",
                                            "illo": "nihil",
                                            "inventore": "libero",
                                        },
                                        map[string]interface{}{
                                            "quasi": "cumque",
                                            "dicta": "harum",
                                        },
                                    },
                                    Phase: test1.String("facere"),
                                    PodName: test1.String("facilis"),
                                    Revision: test1.Int64(105372),
                                    Status: test1.String("cumque"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "expedita": "corrupti",
                                            "rem": "atque",
                                        },
                                        map[string]interface{}{
                                            "cum": "pariatur",
                                            "sapiente": "quo",
                                            "incidunt": "quod",
                                            "minus": "porro",
                                        },
                                        map[string]interface{}{
                                            "excepturi": "occaecati",
                                            "libero": "quo",
                                            "esse": "hic",
                                        },
                                        map[string]interface{}{
                                            "accusantium": "soluta",
                                            "fugit": "pariatur",
                                            "eligendi": "recusandae",
                                            "veritatis": "aut",
                                        },
                                    },
                                    Phase: test1.String("laudantium"),
                                    PodName: test1.String("iusto"),
                                    Revision: test1.Int64(219860),
                                    Status: test1.String("voluptates"),
                                },
                            },
                            Replicas: test1.Int64(274295),
                            Revision: test1.Int64(169935),
                            Status: test1.String("rerum"),
                        },
                        "doloremque": shared.ControllerRequest{
                            Kind: test1.String("voluptatem"),
                            Message: test1.String("eum"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "voluptatum": "blanditiis",
                                            "nihil": "atque",
                                        },
                                        map[string]interface{}{
                                            "deserunt": "atque",
                                            "nostrum": "atque",
                                            "architecto": "est",
                                        },
                                    },
                                    Phase: test1.String("enim"),
                                    PodName: test1.String("rem"),
                                    Revision: test1.Int64(168142),
                                    Status: test1.String("quae"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "enim": "labore",
                                            "sapiente": "saepe",
                                            "delectus": "officia",
                                            "natus": "cumque",
                                        },
                                        map[string]interface{}{
                                            "quaerat": "doloribus",
                                            "quia": "officiis",
                                            "mollitia": "cumque",
                                        },
                                        map[string]interface{}{
                                            "enim": "eum",
                                            "nemo": "illum",
                                        },
                                    },
                                    Phase: test1.String("nesciunt"),
                                    PodName: test1.String("sit"),
                                    Revision: test1.Int64(487148),
                                    Status: test1.String("minus"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "voluptates": "praesentium",
                                            "dicta": "fugit",
                                            "sit": "aliquid",
                                            "necessitatibus": "sed",
                                        },
                                        map[string]interface{}{
                                            "sunt": "nesciunt",
                                            "delectus": "laborum",
                                            "aliquam": "deserunt",
                                        },
                                        map[string]interface{}{
                                            "sunt": "impedit",
                                            "eius": "voluptatum",
                                        },
                                        map[string]interface{}{
                                            "at": "dolorem",
                                        },
                                    },
                                    Phase: test1.String("repellat"),
                                    PodName: test1.String("aspernatur"),
                                    Revision: test1.Int64(80284),
                                    Status: test1.String("sequi"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "hic": "eaque",
                                            "dolorem": "architecto",
                                            "aperiam": "aspernatur",
                                        },
                                    },
                                    Phase: test1.String("nulla"),
                                    PodName: test1.String("enim"),
                                    Revision: test1.Int64(73574),
                                    Status: test1.String("magnam"),
                                },
                            },
                            Replicas: test1.Int64(961842),
                            Revision: test1.Int64(255064),
                            Status: test1.String("optio"),
                        },
                    },
                    "nobis": map[string]shared.ControllerRequest{
                        "repellat": shared.ControllerRequest{
                            Kind: test1.String("quae"),
                            Message: test1.String("deleniti"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "aliquid": "sed",
                                            "beatae": "similique",
                                            "ea": "animi",
                                        },
                                        map[string]interface{}{
                                            "tenetur": "dignissimos",
                                            "esse": "animi",
                                        },
                                        map[string]interface{}{
                                            "esse": "eveniet",
                                            "earum": "velit",
                                            "officiis": "eius",
                                        },
                                        map[string]interface{}{
                                            "itaque": "dignissimos",
                                            "ipsam": "explicabo",
                                            "impedit": "aliquid",
                                        },
                                    },
                                    Phase: test1.String("quis"),
                                    PodName: test1.String("facilis"),
                                    Revision: test1.Int64(218128),
                                    Status: test1.String("ut"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "praesentium": "eveniet",
                                        },
                                        map[string]interface{}{
                                            "expedita": "libero",
                                        },
                                    },
                                    Phase: test1.String("iste"),
                                    PodName: test1.String("illo"),
                                    Revision: test1.Int64(792499),
                                    Status: test1.String("quos"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "iusto": "enim",
                                            "accusamus": "aperiam",
                                            "voluptates": "laudantium",
                                        },
                                        map[string]interface{}{
                                            "quae": "omnis",
                                            "illum": "rem",
                                        },
                                        map[string]interface{}{
                                            "deleniti": "modi",
                                            "earum": "architecto",
                                            "aliquam": "labore",
                                            "maiores": "sequi",
                                        },
                                        map[string]interface{}{
                                            "consequatur": "esse",
                                            "debitis": "facere",
                                            "quisquam": "cumque",
                                            "aliquam": "dolorum",
                                        },
                                    },
                                    Phase: test1.String("deserunt"),
                                    PodName: test1.String("ad"),
                                    Revision: test1.Int64(970411),
                                    Status: test1.String("sequi"),
                                },
                            },
                            Replicas: test1.Int64(785555),
                            Revision: test1.Int64(671528),
                            Status: test1.String("nobis"),
                        },
                        "quibusdam": shared.ControllerRequest{
                            Kind: test1.String("omnis"),
                            Message: test1.String("aut"),
                            Pods: []shared.PodStateRequest{
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "reprehenderit": "quia",
                                            "necessitatibus": "accusantium",
                                            "ad": "nisi",
                                        },
                                        map[string]interface{}{
                                            "quia": "laudantium",
                                            "sed": "odit",
                                        },
                                        map[string]interface{}{
                                            "expedita": "eos",
                                            "repellendus": "nesciunt",
                                        },
                                    },
                                    Phase: test1.String("ipsa"),
                                    PodName: test1.String("sint"),
                                    Revision: test1.Int64(291389),
                                    Status: test1.String("esse"),
                                },
                                shared.PodStateRequest{
                                    ContainerStatuses: []map[string]interface{}{
                                        map[string]interface{}{
                                            "sapiente": "quam",
                                            "est": "aliquam",
                                            "delectus": "culpa",
                                        },
                                    },
                                    Phase: test1.String("voluptatum"),
                                    PodName: test1.String("iusto"),
                                    Revision: test1.Int64(802069),
                                    Status: test1.String("voluptatibus"),
                                },
                            },
                            Replicas: test1.Int64(374414),
                            Revision: test1.Int64(247767),
                            Status: test1.String("ullam"),
                        },
                    },
                },
                Remove: []string{
                    "voluptas",
                    "doloribus",
                    "animi",
                },
                Update: map[string][]shared.UpdateActionRequest{
                    "corporis": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("necessitatibus"),
                            Op: test1.String("distinctio"),
                            Path: test1.String("maiores"),
                            Value: test1.String("laboriosam"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("voluptatem"),
                            Op: test1.String("optio"),
                            Path: test1.String("sequi"),
                            Value: test1.String("sunt"),
                        },
                    },
                    "vitae": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("doloremque"),
                            Op: test1.String("sed"),
                            Path: test1.String("amet"),
                            Value: test1.String("rerum"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("in"),
                            Op: test1.String("nostrum"),
                            Path: test1.String("temporibus"),
                            Value: test1.String("ratione"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("dolor"),
                            Op: test1.String("nisi"),
                            Path: test1.String("dignissimos"),
                            Value: test1.String("reiciendis"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("itaque"),
                            Op: test1.String("vitae"),
                            Path: test1.String("est"),
                            Value: test1.String("accusantium"),
                        },
                    },
                    "quod": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("quos"),
                            Op: test1.String("possimus"),
                            Path: test1.String("maiores"),
                            Value: test1.String("odio"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("provident"),
                            Op: test1.String("sapiente"),
                            Path: test1.String("aperiam"),
                            Value: test1.String("similique"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("nesciunt"),
                            Op: test1.String("provident"),
                            Path: test1.String("ex"),
                            Value: test1.String("repellendus"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("unde"),
                            Op: test1.String("alias"),
                            Path: test1.String("impedit"),
                            Value: test1.String("sequi"),
                        },
                    },
                    "commodi": []shared.UpdateActionRequest{
                        shared.UpdateActionRequest{
                            From: test1.String("expedita"),
                            Op: test1.String("in"),
                            Path: test1.String("quisquam"),
                            Value: test1.String("sunt"),
                        },
                        shared.UpdateActionRequest{
                            From: test1.String("enim"),
                            Op: test1.String("nulla"),
                            Path: test1.String("maiores"),
                            Value: test1.String("distinctio"),
                        },
                    },
                },
            },
            Shared: []shared.UpdateActionRequest{
                shared.UpdateActionRequest{
                    From: test1.String("impedit"),
                    Op: test1.String("accusamus"),
                    Path: test1.String("et"),
                    Value: test1.String("quas"),
                },
                shared.UpdateActionRequest{
                    From: test1.String("blanditiis"),
                    Op: test1.String("cum"),
                    Path: test1.String("dicta"),
                    Value: test1.String("impedit"),
                },
                shared.UpdateActionRequest{
                    From: test1.String("tempora"),
                    Op: test1.String("eveniet"),
                    Path: test1.String("repudiandae"),
                    Value: test1.String("sed"),
                },
            },
        },
        AppID: "impedit",
        OrgID: "quas",
        SetID: "impedit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostOrgsOrgIDAppsAppIDSetsSetID200ApplicationJSONString != nil {
        // handle response
    }
}
```
