package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type subKegiatanController struct {
	subKegiatanService service.SubKegiatanService
}

func NewSubKegiatanController(subKegiatanService service.SubKegiatanService) *subKegiatanController {
	return &subKegiatanController{subKegiatanService}
}

func (c *subKegiatanController) GetSubKegiatans(cntx *gin.Context) {
	var subKegiatans, err = c.subKegiatanService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var subKegiatansResponse []responses.Sub_KegiatanResponse

	for _, subKegiatan := range subKegiatans {
		var subKegiatanResponse = helper.ConvertToSubKegiatanResponse(subKegiatan)
		subKegiatansResponse = append(subKegiatansResponse, subKegiatanResponse)
	}

	cntx.JSON(http.StatusOK, subKegiatansResponse)
}

func (c *subKegiatanController) GetSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var subKegiatan, err = c.subKegiatanService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var subKegiatanResponse = helper.ConvertToSubKegiatanResponse(subKegiatan)

	cntx.JSON(http.StatusOK, subKegiatanResponse)
}

func (c *subKegiatanController) GetSubKegiatanWithRelation(cntx *gin.Context) {
	var subKegiatanRelations, err = c.subKegiatanService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var subKegiatansRelationResponse []responses.Sub_KegiatanWithKegiatanResponse

	for _, subKegiatanRelation := range subKegiatanRelations {
		var subKegiatanRelationResponse = helper.ConvertToSubKegiatanWithKegiatanResponse(subKegiatanRelation)
		subKegiatansRelationResponse = append(subKegiatansRelationResponse, subKegiatanRelationResponse)
	}

	cntx.JSON(http.StatusOK, subKegiatansRelationResponse)
}

func (c *subKegiatanController) CreateSubKegiatan(cntx *gin.Context) {
	var subKegiatanRequest request.CreateSub_KegiatanRequest

	var err = cntx.ShouldBindJSON(&subKegiatanRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	subKegiatan, err := c.subKegiatanService.Create(subKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var subKegiatanResponse = helper.ConvertToSubKegiatanResponse(subKegiatan)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": subKegiatanResponse,
	})
}

func (c *subKegiatanController) UpdateSubKegiatan(cntx *gin.Context) {
	var subKegiatanRequest request.UpdateSub_KegiatanRequest

	var err = cntx.ShouldBindJSON(&subKegiatanRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	subKegiatan, err := c.subKegiatanService.Update(id, subKegiatanRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var subKegiatanResponse = helper.ConvertToSubKegiatanResponse(subKegiatan)

	cntx.JSON(http.StatusOK, gin.H{
		"data": subKegiatanResponse,
	})
}

func (c *subKegiatanController) DeleteSubKegiatan(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.subKegiatanService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}
