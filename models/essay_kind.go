package models

type EssayKindData struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Icon         string `json:"icon" db:"icon"`
	Introduction string `json:"introduction" db:"introduction"`
	EssayCount   int    `json:"essay_count" db:"essay_count"`
}

type EssayKindParam struct {
	Name         string `json:"name" binding:"required" db:"name"`
	Icon         string `json:"icon" binding:"required" db:"icon"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
}

type EssayKindUpdateParam struct {
	ID           int    `json:"id" binding:"required" db:"id"`
	Name         string `json:"name" binding:"required" db:"name"`
	Icon         string `json:"icon" binding:"required" db:"icon"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
}
