package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        string    `gorm:"uniqueIndex" json:"id"`
	Email     string    `json:"email" gorm:"unique:not null"  form:"email"`
	Name      string    `json:"name" gorm:"not null"  form:"name"`
	Address   string    `json:"address" gorm:"not null"  form:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return
}

type GetCustomerResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email" form:"email" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
}

type CustomerPaginate struct {
	Customer     []GetCustomerResponse
	PreviousPage int `json:"previous_page_id,omitempty" example:"10"`
	NextPage     int `json:"next_page_id,omitempty" example:"10"`
}

type CustomerUpdate struct {
	ID      string `gorm:"uniqueIndex" json:"id"`
	Email   string `json:"email" gorm:"unique:not null"  form:"email"`
	Name    string `json:"name" gorm:"not null"  form:"name"`
	Address string `json:"address" gorm:"not null"  form:"address"`
}

func (cu *CustomerUpdate) TableName() string {
	return "customers"
}
