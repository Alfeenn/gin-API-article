package web

type CatResp struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Category   string `json:"category"`
	Status     string `json:"status"`
	Visibility string `json:"visibility"`
	Details    string `json:"details"`
}
