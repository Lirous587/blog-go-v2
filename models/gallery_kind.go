package models

type GalleryKindParams struct {
	Name string `json:"name" binding:"required" db:"name"`
}

type GalleryKindUpdateParams struct {
	GalleryKindParams
	ID int `json:"id" binding:"required" db:"id"`
}

type GalleryKind struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type GalleryKindList struct {
	List []GalleryKind `json:"list"`
}
