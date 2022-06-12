package main

import (
	"github.com/Indra-riswan/basecamp-golang-api/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	sql *gorm.DB = config.Connectionsql()
)

func main() {
	defer config.Closeconnectionsql(sql)
	sql.AutoMigrate()

	r := gin.Default()

	r.Run()
}
