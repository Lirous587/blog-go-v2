package models

type EssayData struct {
	ID           int              `json:"id" db:"id"`
	Name         string           `json:"name" db:"name"`
	LabelList    []EssayLabelData `json:"label_list,omitempty"`
	KindName     string           `json:"kind_name,omitempty" db:"kind_name"`
	KindID       int              `json:"kind_id,omitempty" db:"kind_id"`
	Introduction string           `json:"introduction,omitempty" db:"introduction"`
	CreatedTime  string           `json:"created_time" db:"created_time"`
	VisitedTimes int64            `json:"visited_times,omitempty" db:"visited_times"`
	Content      string           `json:"content,omitempty" db:"content"`
	Keywords     []string         `json:"keywords,omitempty"`
	IfRecommend  bool             `json:"if_recommend" db:"if_recommend"`
	IfTop        bool             `json:"if_top" db:"if_top"`
	Img          `json:"img" binding:"required"`
}

type EssayListAndPage struct {
	EssayList  []EssayData `json:"essay_list,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
}

type EssayQuery struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	LabelID  int `json:"label_id"`
	KindID   int `json:"kind_id"`
}
