// Smoke test for the Paraph Go SDK — calls GET /account and GET /templates
// against a live instance, prints the results. Read-only; safe to run.
//
// Usage:
//
//	PARAPH_API_KEY=your-key go run ./examples/smoke
//
// Optional:
//
//	PARAPH_BASE_URL=http://localhost:8080/api/v1 go run ./examples/smoke

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	paraph "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	apiKey := os.Getenv("PARAPH_API_KEY")
	if apiKey == "" {
		log.Fatal("PARAPH_API_KEY must be set")
	}

	opts := []paraph.SDKOption{
		paraph.WithSecurity(apiKey),
		paraph.WithTimeout(10 * time.Second),
	}
	if baseURL := os.Getenv("PARAPH_BASE_URL"); baseURL != "" {
		opts = append(opts, paraph.WithServerURL(baseURL))
	}

	client := paraph.New(opts...)
	ctx := context.Background()

	fmt.Println("→ GET /account")
	acct, err := client.Accounts.Get(ctx)
	if err != nil {
		log.Fatalf("accounts.Get: %v", err)
	}
	if acct.Object == nil {
		log.Fatalf("accounts.Get returned nil body")
	}
	a := acct.Object.Account
	fmt.Printf("  team=%q plan=%s sandbox=%v\n", a.TeamName, a.Plan, a.SandboxMode)

	fmt.Println("→ GET /templates")
	tmpls, err := client.Templates.List(ctx, nil, nil)
	if err != nil {
		log.Fatalf("templates.List: %v", err)
	}
	if tmpls.TemplateListResponse == nil {
		log.Fatalf("templates.List returned nil body")
	}
	fmt.Printf("  %d template(s)\n", len(tmpls.TemplateListResponse.Templates))
	for i, t := range tmpls.TemplateListResponse.Templates {
		if i >= 5 {
			fmt.Printf("  ... (%d more)\n", len(tmpls.TemplateListResponse.Templates)-5)
			break
		}
		fmt.Printf("  - %s  %q\n", t.ID, t.Name)
	}

	fmt.Println("ok")
}
