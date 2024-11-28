package models

type Img struct {
	Url string `json:"url" db:"img_url"`
	ID  int    `json:"id" db:"img_id"`
}
