package types

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	ID        string    `gorm:"uniqueIndex" json:"id"`
	Email     string    `json:"email" gorm:"unique:not null"  form:"email" binding:"required"`
	Password  string    `gorm:"not null" json:"password" form:"password" binding:"required"`
	Username  string    `json:"username" gorm:"unique:not null"  form:"username" binding:"required"`
	Name      string    `json:"name" gorm:"not null"  form:"name" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (auth *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	username := html.EscapeString(strings.TrimSpace(auth.Username))
	tx.Statement.SetColumn("ID", uuid)
	tx.Statement.SetColumn("Password", hashedPassword)
	tx.Statement.SetColumn("Username", username)
	return
}

type RegisterRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token" form:"token"`
}
