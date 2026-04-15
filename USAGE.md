<!-- Start SDK Example Usage [usage] -->
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