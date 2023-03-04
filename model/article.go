package model

type Article struct {
	Id         string
	Name       string
	Status     string
	Category   string
	Url        string
	Visibility string
	Details    string
}

type CategoryArticle struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Url  string `json:"url"`
}
