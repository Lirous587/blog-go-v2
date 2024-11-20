package models

type Img struct {
	ID  int    `json:"id" db:"img_id"`
	Url string `json:"url" db:"img_url"`
}
