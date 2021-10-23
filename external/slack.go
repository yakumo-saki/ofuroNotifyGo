package external

type SlackMessage struct {
	Channel    string `json:"channel"`    // "General" "@Username"
	Username   string `json:"username"`   // "Ghost"
	Text       string `json:"text"`       // "this is test message"
	Icon_emoji string `json:"icon_emoji"` // ":ghost:",
}
