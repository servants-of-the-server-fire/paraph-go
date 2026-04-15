# DocumentRequestSummaryStatus

## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.DocumentRequestSummaryStatusSuccess

// Open enum: custom values can be created with a direct type cast
custom := components.DocumentRequestSummaryStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `DocumentRequestSummaryStatusSuccess`   | success                                 |
| `DocumentRequestSummaryStatusError`     | error                                   |
| `DocumentRequestSummaryStatusPending`   | pending                                 |
| `DocumentRequestSummaryStatusCancelled` | cancelled                               |