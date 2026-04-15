# Plan

## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.PlanFree

// Open enum: custom values can be created with a direct type cast
custom := components.Plan("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `PlanFree`       | free             |
| `PlanPro`        | pro              |
| `PlanEnterprise` | enterprise       |