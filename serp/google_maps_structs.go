package serp

type GoogleMapsResponse struct {
	General struct {
		Query       string `json:"query"`
		Language    string `json:"language"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		Gl          string `json:"gl"`
	} `json:"general"`
	Organic []GoogleMapsResult `json:"organic"`
	HTML    string             `json:"html"`
}

type GoogleMapsResult struct {
	Title       string `json:"title"`
	DisplayLink string `json:"display_link"`
	Link        string `json:"link"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Category    []struct {
		Title string `json:"title"`
		ID    string `json:"id,omitempty"`
	} `json:"category"`
	Tags []struct {
		GroupID    string      `json:"group_id"`
		GroupTitle string      `json:"group_title"`
		KeyID      interface{} `json:"key_id"`
		ValueTitle string      `json:"value_title"`
	} `json:"tags"`
	Summary       string  `json:"summary"`
	Description   string  `json:"description"`
	Rating        float64 `json:"rating"`
	ReviewsCnt    int     `json:"reviews_cnt"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Claimed       bool    `json:"claimed"`
	Fid           string  `json:"fid"`
	MapIDEncoded  string  `json:"map_id_encoded"`
	MapID         string  `json:"map_id"`
	MapLink       string  `json:"map_link"`
	OriginalImage string  `json:"original_image"`
	Image         string  `json:"image"`
	Thumbnail     string  `json:"thumbnail"`
	Icon          string  `json:"icon"`
	ImageURL      string  `json:"image_url"`
	Rank          int     `json:"rank"`
}
