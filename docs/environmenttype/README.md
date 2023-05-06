# EnvironmentType

## Overview

Environment Types are a way of grouping and managing Environments. Every Environment has exactly 1 Environment Type.

Environment Types can be used with External Resources to manage where resources such as databases are provisioned or even which cluster to deploy to.
<SchemaDefinition schemaRef="#/components/schemas/EnvironmentTypeRequest" />


### Available Operations

* [DeleteOrgsOrgIDEnvTypesEnvTypeID](#deleteorgsorgidenvtypesenvtypeid) - Deletes an Environment Type
* [GetOrgsOrgIDEnvTypes](#getorgsorgidenvtypes) - List all Environment Types
* [GetOrgsOrgIDEnvTypesEnvTypeID](#getorgsorgidenvtypesenvtypeid) - Get an Environment Type
* [PostOrgsOrgIDEnvTypes](#postorgsorgidenvtypes) - Add a new Environment Type

## DeleteOrgsOrgIDEnvTypesEnvTypeID

Deletes a specific Environment Type from an Organization. If there are Environments with this Type in the Organization, the operation will fail.

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
    res, err := s.EnvironmentType.DeleteOrgsOrgIDEnvTypesEnvTypeID(ctx, operations.DeleteOrgsOrgIDEnvTypesEnvTypeIDRequest{
        EnvTypeID: "accusamus",
        OrgID: "impedit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentTypeResponse != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDEnvTypes

Lists all Environment Types in an Organization.

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
    res, err := s.EnvironmentType.GetOrgsOrgIDEnvTypes(ctx, operations.GetOrgsOrgIDEnvTypesRequest{
        OrgID: "hic",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentTypeResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDEnvTypesEnvTypeID

Gets a specific Environment Type within an Organization.

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
    res, err := s.EnvironmentType.GetOrgsOrgIDEnvTypesEnvTypeID(ctx, operations.GetOrgsOrgIDEnvTypesEnvTypeIDRequest{
        EnvTypeID: "necessitatibus",
        OrgID: "asperiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentTypeResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDEnvTypes

Adds a new Environment Type to an Organization.

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
    res, err := s.EnvironmentType.PostOrgsOrgIDEnvTypes(ctx, operations.PostOrgsOrgIDEnvTypesRequest{
        EnvironmentTypeRequest: shared.EnvironmentTypeRequest{
            Description: test1.String("ex"),
            ID: "6ef1caa3-383c-42be-b477-373c8d72f64d",
        },
        OrgID: "inventore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnvironmentTypeResponse != nil {
        // handle response
    }
}
```
