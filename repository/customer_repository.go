package repository

import (
	"database/sql"
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/utils/common"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	BaseRepositoryPaging[model.Customer]
	GetByPhone(phone string) (model.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Create(payload model.Customer) error {
	_, err := c.db.Exec("INSERT INTO customer (id, name, phone_number, address) VALUES ($1, $2, $3, $4)", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	rows, err := c.db.Query("SELECT id, name, phone_number, address FROM customer")
	if err != nil {
		return nil, err
	}

	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *customerRepository) Get(id string) (model.Customer, error) {
	var customer model.Customer
	row := c.db.QueryRow("SELECT id, name, phone_number, address FROM customer WHERE id = $1", id)
	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) GetByPhone(phone string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow("SELECT id, name, phone_number, address FROM customer WHERE phone_number ILIKE $1", "%"+phone+"%").Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) Update(payload model.Customer) error {
	_, err := c.db.Exec("UPDATE customer SET name = $2, phone_number = $3, address = $4 WHERE id = $1", payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Delete(id string) error {
	_, err := c.db.Exec("DELETE FROM customer WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) Paging(requestPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error) {

	paginationQuery := common.GetPaginationParams(requestPaging)

	rows, err := c.db.Query("SELECT id, name, phone_number, address FROM customer LIMIT $1 OFFSET $2", paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		customers = append(customers, customer)
	}

	// count customer = totalRows
	var totalRows int
	row := c.db.QueryRow("SELECT COUNT(*) FROM customer")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	return customers, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
