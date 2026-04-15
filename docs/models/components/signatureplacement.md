# SignaturePlacement


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ID`                                                        | `string`                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `SignerLabel`                                               | `string`                                                    | :heavy_check_mark:                                          | Label identifying which signer this placement belongs to    |
| `PageNumber`                                                | `int64`                                                     | :heavy_check_mark:                                          | N/A                                                         |
| `X`                                                         | `float64`                                                   | :heavy_check_mark:                                          | X coordinate in PDF points from the left edge of the page   |
| `Y`                                                         | `float64`                                                   | :heavy_check_mark:                                          | Y coordinate in PDF points from the bottom edge of the page |
| `Width`                                                     | `float64`                                                   | :heavy_check_mark:                                          | N/A                                                         |
| `Height`                                                    | `float64`                                                   | :heavy_check_mark:                                          | N/A                                                         |