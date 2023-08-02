package model

import "time"

type Bill struct {
	Id          string       `json:"id"`
	BillDate    time.Time    `json:"billDate"`
	EntryDate   time.Time    `json:"entryDate"`
	FinishDate  time.Time    `json:"finishDate"`
	EmployeeId  string       `json:"employeeId"`
	CustomerId  string       `json:"customerId"`
	BillDetails []BillDetail `json:"billDetails"`
}

type BillDetail struct {
	Id           string `json:"id"`
	BillId       string `json:"billId"`
	ProductId    string `json:"productId"`
	ProductPrice int    `json:"productPrice"`
	Qty          int    `json:"qty"`
}
