package brightdatasdk

type BrightDataInput struct {
	OriginalURL string `json:"original_url"`
	RequestID   string `json:"request_id"`
}

type GoogleSearchResult struct {
	Input            BrightDataInput         `json:"input"`
	Organic          []GoogleOrganic         `json:"organic"`
	FeaturedSnippets []GoogleFeaturedSnippet `json:"featured_snippets"`
}

type GoogleExtension struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Rank int    `json:"rank"`
}

type GoogleOrganic struct {
	Link        string            `json:"link"`
	DisplayLink string            `json:"display_link"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Extensions  []GoogleExtension `json:"extensions,omitempty"`
	Image       string            `json:"image,omitempty"`
	ImageAlt    string            `json:"image_alt,omitempty"`
	ImageBase64 string            `json:"image_base64,omitempty"`
	Rank        int               `json:"rank"`
	GlobalRank  int               `json:"global_rank"`
}

type GoogleFeaturedSnippetValue struct {
	Text string `json:"text"`
}

type GoogleFeaturedSnippet struct {
	Type        string                     `json:"type"`
	DisplayLink string                     `json:"display_link"`
	LinkTitle   string                     `json:"link_title"`
	Link        string                     `json:"link"`
	Value       GoogleFeaturedSnippetValue `json:"value"`
	Rank        int                        `json:"rank"`
	GlobalRank  int                        `json:"global_rank"`
}
