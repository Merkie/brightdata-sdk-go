package brightdatasdk

type GoogleSearchResult struct {
	Input struct {
		OriginalURL string `json:"original_url"`
		RequestID   string `json:"request_id"`
	} `json:"input"`
	Organic          []googleOrganic         `json:"organic"`
	FeaturedSnippets []googleFeaturedSnippet `json:"featured_snippets"`
}

type googleExtension struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Rank int    `json:"rank"`
}

type googleOrganic struct {
	Link        string            `json:"link"`
	DisplayLink string            `json:"display_link"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Extensions  []googleExtension `json:"extensions,omitempty"`
	Image       string            `json:"image,omitempty"`
	ImageAlt    string            `json:"image_alt,omitempty"`
	ImageBase64 string            `json:"image_base64,omitempty"`
	Rank        int               `json:"rank"`
	GlobalRank  int               `json:"global_rank"`
}
type googleFeaturedSnippet struct {
	Type        string `json:"type"`
	DisplayLink string `json:"display_link"`
	LinkTitle   string `json:"link_title"`
	Link        string `json:"link"`
	Value       struct {
		Text string `json:"text"`
	} `json:"value"`
	Rank       int `json:"rank"`
	GlobalRank int `json:"global_rank"`
}
