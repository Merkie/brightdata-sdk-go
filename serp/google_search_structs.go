package serp

type GoogleSearchResponse struct {
	General struct {
		SearchEngine string  `json:"search_engine"`
		ResultsCnt   int     `json:"results_cnt"`
		SearchTime   float64 `json:"search_time"`
		Language     string  `json:"language"`
		Mobile       bool    `json:"mobile"`
		BasicView    bool    `json:"basic_view"`
		SearchType   string  `json:"search_type"`
		PageTitle    string  `json:"page_title"`
		Timestamp    string  `json:"timestamp"`
	} `json:"general"`
	Input struct {
		OriginalURL string `json:"original_url"`
		RequestID   string `json:"request_id"`
	} `json:"input"`
	Organic   []GoogleSearchResult `json:"organic"`
	Knowledge struct {
		Name              string `json:"name"`
		Subtitle          string `json:"subtitle"`
		Description       string `json:"description"`
		DescriptionSource string `json:"description_source"`
		DescriptionLink   string `json:"description_link"`
		Images            []struct {
			Link        string `json:"link"`
			Image       string `json:"image"`
			ImageAlt    string `json:"image_alt"`
			ImageBase64 string `json:"image_base64"`
		} `json:"images"`
		Facts []struct {
			Key       string `json:"key"`
			KeyLink   string `json:"key_link"`
			Predicate string `json:"predicate"`
			Value     []struct {
				Text string `json:"text"`
				Link string `json:"link"`
			} `json:"value"`
		} `json:"facts"`
		Widgets []struct {
			Type      string `json:"type"`
			Key       string `json:"key"`
			Predicate string `json:"predicate"`
			KeyLink   string `json:"key_link"`
			Title     string `json:"title"`
			Items     []struct {
				Title       string `json:"title"`
				Link        string `json:"link"`
				Image       string `json:"image"`
				ImageAlt    string `json:"image_alt"`
				ImageBase64 string `json:"image_base64"`
				Rank        int    `json:"rank"`
			} `json:"items"`
			Rank       int `json:"rank"`
			GlobalRank int `json:"global_rank"`
		} `json:"widgets"`
	} `json:"knowledge"`
	Recipes struct {
		Title string `json:"title"`
		Items []struct {
			Title       string   `json:"title"`
			Image       string   `json:"image"`
			ImageURL    string   `json:"image_url"`
			Link        string   `json:"link"`
			Rating      float64  `json:"rating"`
			ReviewsCnt  int      `json:"reviews_cnt"`
			Source      string   `json:"source"`
			CookTime    string   `json:"cook_time"`
			Ingredients []string `json:"ingredients"`
			Rank        int      `json:"rank"`
			GlobalRank  int      `json:"global_rank"`
		} `json:"items"`
	} `json:"recipes"`
	HotelsSelection struct {
		Link        string `json:"link"`
		DateFrom    string `json:"date_from"`
		DateTo      string `json:"date_to"`
		Title       string `json:"title"`
		Suggestions []struct {
			Name       string `json:"name"`
			Link       string `json:"link"`
			Rank       int    `json:"rank"`
			GlobalRank int    `json:"global_rank"`
		} `json:"suggestions"`
	} `json:"hotels_selection"`
	SnackPackMap struct {
		Image       string `json:"image"`
		ImageAlt    string `json:"image_alt"`
		ImageBase64 string `json:"image_base64"`
	} `json:"snack_pack_map"`
	TopAds []struct {
		Link         string `json:"link"`
		DisplayLink  string `json:"display_link"`
		ReferralLink string `json:"referral_link"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		Extensions   []struct {
			Type        string `json:"type"`
			Link        string `json:"link"`
			Text        string `json:"text"`
			Description string `json:"description"`
			Extended    bool   `json:"extended"`
		} `json:"extensions"`
		Rank       int `json:"rank"`
		GlobalRank int `json:"global_rank"`
	} `json:"top_ads"`
	Pagination struct {
		CurrentPage  int    `json:"current_page"`
		NextPageLink string `json:"next_page_link"`
		NextPage     int    `json:"next_page"`
	} `json:"pagination"`
	Related []struct {
		Text       string `json:"text"`
		ListGroup  bool   `json:"list_group"`
		Expanded   bool   `json:"expanded,omitempty"`
		MoreLink   string `json:"more_link,omitempty"`
		MoreText   string `json:"more_text,omitempty"`
		Image      string `json:"image,omitempty"`
		ImageURL   string `json:"image_url,omitempty"`
		Rank       int    `json:"rank"`
		GlobalRank int    `json:"global_rank"`
		Link       string `json:"link,omitempty"`
	} `json:"related"`
	HTML string `json:"html"`
}

type GoogleSearchResult struct {
	Link        string `json:"link"`
	DisplayLink string `json:"display_link"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Extensions  []struct {
		Type     string `json:"type"`
		Extended bool   `json:"extended"`
		Text     string `json:"text"`
		Link     string `json:"link"`
		Rank     int    `json:"rank"`
	} `json:"extensions,omitempty"`
	Rank        int    `json:"rank"`
	GlobalRank  int    `json:"global_rank"`
	Image       string `json:"image,omitempty"`
	ImageAlt    string `json:"image_alt,omitempty"`
	ImageBase64 string `json:"image_base64,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}
