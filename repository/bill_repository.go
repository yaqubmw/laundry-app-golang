package repository

import (
	"database/sql"
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
)

type BillRepository interface {
	Create(payload model.Bill) error
	Get(id string) (dto.BillResponseDto, error)
	List(requestPaging dto.PaginationParam) ([]dto.BillResponseDto, dto.Paging, error)
}

type billRepository struct {
	db *sql.DB
}

// RegisterNewBill implements BillRepository.
func (b *billRepository) Create(payload model.Bill) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}
	// insert bill
	_, err = tx.Exec("INSERT INTO bill (id, bill_date, entry_date, finish_date, employee_id, customer_id) VALUES ($1, $2, $3, $4, $5, $6)", payload.Id, payload.BillDate, payload.EntryDate, payload.FinishDate, payload.EmployeeId, payload.CustomerId)

	if err != nil {
		return err
	}
	// insert bill detail
	for _, item := range payload.BillDetails {
		_, err = tx.Exec("INSERT INTO bill_detail (id, bill_id, product_id, product_price, qty) VALUES ($1, $2, $3, $4, $5)", item.Id, item.BillId, item.ProductId, item.ProductPrice, item.Qty)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// Get implements BillRepository.
func (b *billRepository) Get(id string) (dto.BillResponseDto, error) {
	var billResponseDto dto.BillResponseDto
	sqlBill := `SELECT b.id as bill_id, b.bill_date, b.entry_date, b.finish_date, c.id as customer_id, c.name as customer_name, c.phone_number as customer_phone, c.address as customer_address, e.id as employee_id, e.name as employee_name, e.phone_number as employee_phone, e.address as employee_address
	FROM bill b 
	JOIN customer c ON c.id = b.customer_id 
	JOIN employee e ON e.id = b.employee_id
	WHERE b.id = $1`

	err := b.db.QueryRow(sqlBill, id).Scan(&billResponseDto.Id, &billResponseDto.BillDate, &billResponseDto.EntryDate, &billResponseDto.FinishDate, &billResponseDto.Customer.Id, &billResponseDto.Customer.Name, &billResponseDto.Customer.PhoneNumber, &billResponseDto.Customer.Address, &billResponseDto.Employee.Id, &billResponseDto.Employee.Name, &billResponseDto.Employee.PhoneNumber, &billResponseDto.Employee.Address)
	if err != nil {
		return dto.BillResponseDto{}, err
	}

	sqlBillDetail := `SELECT b.id as bill_id, p.id as product_id, p.name as product_name, p.price, u.id as uom_id, u.name as uom_name, bd.id as bill_detail_id, bd.product_price, bd.qty
	FROM bill b 
	JOIN bill_detail bd ON bd.bill_id = b.id 
	JOIN product p on p.id = bd.product_id
	JOIN uom u ON u.id = p.uom_id
	WHERE b.id = $1`

	rows, err := b.db.Query(sqlBillDetail, id)
	if err != nil {
		return dto.BillResponseDto{}, err
	}

	for rows.Next() {
		var billDetailResponseDto dto.BillDetailResponseDto
		err := rows.Scan(&billDetailResponseDto.BillId, &billDetailResponseDto.Product.Id, &billDetailResponseDto.Product.Name, &billDetailResponseDto.Product.Price, &billDetailResponseDto.Product.Uom.Id, &billDetailResponseDto.Product.Uom.Name, &billDetailResponseDto.Id, &billDetailResponseDto.ProductPrice, &billDetailResponseDto.Qty)
		if err != nil {
			return dto.BillResponseDto{}, err
		}

		billResponseDto.BillDetails = append(billResponseDto.BillDetails, billDetailResponseDto)
	}

	return billResponseDto, nil
}

// Paging implements BillRepository.
func (b *billRepository) List(requestPaging dto.PaginationParam) ([]dto.BillResponseDto, dto.Paging, error) {
	return nil, dto.Paging{}, nil
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{db: db}
}
