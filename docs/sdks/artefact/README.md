# Artefact

## Overview

Artefacts can be registered with Humanitec. Continuous Integration (CI) pipelines notify Humanitec when a new version of an Artefact becomes available. Humanitec tracks the Artefact along with metadata about how it was built.
<SchemaDefinition schemaRef="#/components/schemas/ArtefactRequest" />


### Available Operations

* [DeleteOrgsOrgIDArtefactsArtefactID](#deleteorgsorgidartefactsartefactid) - Delete Artefact and all related Artefact Versions
* [GetOrgsOrgIDArtefacts](#getorgsorgidartefacts) - List all Artefacts.

## DeleteOrgsOrgIDArtefactsArtefactID

The specified Artefact and its Artefact Versions will be permanently deleted. Only Administrators can delete an Artefact.

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
    res, err := s.Artefact.DeleteOrgsOrgIDArtefactsArtefactID(ctx, operations.DeleteOrgsOrgIDArtefactsArtefactIDRequest{
        ArtefactID: "dolores",
        OrgID: "dolorem",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.DeleteOrgsOrgIDArtefactsArtefactIDRequest](../../models/operations/deleteorgsorgidartefactsartefactidrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.DeleteOrgsOrgIDArtefactsArtefactIDResponse](../../models/operations/deleteorgsorgidartefactsartefactidresponse.md), error**


## GetOrgsOrgIDArtefacts

Returns the Artefacts registered with your organization. If no elements are found, an empty list is returned.

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
    res, err := s.Artefact.GetOrgsOrgIDArtefacts(ctx, operations.GetOrgsOrgIDArtefactsRequest{
        Name: test1.String("Rose Rolfson"),
        OrgID: "nemo",
        Type: test1.String("minima"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtefactResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetOrgsOrgIDArtefactsRequest](../../models/operations/getorgsorgidartefactsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetOrgsOrgIDArtefactsResponse](../../models/operations/getorgsorgidartefactsresponse.md), error**

