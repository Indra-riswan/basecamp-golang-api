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
	db               *gorm.DB = config.Connectionsql()
	transaksirepo             = repository.NewTransaksiRepo(db)
	brandrepo                 = repository.NewBrandRepo(db)
	mitrarepo                 = repository.NewMitraRepo(db)
	mobilrepo                 = repository.NewMobilRepo(db)
	transaksiservice          = service.NewTransaksiService(transaksirepo)
	brandservice              = service.NewBrandService(brandrepo)
	mitraservice              = service.NewMitraService(mitrarepo)
	mobilservice              = service.NewMobilService(mobilrepo)
	transaksihandler          = handler.NewTransaksiHandler(transaksiservice)
	brandhandler              = handler.NewBrandHandler(brandservice)
	mitrahandler              = handler.NewMitraHandler(mitraservice)
	mobilhandler              = handler.NewMobilHandler(mobilservice)
)

func main() {
	defer config.Closeconnectionsql(db)
	db.AutoMigrate(&entity.Mitra{}, &entity.Brand{}, &entity.Transaksi{}, &entity.Mobil{})

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
		brand.POST("/login", brandhandler.Login)
	}

	mitra := r.Group("mitra")
	{
		mitra.POST("/create", mitrahandler.Create)
		mitra.PUT("/update/:id", mitrahandler.Update)
		mitra.DELETE("/delete/:id", mitrahandler.Delete)
		mitra.GET("/all", mitrahandler.All)
		mitra.GET("/id/:id", mitrahandler.Find)
		mitra.GET("/penjualan/:id", mitrahandler.Transaksi)
		mitra.POST("/login", mitrahandler.Login)
	}

	mobil := r.Group("mobil")
	{
		mobil.POST("/create", mobilhandler.Create)
		mobil.PUT("/update/:id", mobilhandler.Update)
		mobil.GET("/allmobil", mobilhandler.AllMobil)
		mobil.DELETE("/delete/:id", mobilhandler.DeleteMobil)
	}

	r.Run(":8181")
}
