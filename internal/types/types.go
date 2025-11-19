package types

// WebhookRequest represents the incoming webhook payload
type WebhookRequest struct {
	Service  string                 `json:"service"`  // e.g., "pushover", "telegram", "slack"
	Message  string                 `json:"message"`  // Message content
	Title    string                 `json:"title"`    // Optional title
	Priority int                    `json:"priority"` // Optional priority (for services that support it)
	Extra    map[string]interface{} `json:"extra"`    // Additional service-specific parameters
}
