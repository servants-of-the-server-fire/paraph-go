# Signer


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ID`                                                               | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `SignerLabel`                                                      | `string`                                                           | :heavy_check_mark:                                                 | Role label for this signer (e.g. "Employee", "Manager")            |
| `RecipientEmail`                                                   | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `Status`                                                           | [components.SignerStatus](../../models/components/signerstatus.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `ExpiresAt`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                         | :heavy_minus_sign:                                                 | When the signing link expires                                      |
| `SignedAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                         | :heavy_minus_sign:                                                 | When the signer completed signing                                  |