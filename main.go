package main

import (
	"fmt"
	"rap/dbConfig"
	"rap/repository"
	"rap/service"
	"rap/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	dbConfig.DB, err = gorm.Open("mysql", dbConfig.DbURL(dbConfig.BuildDBConfig()))
	utils.Check(err)
	defer dbConfig.DB.Close()
	service.GenRoleInsertStatement(35, 36)
	service.GenPermissionInsertStatement(35, 36)
	service.GenAPIInsertStatement(35, 36)
	fmt.Println(service.GetStatement())
	fmt.Println(repository.Permissions)
}
