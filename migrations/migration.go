package migration
import (
	"gin-gonic-crud/config"
	"gin-gonic-crud/models/model"
)
// DbMigrate ...
func DbMigrate() {
	config.DB.AutoMigrate(&model.User{})
}