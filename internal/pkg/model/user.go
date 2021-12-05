package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string     `gorm:"column:name" json:"name"`
	Email        string     `gorm:"column:email" json:"email"`
	Phone        string     `gorm:"column:phone" json:"phone"`
	Password     string     `gorm:"column:password" json:"-"`
	Role         string     `gorm:"column:role" json:"role"`
	WxNick       string     `gorm:"column:wx_nick" json:"wx_nick"`
	PersistDay   int        `gorm:"column:persist_day" json:"persist_day"`
	InterruptDay int        `gorm:"column:interrupt_day" json:"interrupt_day"`
	Config       UserConfig `gorm:"column:config;type:string" json:"config"`
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

func (u *User) UpdateComplete(complete bool) error {

	// 如果是新注册的用户，本次是给新用户出今天的题，不用更新坚持或中断了几天
	if u.CreatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02") {
		return nil
	}

	if complete {
		u.PersistDay++
	} else {
		u.InterruptDay++
	}

	return mysqld.Db.Save(u).Error
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
		Name:         param.Name,
		Email:        param.Email,
		Phone:        param.Phone,
		Password:     param.Password,
		Role:         "user",
		PersistDay:   0,
		InterruptDay: 0,
		Config:       UserConfig{ProblemNum: map[string]int{"algorithm": 3}},
	}
	err := mysqld.Db.Create(&user).Error

	return user, err
}

func GetUserByWxNick(nick string) User {
	user := User{}
	if err := mysqld.Db.Where("wx_nick = ?", nick).First(&user).Error; err != nil {
		log.Infof("no user mapping to wx_nick = %s", nick)
		return User{}
	}
	return user
}

func GetAllUsers() ([]User, error) {
	var users []User
	if err := mysqld.Db.Find(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}
