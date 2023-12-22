package domain

import "encoding/json"

type Department struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func (d Department) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}
