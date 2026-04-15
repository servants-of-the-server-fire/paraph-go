# Webhooks

## Overview

Subscribe to events via webhook URLs.


### Available Operations

* [Create](#create) - Create webhook
* [List](#list) - List webhooks
* [Get](#get) - Get webhook
* [Update](#update) - Update webhook
* [Delete](#delete) - Delete webhook
* [Test](#test) - Test webhook

## Create

Registers a URL to receive event notifications.

### Example Usage

<!-- UsageSnippet language="go" operationID="createWebhook" method="post" path="/webhooks" -->
```go
package main

import(
	"context"
	"os"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
	"github.com/servants-of-the-server-fire/paraph-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := paraph.New(
        paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
    )

    res, err := s.Webhooks.Create(ctx, operations.CreateWebhookRequest{
        URL: "https://example.com/webhook",
        Events: []components.WebhookEvent{
            components.WebhookEventRequestSuccess,
            components.WebhookEventSignerSigned,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateWebhookRequest](../../models/operations/createwebhookrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateWebhookResponse](../../models/operations/createwebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 429           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Returns all webhooks registered in your account, newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="listWebhooks" method="get" path="/webhooks" -->
```go
package main

import(
	"context"
	"os"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := paraph.New(
        paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
    )

    res, err := s.Webhooks.List(ctx, paraph.Pointer[int64](1), paraph.Pointer[int64](20))
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `page`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `pageSize`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListWebhooksResponse](../../models/operations/listwebhooksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Returns a single webhook.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWebhook" method="get" path="/webhooks/{id}" -->
```go
package main

import(
	"context"
	"os"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := paraph.New(
        paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
    )

    res, err := s.Webhooks.Get(ctx, "deeb5a05-74d4-40ad-b4be-a9265fd49428")
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetWebhookResponse](../../models/operations/getwebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 404                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Change the webhook URL, subscribed events, or active status. Omitted fields are left unchanged.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateWebhook" method="patch" path="/webhooks/{id}" -->
```go
package main

import(
	"context"
	"os"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"github.com/servants-of-the-server-fire/paraph-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := paraph.New(
        paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
    )

    res, err := s.Webhooks.Update(ctx, "e80a2243-1644-46d7-8f13-7957345de978", operations.UpdateWebhookRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | `string`                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `body`                                                                                     | [operations.UpdateWebhookRequestBody](../../models/operations/updatewebhookrequestbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateWebhookResponse](../../models/operations/updatewebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 404                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Permanently removes the webhook. No further deliveries will be sent.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteWebhook" method="delete" path="/webhooks/{id}" -->
```go
package main

import(
	"context"
	"os"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := paraph.New(
        paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
    )

    res, err := s.Webhooks.Delete(ctx, "2f4cf1de-535d-40b8-9860-de80b52e1022")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteWebhookResponse](../../models/operations/deletewebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 404                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Test

Sends a test event payload to the webhook URL. Use this to verify
your endpoint is correctly receiving and processing webhook deliveries.


### Example Usage

<!-- UsageSnippet language="go" operationID="testWebhook" method="post" path="/webhooks/{id}/test" -->
```go
package main

import(
	"context"
	"os"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := paraph.New(
        paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
    )

    res, err := s.Webhooks.Test(ctx, "4e1ca4d4-efdb-41f2-a630-999c92178d10")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TestWebhookResponse](../../models/operations/testwebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 404                | application/json   |
| apierrors.Error    | 502                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |