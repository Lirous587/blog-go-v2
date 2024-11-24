package models

type LabelData struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name"  db:"name"`
	Introduction string `json:"introduction,omitempty"  db:"introduction"`
	EssayCount   int8   `json:"essay_count" db:"essay_count"`
}

type LabelParams struct {
	Name         string `json:"name" binding:"required" db:"name"`
	Introduction string `json:"introduction" binding:"required" db:"introduction"`
}

type LabelUpdateParams struct {
	LabelParams
	ID int `json:"id" binding:"required" db:"id"`
}
