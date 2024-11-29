package models

type GalleryData struct {
	ImgUrl   string `json:"imgUrl" db:"img_url"`
	KindName string `json:"kindName" db:"name"`
	ID       int    `json:"id" db:"id"`
	KindID   int    `json:"kindID" db:"kind_id"`
}

type GalleryParams struct {
	ImgUrl string `json:"imgUrl" binding:"required" db:"img_url"`
	KindID int    `json:"kindID" binding:"required" db:"kind_id"`
}

type GalleryUpdateParams struct {
	ImgUrl string `json:"imgUrl" binding:"required" db:"img_url"`
	ID     int    `json:"id" binding:"required" db:"id"`
	KindID int    `json:"kindID" binding:"required" db:"kind_id"`
}

type GalleryQuery struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
	KindID   int `form:"kindID" binding:"required" db:"kind_id"`
}

type GalleryListAndPage struct {
	GalleryList []GalleryData `json:"list,omitempty"`
	TotalPages  int           `json:"totalPages,omitempty"`
}
