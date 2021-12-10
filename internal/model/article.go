package model

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id,omitempty"`
	Author  string `json:"author"`
	Tags    []Tage `json:"tags"`
}
