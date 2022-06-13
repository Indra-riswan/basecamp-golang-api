package repository

import (
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"gorm.io/gorm"
)

type BrandRepo interface {
	CreateBrand(brand entity.Brand) entity.Brand
	UpdateBrand(brand entity.Brand) entity.Brand
	Allbrand() []entity.Brand
	DeleteBrand(brand entity.Brand)
	FindBrand(ID uint) entity.Brand
	TransaksiBrand(ID uint) entity.Brand
}

type brandrepo struct {
	db *gorm.DB
}

func NewBrandRepo(db *gorm.DB) *brandrepo {
	return &brandrepo{db}
}

func (r *brandrepo) CreateBrand(brand entity.Brand) entity.Brand {
	r.db.Save(&brand)
	return brand
}

func (r *brandrepo) UpdateBrand(brand entity.Brand) entity.Brand {
	r.db.Save(&brand)
	return brand
}

func (r *brandrepo) Allbrand() []entity.Brand {
	var brands []entity.Brand
	r.db.Find(&brands)
	return brands
}

func (r *brandrepo) DeleteBrand(brand entity.Brand) {
	r.db.Delete(&brand)
}

func (r *brandrepo) FindBrand(ID uint) entity.Brand {
	var brandid entity.Brand
	r.db.Find(&brandid, ID)
	return brandid
}

func (r *brandrepo) TransaksiBrand(ID uint) entity.Brand {
	var brand entity.Brand
	r.db.Preload("Penjualan").Find(&brand, ID)
	return brand
}
