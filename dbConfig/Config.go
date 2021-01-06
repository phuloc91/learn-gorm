package dbConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
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
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
