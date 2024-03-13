package model

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	PushedAt    string `json:"pushed_at"`
	HtmlUrl     string `json:"html_url"`
	Fork        bool   `json:"fork"`
}
