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

type transaksihandler struct {
	service service.TransaksiService
}

func NewTransaksiHandler(service service.TransaksiService) *transaksihandler {
	return &transaksihandler{service}
}

func (h *transaksihandler) Create(ctx *gin.Context) {
	var transaksidto dto.TransaksiDto
	err := ctx.ShouldBindJSON(&transaksidto)
	if err != nil {
		res := helper.Errorsresponse("Failed Create Transaction", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	total := transaksidto.Perhari * transaksidto.Day
	totalop := transaksidto.Op * transaksidto.Day
	transaksidto.Total = total
	transaksidto.TotalOp = totalop
	transaksidto.SisaProfit = total - totalop

	res := h.service.CreateTransaksi(transaksidto)
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(http.StatusOK, result)
}

func (h *transaksihandler) Update(ctx *gin.Context) {
	var transaksiUpdate dto.TransaksiDtoUpdate
	err := ctx.ShouldBindJSON(&transaksiUpdate)
	if err != nil {
		res := helper.Errorsresponse("Failed Update Produk", err.Error(), helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, res)
	}

	total := transaksiUpdate.Perhari * transaksiUpdate.Day
	totalop := transaksiUpdate.Op * transaksiUpdate.Day
	transaksiUpdate.SisaProfit = total - totalop
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.UpdateTransaksi(uint(id), transaksiUpdate)
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}

func (h *transaksihandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.Errorsresponse("Failed Send Request", "No Data Found", helper.Emptyobj{})
		ctx.JSON(http.StatusNotFound, res)
	}
	res := h.service.DeleteTransaksi(uint(id))
	result := helper.Succesresponse(true, "Succes Delete Transaction", res)
	ctx.JSON(200, result)

}

func (h *transaksihandler) Find(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	res := h.service.FindTransaksi(uint(id))
	if (res == entity.Transaksi{}) {
		result := helper.Errorsresponse("Failed Send Request", "No Data Found", helper.Emptyobj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, result)
	}
	result := helper.Succesresponse(true, "Ok!", res)
	ctx.JSON(200, result)
}

func (h *transaksihandler) All(ctx *gin.Context) {
	var transaks []entity.Transaksi = h.service.AllTransaksi()
	res := helper.Succesresponse(true, "Ok!", transaks)
	ctx.JSON(200, res)
}
