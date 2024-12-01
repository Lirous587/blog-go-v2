package models

type EssayKindData struct {
	Name         string `json:"name" db:"name"`
	Icon         string `json:"icon" db:"icon"`
	Introduction string `json:"introduction" db:"introduction"`
	ID           int    `json:"id" db:"id"`
	EssayCount   int    `json:"essayCount" db:"essay_count"`
}

type EssayKindParam struct {
	Name         string `json:"name" binding:"required" db:"name"`
	Icon         string `json:"icon" binding:"required" db:"icon"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
}

type EssayKindUpdateParam struct {
	ID int `json:"id" binding:"required" db:"id"`
	EssayKindParam
}
