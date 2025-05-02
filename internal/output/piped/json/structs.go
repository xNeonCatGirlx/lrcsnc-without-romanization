package json

type JSONOutput struct {
	Text     string `json:"text"`
	Title    string `json:"song"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Player   string `json:"player"`
	Position string `json:"position"`
	Duration string `json:"duration"`
}

type WaybarJSONOutput struct {
	Text       string   `json:"text"`
	Alt        string   `json:"alt"`
	Tooltip    string   `json:"tooltip"`
	Class      []string `json:"class"`
	Percentage float64  `json:"percentage"`
}
