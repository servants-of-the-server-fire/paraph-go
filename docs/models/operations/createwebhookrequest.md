# CreateWebhookRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `URL`                                                                | `string`                                                             | :heavy_check_mark:                                                   | URL to receive webhook POST requests                                 |
| `Events`                                                             | [][components.WebhookEvent](../../models/components/webhookevent.md) | :heavy_check_mark:                                                   | Event types to subscribe to                                          |