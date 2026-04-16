# Templates

## Overview

A template is a PDF you upload once and reuse for every request. Paraph reads the PDF's AcroForm fields (text, date, checkbox, etc.) and returns them so you know what values to pass when creating a request. If the template needs signatures, add signature placements to define where each signer signs on the page. PDFs using XFA forms (Adobe LiveCycle) are not supported.


### Available Operations

* [List](#list) - List templates
* [Create](#create) - Create template
* [Get](#get) - Get template detail
* [Update](#update) - Update template
* [Delete](#delete) - Delete template
* [Download](#download) - Download template PDF

## List

Returns all templates in your account, newest first.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTemplates" method="get" path="/templates" -->
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

    res, err := s.Templates.List(ctx, paraph.Pointer[int64](1), paraph.Pointer[int64](20))
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateListResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
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

**[*operations.ListTemplatesResponse](../../models/operations/listtemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Upload a PDF with named form fields. Paraph detects the fields automatically.

**Provide exactly one of `file` or `file_url`.** Requests with neither or both return `400 Bad Request`. This constraint is validated server-side — OpenAPI doesn't support expressing `oneOf` inside a multipart request body, so it's documented in prose.


### Example Usage: file_url_download_failed

<!-- UsageSnippet language="go" operationID="createTemplate" method="post" path="/templates" example="file_url_download_failed" -->
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

    res, err := s.Templates.Create(ctx, operations.CreateTemplateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```
### Example Usage: invalid_file_url

<!-- UsageSnippet language="go" operationID="createTemplate" method="post" path="/templates" example="invalid_file_url" -->
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

    res, err := s.Templates.Create(ctx, operations.CreateTemplateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```
### Example Usage: name_too_long

<!-- UsageSnippet language="go" operationID="createTemplate" method="post" path="/templates" example="name_too_long" -->
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

    res, err := s.Templates.Create(ctx, operations.CreateTemplateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```
### Example Usage: not_a_pdf

<!-- UsageSnippet language="go" operationID="createTemplate" method="post" path="/templates" example="not_a_pdf" -->
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

    res, err := s.Templates.Create(ctx, operations.CreateTemplateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```
### Example Usage: unsupported_pdf

<!-- UsageSnippet language="go" operationID="createTemplate" method="post" path="/templates" example="unsupported_pdf" -->
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

    res, err := s.Templates.Create(ctx, operations.CreateTemplateRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateTemplateRequest](../../models/operations/createtemplaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateTemplateResponse](../../models/operations/createtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Returns the template with its detected form fields and signature placements.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTemplate" method="get" path="/templates/{id}" -->
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

    res, err := s.Templates.Get(ctx, "a1ca9eb4-6e28-4e86-9856-a51ec50c9041")
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
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

**[*operations.GetTemplateResponse](../../models/operations/gettemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Change the template name, metadata, or signature placements. Omitted fields are left unchanged.

### Example Usage: name_too_long

<!-- UsageSnippet language="go" operationID="updateTemplate" method="patch" path="/templates/{id}" example="name_too_long" -->
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

    res, err := s.Templates.Update(ctx, "f0a2af12-9e30-4dd4-999b-a4512485b634", operations.UpdateTemplateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```
### Example Usage: signer_label_invalid_chars

<!-- UsageSnippet language="go" operationID="updateTemplate" method="patch" path="/templates/{id}" example="signer_label_invalid_chars" -->
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

    res, err := s.Templates.Update(ctx, "9bee6550-0926-4b9b-b4f0-7add978374ef", operations.UpdateTemplateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```
### Example Usage: signer_label_required

<!-- UsageSnippet language="go" operationID="updateTemplate" method="patch" path="/templates/{id}" example="signer_label_required" -->
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

    res, err := s.Templates.Update(ctx, "745d4d0b-c28b-49a8-b4c0-e07c69e4f3b1", operations.UpdateTemplateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```
### Example Usage: signer_label_too_long

<!-- UsageSnippet language="go" operationID="updateTemplate" method="patch" path="/templates/{id}" example="signer_label_too_long" -->
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

    res, err := s.Templates.Update(ctx, "71670bac-1007-46d4-9661-6f0a86fbe3d0", operations.UpdateTemplateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TemplateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `id`                                                                                         | `string`                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `body`                                                                                       | [operations.UpdateTemplateRequestBody](../../models/operations/updatetemplaterequestbody.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateTemplateResponse](../../models/operations/updatetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Permanently removes the template. Existing requests created from it are not affected.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTemplate" method="delete" path="/templates/{id}" -->
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

    res, err := s.Templates.Delete(ctx, "30543acd-947e-422a-9967-d4631e815e08")
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

**[*operations.DeleteTemplateResponse](../../models/operations/deletetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Download

Returns the original PDF that was uploaded when the template was created.

### Example Usage

<!-- UsageSnippet language="go" operationID="downloadTemplate" method="get" path="/templates/{id}/download" -->
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

    res, err := s.Templates.Download(ctx, "c8dc02ab-8924-4437-a2b6-7aad75eecf15")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DownloadTemplateResponse](../../models/operations/downloadtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |