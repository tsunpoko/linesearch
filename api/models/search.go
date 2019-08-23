package models

type Search struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Img         string `json:"img"`
	Num         int    `json:"num"`
}
