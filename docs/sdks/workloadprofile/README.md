# WorkloadProfile

## Overview

Workload Profiles provide the baseline configuration for Workloads in Applications in Humanitec. Developers can configure various features of a workload profile to suit their needs. Examples of features might be `schedules` used in Kubernetes CronJobs or `ingress` which might be used to expose Pods controlled by a Kubernetes Deployment.

Workloads in Humanitec are implemented as Helm Charts which must implement a specific schema.
<SchemaDefinition schemaRef="#/components/schemas/WorkloadProfileRequest" />


### Available Operations

* [DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersion](#deleteorgsorgidworkloadprofilesprofileidversionsversion) - Delete a Workload Profile Version
* [DeleteOrgsOrgIDWorkloadProfilesProfileQid](#deleteorgsorgidworkloadprofilesprofileqid) - Delete a Workload Profile
* [GetOrgsOrgIDWorkloadProfiles](#getorgsorgidworkloadprofiles) - List workload profiles available to the organization.
* [GetOrgsOrgIDWorkloadProfilesProfileQid](#getorgsorgidworkloadprofilesprofileqid) - Get a Workload Profile
* [GetOrgsOrgIDWorkloadProfilesProfileQidVersions](#getorgsorgidworkloadprofilesprofileqidversions) - List versions of the given workload profile with optional constraint.
* [PostOrgsOrgIDWorkloadProfiles](#postorgsorgidworkloadprofiles) - Create new Workload Profile
* [PostOrgsOrgIDWorkloadProfilesProfileQidVersions](#postorgsorgidworkloadprofilesprofileqidversions) - Add new Version of the Workload Profile

## DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersion

Delete a Workload Profile Version

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
    res, err := s.WorkloadProfile.DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersion(ctx, operations.DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersionRequest{
        OrgID: "aut",
        ProfileID: "temporibus",
        Version: "tenetur",
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

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersionRequest](../../models/operations/deleteorgsorgidworkloadprofilesprofileidversionsversionrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersionResponse](../../models/operations/deleteorgsorgidworkloadprofilesprofileidversionsversionresponse.md), error**


## DeleteOrgsOrgIDWorkloadProfilesProfileQid

This will also delete all versions of a workload profile.

It is not possible to delete profiles of other organizations.

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
    res, err := s.WorkloadProfile.DeleteOrgsOrgIDWorkloadProfilesProfileQid(ctx, operations.DeleteOrgsOrgIDWorkloadProfilesProfileQidRequest{
        OrgID: "incidunt",
        ProfileQid: "numquam",
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.DeleteOrgsOrgIDWorkloadProfilesProfileQidRequest](../../models/operations/deleteorgsorgidworkloadprofilesprofileqidrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.DeleteOrgsOrgIDWorkloadProfilesProfileQidResponse](../../models/operations/deleteorgsorgidworkloadprofilesprofileqidresponse.md), error**


## GetOrgsOrgIDWorkloadProfiles

List workload profiles available to the organization.

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
    res, err := s.WorkloadProfile.GetOrgsOrgIDWorkloadProfiles(ctx, operations.GetOrgsOrgIDWorkloadProfilesRequest{
        OrgID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkloadProfileResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetOrgsOrgIDWorkloadProfilesRequest](../../models/operations/getorgsorgidworkloadprofilesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetOrgsOrgIDWorkloadProfilesResponse](../../models/operations/getorgsorgidworkloadprofilesresponse.md), error**


## GetOrgsOrgIDWorkloadProfilesProfileQid

Get a Workload Profile

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
    res, err := s.WorkloadProfile.GetOrgsOrgIDWorkloadProfilesProfileQid(ctx, operations.GetOrgsOrgIDWorkloadProfilesProfileQidRequest{
        OrgID: "similique",
        ProfileQid: "dolore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkloadProfileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.GetOrgsOrgIDWorkloadProfilesProfileQidRequest](../../models/operations/getorgsorgidworkloadprofilesprofileqidrequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.GetOrgsOrgIDWorkloadProfilesProfileQidResponse](../../models/operations/getorgsorgidworkloadprofilesprofileqidresponse.md), error**


## GetOrgsOrgIDWorkloadProfilesProfileQidVersions

List versions of the given workload profile with optional constraint.

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
    res, err := s.WorkloadProfile.GetOrgsOrgIDWorkloadProfilesProfileQidVersions(ctx, operations.GetOrgsOrgIDWorkloadProfilesProfileQidVersionsRequest{
        OrgID: "esse",
        ProfileQid: "reiciendis",
        Version: test1.String("iste"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkloadProfileVersionResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.GetOrgsOrgIDWorkloadProfilesProfileQidVersionsRequest](../../models/operations/getorgsorgidworkloadprofilesprofileqidversionsrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.GetOrgsOrgIDWorkloadProfilesProfileQidVersionsResponse](../../models/operations/getorgsorgidworkloadprofilesprofileqidversionsresponse.md), error**


## PostOrgsOrgIDWorkloadProfiles

Create new Workload Profile

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
    res, err := s.WorkloadProfile.PostOrgsOrgIDWorkloadProfiles(ctx, operations.PostOrgsOrgIDWorkloadProfilesRequest{
        WorkloadProfileRequest: shared.WorkloadProfileRequest{
            ID: "390c5888-0983-4dab-b9ef-3ffdd9f7f079",
        },
        OrgID: "similique",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkloadProfileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostOrgsOrgIDWorkloadProfilesRequest](../../models/operations/postorgsorgidworkloadprofilesrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PostOrgsOrgIDWorkloadProfilesResponse](../../models/operations/postorgsorgidworkloadprofilesresponse.md), error**


## PostOrgsOrgIDWorkloadProfilesProfileQidVersions

Creates a Workload Profile Version from the uploaded Helm chart. The version is retrieved from the chart's metadata (Charts.yaml file).

The request has content type `multipart/form-data` and the request body includes two parts:

1. `file` with `application/x-gzip` content type which is an archive containing a Helm chart.

2. `metadata` with `application/json` content type which defines the version's metadata.

Request body example:

	Content-Type: multipart/form-data; boundary=----boundary 	----boundary 	Content-Disposition: form-data; name="metadata" 	Content-Type: application/json; charset=UTF-8 	{ 	  "features": { 	     "humanitec/service": {}, 	     "humanitec/volumes": {}, 	     "custom": {"schema": {}} 	  }, 	  "notes": "Notes related to this version of the profile" 	} 	----boundary 	Content-Disposition: form-data; name="file"; filename="my-workload-1.0.1.tgz" 	Content-Type: application/x-gzip 	[TGZ_DATA] 	----boundary

**NOTE:**

A Workload Profile must be created before a version can be added to it.

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
    res, err := s.WorkloadProfile.PostOrgsOrgIDWorkloadProfilesProfileQidVersions(ctx, operations.PostOrgsOrgIDWorkloadProfilesProfileQidVersionsRequest{
        RequestBody: operations.PostOrgsOrgIDWorkloadProfilesProfileQidVersionsRequestBody{
            File: &operations.PostOrgsOrgIDWorkloadProfilesProfileQidVersionsRequestBodyFile{
                Content: []byte("asperiores"),
                File: "modi",
            },
            Metadata: &shared.WorkloadProfileVersionRequest{
                Features: map[string]interface{}{
                    "neque": "quis",
                    "in": "sed",
                    "non": "porro",
                    "fugiat": "soluta",
                },
                Notes: test1.String("ipsa"),
            },
        },
        OrgID: "reiciendis",
        ProfileQid: "labore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkloadProfileVersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `request`                                                                                                                                              | [operations.PostOrgsOrgIDWorkloadProfilesProfileQidVersionsRequest](../../models/operations/postorgsorgidworkloadprofilesprofileqidversionsrequest.md) | :heavy_check_mark:                                                                                                                                     | The request object to use for the request.                                                                                                             |


### Response

**[*operations.PostOrgsOrgIDWorkloadProfilesProfileQidVersionsResponse](../../models/operations/postorgsorgidworkloadprofilesprofileqidversionsresponse.md), error**

