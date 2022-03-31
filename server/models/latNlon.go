package models

type result struct {
	results []struct{ results_info } `json:"results"`
	status  string                   `json:"status"`
}

type results_info struct {
	address_component []address_component_info `json:"address_component"`
	formated_address  string                   `json:"formated_address"`
	geometry          []geometry_info          `json:"geometry"`
	place_id          string                   `json:"place_id"`
	types             []string                 `json:"types"`
}

type address_component_info struct {
	long_name  string   `json:"long_name"`
	short_name string   `json:"short_name"`
	types      []string `json:"types"`
}

type geometry_info struct {
	bounds        []bounds_info `json:"bounds"`
	location      []cord_info   `json:"location"`
	location_type string        `json:"location_type"`
	viewport      []bounds_info `json:"viewport"`
}

type bounds_info struct {
	northeast []cord_info `json:"northeast"`
	southwest []cord_info `json:"southwest"`
}

type cord_info struct {
	lat float64 `json:"lat"`
	lng float64 `json:"lng"`
}
