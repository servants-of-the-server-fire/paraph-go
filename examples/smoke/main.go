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

	paraph "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	apiKey := os.Getenv("PARAPH_API_KEY")
	if apiKey == "" {
		log.Fatal("PARAPH_API_KEY must be set")
	}

	cfg := paraph.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	if base := os.Getenv("PARAPH_BASE_URL"); base != "" {
		cfg.Servers = paraph.ServerConfigurations{{URL: base}}
	}

	client := paraph.NewAPIClient(cfg)
	ctx := context.Background()

	fmt.Println("→ GET /account")
	acct, _, err := client.AccountAPI.GetAccount(ctx).Execute()
	if err != nil {
		log.Fatalf("GetAccount: %v", err)
	}
	a := acct.Account
	fmt.Printf("  team=%q plan=%s sandbox=%v\n", a.TeamName, a.Plan, a.SandboxMode)

	fmt.Println("→ GET /templates")
	tmpls, _, err := client.TemplatesAPI.ListTemplates(ctx).Execute()
	if err != nil {
		log.Fatalf("ListTemplates: %v", err)
	}
	fmt.Printf("  %d template(s)\n", len(tmpls.Templates))
	for i, t := range tmpls.Templates {
		if i >= 5 {
			fmt.Printf("  ... (%d more)\n", len(tmpls.Templates)-5)
			break
		}
		fmt.Printf("  - %s  %q\n", t.Id, t.Name)
	}

	fmt.Println("ok")
}
