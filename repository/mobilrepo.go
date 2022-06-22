package repository

import (
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"gorm.io/gorm"
)

type MobilRepo interface {
	Create(mobil entity.Mobil) entity.Mobil
	Update(mobil entity.Mobil) entity.Mobil
	DeleteMobil(mobil entity.Mobil)
	AllMobil() []entity.Mobil
	FindMobil(mobilid uint) entity.Mobil
}

type mobilrepo struct {
	DB *gorm.DB
}

func NewMobilRepo(DB *gorm.DB) *mobilrepo {
	return &mobilrepo{DB}
}

func (r *mobilrepo) Create(mobil entity.Mobil) entity.Mobil {
	r.DB.Save(&mobil)
	return mobil
}

func (r *mobilrepo) Update(mobil entity.Mobil) entity.Mobil {
	r.DB.Save(&mobil)
	return mobil
}

func (r *mobilrepo) DeleteMobil(mobil entity.Mobil) {
	r.DB.Delete(&mobil)
}

func (r *mobilrepo) AllMobil() []entity.Mobil {
	var mobils []entity.Mobil
	r.DB.Preload("Mitra").Find(&mobils)
	return mobils
}

func (r *mobilrepo) FindMobil(mobilid uint) entity.Mobil {
	var mobil entity.Mobil
	r.DB.Preload("Mitra").Find(&mobil, mobilid)
	return mobil
}
