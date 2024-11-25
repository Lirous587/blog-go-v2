package models

type GalleryData struct {
	ID       int    `json:"id" db:"id"`
	ImgUrl   string `json:"img_url" db:"img_url"`
	KindName string `json:"kind_name" db:"name"`
	KindID   int    `json:"kind_id" db:"kind_id"`
}

type GalleryQuery struct {
	Page     int `query:"page" binding:"required"`
	PageSize int `query:"page_size" binding:"required"`
	KindID   int `query:"kind_id" binding:"required" db:"kind_id"`
}

type GalleryListAndPage struct {
	GalleryList []GalleryData `json:"list,omitempty"`
	TotalPages  int           `json:"total_pages,omitempty"`
}
