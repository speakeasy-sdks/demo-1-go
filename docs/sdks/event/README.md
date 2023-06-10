# Event

## Overview

Webhook is a special type of a Job, it performs a HTTPS request to a specified URL with specified headers.
<SchemaDefinition schemaRef="#/components/schemas/WebhookRequest" />


### Available Operations

* [DeleteOrgsOrgIDAppsAppIDJobs](#deleteorgsorgidappsappidjobs) - Deletes all Jobs for the Application
* [DeleteOrgsOrgIDAppsAppIDWebhooksJobID](#deleteorgsorgidappsappidwebhooksjobid) - Delete a Webhook
* [GetOrgsOrgIDAppsAppIDWebhooks](#getorgsorgidappsappidwebhooks) - List Webhooks
* [GetOrgsOrgIDAppsAppIDWebhooksJobID](#getorgsorgidappsappidwebhooksjobid) - Get a Webhook
* [GetOrgsOrgIDEvents](#getorgsorgidevents) - List Events
* [PostOrgsOrgIDAppsAppIDWebhooks](#postorgsorgidappsappidwebhooks) - Create a new Webhook
* [PostOrgsOrgIDAppsAppIDWebhooksJobID](#postorgsorgidappsappidwebhooksjobid) - Update a Webhook

## DeleteOrgsOrgIDAppsAppIDJobs

Deletes all Jobs for the Application

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
    res, err := s.Event.DeleteOrgsOrgIDAppsAppIDJobs(ctx, operations.DeleteOrgsOrgIDAppsAppIDJobsRequest{
        AppID: "facere",
        OrgID: "libero",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.DeleteOrgsOrgIDAppsAppIDJobsRequest](../../models/operations/deleteorgsorgidappsappidjobsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.DeleteOrgsOrgIDAppsAppIDJobsResponse](../../models/operations/deleteorgsorgidappsappidjobsresponse.md), error**


## DeleteOrgsOrgIDAppsAppIDWebhooksJobID

Delete a Webhook

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
    res, err := s.Event.DeleteOrgsOrgIDAppsAppIDWebhooksJobID(ctx, operations.DeleteOrgsOrgIDAppsAppIDWebhooksJobIDRequest{
        AppID: "architecto",
        JobID: "voluptatibus",
        OrgID: "quia",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.DeleteOrgsOrgIDAppsAppIDWebhooksJobIDRequest](../../models/operations/deleteorgsorgidappsappidwebhooksjobidrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |


### Response

**[*operations.DeleteOrgsOrgIDAppsAppIDWebhooksJobIDResponse](../../models/operations/deleteorgsorgidappsappidwebhooksjobidresponse.md), error**


## GetOrgsOrgIDAppsAppIDWebhooks

List Webhooks

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
    res, err := s.Event.GetOrgsOrgIDAppsAppIDWebhooks(ctx, operations.GetOrgsOrgIDAppsAppIDWebhooksRequest{
        AppID: "porro",
        OrgID: "aliquam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetOrgsOrgIDAppsAppIDWebhooksRequest](../../models/operations/getorgsorgidappsappidwebhooksrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDWebhooksResponse](../../models/operations/getorgsorgidappsappidwebhooksresponse.md), error**


## GetOrgsOrgIDAppsAppIDWebhooksJobID

Get a Webhook

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
    res, err := s.Event.GetOrgsOrgIDAppsAppIDWebhooksJobID(ctx, operations.GetOrgsOrgIDAppsAppIDWebhooksJobIDRequest{
        AppID: "velit",
        JobID: "illo",
        OrgID: "accusantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.GetOrgsOrgIDAppsAppIDWebhooksJobIDRequest](../../models/operations/getorgsorgidappsappidwebhooksjobidrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.GetOrgsOrgIDAppsAppIDWebhooksJobIDResponse](../../models/operations/getorgsorgidappsappidwebhooksjobidresponse.md), error**


## GetOrgsOrgIDEvents

List Events

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
    res, err := s.Event.GetOrgsOrgIDEvents(ctx, operations.GetOrgsOrgIDEventsRequest{
        OrgID: "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetOrgsOrgIDEventsRequest](../../models/operations/getorgsorgideventsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetOrgsOrgIDEventsResponse](../../models/operations/getorgsorgideventsresponse.md), error**


## PostOrgsOrgIDAppsAppIDWebhooks

Create a new Webhook

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
    res, err := s.Event.PostOrgsOrgIDAppsAppIDWebhooks(ctx, operations.PostOrgsOrgIDAppsAppIDWebhooksRequest{
        WebhookRequest: shared.WebhookRequest{
            Disabled: test1.Bool(false),
            Headers: map[string]interface{}{
                "beatae": "vero",
                "excepturi": "eum",
            },
            ID: test1.String("349e1cf9-e06e-43a4-b700-0ae6b6bc9b8f"),
            Payload: map[string]interface{}{
                "ullam": "unde",
                "necessitatibus": "animi",
            },
            Triggers: []shared.EventBaseRequest{
                shared.EventBaseRequest{
                    Scope: test1.String("ipsam"),
                    Type: test1.String("corporis"),
                },
                shared.EventBaseRequest{
                    Scope: test1.String("est"),
                    Type: test1.String("error"),
                },
                shared.EventBaseRequest{
                    Scope: test1.String("esse"),
                    Type: test1.String("labore"),
                },
                shared.EventBaseRequest{
                    Scope: test1.String("veritatis"),
                    Type: test1.String("vero"),
                },
            },
            URL: test1.String("consectetur"),
        },
        AppID: "vitae",
        OrgID: "inventore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostOrgsOrgIDAppsAppIDWebhooksRequest](../../models/operations/postorgsorgidappsappidwebhooksrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.PostOrgsOrgIDAppsAppIDWebhooksResponse](../../models/operations/postorgsorgidappsappidwebhooksresponse.md), error**


## PostOrgsOrgIDAppsAppIDWebhooksJobID

Update a Webhook

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
    res, err := s.Event.PostOrgsOrgIDAppsAppIDWebhooksJobID(ctx, operations.PostOrgsOrgIDAppsAppIDWebhooksJobIDRequest{
        WebhookRequest: shared.WebhookRequest{
            Disabled: test1.Bool(false),
            Headers: map[string]interface{}{
                "ad": "qui",
            },
            ID: test1.String("965bb8a7-2026-4114-b5e1-39dbc2259b1a"),
            Payload: map[string]interface{}{
                "fugiat": "officia",
                "quos": "placeat",
                "sit": "iusto",
            },
            Triggers: []shared.EventBaseRequest{
                shared.EventBaseRequest{
                    Scope: test1.String("voluptates"),
                    Type: test1.String("inventore"),
                },
            },
            URL: test1.String("aperiam"),
        },
        AppID: "totam",
        JobID: "dolore",
        OrgID: "eligendi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.PostOrgsOrgIDAppsAppIDWebhooksJobIDRequest](../../models/operations/postorgsorgidappsappidwebhooksjobidrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.PostOrgsOrgIDAppsAppIDWebhooksJobIDResponse](../../models/operations/postorgsorgidappsappidwebhooksjobidresponse.md), error**

