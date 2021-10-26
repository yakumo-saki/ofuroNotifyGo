package external

type WebhookApi struct {
	InOut       string `json:"type"`         // bath{inOut}
	DurationSec int    `json:"duration_sec"` // duration_sec
	ClickType   string `json:"clickType"`    // SINGLE DOUBLE LONG
	Message     string `json:"message"`      // not used by API
}
