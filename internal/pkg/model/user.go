package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `gorm:"column:name" json:"name"`
	Email    string     `gorm:"column:email" json:"email"`
	Phone    string     `gorm:"column:phone" json:"phone"`
	Password string     `gorm:"column:password" json:"-"`
	Role     string     `gorm:"column:role" json:"role"`
	Config   UserConfig `gorm:"column:config;type:string" json:"config"`
}

type UserConfig struct {
	ProblemNum map[string]int `json:"problem_num"`
}

// 怎么从数据库取
func (uc *UserConfig) Scan(value interface{}) error {
	buf, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal user config, value: %+v", value)
	}

	return json.Unmarshal(buf, uc)
}

// 怎么存在数据库中
func (uc UserConfig) Value() (driver.Value, error) {
	buf, err := json.Marshal(uc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user config: %+v", uc)
	}

	return string(buf), nil
}

func LoginCheck(param LoginParam) (User, error) {
	user := User{
		Name:     param.Name,
		Password: param.Password,
	}

	err := mysqld.Db.Where(&user).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, fmt.Errorf("user not exist or password wrong")
	}

	return user, err
}

func Register(param RegisterParam) (User, error) {
	user := User{}

	if param.Name == "" || param.Password == "" {
		return user, fmt.Errorf("params should't empty")
	}
	if err := mysqld.Db.Where("name = ?", param.Name).First(&user).Error; err == nil {
		return user, fmt.Errorf("username already exist")
	}
	if param.Email != "" {
		if err := mysqld.Db.Where("email = ?", param.Email).First(&user).Error; err != nil {
			return user, fmt.Errorf("email already exist")
		}
	}
	if param.Phone != "" {
		if err := mysqld.Db.Where("phone = ?", param.Phone).First(&user).Error; err != nil {
			return user, fmt.Errorf("phone already exist")
		}
	}

	user = User{
		Name:     param.Name,
		Email:    param.Email,
		Phone:    param.Phone,
		Password: param.Password,
		Role:     "user",
		Config:   UserConfig{ProblemNum: map[string]int{"algorithm": 3}},
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
