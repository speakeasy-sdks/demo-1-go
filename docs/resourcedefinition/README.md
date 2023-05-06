# ResourceDefinition

## Overview

A Resource Definitions describes how and when a resource should be provisioned. It links a driver (the how) along with a Matching Criteria (the when) to a Resource Type. This allows Humanitec to invoke a particular driver for the required Resource Type in the context of a particular Application and Environment.

The schema for the `driver_inputs` is defined by the `input_schema` property on the DriverDefinition identified by the `driver_type` property.
<SchemaDefinition schemaRef="#/components/schemas/ResourceDefinitionRequest" />


### Available Operations

* [DeleteOrgsOrgIDResourcesDefsDefID](#deleteorgsorgidresourcesdefsdefid) - Delete a Resource Definition.
* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [GetOrgsOrgIDResourcesDefs](#getorgsorgidresourcesdefs) - List Resource Definitions.
* [GetOrgsOrgIDResourcesDefsDefID](#getorgsorgidresourcesdefsdefid) - Get a specific Resource Definition.
* [GetOrgsOrgIDResourcesDefsDefIDResources](#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.
* [PatchOrgsOrgIDResourcesDefsDefID](#patchorgsorgidresourcesdefsdefid) - Update a Resource Definition.
* [PostOrgsOrgIDResourcesDefs](#postorgsorgidresourcesdefs) - Create a new Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.
* [PutOrgsOrgIDResourcesDefsDefID](#putorgsorgidresourcesdefsdefid) - Update a Resource Definition.

## DeleteOrgsOrgIDResourcesDefsDefID

If there **are no** Active Resources provisioned via the current definition, the Resource Definition is deleted immediately.

If there **are** Active Resources provisioned via the current definition, the request fails. The response will describe the changes to the affected Active Resources if operation is forced.

The request can take an optional `force` query parameter. If set to `true`, the current Resource Definition is **marked as** pending deletion and will be deleted (purged) as soon as no existing Active Resources reference it. With the next deployment matching criteria for Resources will be re-evaluated, and current Active Resources for the target environment would be either linked to another matching Resource Definition or decommissioned and created using the new or default Resource Definition (when available).

The Resource Definition that has been marked for deletion cannot be used to provision new resources.

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
    res, err := s.ResourceDefinition.DeleteOrgsOrgIDResourcesDefsDefID(ctx, operations.DeleteOrgsOrgIDResourcesDefsDefIDRequest{
        DefID: "asperiores",
        Force: test1.Bool(false),
        OrgID: "veniam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID

If there **are no** Active Resources that would match to a different Resource Definition when the current Matching Criteria is deleted, the Matching Criteria is deleted immediately.

If there **are** Active Resources that would match to a different Resource Definition, the request fails with HTTP status code 409 (Conflict). The response content will list all of affected Active Resources and their new matches.

The request can take an optional `force` query parameter. If set to `true`, the Matching Criteria is deleted immediately. Referenced Active Resources would match to a different Resource Definition during the next deployment in the target environment.

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
    res, err := s.ResourceDefinition.DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID(ctx, operations.DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaIDRequest{
        CriteriaID: "consequuntur",
        DefID: "facere",
        Force: test1.Bool(false),
        OrgID: "laudantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetOrgsOrgIDResourcesDefs

Filter criteria can be applied to obtain all the resource definitions that could match the filters, grouped by type and sorted by matching rank.

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
    res, err := s.ResourceDefinition.GetOrgsOrgIDResourcesDefs(ctx, operations.GetOrgsOrgIDResourcesDefsRequest{
        App: test1.String("odit"),
        Env: test1.String("pariatur"),
        EnvType: test1.String("amet"),
        OrgID: "exercitationem",
        Res: test1.String("ab"),
        ResType: test1.String("velit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceDefinitionResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDResourcesDefsDefID

Get a specific Resource Definition.

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
    res, err := s.ResourceDefinition.GetOrgsOrgIDResourcesDefsDefID(ctx, operations.GetOrgsOrgIDResourcesDefsDefIDRequest{
        DefID: "facilis",
        OrgID: "tempore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceDefinitionResponse != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDResourcesDefsDefIDResources

List Active Resources provisioned via a specific Resource Definition.

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
    res, err := s.ResourceDefinition.GetOrgsOrgIDResourcesDefsDefIDResources(ctx, operations.GetOrgsOrgIDResourcesDefsDefIDResourcesRequest{
        DefID: "nisi",
        OrgID: "voluptatibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ActiveResourceResponses != nil {
        // handle response
    }
}
```

## PatchOrgsOrgIDResourcesDefsDefID

Update a Resource Definition.

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
    res, err := s.ResourceDefinition.PatchOrgsOrgIDResourcesDefsDefID(ctx, operations.PatchOrgsOrgIDResourcesDefsDefIDRequest{
        PatchResourceDefinitionRequestRequest: shared.PatchResourceDefinitionRequestRequest{
            DriverAccount: test1.String("quaerat"),
            DriverInputs: &shared.ValuesSecretsRequest{
                Secrets: map[string]interface{}{
                    "distinctio": "nisi",
                    "quis": "nisi",
                    "libero": "minus",
                },
                Values: map[string]interface{}{
                    "facilis": "ipsum",
                    "ad": "voluptatibus",
                    "voluptatibus": "consequuntur",
                    "debitis": "labore",
                },
            },
            Name: test1.String("Craig Kiehn"),
        },
        DefID: "iusto",
        OrgID: "est",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceDefinitionResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDResourcesDefs

Create a new Resource Definition.

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
    res, err := s.ResourceDefinition.PostOrgsOrgIDResourcesDefs(ctx, operations.PostOrgsOrgIDResourcesDefsRequest{
        CreateResourceDefinitionRequestRequest: shared.CreateResourceDefinitionRequestRequest{
            Criteria: []shared.MatchingCriteriaRequest{
                shared.MatchingCriteriaRequest{
                    AppID: test1.String("eligendi"),
                    EnvID: test1.String("fugiat"),
                    EnvType: test1.String("unde"),
                    ID: test1.String("e7319c17-7d52-45f7-bb11-4eeb52ff785f"),
                    ResID: test1.String("quisquam"),
                },
                shared.MatchingCriteriaRequest{
                    AppID: test1.String("sequi"),
                    EnvID: test1.String("nihil"),
                    EnvType: test1.String("deleniti"),
                    ID: test1.String("14d4c98e-0c2b-4b89-ab75-dad636c60050"),
                    ResID: test1.String("amet"),
                },
                shared.MatchingCriteriaRequest{
                    AppID: test1.String("illum"),
                    EnvID: test1.String("praesentium"),
                    EnvType: test1.String("quidem"),
                    ID: test1.String("b31180f7-39ae-49e0-97eb-809e2810331f"),
                    ResID: test1.String("dolor"),
                },
            },
            DriverAccount: test1.String("occaecati"),
            DriverInputs: &shared.ValuesSecretsRequest{
                Secrets: map[string]interface{}{
                    "beatae": "at",
                    "labore": "minus",
                    "esse": "voluptatem",
                },
                Values: map[string]interface{}{
                    "rerum": "ea",
                },
            },
            DriverType: "aperiam",
            ID: "7f3c93c7-3b9d-4a3f-aced-a7e23f225741",
            Name: "May Nolan",
            Type: "distinctio",
        },
        OrgID: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceDefinitionResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDResourcesDefsDefIDCriteria

Matching Criteria are combined with Resource Type to select a specific definition. Matching Criteria can be set for any combination of Application ID, Environment ID, Environment Type, and Resource ID. In the event of multiple matches, the most specific match is chosen.

For example, given 3 sets of matching criteria for the same type:

```
 1. {"env_type":"test"}
 2. {"env_type":"development"}
 3. {"env_type":"test", "app_id":"my-app"}
```

If, a resource of that time was needed in an Application `my-app`, Environment `qa-team` with Type `test` and Resource ID `modules.my-module-externals.my-resource`, there would be two resources definitions matching the criteria: #1 & #3. Definition #3 will be chosen because it's matching criteria is the most specific.

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
    res, err := s.ResourceDefinition.PostOrgsOrgIDResourcesDefsDefIDCriteria(ctx, operations.PostOrgsOrgIDResourcesDefsDefIDCriteriaRequest{
        MatchingCriteriaRuleRequest: shared.MatchingCriteriaRuleRequest{
            AppID: test1.String("exercitationem"),
            EnvID: test1.String("labore"),
            EnvType: test1.String("numquam"),
            ResID: test1.String("repudiandae"),
        },
        DefID: "modi",
        OrgID: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MatchingCriteriaResponse != nil {
        // handle response
    }
}
```

## PutOrgsOrgIDResourcesDefsDefID

Update a Resource Definition.

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
    res, err := s.ResourceDefinition.PutOrgsOrgIDResourcesDefsDefID(ctx, operations.PutOrgsOrgIDResourcesDefsDefIDRequest{
        UpdateResourceDefinitionRequestRequest: shared.UpdateResourceDefinitionRequestRequest{
            DriverAccount: test1.String("explicabo"),
            DriverInputs: &shared.ValuesSecretsRequest{
                Secrets: map[string]interface{}{
                    "rem": "aperiam",
                    "odit": "deleniti",
                    "enim": "voluptate",
                    "similique": "minima",
                },
                Values: map[string]interface{}{
                    "magnam": "sit",
                    "modi": "eum",
                    "nesciunt": "mollitia",
                },
            },
            Name: "Hope Hegmann",
        },
        DefID: "reiciendis",
        OrgID: "ab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ResourceDefinitionResponse != nil {
        // handle response
    }
}
```
