package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"novel_spider/config"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	var (
		host     = config.GetConf().Mysql.Host
		port     = config.GetConf().Mysql.Port
		database = config.GetConf().Mysql.Database
		username = config.GetConf().Mysql.Username
		password = config.GetConf().Mysql.Password
		charset  = config.GetConf().Mysql.Charset

		err error
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=utf8mb4_unicode_ci", username, password, host, port, database, charset)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 不记录SQL慢日志
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 表名不加s
		},
	})

	if err != nil {
		fmt.Println("Mysql 连接异常: ")
		panic(err)
	}

	//设置连接池
	db, _ := DB.DB()
	//空闲
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(150)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(300)

	err = db.Ping()
	if err != nil {
		fmt.Println("Mysql 无法Ping通: ")
		panic(err)
	}

	return DB
}
