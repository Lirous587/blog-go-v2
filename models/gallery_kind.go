package models

type GalleryKindData struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GalleryKindParams struct {
	Name string `json:"name" binding:"required" db:"name"`
}

type GalleryKindUpdateParams struct {
	ID int `json:"id" binding:"required" db:"id"`
	GalleryKindParams
}

type GalleryKindQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type GalleryKindList struct {
	List       []GalleryKindData `json:"list"`
	TotalPages int               `json:"totalPages"`
}
