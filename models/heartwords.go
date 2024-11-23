package models

type HeartWordsData struct {
	ID          int    `json:"id" db:"id"`
	Content     string `json:"content"  db:"content"`
	Source      string `json:"source"  db:"source"`
	IfCouldType bool   `json:"if_could_type" db:"if_could_type"`
	Img         `json:"img" binding:"required"`
}

type HeartWordsQuery struct {
	Page        int
	PageSize    int
	IfCouldType bool
}

type HeartWordsListAndPage struct {
	HeartWordsList []HeartWordsData `json:"list,omitempty"`
	TotalPage      int              `json:"totalPage,omitempty"`
}
