package models

type GalleryData struct {
	ImgUrl   string `json:"img_url" db:"img_url"`
	KindName string `json:"kind_name" db:"name"`
	ID       int    `json:"id" db:"id"`
	KindID   int    `json:"kind_id" db:"kind_id"`
}

type GalleryParams struct {
	ImgUrl string `json:"img_url" binding:"required" db:"img_url"`
	KindID int    `json:"kind_id" binding:"required" db:"kind_id"`
}

type GalleryUpdateParams struct {
	ImgUrl string `json:"img_url" binding:"required" db:"img_url"`
	ID     int    `json:"id" binding:"required" db:"id"`
	KindID int    `json:"kind_id" binding:"required" db:"kind_id"`
}

type GalleryQuery struct {
	Page     int `query:"page" form:"page"`
	PageSize int `query:"page_size" form:"page_size"`
	KindID   int `query:"kind_id" form:"kind_id" binding:"required" db:"kind_id"`
}

type GalleryListAndPage struct {
	GalleryList []GalleryData `json:"list,omitempty"`
	TotalPages  int           `json:"total_pages,omitempty"`
}
