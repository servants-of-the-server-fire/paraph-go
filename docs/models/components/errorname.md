# ErrorName

## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.ErrorNameUnauthorized

// Open enum: custom values can be created with a direct type cast
custom := components.ErrorName("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ErrorNameUnauthorized`  | unauthorized             |
| `ErrorNameNotFound`      | not_found                |
| `ErrorNameBadRequest`    | bad_request              |
| `ErrorNameExceededQuota` | exceeded_quota           |
| `ErrorNameExceededRate`  | exceeded_rate            |
| `ErrorNameUnknown`       | unknown                  |
| `ErrorNameMergeFailed`   | merge_failed             |
| `ErrorNameInvalidStatus` | invalid_status           |