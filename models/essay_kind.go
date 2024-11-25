package models

type EssayKindData struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" binding:"required" db:"name"`
	Icon         string `json:"icon" binding:"required" db:"icon"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
	EssayCount   int    `json:"essay_count" db:"essay_count"`
}
