package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"project/shop/basic"
	basicconfig "project/shop/basic/config"
	"sync"
	"time"
)

const (
	dialect = "mysql"
)

var (
	once sync.Once
	s    *dbService
)

type DataBaseConfig struct {
	MysqlConfig MySqlConfig `json:"mysql"`
}

type MySqlConfig struct {
	Enable            bool   `json:"enabled"`
	Address           string `json:"address"`
	Port              int    `json:"port"`
	User              string `json:"user"`
	Password          string `json:"password"`
	DbName            string `json:"dbName"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

type DBService interface {
	DB() *gorm.DB
}

type dbService struct {
	db *gorm.DB
}

func (rs *dbService) DB() *gorm.DB {
	return rs.db
}

func initDataBase() {
	once.Do(func() {
		log.Info("common: db初始化...")

		config := &DataBaseConfig{}
		err := basicconfig.GetConfigurator().App("db", config)
		if err != nil {
			log.Fatalf("common: db加载配置失败, error: %v", dialect, err)
			return
		}

		if !config.MysqlConfig.Enable {
			log.Info("未启用数据库")
			return
		}

		url := config.MysqlConfig.Address + ":" + fmt.Sprint(config.MysqlConfig.Port)
		dbArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", config.MysqlConfig.User, config.MysqlConfig.Password, url, config.MysqlConfig.DbName)
		db, err := gorm.Open(dialect, dbArgs)
		if err != nil {
			log.Fatalf("common: %s连接失败, error: %v", dialect, err)
			return
		}

		// 最大连接数
		db.DB().SetMaxIdleConns(config.MysqlConfig.MaxIdleConnection)
		// 最大闲置数
		db.DB().SetMaxOpenConns(config.MysqlConfig.MaxOpenConnection)

		// 最大连接时间
		db.DB().SetConnMaxLifetime(time.Duration(int64(time.Second) * 50))
		db.LogMode(true)

		s = &dbService{
			db: db,
		}
	})

	return
}

func init() {
	basic.Register(initDataBase)
}

func GetDB() DBService {
	return s
}
