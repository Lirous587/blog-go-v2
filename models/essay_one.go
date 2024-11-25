package models

type EssayContent struct {
	Name          string      `json:"name" db:"name"`
	KindID        int         `json:"kind_id,omitempty" db:"kind_id"`
	KindName      string      `json:"kind_name" db:"kind_name"`
	LabelList     []LabelData `json:"label_list"`
	Id            int         `json:"id" db:"id"`
	Introduction  string      `json:"introduction" db:"introduction"`
	Content       string      `json:"content" db:"content"`
	VisitedTimes  int64       `json:"visited_times" db:"visited_times"`
	CreatedTime   string      `json:"created_time" db:"created_time"`
	Keywords      []string    `json:"keywords,omitempty"`
	NearEssayList []EssayData `json:"near_essay_list,omitempty"`
	Img           `json:"img" binding:"required"`
}

type EssayParams struct {
	ID           int      `json:"id" db:"id"`
	Name         string   `json:"name" binding:"required" db:"name"`
	KindID       int      `json:"kind_id" binding:"required" db:"kind_id"`
	LabelIds     []int    `json:"label_ids" bind:"required"`
	Introduction string   `json:"introduction" binding:"required" db:"introduction"`
	CreatedTime  string   `json:"created_time" db:"created_time"`
	Content      string   `json:"content" binding:"required" db:"content"`
	IfTop        bool     `json:"if_top" binging:"required" db:"if_top"`
	IfRecommend  bool     `json:"if_recommend"  binging:"required" db:"if_recommend"`
	Keywords     []string `json:"keywords"`
	Img          `json:"img" binding:"required"`
}

type EssayUpdateParams struct {
	EssayParams
	OldLabelIds []int `json:"old_label_ids" binding:"required"`
}
