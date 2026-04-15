# Type

AcroForm field type

## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.TypeText

// Open enum: custom values can be created with a direct type cast
custom := components.Type("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `TypeText`       | text             |
| `TypeDate`       | date             |
| `TypeCheckbox`   | checkbox         |
| `TypeCombobox`   | combobox         |
| `TypeListbox`    | listbox          |
| `TypeRadioGroup` | radio_group      |