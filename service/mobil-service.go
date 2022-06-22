package service

import (
	"log"

	"github.com/Indra-riswan/basecamp-golang-api/dto"
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"github.com/Indra-riswan/basecamp-golang-api/repository"
	"github.com/mashingan/smapping"
)

type MobilService interface {
	Create(mobildto dto.MobilDto) entity.Mobil
	Update(mobilid uint, mobildto dto.MobilDtoUpdate) entity.Mobil
	DeleteMobil(mobilid uint) entity.Mobil
	FindMobil(mobilid uint) entity.Mobil
	AllMobil() []entity.Mobil
}

type mobilservice struct {
	repository repository.MobilRepo
}

func NewMobilService(repository repository.MobilRepo) *mobilservice {
	return &mobilservice{repository}
}

func (s *mobilservice) Create(mobildto dto.MobilDto) entity.Mobil {
	var mobil = entity.Mobil{}
	err := smapping.FillStruct(&mobil, smapping.MapFields(&mobildto))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.Create(mobil)
	return mobil
}

func (s *mobilservice) Update(mobilid uint, mobildto dto.MobilDtoUpdate) entity.Mobil {
	mobil := s.repository.FindMobil(mobilid)
	err := smapping.FillStruct(&mobil, smapping.MapFields(&mobildto))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.Update(mobil)
	return mobil
}

func (s *mobilservice) DeleteMobil(mobilid uint) entity.Mobil {
	mobil := s.repository.FindMobil(mobilid)
	s.repository.DeleteMobil(mobil)
	return mobil
}

func (s *mobilservice) FindMobil(mobilid uint) entity.Mobil {
	return s.repository.FindMobil(mobilid)
}

func (s *mobilservice) AllMobil() []entity.Mobil {
	return s.repository.AllMobil()
}
