# Requests

## Overview

A request fills a template's fields and optionally sends the document for signing. Requests without signers produce a filled PDF immediately (status `success`). Requests with signers send each signer an email containing a signing link (status `pending` until all sign). Use the download endpoint to get the final PDF.

**Data retention:** Requests are subject to plan-based retention limits. On the Free plan, requests older than 30 days are no longer accessible. Pro plans retain requests for 5 years. Enterprise plans have unlimited retention.


### Available Operations

* [Create](#create) - Create request
* [List](#list) - List requests
* [Get](#get) - Get request detail
* [Download](#download) - Download request PDF
* [Cancel](#cancel) - Cancel request

## Create

Fills the template's form fields and creates a request.

If `signers` is provided and the template has signature placements
configured, each signer receives a signing link via email. To provide
a signer's signature directly (skipping the email), use
`override_signature_url` in the signer object.

Omitting `signers` produces a fill-only PDF with no signing flow.

Use `GET /requests/{id}/download` to retrieve the filled or signed PDF.


### Example Usage: bad_request

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="bad_request" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "1f5adbea-4edb-4e7a-97f7-26c31fd11d09",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: fill_only

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="fill_only" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "d290f1ee-6c54-4b01-90e6-d701748f0851",
        Fields: map[string]string{
            "employee_name": "Jane Doe",
            "start_date": "2026-04-01",
            "department": "Engineering",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: invalid_email

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="invalid_email" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "1f5adbea-4edb-4e7a-97f7-26c31fd11d09",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: invalid_override_url

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="invalid_override_url" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "1f5adbea-4edb-4e7a-97f7-26c31fd11d09",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: message_too_long

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="message_too_long" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "1f5adbea-4edb-4e7a-97f7-26c31fd11d09",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: multi_signer

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="multi_signer" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "d290f1ee-6c54-4b01-90e6-d701748f0851",
        Fields: map[string]string{
            "employee_name": "Jane Doe",
            "start_date": "2026-04-01",
        },
        Signers: map[string]components.SignerInput{
            "Employee": components.SignerInput{
                Email: "jane.doe@example.com",
            },
            "Manager": components.SignerInput{
                Email: "bob.smith@example.com",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: pre_signed

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="pre_signed" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "d290f1ee-6c54-4b01-90e6-d701748f0851",
        Fields: map[string]string{
            "employee_name": "Jane Doe",
            "start_date": "2026-04-01",
        },
        Signers: map[string]components.SignerInput{
            "Employee": components.SignerInput{
                Email: "jane.doe@example.com",
                OverrideSignatureURL: paraph.Pointer("https://example.com/signatures/jane.png"),
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: single_signer

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="single_signer" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "d290f1ee-6c54-4b01-90e6-d701748f0851",
        Fields: map[string]string{
            "employee_name": "Jane Doe",
            "start_date": "2026-04-01",
        },
        Signers: map[string]components.SignerInput{
            "Employee": components.SignerInput{
                Email: "jane.doe@example.com",
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: title_too_long

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="title_too_long" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "1f5adbea-4edb-4e7a-97f7-26c31fd11d09",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: unfillable

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="unfillable" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "1f5adbea-4edb-4e7a-97f7-26c31fd11d09",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```
### Example Usage: with_signing

<!-- UsageSnippet language="go" operationID="createRequest" method="post" path="/requests" example="with_signing" -->
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

    res, err := s.Requests.Create(ctx, operations.CreateRequestRequestBody{
        TemplateID: "1f5adbea-4edb-4e7a-97f7-26c31fd11d09",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                   | Type                                                                                                                                                                                                                                                        | Required                                                                                                                                                                                                                                                    | Description                                                                                                                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                          | The context to use for the request.                                                                                                                                                                                                                         |
| `body`                                                                                                                                                                                                                                                      | [operations.CreateRequestRequestBody](../../models/operations/createrequestrequestbody.md)                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                          | N/A                                                                                                                                                                                                                                                         |
| `idempotencyKey`                                                                                                                                                                                                                                            | `*string`                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                          | Optional. If provided, ensures the request is processed at most once within 24 hours. Retries with the same key and body return the original response without creating a duplicate. Using the same key with a different request body returns 409 Conflict.<br/> |
| `opts`                                                                                                                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                          | The options for this request.                                                                                                                                                                                                                               |

### Response

**[*operations.CreateRequestResponse](../../models/operations/createrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 409, 429      | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Returns all requests in your account, newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="listRequests" method="get" path="/requests" -->
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

    res, err := s.Requests.List(ctx, operations.ListRequestsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListRequestsRequest](../../models/operations/listrequestsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListRequestsResponse](../../models/operations/listrequestsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Returns the request with its field inputs and signer statuses.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRequest" method="get" path="/requests/{id}" -->
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

    res, err := s.Requests.Get(ctx, "623258bc-9d37-4825-8e7d-4c6c9aaac789")
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Request ID                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetRequestResponse](../../models/operations/getrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 404                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Download

Returns the filled PDF. If all signers have completed, signatures are
applied to the PDF. If signing is still pending, you get the filled
PDF without signatures.


### Example Usage

<!-- UsageSnippet language="go" operationID="downloadRequest" method="get" path="/requests/{id}/download" -->
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

    res, err := s.Requests.Download(ctx, "1ebbb2b9-2ec9-4aea-92bc-5466ac2de6b8")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Request ID                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DownloadRequestResponse](../../models/operations/downloadrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 404                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Cancel

Cancels all pending signers on the request and sets the request status
to `cancelled`. Signers who have already signed are not affected.
Only requests with status `pending` can be cancelled.


### Example Usage

<!-- UsageSnippet language="go" operationID="cancelRequest" method="post" path="/requests/{id}/cancel" -->
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

    res, err := s.Requests.Cancel(ctx, "a0ee08a0-c079-46b4-a5ed-e042adc1b586")
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Request ID                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CancelRequestResponse](../../models/operations/cancelrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 404, 409           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |