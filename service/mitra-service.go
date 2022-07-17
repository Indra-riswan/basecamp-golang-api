package service

import (
	"log"

	"github.com/Indra-riswan/basecamp-golang-api/dto"
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"github.com/Indra-riswan/basecamp-golang-api/repository"
	"github.com/mashingan/smapping"
)

type MitraService interface {
	CreateMitra(mitra dto.MitraDto) entity.Mitra
	UpdateMitra(mitraid uint, mitra dto.MitraDtoUpdate) entity.Mitra
	DeleteMitra(mitraid uint) entity.Mitra
	FindMitra(mitraid uint) entity.Mitra
	AllMitra() []entity.Mitra
	TransaksiMitra(mitraid uint) entity.Mitra
	VerifyMitra(nama, password string) interface{}
	Duplicatemitra(nama string) bool
}

type mitraservice struct {
	repository repository.MitraRepo
}

func NewMitraService(repository repository.MitraRepo) *mitraservice {
	return &mitraservice{repository}
}

func (s *mitraservice) CreateMitra(mitra dto.MitraDto) entity.Mitra {
	var mit = entity.Mitra{}
	err := smapping.FillStruct(&mit, smapping.MapFields(&mitra))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.CreateMitra(mit)
	return mit
}

func (s *mitraservice) UpdateMitra(mitraid uint, mitra dto.MitraDtoUpdate) entity.Mitra {
	mit := s.repository.FindMitra(mitraid)
	err := smapping.FillStruct(&mit, smapping.MapFields(&mitra))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.UpdateMitra(mit)
	return mit
}

func (s *mitraservice) DeleteMitra(mitraid uint) entity.Mitra {
	mitra := s.repository.FindMitra(mitraid)
	s.repository.DeleteMitra(mitra)
	return mitra
}

func (s *mitraservice) FindMitra(mitraid uint) entity.Mitra {
	return s.repository.FindMitra(mitraid)
}

func (s *mitraservice) AllMitra() []entity.Mitra {
	return s.repository.AllMitra()
}

func (s *mitraservice) TransaksiMitra(mitraid uint) entity.Mitra {
	mit := s.repository.TransaksiMitra(mitraid)
	return mit
}

func (s *mitraservice) VerifyMitra(nama, password string) interface{} {
	res := s.repository.VerifyMitra(nama, password)
	if v, ok := res.(entity.Mitra); ok {
		compare := ComparePasswd([]byte(v.Password), []byte(password))
		if v.Nama == nama && compare {
			return res
		}
		return false
	}
	return false
}

func (s *mitraservice) Duplicatemitra(nama string) bool {
	res := s.repository.Duplicatemitra(nama)
	return res.Error != nil
}
