package dbConfig

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// DB db connection
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	DBName   string `json:"database"`
	Password string `json:"password"`
}

func InitConnection() {
	db, err := gorm.Open(mysql.Open(DbURL(BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Khong ket noi duoc den DB thi chay lam d gi")
		return
	}
	DB = db
}

func BuildDBConfig() *DBConfig {
	configFile, err := ioutil.ReadFile("db-config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var dbConfig DBConfig
	err = json.Unmarshal(configFile, &dbConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
