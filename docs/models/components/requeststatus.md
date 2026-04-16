# RequestStatus

Lifecycle status of a document request.
- `success` — fill completed (and all signers signed, if any)
- `error` — terminal failure
- `pending` — awaiting signer action
- `cancelled` — caller cancelled before completion


## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.RequestStatusSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.RequestStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `RequestStatusSuccess`   | success                  |
| `RequestStatusError`     | error                    |
| `RequestStatusPending`   | pending                  |
| `RequestStatusCancelled` | cancelled                |