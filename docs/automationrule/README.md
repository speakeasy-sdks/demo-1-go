# AutomationRule

## Overview

An Automation Rule defining how and when artefacts in an environment should be updated.
<SchemaDefinition schemaRef="#/components/schemas/AutomationRuleRequest" />


### Available Operations

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](#deleteorgsorgidappsappidenvsenvidrulesruleid) - Delete Automation Rule from an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRules](#getorgsorgidappsappidenvsenvidrules) - List all Automation Rules in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](#getorgsorgidappsappidenvsenvidrulesruleid) - Get a specific Automation Rule for an Environment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDRules](#postorgsorgidappsappidenvsenvidrules) - Create a new Automation Rule for an Environment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](#putorgsorgidappsappidenvsenvidrulesruleid) - Update an existing Automation Rule for an Environment.

## DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID

Delete Automation Rule from an Environment.

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
    res, err := s.AutomationRule.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID(ctx, operations.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleIDRequest{
        AppID: "ipsam",
        EnvID: "id",
        OrgID: "possimus",
        RuleID: "aut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDEnvsEnvIDRules

List all Automation Rules in an Environment.

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
    res, err := s.AutomationRule.GetOrgsOrgIDAppsAppIDEnvsEnvIDRules(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRequest{
        AppID: "quasi",
        EnvID: "error",
        OrgID: "temporibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AutomationRuleResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID

Get a specific Automation Rule for an Environment.

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
    res, err := s.AutomationRule.GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleIDRequest{
        AppID: "laborum",
        EnvID: "quasi",
        OrgID: "reiciendis",
        RuleID: "voluptatibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AutomationRuleResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDEnvsEnvIDRules

Items marked as deprecated are still supported (however not recommended) for use and are incompatible with properties of the latest api version. In particular an error is raised if  `images_filter` (deprecated) and `artefacts_filter` are used in the same payload. The same is true for `exclude_images_filter` (deprecated) and `exclude_artefacts_filter`. `match` and `update_to` are still supported but will trigger an error if combined with `match_ref`.

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
    res, err := s.AutomationRule.PostOrgsOrgIDAppsAppIDEnvsEnvIDRules(ctx, operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDRulesRequest{
        AutomationRuleRequest: shared.AutomationRuleRequest{
            Active: test1.Bool(false),
            ArtefactsFilter: []string{
                "nihil",
                "praesentium",
                "voluptatibus",
                "ipsa",
            },
            ExcludeArtefactsFilter: test1.Bool(false),
            ExcludeImagesFilter: test1.Bool(false),
            ImagesFilter: []string{
                "voluptate",
                "cum",
                "perferendis",
            },
            Match: test1.String("doloremque"),
            MatchRef: test1.String("reprehenderit"),
            Type: "ut",
            UpdateTo: test1.String("maiores"),
        },
        AppID: "dicta",
        EnvID: "corporis",
        OrgID: "dolore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AutomationRuleResponse != nil {
        // handle response
    }
}
```

## PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID

Items marked as deprecated are still supported (however not recommended) for use and are incompatible with properties of the latest api version. In particular an error is raised if  `images_filter` (deprecated) and `artefacts_filter` are used in the same payload. The same is true for `exclude_images_filter` (deprecated) and `exclude_artefacts_filter`. `match` and `update_to` are still supported but will trigger an error if combined with `match_ref`.

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
    res, err := s.AutomationRule.PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID(ctx, operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleIDRequest{
        AutomationRuleRequest: shared.AutomationRuleRequest{
            Active: test1.Bool(false),
            ArtefactsFilter: []string{
                "dicta",
                "harum",
            },
            ExcludeArtefactsFilter: test1.Bool(false),
            ExcludeImagesFilter: test1.Bool(false),
            ImagesFilter: []string{
                "accusamus",
                "commodi",
            },
            Match: test1.String("repudiandae"),
            MatchRef: test1.String("quae"),
            Type: "ipsum",
            UpdateTo: test1.String("quidem"),
        },
        AppID: "molestias",
        EnvID: "excepturi",
        OrgID: "pariatur",
        RuleID: "modi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AutomationRuleResponse != nil {
        // handle response
    }
}
```
