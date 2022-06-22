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

type mobilhandler struct {
	service service.MobilService
}

func NewMobilHandler(service service.MobilService) *mobilhandler {
	return &mobilhandler{service}
}

func (h *mobilhandler) Create(ctx *gin.Context) {
	var mobildto dto.MobilDto
	err := ctx.ShouldBindJSON(&mobildto)
	if err != nil {
		res := helper.Errorsresponse("Failed To Create", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	res := h.service.Create(mobildto)
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}

func (h *mobilhandler) Update(ctx *gin.Context) {
	var mobildtoupdte dto.MobilDtoUpdate
	err := ctx.ShouldBindJSON(&mobildtoupdte)
	if err != nil {
		res := helper.Errorsresponse("Failed To Update", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.Update(uint(id), mobildtoupdte)
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}

func (h *mobilhandler) AllMobil(ctx *gin.Context) {
	var mobils []entity.Mobil = h.service.AllMobil()
	res := helper.Succesresponse(true, "Ok!", mobils)
	ctx.JSON(200, res)
}

func (h *mobilhandler) DeleteMobil(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.DeleteMobil(uint(id))
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}
