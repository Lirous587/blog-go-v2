package models

type HeartWordsData struct {
	ID          int    `json:"id" db:"id"`
	Content     string `json:"content"  db:"content"`
	Source      string `json:"source"  db:"source"`
	IfCouldType bool   `json:"ifCouldType" db:"if_could_type"`
	Img         `json:"img" binding:"required"`
}

type HeartWordsParams struct {
	Content     string `json:"content"  db:"content"`
	Source      string `json:"source"  db:"source"`
	IfCouldType bool   `json:"ifCouldType" db:"if_could_type"` //如果为false binding:"required" 会导致绑定失败 所以不要加
	Img         `json:"img" binding:"required"`
}

type HeartWordsUpdateParams struct {
	HeartWordsParams
	ID int `json:"id" binding:"required" db:"id"`
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
