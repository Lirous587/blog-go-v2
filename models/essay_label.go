package models

type EssayLabelData struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
	EssayCount   int8   `json:"essay_count,omitempty" db:"essay_count"`
}
