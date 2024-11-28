package models

type EssayData struct {
	Name          string           `json:"name" db:"name"`
	KindName      string           `json:"kind_name,omitempty" db:"kind_name"`
	Introduction  string           `json:"introduction,omitempty" db:"introduction"`
	CreatedTime   string           `json:"created_time" db:"created_time"`
	Content       string           `json:"content,omitempty" db:"content"`
	VisitedTimes  int64            `json:"visited_times,omitempty" db:"visited_times"`
	LabelList     []EssayLabelData `json:"label_list,omitempty"`
	Keywords      []string         `json:"keywords,omitempty"`
	NearEssayList []EssayData      `json:"near_essay_list,omitempty"`
	Img
	ID          int  `json:"id" db:"id"`
	KindID      int  `json:"kind_id,omitempty" db:"kind_id"`
	IfRecommend bool `json:"if_recommend" db:"if_recommend"`
	IfTop       bool `json:"if_top" db:"if_top"`
}

type EssayParams struct {
	Name         string   `json:"name" binding:"required" db:"name"`
	Introduction string   `json:"introduction" binding:"required" db:"introduction"`
	CreatedTime  string   `json:"created_time" db:"created_time"`
	Content      string   `json:"content" binding:"required" db:"content"`
	LabelIds     []int    `json:"label_ids" bind:"required"`
	Keywords     []string `json:"keywords"`
	Img          `binding:"required"`
	KindID       int  `json:"kind_id" binding:"required" db:"kind_id"`
	IfTop        bool `json:"if_top" binging:"required" db:"if_top"`
	IfRecommend  bool `json:"if_recommend"  binging:"required" db:"if_recommend"`
}

type EssayUpdateParams struct {
	EssayParams
	OldLabelIds []int `json:"old_label_ids" binding:"required"`
	ID          int   `json:"id" db:"id"`
}

type EssayQuery struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
	LabelID  int `json:"label_id" form:"label_id"`
	KindID   int `json:"kind_id" form:"kind_id"`
}

type EssayListAndPage struct {
	EssayList  []EssayData `json:"essay_list,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
}

type SearchParam struct {
	Keyword string `json:"keyword" binging:"required"`
	IfAdd   bool   `json:"if_add"`
}

type EssayIdAndKeyword struct {
	Keywords []string `json:"keywords"`
	EssayId  int      `json:"essay_id" binging:"required"`
}
