// Smoke test for the Paraph Go SDK.
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

	paraph "github.com/servants-of-the-server-fire/paraph-go"
)

func main() {
	client, err := paraph.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	fmt.Println("→ GET /account")
	acct, _, err := client.Account.GetAccount(ctx).Execute()
	if err != nil {
		log.Fatalf("GetAccount: %v", err)
	}
	a := acct.Account
	fmt.Printf("  team=%q plan=%s sandbox=%v\n", a.TeamName, a.Plan, a.SandboxMode)

	fmt.Println("→ GET /templates")
	tmpls, _, err := client.Templates.ListTemplates(ctx).Execute()
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
