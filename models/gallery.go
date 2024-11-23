package models

type GalleryData struct {
	ID       int    `json:"id" db:"id"`
	ImgUrl   string `json:"img_url" db:"img_url"`
	KindName string `json:"k_name" db:"name"`
	KindID   int    `json:"k_id" db:"kind_id"`
}

type GalleryQuery struct {
	Page     int
	PageSize int
	KindID   int `json:"k_id" db:"kind_id"`
}

type GalleryListAndPage struct {
	GalleryList []GalleryData `json:"list,omitempty"`
	TotalPage   int           `json:"totalPage,omitempty"`
}
