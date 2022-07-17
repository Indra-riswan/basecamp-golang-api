package repository

import (
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"gorm.io/gorm"
)

type MitraRepo interface {
	CreateMitra(mitra entity.Mitra) entity.Mitra
	UpdateMitra(mitra entity.Mitra) entity.Mitra
	DeleteMitra(mitra entity.Mitra)
	AllMitra() []entity.Mitra
	FindMitra(ID uint) entity.Mitra
	TransaksiMitra(ID uint) entity.Mitra
	VerifyMitra(nama, password string) interface{}
	Duplicatemitra(nama string) (tx *gorm.DB)
}

type mitrarepo struct {
	db *gorm.DB
}

func NewMitraRepo(db *gorm.DB) *mitrarepo {
	return &mitrarepo{db}
}

func (r *mitrarepo) CreateMitra(mitra entity.Mitra) entity.Mitra {
	mitra.Password = HashAndSalt([]byte(mitra.Password))
	r.db.Save(&mitra)
	return mitra
}

func (r *mitrarepo) UpdateMitra(mitra entity.Mitra) entity.Mitra {
	r.db.Save(&mitra)
	return mitra
}

func (r *mitrarepo) DeleteMitra(mitra entity.Mitra) {
	r.db.Delete(&mitra)
}

func (r *mitrarepo) AllMitra() []entity.Mitra {
	var mitras []entity.Mitra
	r.db.Find(&mitras)
	return mitras
}

func (r *mitrarepo) FindMitra(ID uint) entity.Mitra {
	var mit entity.Mitra
	r.db.Find(&mit, ID)
	return mit
}

func (r *mitrarepo) TransaksiMitra(ID uint) entity.Mitra {
	var mit entity.Mitra
	r.db.Preload("Penjualan").Find(&mit, ID)
	return mit
}

func (r *mitrarepo) VerifyMitra(nama, password string) interface{} {
	var mitra entity.Mitra
	res := r.db.Where("nama = ?", nama).Take(&mitra)
	if res.Error == nil {
		return mitra
	}
	return nil
}

func (r *mitrarepo) Duplicatemitra(nama string) (tx *gorm.DB) {
	var mitra entity.Mitra
	return r.db.Where("nama = ?", nama).Take(&mitra)
}
