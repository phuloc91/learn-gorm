package main

import (
	"fmt"
	"rap/dbConfig"
	"rap/repository"
	"rap/service"
)

func main() {
	dbConfig.InitConnection()
	defer dbConfig.DB.Close()
	service.GenRoleInsertStatement(35, 36)
	service.GenPermissionInsertStatement(35, 36)
	service.GenAPIInsertStatement(35, 36)
	fmt.Println(service.GetStatement())
	fmt.Println(repository.Permissions)
}
