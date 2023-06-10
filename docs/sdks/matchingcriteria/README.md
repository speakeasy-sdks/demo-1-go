# MatchingCriteria

## Overview

Matching Criteria are a set of rules used to choose which Resource Definition to use to provision a particular Resource Type.

Matching criteria are made up in order of specificity with least specific first:

- Environment Type (`env_type`)

- Application (`app_id`)

- Environment (`env_id`)

- Resource (`res_id`)

When selecting matching criteria, the most specific one is selected. In general, this means of all the Matching Criteria fully matching the context, the Matching Criteria Rule with the most specific element filled is chosen. If there is a tie, the next most specific elements are compared and so on until one is chosen.

**NOTE:**

Humanitec will reject the registration of matching criteria rules that duplicate rules already present for a Resource Type.
<SchemaDefinition schemaRef="#/components/schemas/MatchingCriteriaRequest" />


### Available Operations

* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.

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
    res, err := s.MatchingCriteria.DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID(ctx, operations.DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaIDRequest{
        CriteriaID: "sint",
        DefID: "ea",
        Force: test1.Bool(false),
        OrgID: "autem",
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

| Parameter                                                                                                                                                      | Type                                                                                                                                                           | Required                                                                                                                                                       | Description                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                                          | :heavy_check_mark:                                                                                                                                             | The context to use for the request.                                                                                                                            |
| `request`                                                                                                                                                      | [operations.DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaIDRequest](../../models/operations/deleteorgsorgidresourcesdefsdefidcriteriacriteriaidrequest.md) | :heavy_check_mark:                                                                                                                                             | The request object to use for the request.                                                                                                                     |


### Response

**[*operations.DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaIDResponse](../../models/operations/deleteorgsorgidresourcesdefsdefidcriteriacriteriaidresponse.md), error**


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
    res, err := s.MatchingCriteria.PostOrgsOrgIDResourcesDefsDefIDCriteria(ctx, operations.PostOrgsOrgIDResourcesDefsDefIDCriteriaRequest{
        MatchingCriteriaRuleRequest: shared.MatchingCriteriaRuleRequest{
            AppID: test1.String("ipsam"),
            EnvID: test1.String("rerum"),
            EnvType: test1.String("laudantium"),
            ResID: test1.String("corporis"),
        },
        DefID: "officiis",
        OrgID: "voluptatibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MatchingCriteriaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.PostOrgsOrgIDResourcesDefsDefIDCriteriaRequest](../../models/operations/postorgsorgidresourcesdefsdefidcriteriarequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.PostOrgsOrgIDResourcesDefsDefIDCriteriaResponse](../../models/operations/postorgsorgidresourcesdefsdefidcriteriaresponse.md), error**

