package dto

type AddDepartmentDTO struct {
	Title string `json:"title"`
}

type IdDerartmentDTO struct {
	Id int64 `json:"id"`
}

type RenameDerartmentDTO struct {
	Id       int64  `json:"id"`
	NewTitle string `json:"new_title"`
}
