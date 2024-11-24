package models

type GalleryKindData struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GalleryKindList struct {
	List       []GalleryKindData `json:"list"`
	TotalPages int               `json:"total_pages"`
}
