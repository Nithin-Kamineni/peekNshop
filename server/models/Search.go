package models

type Stores struct {
	html_attributions []string                   `json:"html_attributions"`
	results           []struct{ nearest_stores } `json:"results"`
	status            string                     `json:"status"`
}

type nearest_stores struct {
	business_status       string                     `json:"business_status"`
	geometry              struct{ precise_location } `json:"geometry"`
	icon                  string                     `json:"icon"`
	icon_background_color string                     `json:"icon_background_color"`
	icon_mask_base_uri    string                     `json:"icon_mask_base_uri"`
	name                  string                     `json:"name"`
	opening_hours         struct{ open_status }      `json:"name"`
	photos                []struct{ photo_detail }   `json:"photos"`
	place_id              string                     `json:"place_id"`
	plus_code             struct{ other_codes }      `json:"plus_code"`
	price_level           int                        `json:"price_level"`
	rating                float32                    `json:"rating"`
	reference             string                     `json:"reference"`
	scope                 string                     `json:"scope"`
	types                 []string                   `json:"types"`
	user_ratings_total    int                        `json:"user_ratings_total"`
	vicinity              string                     `json:"vicinity"`
}

type precise_location struct {
	location struct{ latlong }    `json:"location"`
	viewport struct{ directions } `json:"viewport"`
}

type latlong struct {
	lat float32 `json:"lat"`
	lng float32 `json:"lng"`
}

type directions struct {
	northeast struct{ latlong } `json:"northeast"`
	southwest struct{ latlong } `json:"southwest"`
}

type open_status struct {
	open_now bool `json:"open_now"`
}

type photo_detail struct {
	height            int      `json:"height"`
	html_attributions []string `json:"html_attributions"`
	photo_reference   string   `json:"string"`
	width             int      `json:"width"`
}

type other_codes struct {
	compound_code string `json:"compound_code"`
	global_code   string `json:"global_code"`
}
