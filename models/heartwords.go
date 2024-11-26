package models

type HeartWordsData struct {
	ID          int    `json:"id" db:"id"`
	Content     string `json:"content" binding:"required"  db:"content"`
	Source      string `json:"source" binding:"required"  db:"source"`
	IfCouldType bool   `json:"if_could_type" binding:"required" db:"if_could_type"`
	Img         `json:"img" binding:"required"`
}

type HeartWordsQuery struct {
	Page        int  `form:"page" binding:"required"`
	PageSize    int  `form:"page_size" binding:"required"`
	IfCouldType bool `form:"if_could_type" db:"if_could_type"`
}

type HeartWordsListAndPage struct {
	HeartWordsList []HeartWordsData `json:"list,omitempty"`
	TotalPages     int              `json:"total_pages,omitempty"`
}
