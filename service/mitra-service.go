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
	return s.repository.FindMitra(mitraid)

}

func (s *mitraservice) FindMitra(mitraid uint) entity.Mitra {
	return s.repository.FindMitra(mitraid)
}

func (s *mitraservice) AllMitra() []entity.Mitra {
	return s.repository.AllMitra()
}

func (s *mitraservice) TransaksiMitra(mitraid uint) entity.Mitra {
	mit := s.repository.FindMitra(mitraid)
	return mit
}
