package mysqld

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init(ctx context.Context) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.Get("mysql.username"),
		viper.Get("mysql.password"),
		viper.Get("mysql.host"),
		viper.Get("mysql.database"))

	dialector := mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})

	if Db, err = gorm.Open(dialector, &gorm.Config{}); err != nil {
		return fmt.Errorf("fail to init mysqld client, err: %+v", err)
	}

	return nil
}
