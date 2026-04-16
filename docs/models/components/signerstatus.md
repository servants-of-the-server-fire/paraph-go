# SignerStatus

Per-signer status.
- `pending` — signing link sent, awaiting signer
- `signed` — signer completed
- `expired` — signer did not sign before `expires_at`
- `error` — terminal failure for this signer
- `cancelled` — parent request was cancelled


## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.SignerStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.SignerStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `SignerStatusPending`   | pending                 |
| `SignerStatusSigned`    | signed                  |
| `SignerStatusExpired`   | expired                 |
| `SignerStatusError`     | error                   |
| `SignerStatusCancelled` | cancelled               |