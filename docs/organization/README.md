# Organization

## Overview

An Organization is the top level object in Humanitec. All other objects belong to an Organization.
<SchemaDefinition schemaRef="#/components/schemas/OrganizationRequest" />


### Available Operations

* [GetOrgs](#getorgs) - List active organizations the user has access to.
* [GetOrgsOrgID](#getorgsorgid) - Get the specified Organization.

## GetOrgs

List active organizations the user has access to.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"test-1"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.Organization.GetOrgs(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.OrganizationResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgID

Get the specified Organization.

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
    res, err := s.Organization.GetOrgsOrgID(ctx, operations.GetOrgsOrgIDRequest{
        OrgID: "cum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OrganizationResponse != nil {
        // handle response
    }
}
```
