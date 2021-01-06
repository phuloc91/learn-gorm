package main

import (
	"fmt"
	"log"
	"os"
	"rap/dbConfig"
	"rap/repository"
	"rap/service"
	"rap/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	dbConfig.DB, err = gorm.Open("mysql", dbConfig.DbURL(dbConfig.BuildDBConfig()), &gorm.Config{
		Logger: newLogger,
	})
	utils.Check(err)
	defer dbConfig.DB.Close()
	service.GenRoleInsertStatement(35, 36)
	service.GenPermissionInsertStatement(35, 36)
	service.GenAPIInsertStatement(35, 36)
	fmt.Println(service.GetStatement())
	fmt.Println(repository.Permissions)
}
