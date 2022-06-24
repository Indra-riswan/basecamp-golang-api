package handler

import (
	"net/http"
	"strconv"

	"github.com/Indra-riswan/basecamp-golang-api/dto"
	"github.com/Indra-riswan/basecamp-golang-api/entity"
	"github.com/Indra-riswan/basecamp-golang-api/helper"
	"github.com/Indra-riswan/basecamp-golang-api/service"
	"github.com/gin-gonic/gin"
)

type brandhandler struct {
	service service.BrandService
}

func NewBrandHandler(service service.BrandService) *brandhandler {
	return &brandhandler{service}
}

func (h *brandhandler) Login(ctx *gin.Context) {
	var brand dto.BrandDtoLogin
	err := ctx.ShouldBindJSON(&brand)
	if err != nil {
		res := helper.Errorsresponse("Failed Login", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	auth := h.service.VerifiCredential(brand.Nama, brand.Password)
	if v, ok := auth.(entity.Brand); ok {
		res := helper.Succesresponse(true, "Ok!", v)
		ctx.JSON(200, res)
		return
	}

	res := helper.Errorsresponse("Denied", "Please Check Email Or Password", helper.Emptyobj{})
	ctx.AbortWithStatusJSON(http.StatusConflict, res)
}

func (h *brandhandler) Create(ctx *gin.Context) {
	var brand dto.BrandDto
	err := ctx.ShouldBindJSON(&brand)
	if err != nil {
		res := helper.Errorsresponse("Failed Create New Brand", err.Error(), helper.Emptyobj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	if !h.service.Isduplicate(brand.Nama) {
		res := helper.Errorsresponse("Failed to Create", "Brand Already Exist", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	} else {
		res := h.service.CreateBrand(brand)
		result := helper.Succesresponse(true, "Ok!", res)
		ctx.JSON(200, result)
	}

}

func (h *brandhandler) Update(ctx *gin.Context) {
	var brandupdate dto.BrandDtoUpdate
	err := ctx.ShouldBindJSON(&brandupdate)
	if err != nil {
		res := helper.Errorsresponse("Failed Update Brand", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
		return
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.UpdateBrand(uint(id), brandupdate)
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}

func (h *brandhandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.DeleteBrand(uint(id))
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}

func (h *brandhandler) Find(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := h.service.FindBrand(uint(id))
	res := helper.Succesresponse(true, "Ok!", result)
	ctx.JSON(200, res)
}

func (h *brandhandler) All(ctx *gin.Context) {
	var brands []entity.Brand = h.service.AllBrand()
	res := helper.Succesresponse(true, "Ok!", brands)
	ctx.JSON(200, res)
}

func (h *brandhandler) Transaksi(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.TransaksiBrand(uint(id))
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}
