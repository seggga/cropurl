package datatype

import "github.com/seggga/cropurl/internal/cropurl/user/datatype"

type ShortURL struct {
	ShortURL   *string        `json:"short_url"`
	LongURL    *string        `json:"long_url"`
	Statistics *int64         `json:"statistics"`
	Owner      *datatype.User `json:"owner"`
}
