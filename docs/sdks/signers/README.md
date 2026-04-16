# Signers

## Overview

Each signer on a request has a label (matching a signature placement on the template), an email address, and a status. Signers receive an email with a link to sign. You can resend the email or download a signer's signature image after they complete signing.


### Available Operations

* [ResendSigningLink](#resendsigninglink) - Resend signing link
* [DownloadSignature](#downloadsignature) - Download signer signature

## ResendSigningLink

Generates a new signing link and resends the signing email.
The previous link is invalidated. Only works for signers
with status `pending`.


### Example Usage

<!-- UsageSnippet language="go" operationID="resendSigningLink" method="post" path="/requests/{id}/signers/{sid}/resend" -->
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

    res, err := s.Signers.ResendSigningLink(ctx, "610c63cc-9f32-4b63-b8cd-f650962a8ec1", "28e1aac3-fe08-4644-b6b7-da78a379ee8d")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Request ID                                               |
| `sid`                                                    | `string`                                                 | :heavy_check_mark:                                       | Signer ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ResendSigningLinkResponse](../../models/operations/resendsigninglinkresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.Error         | 400, 401, 404, 409, 429 | application/json        |
| apierrors.Error         | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DownloadSignature

Returns the stored signature image (PNG) for a signer who has completed
signing, either via the signing page or a pre-provided override_signature_url.


### Example Usage

<!-- UsageSnippet language="go" operationID="downloadSignerSignature" method="get" path="/requests/{id}/signers/{sid}/signature" -->
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

    res, err := s.Signers.DownloadSignature(ctx, "6ffae6c9-b65c-442c-b3be-8fa64c038782", "318c43ea-b47d-46b3-9643-143a7b192f68")
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
| `sid`                                                    | `string`                                                 | :heavy_check_mark:                                       | Signer ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DownloadSignerSignatureResponse](../../models/operations/downloadsignersignatureresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |