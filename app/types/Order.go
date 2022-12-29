package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID            string   `gorm:"uniqueIndex" json:"id"`
	CustomerID    string   `json:"customer_id" gorm:"not null;size:191"  form:"customer_id" binding:"required"`
	Pcs           string   `json:"pcs" gorm:"not null"  form:"pcs" binding:"required"`
	NameProduct   string   `json:"name_product" gorm:"not null"  form:"name_product" binding:"required"`
	CustomerOrder Customer `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CustomerID;references:ID"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}

type GetOrderResponse struct {
	ID          string `json:"id"`
	CustomerID  string `json:"customer_id" form:"customer_id" binding:"required"`
	Pcs         string `json:"pcs" form:"pcs" binding:"required"`
	NameProduct string `json:"name_product" form:"name_product" binding:"required"`
}

type OrderPaginate struct {
	Order        []GetOrderResponse
	PreviousPage int `json:"previous_page_id,omitempty" example:"10"`
	NextPage     int `json:"next_page_id,omitempty" example:"10"`
}

type OrderUpdate struct {
	ID          string `gorm:"uniqueIndex" json:"id"`
	CustomerID  string `json:"customer_id" gorm:"not null"  form:"customer_id"`
	Pcs         string `json:"pcs" gorm:"not null"  form:"pcs"`
	NameProduct string `json:"name_product" gorm:"not null"  form:"name_product"`
}

func (cu *OrderUpdate) TableName() string {
	return "Orders"
}
