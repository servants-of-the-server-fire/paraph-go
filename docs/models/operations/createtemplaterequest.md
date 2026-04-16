# CreateTemplateRequest

Upload a PDF template with named form fields. Provide exactly one of `file` or `file_url`.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Name`                                                             | `string`                                                           | :heavy_check_mark:                                                 | Display name for the template                                      |
| `File`                                                             | [*operations.File](../../models/operations/file.md)                | :heavy_minus_sign:                                                 | PDF file with AcroForm fields (mutually exclusive with `file_url`) |
| `FileURL`                                                          | `*string`                                                          | :heavy_minus_sign:                                                 | URL to fetch the PDF from (mutually exclusive with `file`)         |