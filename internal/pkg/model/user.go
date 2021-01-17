package model

import (
	"fmt"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"role"`
}

func LoginCheck(param LoginParam) (User, error) {
	user := User{
		Name:     param.Name,
		Password: param.Password,
	}

	err := mysqld.Db.First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, fmt.Errorf("password not match")
	}

	return user, err
}

func Register(param RegisterParam) (User, error) {
	user := User{}

	if param.Name == "" || param.Password == "" {
		return user, fmt.Errorf("params should't empty")
	}
	if err := mysqld.Db.Where("name = ?", param.Name).First(&user).Error; err == nil {
		return user, fmt.Errorf("username exist")
	}
	if param.Email != "" {
		if err := mysqld.Db.Where("email = ?", param.Email).First(&user).Error; err != nil {
			return user, fmt.Errorf("email exist")
		}
	}
	if param.Phone != "" {
		if err := mysqld.Db.Where("phone = ?", param.Phone).First(&user).Error; err != nil {
			return user, fmt.Errorf("phone exist")
		}
	}

	user = User{
		Name:     param.Name,
		Email:    param.Email,
		Phone:    param.Phone,
		Password: param.Password,
		Role:     "user",
	}
	err := mysqld.Db.Create(&user).Error

	return user, err
}

func GetAllUsers() ([]User, error) {
	var users []User
	if err := mysqld.Db.Find(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}
