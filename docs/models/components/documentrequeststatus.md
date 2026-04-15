# DocumentRequestStatus

Request status

## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.DocumentRequestStatusSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.DocumentRequestStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `DocumentRequestStatusSuccess`   | success                          |
| `DocumentRequestStatusError`     | error                            |
| `DocumentRequestStatusPending`   | pending                          |
| `DocumentRequestStatusCancelled` | cancelled                        |