package service

import (
	"fmt"
	"log"

	"github.com/Indra-riswan/basecamp-golang-api/dto"
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"github.com/Indra-riswan/basecamp-golang-api/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type BrandService interface {
	CreateBrand(brand dto.BrandDto) entity.Brand
	UpdateBrand(brandid uint, brand dto.BrandDtoUpdate) entity.Brand
	DeleteBrand(ID uint) entity.Brand
	FindBrand(ID uint) entity.Brand
	AllBrand() []entity.Brand
	TransaksiBrand(ID uint) entity.Brand
	VerifiCredential(nama, passwd string) interface{}
	Isduplicate(nama string) bool
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

func (s *brandservice) UpdateBrand(brandid uint, brand dto.BrandDtoUpdate) entity.Brand {
	brands := s.repository.FindBrand(brandid)
	err := smapping.FillStruct(&brands, smapping.MapFields(&brand))
	if err != nil {
		log.Panic(err.Error())
	}

	s.repository.UpdateBrand(brands)
	return brands
}

func (s *brandservice) DeleteBrand(ID uint) entity.Brand {
	brands := s.repository.FindBrand(ID)
	s.repository.DeleteBrand(brands)
	return brands
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

func ComparePasswd(pass, plainpass []byte) bool {
	psswd := []byte(pass)
	err := bcrypt.CompareHashAndPassword(psswd, plainpass)
	if err != nil {
		fmt.Println("Failed to Compare :", err.Error())
		return false
	}
	return true
}

func (s *brandservice) VerifiCredential(nama, passwd string) interface{} {
	res := s.repository.VerifyCredential(nama, passwd)
	if v, ok := res.(entity.Brand); ok {
		compare := ComparePasswd([]byte(v.Password), []byte(passwd))
		if v.Nama == nama && compare {
			return res
		}
		return false
	}
	return false

}

func (s *brandservice) Isduplicate(nama string) bool {
	res := s.repository.Isduplicate(nama)
	return res.Error != nil

}
