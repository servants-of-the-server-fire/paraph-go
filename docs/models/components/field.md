# Field


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ID`                                                                | `string`                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `Name`                                                              | `string`                                                            | :heavy_check_mark:                                                  | Field name as defined in the PDF form                               |
| `Type`                                                              | [components.Type](../../models/components/type.md)                  | :heavy_check_mark:                                                  | AcroForm field type                                                 |
| `Options`                                                           | []`string`                                                          | :heavy_check_mark:                                                  | Available options for ComboBox, ListBox, or RadioButtonGroup fields |