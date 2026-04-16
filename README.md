# paraph

Developer-friendly & type-safe Go SDK specifically catered to leverage *paraph* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=paraph&utm_campaign=go)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/paraph/paraph). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Paraph API: Document API for developers. Upload PDF templates, fill them via API, and send for signing when you need to. Both workflows use the same endpoint.

## Quick Start

Upload a PDF template, then create requests against it.

All requests require `Authorization: Bearer <your-api-key>`.

### Fill a PDF

1. [Upload a template](#tag/templates/POST/templates) with a PDF that has named form fields. The response includes the template `id` and the detected fields.
2. [Create a request](#tag/requests/POST/requests) with the `template_id` and a `fields` map. Status is `success` immediately.
3. [Download the filled PDF](#tag/requests/GET/requests/{id}/download).

---

### Fill + send for signing

Same as above, but add signers. You'll need to set up signature
placements on the template first (one-time).

1. [Upload a template](#tag/templates/POST/templates).
2. [Update the template](#tag/templates/PUT/templates/{id}) with `signature_placements` (page, coordinates, and which signer label goes where).
3. [Create a request](#tag/requests/POST/requests) with a `fields` map and a `signers` map keyed by label. Each signer gets an email with a signing link. Status is `pending` until everyone signs.
4. [Download the signed PDF](#tag/requests/GET/requests/{id}/download) once all signers have completed.

Use [webhooks](#tag/webhooks) to get notified instead of polling.

---

### Pre-signed documents

If you already have a signer's signature image (captured in your own
app, for example), pass `override_signature_url` in the `SignerInput`
when [creating a request](#tag/requests/POST/requests). That signer is
marked as signed immediately and no email is sent. You can mix this
with regular email-based signing on the same request.

---

### Webhooks

1. [Create a webhook](#tag/webhooks/POST/webhooks) with a URL and the events you care about.
2. Each delivery is a POST with `Content-Type: application/json`. See the `WebhookDelivery` schema for the payload format.
3. Verify deliveries using the `secret` returned at creation time. Each delivery includes an `X-Signature-256` header with an HMAC-SHA256 signature.

---

### Metadata

Pass a `metadata` object (up to 10 key-value pairs) when
[creating a request](#tag/requests/POST/requests) to tag it with your
own identifiers. Metadata is returned on all
[request endpoints](#tag/requests) and in webhook deliveries.

---

### Salesforce integration

Paraph offers a managed package on the Salesforce AppExchange. The SF
package talks to this same API — no special endpoints, no separate auth.

**Setup (one-time per SF org):**

1. In Paraph: create an API key (Admin → API Keys). Name it something
   like "Salesforce production".
2. In Salesforce: open the paraph managed package's setup wizard and
   paste the API key. The package stores it in a Named Credential.
3. Done. Every Apex callout auto-injects `Authorization: Bearer <key>`
   via the Named Credential's custom header.

**What Apex code looks like:**

```apex
HttpRequest req = new HttpRequest();
req.setEndpoint('callout:Paraph_Named_Credential/api/v1/requests');
req.setMethod('POST');
req.setHeader('Content-Type', 'application/json');
req.setBody(payload);
HttpResponse res = new Http().send(req);
```

**Rotation:** create a second API key, paste the new one into SF,
delete the old one from Paraph. Both are valid simultaneously so there
is no downtime window.

**Revocation:** delete the key in Paraph. The next SF callout fails
with `401`; the admin pastes a fresh key.

---

### Common headers

All API responses include:
- `X-Request-Id` — unique identifier for the request, useful for debugging and support.
- `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` — rate limit status.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [paraph](#paraph)
  * [Quick Start](#quick-start)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/servants-of-the-server-fire/paraph-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paraph.New(
		paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
	)

	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      | Environment Variable |
| ------------ | ---- | ----------- | -------------------- |
| `BearerAuth` | http | HTTP Bearer | `PARAPH_BEARER_AUTH` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paraph.New(
		paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
	)

	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Accounts](docs/sdks/accounts/README.md)

* [Get](docs/sdks/accounts/README.md#get) - Get account info

### [Requests](docs/sdks/requests/README.md)

* [Create](docs/sdks/requests/README.md#create) - Create request
* [List](docs/sdks/requests/README.md#list) - List requests
* [Get](docs/sdks/requests/README.md#get) - Get request detail
* [Download](docs/sdks/requests/README.md#download) - Download request PDF
* [Cancel](docs/sdks/requests/README.md#cancel) - Cancel request

### [Signers](docs/sdks/signers/README.md)

* [ResendSigningLink](docs/sdks/signers/README.md#resendsigninglink) - Resend signing link
* [DownloadSignature](docs/sdks/signers/README.md#downloadsignature) - Download signer signature

### [Templates](docs/sdks/templates/README.md)

* [List](docs/sdks/templates/README.md#list) - List templates
* [Create](docs/sdks/templates/README.md#create) - Create template
* [Get](docs/sdks/templates/README.md#get) - Get template detail
* [Update](docs/sdks/templates/README.md#update) - Update template
* [Delete](docs/sdks/templates/README.md#delete) - Delete template
* [Download](docs/sdks/templates/README.md#download) - Download template PDF

### [Webhooks](docs/sdks/webhooks/README.md)

* [Create](docs/sdks/webhooks/README.md#create) - Create webhook
* [List](docs/sdks/webhooks/README.md#list) - List webhooks
* [Get](docs/sdks/webhooks/README.md#get) - Get webhook
* [Update](docs/sdks/webhooks/README.md#update) - Update webhook
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete webhook
* [Test](docs/sdks/webhooks/README.md#test) - Test webhook

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
	"os"
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"github.com/servants-of-the-server-fire/paraph-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := paraph.New(
		paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
	)

	res, err := s.Accounts.Get(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"github.com/servants-of-the-server-fire/paraph-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paraph.New(
		paraph.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
	)

	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Get` function may return the following errors:

| Error Type         | Status Code        | Content Type     |
| ------------------ | ------------------ | ---------------- |
| apierrors.Error    | 400, 401, 404, 429 | application/json |
| apierrors.Error    | 500                | application/json |
| apierrors.APIError | 4XX, 5XX           | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"github.com/servants-of-the-server-fire/paraph-go/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paraph.New(
		paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
	)

	res, err := s.Accounts.Get(ctx)
	if err != nil {

		var e *apierrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	paraph "github.com/servants-of-the-server-fire/paraph-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paraph.New(
		paraph.WithServerURL("https://paraph.dev/api/v1"),
		paraph.WithSecurity(os.Getenv("PARAPH_BEARER_AUTH")),
	)

	res, err := s.Accounts.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/servants-of-the-server-fire/paraph-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = paraph.New(paraph.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=paraph&utm_campaign=go)
