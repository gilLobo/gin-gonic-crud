package main

import (
	"fmt"
	"gin-gonic-crud/config"
	migration "gin-gonic-crud/migrations"
	"gin-gonic-crud/routes"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	config.DB, err = gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Falha ao conectar no bd: ", err)
	}

	defer config.DB.Close()
	migration.DbMigrate()
	// config.DB.AutoMigrate(&model.User{})

	r := routes.SetupRouter()
	// running
	r.Run()
}
