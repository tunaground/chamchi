package util

type Pagination struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}
