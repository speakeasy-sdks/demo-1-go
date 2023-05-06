# Application

## Overview

An Application is a collection of Workloads that work together. When deployed, all Workloads in an Application are deployed to the same namespace.

Apps are the root of the configuration tree holding Environments, Deployments, Shared Values, and Secrets.
<SchemaDefinition schemaRef="#/components/schemas/ApplicationRequest" />


### Available Operations

* [DeleteOrgsOrgIDAppsAppID](#deleteorgsorgidappsappid) - Delete an Application
* [GetOrgsOrgIDApps](#getorgsorgidapps) - List all Applications in an Organization.
* [GetOrgsOrgIDAppsAppID](#getorgsorgidappsappid) - Get an existing Application
* [PostOrgsOrgIDApps](#postorgsorgidapps) - Add a new Application to an Organization

## DeleteOrgsOrgIDAppsAppID

Deleting an Application will also delete everything associated with it. This includes Environments, Deployment history on those Environments, and any shared values and secrets associated.

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
    res, err := s.Application.DeleteOrgsOrgIDAppsAppID(ctx, operations.DeleteOrgsOrgIDAppsAppIDRequest{
        AppID: "iure",
        OrgID: "magnam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetOrgsOrgIDApps

Listing or lists of all Applications that exist within a specific Organization.

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
    res, err := s.Application.GetOrgsOrgIDApps(ctx, operations.GetOrgsOrgIDAppsRequest{
        OrgID: "debitis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppID

Gets a specific Application in the specified Organization by ID.

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
    res, err := s.Application.GetOrgsOrgIDAppsAppID(ctx, operations.GetOrgsOrgIDAppsAppIDRequest{
        AppID: "ipsa",
        OrgID: "delectus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDApps

Creates a new Application, then adds it to the specified Organization.

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
    res, err := s.Application.PostOrgsOrgIDApps(ctx, operations.PostOrgsOrgIDAppsRequest{
        ApplicationCreationRequest: shared.ApplicationCreationRequest{
            Env: &shared.EnvironmentBaseRequest{
                ID: "467cc879-6ed1-451a-85df-c2ddf7cc78ca",
                Name: "Antoinette Nikolaus",
                Type: "deleniti",
            },
            ID: "fc816742-cb73-4920-9929-396fea7596eb",
            Name: "Brenda Wisozk",
        },
        OrgID: "laborum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationResponse != nil {
        // handle response
    }
}
```
