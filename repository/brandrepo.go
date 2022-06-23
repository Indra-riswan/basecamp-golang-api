package repository

import (
	"fmt"

	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type BrandRepo interface {
	CreateBrand(brand entity.Brand) entity.Brand
	UpdateBrand(brand entity.Brand) entity.Brand
	Allbrand() []entity.Brand
	DeleteBrand(brand entity.Brand)
	FindBrand(ID uint) entity.Brand
	TransaksiBrand(ID uint) entity.Brand
	VerifyCredential(nama, password string) interface{}
	Isduplicate(nama string) (tx *gorm.DB)
}

type brandrepo struct {
	db *gorm.DB
}

func NewBrandRepo(db *gorm.DB) *brandrepo {
	return &brandrepo{db}
}

func HashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		fmt.Print(err.Error())
	}
	return string(hash)
}

func (r *brandrepo) CreateBrand(brand entity.Brand) entity.Brand {
	brand.Password = HashAndSalt([]byte(brand.Password))
	r.db.Save(&brand)
	return brand
}

func (r *brandrepo) UpdateBrand(brand entity.Brand) entity.Brand {
	if brand.Password != "" {
		brand.Password = HashAndSalt([]byte(brand.Password))
	} else {
		var tempbrand entity.Brand
		r.db.Find(tempbrand, brand.ID)
		brand.Password = tempbrand.Password
	}
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

func (r *brandrepo) VerifyCredential(nama, password string) interface{} {
	var brand entity.Brand
	res := r.db.Where("nama = ?", nama).Take(&brand)
	if res.Error == nil {
		return brand
	}
	return nil
}

func (r *brandrepo) Isduplicate(nama string) (tx *gorm.DB) {
	var brand entity.Brand
	return r.db.Where("nama = ?", nama).Take(&brand)
}
