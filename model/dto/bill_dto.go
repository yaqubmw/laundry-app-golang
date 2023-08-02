package dto

import (
	"enigma-laundry-apps/model"
	"time"
)

type BillResponseDto struct {
	Id          string                  `json:"id"`
	BillDate    time.Time               `json:"billDate"`
	EntryDate   time.Time               `json:"entryDate"`
	FinishDate  time.Time               `json:"finishDate"`
	Employee    model.Employee          `json:"employee"`
	Customer    model.Customer          `json:"customer"`
	BillDetails []BillDetailResponseDto `json:"billDetails"`
	TotalBill   int                     `json:"totalBill"`
}

type BillDetailResponseDto struct {
	Id           string        `json:"id"`
	BillId       string        `json:"billId"`
	Product      model.Product `json:"product"`
	ProductPrice int           `json:"productPrice"`
	Qty          int           `json:"qty"`
}

