package web

type UpdateRequest struct {
	Id   string
	Name string `json:"name" binding:"required"`
}
