package models

import (
	"fmt"

	"github.com/tryhd/dbo-test/app/config"
	"github.com/tryhd/dbo-test/app/types"
	"gorm.io/gorm"
)

type OrderModel interface {
	Init()
	RegisterOrder(req types.Order) (res types.GetOrderResponse, err error)
	GetAllOrder(l, p int, sort string) (res types.OrderPaginate, err error)
	DetailOrder(uuid string) (res types.GetOrderResponse, err error)
	UpdateOrder(req types.OrderUpdate) (res types.GetOrderResponse, err error)
	DeleteOrder(uuid string) (res types.GetOrderResponse, err error)
	FindOrder(search string) (res []types.GetOrderResponse, err error)
}

type orderModel struct {
	db *gorm.DB
}

func NewOrderModels() OrderModel {
	return &orderModel{}
}

func (m *orderModel) Init() {
	m.db = config.SetupDatabaseConnection()
}

func (m *orderModel) RegisterOrder(req types.Order) (res types.GetOrderResponse, err error) {

	res = types.GetOrderResponse{}

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

func (m *orderModel) GetAllOrder(l, p int, s string) (res types.OrderPaginate, err error) {

	var result []types.GetOrderResponse
	offset := (p - 1) * l
	tx2 := m.db.Debug().Model(&types.Order{}).Limit(l).Offset(offset).Order(s).Scan(&result)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	for _, v := range result {
		res.Order = append(res.Order, types.GetOrderResponse{
			ID:          v.ID,
			CustomerID:  v.CustomerID,
			Pcs:         v.Pcs,
			NameProduct: v.NameProduct,
		})

	}
	if len(res.Order) != 0 {
		res.NextPage = offset + 2
	}
	if offset > 0 {
		res.PreviousPage = offset
	}

	return res, err
}

func (m *orderModel) DetailOrder(uuid string) (res types.GetOrderResponse, err error) {

	res = types.GetOrderResponse{}
	tx2 := m.db.Debug().Model(&types.Order{}).Where("id", uuid).Scan(&res)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	return res, err
}

func (m *orderModel) FindOrder(search string) (res []types.GetOrderResponse, err error) {

	res = []types.GetOrderResponse{}
	var result []types.GetOrderResponse

	tx2 := m.db.Debug().Model(&types.Order{}).Joins("JOIN customers on customers.id = orders.customer_id").Where("name_product LIKE ?", search+"%").Or("customers.name LIKE ? ", search+"%").Scan(&result)
	if tx2.Error != nil {
		return res, tx2.Error
	}
	for _, v := range result {
		res = append(res, types.GetOrderResponse{
			ID:          v.ID,
			CustomerID:  v.CustomerID,
			Pcs:         v.Pcs,
			NameProduct: v.NameProduct,
		})

	}
	return res, err
}

func (m *orderModel) UpdateOrder(req types.OrderUpdate) (res types.GetOrderResponse, err error) {

	res = types.GetOrderResponse{}
	tx := m.db.Debug().Model(&types.Order{}).Where("id", req.ID).Updates(&req)
	if tx.Error != nil {
		return res, tx.Error
	}

	tx2 := m.db.Debug().Model(&types.Order{}).Where("id", req.ID).Preload("Customer").Scan(&res)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	return res, err
}

func (m *orderModel) DeleteOrder(uuid string) (res types.GetOrderResponse, err error) {

	res = types.GetOrderResponse{}
	req := types.Order{
		ID: uuid,
	}

	tx := m.db.Debug().Model(&types.Order{}).Delete(&req)
	if tx.Error != nil {
		return res, tx.Error
	}
	tx2 := m.db.Debug().Model(&types.Order{}).Where("id", uuid).Scan(&res)
	if tx2.Error != nil {
		return res, tx2.Error
	}

	return res, err
}
