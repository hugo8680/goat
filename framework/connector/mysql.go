package connector

import (
	"forum-service/framework/config"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func ConnectToMySQL() {
	var err error
	conf := config.GetSetting()
	dsn := conf.DB.Username + ":" + conf.DB.Password + "@tcp(" + conf.DB.Host + ":" + strconv.Itoa(conf.DB.Port) + ")/" + conf.DB.Database + "?charset=" + conf.DB.Charset + "&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(log.Default(), logger.Config{
			LogLevel:                  logger.Error, // 打印错误日志
			IgnoreRecordNotFoundError: true,
		}),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(conf.DB.MaxOpenConn)
	sqlDB.SetMaxIdleConns(conf.DB.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
