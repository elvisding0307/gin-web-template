package dao

import (
	"errors"
	"gin-web-template/internal/config"
	"gin-web-template/internal/model"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	MYSQL_INSTANCE         *gorm.DB
	onceMysqlInitilization sync.Once
)

// InitMysqlDB 初始化数据库连接。
func InitMysqlDB() error {
	var err error
	var cfg *config.Config
	onceMysqlInitilization.Do(func() {
		if cfg, err = config.ServerConfig(); err != nil {
			log.Println("无法获取服务器配置: ", err)
			return 
		}
		dsn := cfg.GetMysqlDSN()
		db, err := connect(dsn)
		if err != nil {
			log.Println("Mysql数据库连接失败: ", err)
			return
		}

		log.Println("Mysql数据库连接成功")
		MYSQL_INSTANCE = db
		err = autoMigrateModels(db)
		if err != nil {
			log.Println("数据库模型迁移失败: ", err)
			return 
		}
	})

	return err
}

// connect 建立数据库连接
func connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 可选：开启 GORM 日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数（可按需调整）
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

// autoMigrateModels 自动迁移所有模型
func autoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&model.Hello{}, &model.User{})
}

// GetMysqlInstance 获取数据库实例
func GetMysqlInstance() (*gorm.DB, error) {
	if MYSQL_INSTANCE == nil {
		log.Println("Mysql实例未初始化")
		return nil, errors.New("Mysql实例未初始化")
	}
	return MYSQL_INSTANCE, nil
}
