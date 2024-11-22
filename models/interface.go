package models

type OneCruder[T any] interface {
	Create(data *T) error
	Read(id int) (*T, error)
	Update(data *T) error
	Delete(id int) error
}

// ListAndPage is a generic struct that contains a list of items and page information.
type ListAndPage[T any] struct {
	List      []T `json:"list"`
	TotalPage int `json:"totalPage"`
}

// ListReader is a generic interface for reading a list of items with query parameters.
type ListReader[T any, Q any] interface {
	ReadList(query Q) (*ListAndPage[T], error)
}
