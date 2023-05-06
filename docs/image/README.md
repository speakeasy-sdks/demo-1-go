# Image

## Overview

DEPRECATED: This type exists for historical compatibility and should not be used. Please use the [Artefact API](https://api-docs.humanitec.com/#tag/Artefact) instead.

Container Images (known simply as Images) can be registered with Humanitec. Continuous Integration (CI) pipelines can then notify Humanitec when a new build of a Container Image becomes available. Humanitec tracks the Image along with metadata about how it was built.
<SchemaDefinition schemaRef="#/components/schemas/ImageRequest" />


### Available Operations

* [GetOrgsOrgIDImages](#getorgsorgidimages) - List all Container Images
* [GetOrgsOrgIDImagesImageID](#getorgsorgidimagesimageid) - Get a specific Image Object
* [GetOrgsOrgIDImagesImageIDBuilds](#getorgsorgidimagesimageidbuilds) - Lists all the Builds of an Image
* [PostOrgsOrgIDImagesImageIDBuilds](#postorgsorgidimagesimageidbuilds) - Add a new Image Build

## GetOrgsOrgIDImages

DEPRECATED: This endpoint exists for historical compatibility and should not be used. Please use the [Artefact API](https://api-docs.humanitec.com/#tag/Artefact) instead.

Lists all of the Container Images registered for this organization.

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
    res, err := s.Image.GetOrgsOrgIDImages(ctx, operations.GetOrgsOrgIDImagesRequest{
        OrgID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDImagesImageID

DEPRECATED: This endpoint exists for historical compatibility and should not be used. Please use the [Artefact API](https://api-docs.humanitec.com/#tag/Artefact) instead.

The response includes a list of Image Builds as well as some metadata about the Image such as its Image Source.

Note, `imageId` may not be the same as the container name. `imageId` is determined by the system making notifications about new builds.

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
    res, err := s.Image.GetOrgsOrgIDImagesImageID(ctx, operations.GetOrgsOrgIDImagesImageIDRequest{
        ImageID: "voluptatem",
        OrgID: "autem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageResponse != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDImagesImageIDBuilds

DEPRECATED: This endpoint exists for historical compatibility and should not be used. Please use the [Artefact API](https://api-docs.humanitec.com/#tag/Artefact) instead.

The response lists all available Image Builds of an Image.

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
    res, err := s.Image.GetOrgsOrgIDImagesImageIDBuilds(ctx, operations.GetOrgsOrgIDImagesImageIDBuildsRequest{
        ImageID: "esse",
        OrgID: "dolores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImageBuildResponses != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDImagesImageIDBuilds

DEPRECATED: This endpoint exists for historical compatibility and should not be used. Please use the [Artefact API](https://api-docs.humanitec.com/#tag/Artefact) instead.

This endpoint is used by Continuous Integration (CI) pipelines to notify Humanitec that a new Image Build is available.

If there is no Image with ID `imageId`, it will be automatically created.

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
    res, err := s.Image.PostOrgsOrgIDImagesImageIDBuilds(ctx, operations.PostOrgsOrgIDImagesImageIDBuildsRequest{
        ImageBuildRequest: shared.ImageBuildRequest{
            Branch: test1.String("assumenda"),
            Commit: test1.String("beatae"),
            Image: test1.String("est"),
            Tags: []string{
                "corrupti",
                "molestiae",
                "provident",
                "accusamus",
            },
        },
        ImageID: "necessitatibus",
        OrgID: "tempore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
