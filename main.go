package main

import (
	"github.com/Indra-riswan/basecamp-golang-api/config"
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"github.com/Indra-riswan/basecamp-golang-api/handler"
	"github.com/Indra-riswan/basecamp-golang-api/repository"
	"github.com/Indra-riswan/basecamp-golang-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	sql              *gorm.DB = config.Connectionsql()
	transaksirepo             = repository.NewTransaksiRepo(sql)
	brandrepo                 = repository.NewBrandRepo(sql)
	mitrarepo                 = repository.NewMitraRepo(sql)
	transaksiservice          = service.NewTransaksiService(transaksirepo)
	brandservice              = service.NewBrandService(brandrepo)
	mitraservice              = service.NewMitraService(mitrarepo)
	transaksihandler          = handler.NewTransaksiHandler(transaksiservice)
	brandhandler              = handler.NewBrandHandler(brandservice)
	mitrahandler              = handler.NewMitraHandler(mitraservice)
)

func main() {
	defer config.Closeconnectionsql(sql)
	sql.AutoMigrate(&entity.Transaksi{}, &entity.Brand{}, &entity.Mitra{})

	r := gin.Default()

	transaksi := r.Group("transaksi")
	{
		transaksi.POST("/create", transaksihandler.Create)
		transaksi.PUT("/update/:id", transaksihandler.Update)
		transaksi.DELETE("/delete/:id", transaksihandler.Delete)
		transaksi.GET("/all", transaksihandler.All)
		transaksi.GET("/id/:id", transaksihandler.Find)
	}

	brand := r.Group("brand")
	{
		brand.POST("/create", brandhandler.Create)
		brand.GET("/all", brandhandler.All)
		brand.PUT("/update/:id", brandhandler.Update)
		brand.DELETE("/delete/:id", brandhandler.Delete)
		brand.GET("/id/:id", brandhandler.Find)
		brand.GET("/penjualan/:id", brandhandler.Transaksi)
	}

	mitra := r.Group("mitra")
	{
		mitra.POST("/create", mitrahandler.Create)
		mitra.PUT("/update/:id", mitrahandler.Update)
		mitra.DELETE("/delete/:id", mitrahandler.Delete)
		mitra.GET("/all", mitrahandler.All)
		mitra.GET("/id/:id", mitrahandler.Find)
		mitra.GET("/penjualan/:id", mitrahandler.Transaksi)
	}

	r.Run()
}
