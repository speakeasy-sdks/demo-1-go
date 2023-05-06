# Value

## Overview

Shared Values can be used to manage variables and configuration that might vary between environments. They are also the way that secrets can be stored securely.

Shared Values are by default shared across all environments in an application. However, they can be overridden on an Environment by Environment basis.

For example: There might be 2 API keys that are used in an application. One development key used in the development and staging environments and another used for production. The development API key would be set at the Application level. The value would then be overridden at the Environment level for the production Environment.
<SchemaDefinition schemaRef="#/components/schemas/ValueRequest" />


### Available Operations

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValues](#deleteorgsorgidappsappidenvsenvidvalues) - Delete all Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](#deleteorgsorgidappsappidenvsenvidvalueskey) - Delete Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDValues](#deleteorgsorgidappsappidvalues) - Delete all Shared Value for an App
* [DeleteOrgsOrgIDAppsAppIDValuesKey](#deleteorgsorgidappsappidvalueskey) - Delete Shared Value for an Application
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValues](#getorgsorgidappsappidenvsenvidvalues) - List Shared Values in an Environment
* [GetOrgsOrgIDAppsAppIDValues](#getorgsorgidappsappidvalues) - List Shared Values in an Application
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](#patchorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PatchOrgsOrgIDAppsAppIDValuesKey](#patchorgsorgidappsappidvalueskey) - Update Shared Value for an Application
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValues](#postorgsorgidappsappidenvsenvidvalues) - Create a Shared Value for an Environment
* [PostOrgsOrgIDAppsAppIDValues](#postorgsorgidappsappidvalues) - Create a Shared Value for an Application
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](#putorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PutOrgsOrgIDAppsAppIDValuesKey](#putorgsorgidappsappidvalueskey) - Update Shared Value for an Application

## DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValues

All Shared Values will be deleted. If the Shared Values are marked as a secret, they will also be deleted.

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
    res, err := s.Value.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValues(ctx, operations.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesRequest{
        AppID: "quo",
        EnvID: "veniam",
        OrgID: "aliquam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey

The specified Shared Value will be permanently deleted. If the Shared Value is marked as a secret, it will also be permanently deleted.

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
    res, err := s.Value.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey(ctx, operations.DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKeyRequest{
        AppID: "provident",
        EnvID: "vero",
        Key: "earum",
        OrgID: "doloremque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DeleteOrgsOrgIDAppsAppIDValues

All Shared Values will be deleted. If the Shared Values are marked as a secret, they will also be deleted.

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
    res, err := s.Value.DeleteOrgsOrgIDAppsAppIDValues(ctx, operations.DeleteOrgsOrgIDAppsAppIDValuesRequest{
        AppID: "ipsum",
        OrgID: "alias",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DeleteOrgsOrgIDAppsAppIDValuesKey

The specified Shared Value will be permanently deleted. If the Shared Value is marked as a secret, it will also be permanently deleted.

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
    res, err := s.Value.DeleteOrgsOrgIDAppsAppIDValuesKey(ctx, operations.DeleteOrgsOrgIDAppsAppIDValuesKeyRequest{
        AppID: "doloremque",
        Key: "tempora",
        OrgID: "perspiciatis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDEnvsEnvIDValues

The returned values will be the base Application values with the Environment overrides where applicable. The `source` field will specify the level from which the value is from.

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
    res, err := s.Value.GetOrgsOrgIDAppsAppIDEnvsEnvIDValues(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDValuesRequest{
        AppID: "quam",
        EnvID: "atque",
        OrgID: "officia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDValues

The returned values will be the "base" values for the Application. The overridden value for the Environment can be retrieved via the `/orgs/{orgId}/apps/{appId}/envs/{envId}/values` endpoint.

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
    res, err := s.Value.GetOrgsOrgIDAppsAppIDValues(ctx, operations.GetOrgsOrgIDAppsAppIDValuesRequest{
        AppID: "ex",
        OrgID: "architecto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponses != nil {
        // handle response
    }
}
```

## PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey

Update the value or description of the Shared Value. Shared Values marked as secret can also be updated.

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
    res, err := s.Value.PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey(ctx, operations.PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKeyRequest{
        ValuePatchPayloadRequest: shared.ValuePatchPayloadRequest{
            Description: test1.String("a"),
            Value: test1.String("laborum"),
        },
        AppID: "veritatis",
        EnvID: "quod",
        Key: "a",
        OrgID: "qui",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponse != nil {
        // handle response
    }
}
```

## PatchOrgsOrgIDAppsAppIDValuesKey

Update the value or description of the Shared Value. Shared Values marked as secret can also be updated.

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
    res, err := s.Value.PatchOrgsOrgIDAppsAppIDValuesKey(ctx, operations.PatchOrgsOrgIDAppsAppIDValuesKeyRequest{
        ValuePatchPayloadRequest: shared.ValuePatchPayloadRequest{
            Description: test1.String("accusantium"),
            Value: test1.String("commodi"),
        },
        AppID: "atque",
        Key: "totam",
        OrgID: "tenetur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDEnvsEnvIDValues

The Shared Value created will only be available to the specific Environment.

If a Value is marked as a secret, it will be securely stored. It will not be possible to retrieve the value again through the API. The value of the secret can however be updated.

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
    res, err := s.Value.PostOrgsOrgIDAppsAppIDEnvsEnvIDValues(ctx, operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDValuesRequest{
        ValueCreatePayloadRequest: shared.ValueCreatePayloadRequest{
            Description: "voluptate",
            IsSecret: test1.Bool(false),
            Key: "quam",
            Value: "quod",
        },
        AppID: "vitae",
        EnvID: "sapiente",
        OrgID: "reiciendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDValues

The Shared Value created will be available to all Environments in that Application.

If a Value is marked as a secret, it will be securely stored. It will not be possible to retrieve the value again through the API. The value of the secret can however be updated.

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
    res, err := s.Value.PostOrgsOrgIDAppsAppIDValues(ctx, operations.PostOrgsOrgIDAppsAppIDValuesRequest{
        ValueCreatePayloadRequest: shared.ValueCreatePayloadRequest{
            Description: "quod",
            IsSecret: test1.Bool(false),
            Key: "voluptate",
            Value: "inventore",
        },
        AppID: "facere",
        OrgID: "maxime",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponse != nil {
        // handle response
    }
}
```

## PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey

Update the value or description of the Shared Value. Shared Values marked as secret can also be updated.

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
    res, err := s.Value.PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey(ctx, operations.PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKeyRequest{
        ValueEditPayloadRequest: shared.ValueEditPayloadRequest{
            Description: "fuga",
            IsSecret: test1.Bool(false),
            Key: test1.String("ab"),
            Value: "ex",
        },
        AppID: "consectetur",
        EnvID: "maiores",
        Key: "sed",
        OrgID: "animi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponse != nil {
        // handle response
    }
}
```

## PutOrgsOrgIDAppsAppIDValuesKey

Update the value or description of the Shared Value. Shared Values marked as secret can also be updated.

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
    res, err := s.Value.PutOrgsOrgIDAppsAppIDValuesKey(ctx, operations.PutOrgsOrgIDAppsAppIDValuesKeyRequest{
        ValueEditPayloadRequest: shared.ValueEditPayloadRequest{
            Description: "sequi",
            IsSecret: test1.Bool(false),
            Key: test1.String("eligendi"),
            Value: "voluptatum",
        },
        AppID: "perferendis",
        Key: "laborum",
        OrgID: "omnis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueResponse != nil {
        // handle response
    }
}
```
