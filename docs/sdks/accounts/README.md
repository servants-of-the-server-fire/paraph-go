# Accounts

## Overview

### Available Operations

* [Get](#get) - Get account info

## Get

Returns the current team's plan, usage counters, and limits for the API key used in the request.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccount" method="get" path="/account" -->
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

    res, err := s.Accounts.Get(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAccountResponse](../../models/operations/getaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404, 429 | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |