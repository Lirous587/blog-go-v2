package models

type EssayData struct {
	EssayParams
	Img           `json:"img"`
	KindName      string           `json:"kindName,omitempty" db:"kind_name"`
	LabelList     []EssayLabelData `json:"labelList,omitempty"`
	NearEssayList []EssayData      `json:"nearEssayList,omitempty"`
	VisitedTimes  int64            `json:"visitedTimes,omitempty" db:"visited_times"`
	ID            int              `json:"id" db:"id"`
}

type EssayParams struct {
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
	CreatedTime  string `json:"createdTime" db:"created_time"`
	Content      string `json:"content" binding:"required" db:"content"`
	Keywords     string `json:"keywords" binding:"required" db:"keywords"`
	LabelIds     []int  `json:"labelIds" bind:"required"`
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
	PageSize int `json:"pageSize" form:"pageSize"`
	LabelID  int `json:"labelID" form:"labelID"`
	KindID   int `json:"kindID" form:"kindID"`
}

type EssayListAndPage struct {
	EssayList  []EssayData `json:"list,omitempty"`
	TotalPages int         `json:"totalPages,omitempty"`
}

type SearchParam struct {
	Keyword string `json:"keyword" binging:"required"`
	IfAdd   bool   `json:"ifAdd"`
}

type EssayDesc struct {
	Name         string           `json:"name" db:"name"`
	KindName     string           `json:"kindName,omitempty" db:"kind_name"`
	Introduction string           `json:"introduction,omitempty" db:"introduction"`
	CreatedTime  string           `json:"createdTime" db:"created_time"`
	VisitedTimes int64            `json:"visitedTimes,omitempty" db:"visited_times"`
	Keywords     string           `json:"keywords"`
	LabelList    []EssayLabelData `json:"labelList,omitempty"`
	Img          `json:"img"`
	ID           int  `json:"id" db:"id"`
	KindID       int  `json:"kindID,omitempty" db:"kind_id"`
	IfRecommend  bool `json:"ifRecommend" db:"if_recommend"`
	IfTop        bool `json:"ifTop" db:"if_top"`
}
