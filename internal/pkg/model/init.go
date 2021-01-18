package model

import (
	"context"
	"fmt"
	"github.com/chaosi-zju/daily-problem/internal/pkg/mysqld"
	"github.com/spf13/viper"
)

func Init(ctx context.Context) error {
	mysqld.Db.AutoMigrate(User{})
	mysqld.Db.AutoMigrate(Problem{})
	mysqld.Db.AutoMigrate(UserProblem{})
	return initMultiFunc(ctx, initJWTSignKey)
}

func initMultiFunc(_ context.Context, fs ...func() error) error {
	for _, f := range fs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func initJWTSignKey() error {
	conf := Conf{Name: "JWT-SignKey"}
	if err := mysqld.Db.Where(conf).First(&conf).Error; err != nil {
		return fmt.Errorf("jwt signkey not found")
	} else {
		viper.Set("JWT-SignKey", conf.Value)
		return nil
	}
}
