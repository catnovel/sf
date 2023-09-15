package sf

type putReadingTimePayload struct {
	OriginUA  string `json:"originUA"`
	WebViewUA string `json:"webViewUA"`
	Action    string `json:"action"`
	OaID      string `json:"oaID"`
}

type putSignInfoPayload struct {
	SignDate   string `json:"signDate"`
	EntityId   string `json:"entityId"`
	EntityType string `json:"entityType"`
}

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
