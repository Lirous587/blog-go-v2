package models

type HeartWordsData struct {
	Content     string `json:"content" db:"content"`
	Source      string `json:"source" db:"source"`
	Img         `json:"img" binding:"required"`
	ID          int  `json:"id" db:"id"`
	IfCouldType bool `json:"ifCouldType" db:"if_could_type"`
}

type HeartWordsParam struct {
	Content     string `json:"content" binding:"required"  db:"content"`
	Source      string `json:"source" binding:"required"  db:"source"`
	Img         `json:"img" binding:"required"`
	IfCouldType bool `json:"ifCouldType" db:"if_could_type"`
}

type HeartWordsUpdateParam struct {
	Content     string `json:"content" binding:"required"  db:"content"`
	Source      string `json:"source" binding:"required"  db:"source"`
	ID          int    `json:"id" binding:"required" db:"id"`
	Img         `json:"img" binding:"required"`
	IfCouldType bool `json:"ifCouldType" db:"if_could_type"`
}

type HeartWordsQuery struct {
	Page        int  `form:"page" binding:"required"`
	PageSize    int  `form:"pageSize" binding:"required"`
	IfCouldType bool `form:"ifCouldType" db:"if_could_type"`
}

type HeartWordsListAndPage struct {
	HeartWordsList []HeartWordsData `json:"list,omitempty"`
	TotalPages     int              `json:"totalPages,omitempty"`
}
