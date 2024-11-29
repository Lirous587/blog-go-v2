package models

type EssayLabelData struct {
	Name         string `json:"name" db:"name"`
	Introduction string `json:"introduction" db:"introduction"`
	ID           int    `json:"id" db:"id"`
	EssayCount   int8   `json:"essayCount,omitempty" db:"essay_count"`
}

type EssayLabelParam struct {
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
	ID           int    `json:"id" db:"id"`
}

type EssayLabelUpdateParam struct {
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
	ID           int    `json:"id"  binding:"required" db:"id"`
}
