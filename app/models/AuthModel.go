package models

import (
	"fmt"

	"github.com/tryhd/dbo-test/app/config"
	"github.com/tryhd/dbo-test/app/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthModel interface {
	Init()
	Login(req types.LoginRequest) (res types.LoginResponse, err error)
	RegisterAuth(req types.RegisterRequest) (res bool, err error)
}

type authModel struct {
	db *gorm.DB
}

func NewAuthModels() AuthModel {
	return &authModel{}
}

func (m *authModel) Init() {
	m.db = config.SetupDatabaseConnection()
}

func (m *authModel) RegisterAuth(req types.RegisterRequest) (res bool, err error) {

	re := types.Auth{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
		Name:     req.Name,
	}
	tx := m.db.Debug().Create(&re)
	fmt.Println("tx", tx)
	if tx.Error != nil {
		return false, tx.Error
	}
	tx2 := m.db.Debug().Find(&re)
	if tx2.Error != nil {
		return false, tx2.Error
	}
	return true, err
}

func (m *authModel) Login(req types.LoginRequest) (res types.LoginResponse, err error) {
	token, err := m.LoginCheck(req.Username, req.Password)

	if err != nil {
		return res, err
	}
	res = types.LoginResponse{
		Token: token,
	}
	return res, err
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (m *authModel) LoginCheck(username string, password string) (string, error) {

	var err error

	u := types.Auth{}

	err = m.db.Table("auths").Debug().Where("email = ?", username).Or("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := config.GenerateToken(u.ID, u.Name)

	if err != nil {
		return "", err
	}
	return token, nil

}
