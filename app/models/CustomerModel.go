package models

import (
	"fmt"

	"github.com/tryhd/dbo-test/app/config"
	"github.com/tryhd/dbo-test/app/types"
	"gorm.io/gorm"
)

type CustomerModel interface {
	Init()
	RegisterCustomer(req types.Customer) (res types.GetCustomerResponse, err error)
	GetAllCustomer(l, p int, sort string) (res types.CustomerPaginate, err error)
	DetailCustomer(uuid string) (res types.GetCustomerResponse, err error)
	UpdateCustomer(req types.CustomerUpdate) (res types.GetCustomerResponse, err error)
	DeleteCustomer(uuid string) (res types.GetCustomerResponse, err error)
	FindCustomer(search string) (res []types.GetCustomerResponse, err error)
}

type customerModel struct {
	db *gorm.DB
}

func NewCustomerModels() CustomerModel {
	return &customerModel{}
}

func (m *customerModel) Init() {
	m.db = config.SetupDatabaseConnection()
}

func (m *customerModel) RegisterCustomer(req types.Customer) (res types.GetCustomerResponse, err error) {

	res = types.GetCustomerResponse{}

	tx := m.db.Debug().Create(&req)
	fmt.Println("tx", tx)
	if tx.Error != nil {
		return res, tx.Error
	}
	tx2 := m.db.Debug().Find(&req).Scan(&res)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	return res, err
}

func (m *customerModel) GetAllCustomer(l, p int, s string) (res types.CustomerPaginate, err error) {

	var result []types.GetCustomerResponse
	offset := (p - 1) * l
	tx2 := m.db.Debug().Model(&types.Customer{}).Limit(l).Offset(offset).Order(s).Scan(&result)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	for _, v := range result {
		res.Customer = append(res.Customer, types.GetCustomerResponse{
			ID:      v.ID,
			Email:   v.Email,
			Address: v.Address,
			Name:    v.Name,
		})

	}
	if len(res.Customer) != 0 {
		res.NextPage = offset + 2
	}
	if offset > 0 {
		res.PreviousPage = offset
	}

	return res, err
}

func (m *customerModel) DetailCustomer(uuid string) (res types.GetCustomerResponse, err error) {

	res = types.GetCustomerResponse{}
	tx2 := m.db.Debug().Model(&types.Customer{}).Where("id", uuid).Scan(&res)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	return res, err
}

func (m *customerModel) FindCustomer(search string) (res []types.GetCustomerResponse, err error) {

	res = []types.GetCustomerResponse{}
	var result []types.GetCustomerResponse

	tx2 := m.db.Debug().Model(&types.Customer{}).Where("name LIKE ?", search+"%").Or("email LIKE ? ", search+"%").Or("address LIKE ? ", search+"%").Scan(&result)
	if tx2.Error != nil {
		return res, tx2.Error
	}
	for _, v := range result {
		res = append(res, types.GetCustomerResponse{
			ID:      v.ID,
			Email:   v.Email,
			Address: v.Address,
			Name:    v.Name,
		})

	}
	return res, err
}

func (m *customerModel) UpdateCustomer(req types.CustomerUpdate) (res types.GetCustomerResponse, err error) {

	res = types.GetCustomerResponse{}
	tx := m.db.Debug().Model(&types.Customer{}).Where("id", req.ID).Save(&req)
	if tx.Error != nil {
		return res, tx.Error
	}

	tx2 := m.db.Debug().Model(&types.Customer{}).Where("id", req.ID).Scan(&res)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	return res, err
}

func (m *customerModel) DeleteCustomer(uuid string) (res types.GetCustomerResponse, err error) {

	res = types.GetCustomerResponse{}
	req := types.Customer{
		ID: uuid,
	}

	tx := m.db.Debug().Model(&types.Customer{}).Delete(&req)
	if tx.Error != nil {
		return res, tx.Error
	}
	tx2 := m.db.Debug().Model(&types.Customer{}).Where("id", uuid).Scan(&res)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	return res, err
}
