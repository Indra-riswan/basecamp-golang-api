package repository

import (
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"gorm.io/gorm"
)

type TransaksiRepo interface {
	CreateTransaksi(transaksi entity.Transaksi) entity.Transaksi
	UpdateTransaksi(transaksi entity.Transaksi) entity.Transaksi
	DeleteTransaksi(transaksi entity.Transaksi) entity.Transaksi
	FindTransaksi(ID uint) entity.Transaksi
	AllTransaksi() []entity.Transaksi
}

type transaksirepo struct {
	db *gorm.DB
}

func NewTransaksiRepo(db *gorm.DB) *transaksirepo {
	return &transaksirepo{db}
}

func (r *transaksirepo) CreateTransaksi(transaksi entity.Transaksi) entity.Transaksi {
	r.db.Save(&transaksi)
	return transaksi
}

func (r *transaksirepo) UpdateTransaksi(transaksi entity.Transaksi) entity.Transaksi {
	r.db.Save(&transaksi)
	return transaksi
}

func (r *transaksirepo) DeleteTransaksi(transaksi entity.Transaksi) entity.Transaksi {
	r.db.Delete(&transaksi)
	return transaksi
}

func (r *transaksirepo) FindTransaksi(ID uint) entity.Transaksi {
	var transaksiid entity.Transaksi
	r.db.Find(&transaksiid, ID)
	return transaksiid
}

func (r *transaksirepo) AllTransaksi() []entity.Transaksi {
	var transaksis []entity.Transaksi
	r.db.Find(&transaksis)
	return transaksis
}
