package models

type GalleryKindData struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GalleryKindParams struct {
	Name string `json:"name" binding:"required" db:"name"`
}

type GalleryKindUpdateParams struct {
	Name string `json:"name" binding:"required" db:"name"`
	ID   int    `json:"id" binding:"required" db:"id"`
}

type GalleryKindQuery struct {
	Page     int `query:"page" form:"page"`
	PageSize int `query:"page_size" form:"page_size"`
}

type GalleryKindList struct {
	List       []GalleryKindData `json:"list"`
	TotalPages int               `json:"total_pages"`
}
