package models

type EssayLabelData struct {
	Name         string `json:"name" db:"name"`
	Introduction string `json:"introduction,omitempty" db:"introduction"`
	ID           int    `json:"id" db:"id"`
	EssayCount   int8   `json:"essayCount,omitempty" db:"essay_count"`
}

type EssayLabelParam struct {
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
}

type EssayLabelUpdateParam struct {
	ID int `json:"id"  binding:"required" db:"id"`
	EssayLabelParam
}
