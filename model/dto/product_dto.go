package dto

type ProductRequestDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	UomId string `json:"uomId"`
}
