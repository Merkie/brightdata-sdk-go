package brightdatasdk

type GoogleSearchResponse struct {
	General struct {
		SearchEngine string  `json:"search_engine"`
		ResultsCount int     `json:"results_cnt"`
		SearchTime   float32 `json:"search_time"`
		Language     string  `json:"language"`
		Mobile       bool    `json:"mobile"`
		BasicView    bool    `json:"basic_view"`
		SearchType   string  `json:"search_type"`
		PageTitle    string  `json:"page_title"`
		Timestamp    string  `json:"timestamp"`
	}
	Input struct {
		OriginalURL string `json:"original_url"`
		RequestID   string `json:"request_id"`
	} `json:"input"`
	Organic          []GoogleSearchResult    `json:"organic"`
	Html             string                  `json:"html"`
	FeaturedSnippets []GoogleFeaturedSnippet `json:"featured_snippets"`
}

type GoogleSearchResult struct {
	Link        string  `json:"link"`
	DisplayLink string  `json:"display_link"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Extensions  *[]struct {
		Type string `json:"type"`
		Text string `json:"text"`
		Rank int    `json:"rank"`
	} `json:"extensions,omitempty"`
	Image       *string `json:"image,omitempty"`
	ImageAlt    *string `json:"image_alt,omitempty"`
	ImageBase64 *string `json:"image_base64,omitempty"`
	Rank        int     `json:"rank"`
	GlobalRank  int     `json:"global_rank"`
	Subresults  *[]struct {
		Link        string  `json:"link"`
		DisplayLink string  `json:"display_link"`
		Title       string  `json:"title"`
		Description *string `json:"description,omitempty"`
		Extensions  *[]struct {
			Type string `json:"type"`
			Text string `json:"text"`
			Rank int    `json:"rank"`
		} `json:"extensions,omitempty"`
		Image       *string `json:"image,omitempty"`
		ImageAlt    *string `json:"image_alt,omitempty"`
		ImageBase64 *string `json:"image_base64,omitempty"`
		Rank        int     `json:"rank"`
		GlobalRank  int     `json:"global_rank"`
	} `json:"subresults,omitempty"`
}
type GoogleFeaturedSnippet struct {
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
