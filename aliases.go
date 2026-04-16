package paraph

// Response type aliases for auto-generated names that include HTTP status codes.
// Use these in your code instead of GetAccount200Response etc.
type (
	AccountInfo         = GetAccount200Response
	WebhookCreated      = CreateWebhook201Response
	WebhookCreatedInner = CreateWebhook201ResponseWebhook
	ResendResult        = ResendSigningLink200Response
)
