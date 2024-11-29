package models

type EssayData struct {
	Name          string           `json:"name" db:"name"`
	KindName      string           `json:"kindName,omitempty" db:"kind_name"`
	Introduction  string           `json:"introduction,omitempty" db:"introduction"`
	CreatedTime   string           `json:"createdTime" db:"created_time"`
	Content       string           `json:"content,omitempty" db:"content"`
	VisitedTimes  int64            `json:"visitedTimes,omitempty" db:"visited_times"`
	LabelList     []EssayLabelData `json:"labelList,omitempty"`
	Keywords      []string         `json:"keywords,omitempty"`
	NearEssayList []EssayData      `json:"nearEssayList,omitempty"`
	Img           `json:"img"`
	ID            int  `json:"id" db:"id"`
	KindID        int  `json:"kindID,omitempty" db:"kind_id"`
	IfRecommend   bool `json:"ifRecommend" db:"if_recommend"`
	IfTop         bool `json:"ifTop" db:"if_top"`
}

type EssayParams struct {
	Name         string   `json:"name" binding:"required" db:"name"`
	Introduction string   `json:"introduction" binding:"required" db:"introduction"`
	CreatedTime  string   `json:"createdTime" db:"created_time"`
	Content      string   `json:"content" binding:"required" db:"content"`
	LabelIds     []int    `json:"labelIds" bind:"required"`
	Keywords     []string `json:"keywords"`
	Img          `json:"img" binding:"required"`
	KindID       int  `json:"kindID" binding:"required" db:"kind_id"`
	IfTop        bool `json:"ifTop" binging:"required" db:"if_top"`
	IfRecommend  bool `json:"ifRecommend"  binging:"required" db:"if_recommend"`
}

type EssayUpdateParams struct {
	EssayParams
	OldLabelIds []int `json:"oldLabelIds" binding:"required"`
	ID          int   `json:"id" db:"id"`
}

type EssayQuery struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"page_size"`
	LabelID  int `json:"labelID" form:"label_id"`
	KindID   int `json:"kindID" form:"kind_id"`
}

type EssayListAndPage struct {
	EssayList  []EssayData `json:"essayList,omitempty"`
	TotalPages int         `json:"totalPages,omitempty"`
}

type SearchParam struct {
	Keyword string `json:"keyword" binging:"required"`
	IfAdd   bool   `json:"ifAdd"`
}

type EssayIdAndKeyword struct {
	Keywords []string `json:"keywords"`
	EssayId  int      `json:"id" binging:"required"`
}
