package service

import (
	"log"

	"github.com/Indra-riswan/basecamp-golang-api/dto"
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"github.com/Indra-riswan/basecamp-golang-api/repository"
	"github.com/mashingan/smapping"
)

type BrandService interface {
	CreateBrand(brand dto.BrandDto) entity.Brand
	UpdateBrand(brandid uint, brand dto.BrandDto) entity.Brand
	DeleteBrand(ID uint)
	FindBrand(ID uint) entity.Brand
	AllBrand() []entity.Brand
	TransaksiBrand(ID uint) entity.Brand
}

type brandservice struct {
	repository repository.BrandRepo
}

func NewBrandService(repository repository.BrandRepo) *brandservice {
	return &brandservice{repository}
}

func (s *brandservice) CreateBrand(brand dto.BrandDto) entity.Brand {
	var brands = entity.Brand{}
	err := smapping.FillStruct(&brands, smapping.MapFields(&brand))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.CreateBrand(brands)
	return brands
}

func (s *brandservice) UpdateBrand(brandid uint, brand dto.BrandDto) entity.Brand {
	brands := s.repository.FindBrand(brandid)
	err := smapping.FillStruct(&brands, smapping.MapFields(&brand))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.UpdateBrand(brands)
	return brands
}

func (s *brandservice) DeleteBrand(ID uint) {
	brands := s.repository.FindBrand(ID)
	s.repository.DeleteBrand(brands)
}

func (s *brandservice) FindBrand(ID uint) entity.Brand {
	return s.repository.FindBrand(ID)
}

func (s *brandservice) AllBrand() []entity.Brand {
	return s.repository.Allbrand()
}

func (s *brandservice) TransaksiBrand(ID uint) entity.Brand {
	return s.repository.TransaksiBrand(ID)
}
