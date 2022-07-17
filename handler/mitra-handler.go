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

type mitrahandler struct {
	service service.MitraService
}

func NewMitraHandler(service service.MitraService) *mitrahandler {
	return &mitrahandler{service}
}

func (h *mitrahandler) Login(ctx *gin.Context) {
	var Mitra dto.MitraDtoLogin
	err := ctx.ShouldBindJSON(&Mitra)
	if err != nil {
		res := helper.Errorsresponse("Failed Login", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	auth := h.service.VerifyMitra(Mitra.Nama, Mitra.Password)
	if v, ok := auth.(entity.Mitra); ok {
		res := helper.Succesresponse(true, "Ok!", v)
		ctx.JSON(200, res)
	} else {
		res := helper.Errorsresponse("Failed Login", "Please check Email Or Password", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	}
}

func (h *mitrahandler) Create(ctx *gin.Context) {
	var mitra dto.MitraDto
	err := ctx.ShouldBindJSON(&mitra)
	if err != nil {
		res := helper.Errorsresponse("Failed Create Mitra", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	if !h.service.Duplicatemitra(mitra.Nama) {
		res := helper.Errorsresponse("Failed to Create", "Brand Already Exist", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	} else {
		res := h.service.CreateMitra(mitra)
		result := helper.Succesresponse(true, "Ok!", res)
		ctx.JSON(http.StatusCreated, result)
	}
}

func (h *mitrahandler) Update(ctx *gin.Context) {
	var mitraupdate dto.MitraDtoUpdate
	err := ctx.ShouldBindJSON(&mitraupdate)
	if err != nil {
		res := helper.Errorsresponse("Failed Update Mitra", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.UpdateMitra(uint(id), mitraupdate)
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(http.StatusOK, result)
}

func (h *mitrahandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.DeleteMitra(uint(id))
	result := helper.Succesresponse(true, "Succesed Delete!", res)
	ctx.JSON(200, result)
}

func (h *mitrahandler) Find(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.FindMitra(uint(id))
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}

func (h *mitrahandler) All(ctx *gin.Context) {
	var mitras []entity.Mitra = h.service.AllMitra()
	res := helper.Succesresponse(true, "Ok!", mitras)
	ctx.JSON(200, res)
}

func (h *mitrahandler) Transaksi(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.TransaksiMitra(uint(id))
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}
