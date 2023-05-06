# ValueSetVersion

## Overview

A Value Set Version can be used as a track record of Shared Values changes, to restore a previous version of a Shared Value or Value Set, or to purge a Shared Value if it shouldn't be accessible anymore.
<SchemaDefinition schemaRef="#/components/schemas/ValueSetVersionResponse" />


### Available Operations

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersions](#getorgsorgidappsappidenvsenvidvaluesetversions) - List Value Set Versions in an Environment of an App
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionID](#getorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionid) - Get a single Value Set Version in an Environment of an App
* [GetOrgsOrgIDAppsAppIDValueSetVersions](#getorgsorgidappsappidvaluesetversions) - List Value Set Versions in the App
* [GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionID](#getorgsorgidappsappidvaluesetversionsvaluesetversionid) - Get a single Value Set Version from the App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKey](#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Environment Version history.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestore](#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKey](#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKey](#postorgsorgidappsappidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Version history.
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestore](#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKey](#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an App

## GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersions

A new Value Set Version is created on every modification of a Value inside the an Environment of an App. In case this environment has no overrides the response is the same as the App level endpoint.

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
    res, err := s.ValueSetVersion.GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersions(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsRequest{
        AppID: "nihil",
        EnvID: "tenetur",
        KeyChanged: test1.String("sapiente"),
        OrgID: "velit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionID

Get a single Value Set Version in an Environment of an App

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
    res, err := s.ValueSetVersion.GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionID(ctx, operations.GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRequest{
        AppID: "adipisci",
        EnvID: "non",
        OrgID: "optio",
        ValueSetVersionID: "ddf857a9-e618-476c-aab2-1d29dfc94d6f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponse != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDValueSetVersions

A new Value Set Version is created on every modification of a Value inside the app.

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
    res, err := s.ValueSetVersion.GetOrgsOrgIDAppsAppIDValueSetVersions(ctx, operations.GetOrgsOrgIDAppsAppIDValueSetVersionsRequest{
        AppID: "recusandae",
        KeyChanged: test1.String("quisquam"),
        OrgID: "facere",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponses != nil {
        // handle response
    }
}
```

## GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionID

Get a single Value Set Version from the App

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
    res, err := s.ValueSetVersion.GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionID(ctx, operations.GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRequest{
        AppID: "dignissimos",
        OrgID: "iste",
        ValueSetVersionID: "9390066a-6d2d-4000-b553-38cec086fa21",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKey

Purging permanently removes the value of a specific Shared Value in an application. A purged value is no longer accessible, can't be restored and can't be used
by deployments referencing a Value Set Version where the value was present.

Learn more about purging in our [docs](https://docs.humanitec.com/reference/concepts/app-config/shared-app-values#purge).


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
    res, err := s.ValueSetVersion.PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKey(ctx, operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKeyRequest{
        ValueSetActionPayloadRequest: shared.ValueSetActionPayloadRequest{
            Comment: test1.String("necessitatibus"),
        },
        AppID: "iste",
        EnvID: "dicta",
        Key: "ipsam",
        OrgID: "consequuntur",
        ValueSetVersionID: "cb311916-7b8e-43c8-9b03-408d6d364ffd",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestore

Restore the values of all Shared Values in an environment from a specific version. Keys not existing in the selected version are deleted.

Learn more about reverting in our [docs](https://docs.humanitec.com/reference/concepts/app-config/shared-app-values#revert).


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
    res, err := s.ValueSetVersion.PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestore(ctx, operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreRequest{
        ValueSetActionPayloadRequest: shared.ValueSetActionPayloadRequest{
            Comment: test1.String("dolore"),
        },
        AppID: "enim",
        EnvID: "ullam",
        OrgID: "perspiciatis",
        ValueSetVersionID: "06d1263d-48e9-435c-ac9e-81f30be3e432",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKey

Restore the values of a single Shared Value in an Environment from a specific version.

Learn more about reverting in our [docs](https://docs.humanitec.com/reference/concepts/app-config/shared-app-values#revert).


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
    res, err := s.ValueSetVersion.PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKey(ctx, operations.PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKeyRequest{
        ValueSetActionPayloadRequest: shared.ValueSetActionPayloadRequest{
            Comment: test1.String("aperiam"),
        },
        AppID: "dolores",
        EnvID: "illum",
        Key: "iusto",
        OrgID: "magni",
        ValueSetVersionID: "16576506-6418-470d-9d21-f9ad030c4ecc",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKey

Purging permanently removes the value of a specific Shared Value in an Application. A purged value is no longer accessible, can't be restored and can't be used
by deployments referencing a Value Set Version where the value was present.

Learn more about purging in our [docs](https://docs.humanitec.com/reference/concepts/app-config/shared-app-values#purge).


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
    res, err := s.ValueSetVersion.PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKey(ctx, operations.PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKeyRequest{
        ValueSetActionPayloadRequest: shared.ValueSetActionPayloadRequest{
            Comment: test1.String("et"),
        },
        AppID: "beatae",
        Key: "id",
        OrgID: "consequatur",
        ValueSetVersionID: "83642906-8b85-402a-95e7-f73bc845e320",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestore

Restore the values of all Shared Values in an application from a specific version. Keys not existing in the selected version are deleted.

Learn more about reverting in our [docs](https://docs.humanitec.com/reference/concepts/app-config/shared-app-values#revert).


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
    res, err := s.ValueSetVersion.PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestore(ctx, operations.PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreRequest{
        ValueSetActionPayloadRequest: shared.ValueSetActionPayloadRequest{
            Comment: test1.String("est"),
        },
        AppID: "amet",
        OrgID: "veritatis",
        ValueSetVersionID: "9f4badf9-47c9-4a86-bbc4-2426665816dd",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponse != nil {
        // handle response
    }
}
```

## PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKey

Restore the values of a single Shared Value in an application from a specific version.

Learn more about reverting in our [docs](https://docs.humanitec.com/reference/concepts/app-config/shared-app-values#revert).


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
    res, err := s.ValueSetVersion.PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKey(ctx, operations.PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKeyRequest{
        ValueSetActionPayloadRequest: shared.ValueSetActionPayloadRequest{
            Comment: test1.String("impedit"),
        },
        AppID: "culpa",
        Key: "atque",
        OrgID: "voluptates",
        ValueSetVersionID: "f51fcb4c-593e-4c12-8daa-d0ec7afedbd8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValueSetVersionResponse != nil {
        // handle response
    }
}
```
