package web

type CatRequest struct {
	Id         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Status     string `json:"status"`
	Visibility string `json:"visibility"`
	Details    string `json:"details"`
}
