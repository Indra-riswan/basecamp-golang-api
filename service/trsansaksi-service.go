package service

import (
	"log"

	"github.com/Indra-riswan/basecamp-golang-api/dto"
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"github.com/Indra-riswan/basecamp-golang-api/repository"
	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	CreateTransaksi(transaksi dto.TransaksiDto) entity.Transaksi
	UpdateTransaksi(ID uint, transaksi dto.TransaksiDtoUpdate) entity.Transaksi
	DeleteTransaksi(transaksiid uint)
	FindTransaksi(ID uint) entity.Transaksi
	AllTransaksi() []entity.Transaksi
}

type transaksiservice struct {
	repository repository.TransaksiRepo
}

func NewTransaksiService(repository repository.TransaksiRepo) *transaksiservice {
	return &transaksiservice{repository}
}

func (s *transaksiservice) CreateTransaksi(transaksi dto.TransaksiDto) entity.Transaksi {
	var trans = entity.Transaksi{}
	err := smapping.FillStruct(&trans, smapping.MapFields(&transaksi))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.CreateTransaksi(trans)

	return trans
}

func (s *transaksiservice) UpdateTransaksi(ID uint, transaksi dto.TransaksiDtoUpdate) entity.Transaksi {
	transaksiid := s.repository.FindTransaksi(ID)
	err := smapping.FillStruct(&transaksiid, smapping.MapFields(&transaksi))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.UpdateTransaksi(transaksiid)
	return transaksiid

}

func (s *transaksiservice) DeleteTransaksi(transaksiid uint) {
	transaksi := s.repository.FindTransaksi(transaksiid)
	s.repository.DeleteTransaksi(transaksi)
}

func (s *transaksiservice) FindTransaksi(ID uint) entity.Transaksi {
	return s.repository.FindTransaksi(ID)
}

func (s *transaksiservice) AllTransaksi() []entity.Transaksi {
	return s.repository.AllTransaksi()
}
