# ArtefactVersion

## Overview

An Artefact Version represents a particular version of an Artefact that can be added to an Application.
<SchemaDefinition schemaRef="#/components/schemas/ArtefactVersionRequest" />


### Available Operations

* [GetOrgsOrgIDArtefactVersions](#getorgsorgidartefactversions) - List all Artefacts Versions.
* [GetOrgsOrgIDArtefactVersionsArtefactVersionID](#getorgsorgidartefactversionsartefactversionid) - Get an Artefacts Versions.
* [GetOrgsOrgIDArtefactsArtefactIDVersions](#getorgsorgidartefactsartefactidversions) - List all Artefact Versions of an Artefact.
* [PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionID](#patchorgsorgidartefactsartefactidversionsversionid) - Update Version of an Artefact.
* [PostOrgsOrgIDArtefactVersions](#postorgsorgidartefactversions) - Register a new Artefact Version with your organization.

## GetOrgsOrgIDArtefactVersions

Returns the Artefact Versions registered with your organization. If no elements are found, an empty list is returned.

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
    res, err := s.ArtefactVersion.GetOrgsOrgIDArtefactVersions(ctx, operations.GetOrgsOrgIDArtefactVersionsRequest{
        Archived: test1.String("excepturi"),
        Name: test1.String("Charlene Nicolas"),
        OrgID: "architecto",
        Reference: test1.String("mollitia"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtefactVersionResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetOrgsOrgIDArtefactVersionsRequest](../../models/operations/getorgsorgidartefactversionsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetOrgsOrgIDArtefactVersionsResponse](../../models/operations/getorgsorgidartefactversionsresponse.md), error**


## GetOrgsOrgIDArtefactVersionsArtefactVersionID

Returns a specific Artefact Version.

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
    res, err := s.ArtefactVersion.GetOrgsOrgIDArtefactVersionsArtefactVersionID(ctx, operations.GetOrgsOrgIDArtefactVersionsArtefactVersionIDRequest{
        ArtefactVersionID: "dolorem",
        OrgID: "culpa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtefactVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.GetOrgsOrgIDArtefactVersionsArtefactVersionIDRequest](../../models/operations/getorgsorgidartefactversionsartefactversionidrequest.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |


### Response

**[*operations.GetOrgsOrgIDArtefactVersionsArtefactVersionIDResponse](../../models/operations/getorgsorgidartefactversionsartefactversionidresponse.md), error**


## GetOrgsOrgIDArtefactsArtefactIDVersions

Returns the Artefact Versions of a specified Artefact registered with your organization. If no elements are found, an empty list is returned.

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
    res, err := s.ArtefactVersion.GetOrgsOrgIDArtefactsArtefactIDVersions(ctx, operations.GetOrgsOrgIDArtefactsArtefactIDVersionsRequest{
        Archived: test1.String("consequuntur"),
        ArtefactID: "repellat",
        Limit: test1.String("mollitia"),
        OrgID: "occaecati",
        Reference: test1.String("numquam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtefactVersionResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GetOrgsOrgIDArtefactsArtefactIDVersionsRequest](../../models/operations/getorgsorgidartefactsartefactidversionsrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GetOrgsOrgIDArtefactsArtefactIDVersionsResponse](../../models/operations/getorgsorgidartefactsartefactidversionsresponse.md), error**


## PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionID

Update the version of a specified Artefact registered with your organization".

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
    res, err := s.ArtefactVersion.PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionID(ctx, operations.PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionIDRequest{
        UpdateArtefactVersionPayloadRequest: shared.UpdateArtefactVersionPayloadRequest{
            Archived: false,
        },
        ArtefactID: "commodi",
        OrgID: "quam",
        VersionID: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtefactVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                    | Type                                                                                                                                                         | Required                                                                                                                                                     | Description                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                        | :heavy_check_mark:                                                                                                                                           | The context to use for the request.                                                                                                                          |
| `request`                                                                                                                                                    | [operations.PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionIDRequest](../../models/operations/patchorgsorgidartefactsartefactidversionsversionidrequest.md) | :heavy_check_mark:                                                                                                                                           | The request object to use for the request.                                                                                                                   |


### Response

**[*operations.PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionIDResponse](../../models/operations/patchorgsorgidartefactsartefactidversionsversionidresponse.md), error**


## PostOrgsOrgIDArtefactVersions

Register a new Artefact Version with your organization.

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
    res, err := s.ArtefactVersion.PostOrgsOrgIDArtefactVersions(ctx, operations.PostOrgsOrgIDArtefactVersionsRequest{
        AddArtefactVersionPayloadRequest: shared.AddArtefactVersionPayloadRequest{
            Commit: test1.String("velit"),
            Digest: test1.String("error"),
            Name: "Beatrice Brown",
            Ref: test1.String("enim"),
            Type: "odit",
            Version: test1.String("quo"),
        },
        OrgID: "sequi",
        Vcs: test1.String("tenetur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ArtefactVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostOrgsOrgIDArtefactVersionsRequest](../../models/operations/postorgsorgidartefactversionsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PostOrgsOrgIDArtefactVersionsResponse](../../models/operations/postorgsorgidartefactversionsresponse.md), error**

