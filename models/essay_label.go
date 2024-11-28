package models

type EssayLabelData struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Introduction string `json:"introduction" db:"introduction"`
	EssayCount   int8   `json:"essay_count,omitempty" db:"essay_count"`
}

type EssayLabelParam struct {
	ID           int    `json:"id"  binding:"required" db:"id"`
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
}

type EssayLabelUpdateParam struct {
	ID           int    `json:"id"  binding:"required" db:"id"`
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
}
