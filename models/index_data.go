package models

type IndexData struct {
	KindList       []EssayKindData  `json:"kindList"`
	LabelList      []EssayLabelData `json:"labelList"`
	EssayList      []EssayData      `json:"essayList"`
	HeartWordsList []HeartWordsData `json:"heartWordsList"`
}
