# Registry

## Overview

Humanitec can be used to manage registry credentials. The Registry object represents how to match credentials to a particular registry.

Humanitec supports all Docker compatible registries as well as the custom authentication formats used by AWS Elastic Container Registry and Google Container Registry. It also supports the injection of a specific secret to be copied from an existing namespace in the cluster.
<SchemaDefinition schemaRef="#/components/schemas/RegistryRequest" />


### Available Operations

* [DeleteOrgsOrgIDRegistriesRegID](#deleteorgsorgidregistriesregid) - Deletes an existing registry record and all associated credentials and secrets.
* [GetOrgsOrgIDRegistries](#getorgsorgidregistries) - Lists available registries for the organization.
* [GetOrgsOrgIDRegistriesRegID](#getorgsorgidregistriesregid) - Loads a registry record details.
* [GetOrgsOrgIDRegistriesRegIDCreds](#getorgsorgidregistriesregidcreds) - Returns current account credentials or secret details for the registry.
* [PatchOrgsOrgIDRegistriesRegID](#patchorgsorgidregistriesregid) - Updates (patches) an existing registry record.
* [PostOrgsOrgIDRegistries](#postorgsorgidregistries) - Creates a new registry record.

## DeleteOrgsOrgIDRegistriesRegID

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
    res, err := s.Registry.DeleteOrgsOrgIDRegistriesRegID(ctx, operations.DeleteOrgsOrgIDRegistriesRegIDRequest{
        OrgID: "at",
        RegID: "alias",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetOrgsOrgIDRegistries

Lists available registries for the organization.

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
    res, err := s.Registry.GetOrgsOrgIDRegistries(ctx, operations.GetOrgsOrgIDRegistriesRequest{
        OrgID: "quia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RegistryResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDRegistriesRegID

Loads a registry record details.

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
    res, err := s.Registry.GetOrgsOrgIDRegistriesRegID(ctx, operations.GetOrgsOrgIDRegistriesRegIDRequest{
        OrgID: "quidem",
        RegID: "fuga",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RegistryResponse != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDRegistriesRegIDCreds

Returns current account credentials or secret details for the registry.

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
    res, err := s.Registry.GetOrgsOrgIDRegistriesRegIDCreds(ctx, operations.GetOrgsOrgIDRegistriesRegIDCredsRequest{
        OrgID: "repudiandae",
        RegID: "accusantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RegistryCredsResponse != nil {
        // handle response
    }
}
```

## PatchOrgsOrgIDRegistriesRegID

Updates (patches) an existing registry record.

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
    res, err := s.Registry.PatchOrgsOrgIDRegistriesRegID(ctx, operations.PatchOrgsOrgIDRegistriesRegIDRequest{
        RegistryRequest: shared.RegistryRequest{
            Creds: &shared.AccountCredsRequest{
                Expires: test1.String("2020-06-22T09:37:23.523Z"),
                Password: "expedita",
                Username: "Shea.Dare51",
            },
            EnableCi: test1.Bool(false),
            ID: "2259e3ea-4b51-497f-9244-3da7ce52b895",
            Registry: "placeat",
            Secrets: map[string]shared.ClusterSecretRequest{
                "neque": shared.ClusterSecretRequest{
                    Namespace: "in",
                    Secret: "minus",
                },
                "eum": shared.ClusterSecretRequest{
                    Namespace: "modi",
                    Secret: "corporis",
                },
            },
            Type: "magnam",
        },
        OrgID: "voluptates",
        RegID: "maiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RegistryResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDRegistries

Creates a new registry record.

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
    res, err := s.Registry.PostOrgsOrgIDRegistries(ctx, operations.PostOrgsOrgIDRegistriesRequest{
        RegistryRequest: shared.RegistryRequest{
            Creds: &shared.AccountCredsRequest{
                Expires: test1.String("2020-06-22T09:37:23.523Z"),
                Password: "tempore",
                Username: "Alvis28",
            },
            EnableCi: test1.Bool(false),
            ID: "896c3ca5-acfb-4e2f-9570-7577929177de",
            Registry: "similique",
            Secrets: map[string]shared.ClusterSecretRequest{
                "ex": shared.ClusterSecretRequest{
                    Namespace: "quaerat",
                    Secret: "commodi",
                },
                "officiis": shared.ClusterSecretRequest{
                    Namespace: "placeat",
                    Secret: "quidem",
                },
                "exercitationem": shared.ClusterSecretRequest{
                    Namespace: "quam",
                    Secret: "dolorem",
                },
                "modi": shared.ClusterSecretRequest{
                    Namespace: "ipsa",
                    Secret: "sint",
                },
            },
            Type: "vero",
        },
        OrgID: "sequi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RegistryResponse != nil {
        // handle response
    }
}
```
