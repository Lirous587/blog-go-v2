package models

type IndexData struct {
	KindList       []EssayKindData  `json:"kind_list"`
	LabelList      []EssayLabelData `json:"label_list"`
	EssayList      []EssayData      `json:"essay_list"`
	HeartWordsList []HeartWordsData `json:"heart_words_list"`
}
