package web

type CatRequest struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	Url        string `json:"url"`
	Status     string `json:"status"`
	Visibility string `json:"visibility"`
	Details    string `json:"details"`
}
