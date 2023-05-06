# ActiveResource

## Overview

Active Resources represent the concrete resources provisioned for an Environment. They are provisioned on the first deployment after a dependency on a particular resource type is introduced into an Environment. In general, Active Resources are only deleted when their introductory Environment is deleted.

Active Resources are provisioned based on a Resource Definition. The Resource Definition describes how to provision a concrete resource based on a Resource Type and metadata about the Environment (for example the Environment Type or the Application ID). The criteria for how to choose a Resource Definition is known as a Matching Criteria. If the Matching Criteria changes or the Resource Definition is deleted, the concrete resource represented by an Active Resource might be deleted and reprovisioned when a deployment occurs in the Environment.
<SchemaDefinition schemaRef="#/components/schemas/ActiveResourceRequest" />


### Available Operations

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResID](#deleteorgsorgidappsappidenvsenvidresourcestyperesid) - Delete Active Resources.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDResources](#getorgsorgidappsappidenvsenvidresources) - List Active Resources provisioned in an environment.
* [GetOrgsOrgIDResourcesDefsDefIDResources](#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.

## DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResID

Delete Active Resources.

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
    res, err := s.ActiveResource.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResID(ctx, operations.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResIDRequest{
        AppID: "distinctio",
        EnvID: "quibusdam",
        OrgID: "unde",
        ResID: "nulla",
        Type: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDEnvsEnvIDResources

List Active Resources provisioned in an environment.

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
    res, err := s.ActiveResource.GetOrgsOrgIDAppsAppIDEnvsEnvIDResources(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDResourcesRequest{
        AppID: "illum",
        EnvID: "vel",
        OrgID: "error",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ActiveResourceResponses != nil {
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
    res, err := s.ActiveResource.GetOrgsOrgIDResourcesDefsDefIDResources(ctx, operations.GetOrgsOrgIDResourcesDefsDefIDResourcesRequest{
        DefID: "deserunt",
        OrgID: "suscipit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ActiveResourceResponses != nil {
        // handle response
    }
}
```
