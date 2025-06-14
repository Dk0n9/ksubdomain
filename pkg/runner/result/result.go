package result

type Answer struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Result struct {
	Subdomain string   `json:"subdomain"`
	Answers   []Answer `json:"answers"`
}
