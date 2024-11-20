package models

type GalleryData struct {
	ID       int    `json:"id" db:"id"`
	Url      string `json:"url" db:"img_url"`
	KindName string `json:"kindName" db:"name"`
	KindID   int    `json:"kindID" db:"kind_id"`
}

type GalleryParams struct {
	Url    string `json:"url" binding:"required" db:"img_url"`
	KindID int    `json:"kindID" binding:"required" db:"kind_id"`
}

type GalleryUpdateParams struct {
	GalleryParams
	ID int `json:"id" binding:"required" db:"id"`
}

type GalleryQuery struct {
	Page     int
	PageSize int
	KindID   int `json:"kindID" db:"kind_id"`
}

type GalleryListAndPage struct {
	GalleryList []GalleryData `json:"list,omitempty"`
	TotalPage   int           `json:"totalPage,omitempty"`
}
