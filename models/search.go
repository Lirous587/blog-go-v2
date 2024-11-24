package models

type SearchParam struct {
	Keyword string `json:"keyword" binging:"required"`
	IfAdd   bool   `json:"if_add"`
}

type EssayIdAndKeyword struct {
	EssayId  int      `json:"essay_id" binging:"required"`
	Keywords []string `json:"keywords"`
}
