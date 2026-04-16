# paraph-go

Go client for the [Paraph API](https://paraph.dev). Upload PDF templates, fill fields, send for signing.

## Install

```bash
go get github.com/servants-of-the-server-fire/paraph-go@v0.3.0
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	paraph "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	cfg := paraph.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer YOUR_API_KEY")
	client := paraph.NewAPIClient(cfg)

	ctx := context.Background()

	// List templates
	resp, _, err := client.TemplatesAPI.ListTemplates(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range resp.Templates {
		fmt.Printf("%s  %s\n", t.Id, t.Name)
	}

	// Create a request (fill a PDF)
	req := *paraph.NewCreateRequestRequest("TEMPLATE_ID")
	req.Fields = &map[string]string{"name": "Jane Doe", "date": "2026-04-15"}
	result, _, err := client.RequestsAPI.CreateRequest(ctx).CreateRequestRequest(req).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Request.Id)
}
```

## Auth

All requests need a Bearer token. Get an API key from your Paraph dashboard under Admin > API Keys.

```go
cfg := paraph.NewConfiguration()
cfg.AddDefaultHeader("Authorization", "Bearer " + os.Getenv("PARAPH_API_KEY"))
```

## Docs

Full API reference at [paraph.dev/docs](https://paraph.dev/docs).

## License

MIT
