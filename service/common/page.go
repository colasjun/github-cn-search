package common

type PageDataStructure struct {
	Total int `json:"total"`
	TotalPage int `json:"totalSize"`
	CurrentPage int `json:"currentPage"`
	PageSize int `json:"pagesize"`
}