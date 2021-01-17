package model

import "gorm.io/gorm"

type Conf struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Value string `gorm:"column:value"`
}
