package web

type CatRequest struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Visibility string `json:"visibility"`
	Details    string `json:"details"`
}
