# WebhookEvent

Event types. `request.created`, `request.success`, `request.cancelled`, `request.error`, and `signer.signed` are dispatched for their respective lifecycle events. `webhook.test` is sent when you use the test endpoint to verify your webhook URL.


## Example Usage

```go
import (
	"github.com/servants-of-the-server-fire/paraph-go/models/components"
)

value := components.WebhookEventRequestCreated

// Open enum: custom values can be created with a direct type cast
custom := components.WebhookEvent("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `WebhookEventRequestCreated`   | request.created                |
| `WebhookEventRequestSuccess`   | request.success                |
| `WebhookEventRequestCancelled` | request.cancelled              |
| `WebhookEventRequestError`     | request.error                  |
| `WebhookEventSignerSigned`     | signer.signed                  |
| `WebhookEventWebhookTest`      | webhook.test                   |