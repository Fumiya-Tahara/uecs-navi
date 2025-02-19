package dto

type M304IDRequest int

func NewM304IDRequest(id int) *M304IDRequest {
	request := M304IDRequest(id)
	return &request
}
