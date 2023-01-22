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

type mahasiswaController struct {
	mahasiswaService service.MahasiswaService
}

func NewMahasiswaController(mahasiswaService service.MahasiswaService) *mahasiswaController {
	return &mahasiswaController{mahasiswaService}
}

func (c *mahasiswaController) GetMahasiswas(cntx *gin.Context) {
	var mahasiswas, err = c.mahasiswaService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var mahasiswasResponse []responses.MahasiswaResponse

	for _, mahasiswa := range mahasiswas {
		var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)
		mahasiswasResponse = append(mahasiswasResponse, mahasiswaResponse)
	}

	cntx.JSON(http.StatusOK, mahasiswasResponse)
}

func (c *mahasiswaController) GetMahasiswa(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var mahasiswa, err = c.mahasiswaService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)

	cntx.JSON(http.StatusOK, mahasiswaResponse)
}

func (c *mahasiswaController) GetMahasiswaWithRelation(cntx *gin.Context) {
	var mahasiswas, err = c.mahasiswaService.FindAllRelation()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var mahasiswasResponse []responses.MahasiswaWithRelationResponse

	for _, mahasiswa := range mahasiswas {
		var mahasiswaResponse = helper.ConvertToMahasiswaWithRelationResponse(mahasiswa)
		mahasiswasResponse = append(mahasiswasResponse, mahasiswaResponse)
	}

	cntx.JSON(http.StatusOK, mahasiswasResponse)
}

func (c *mahasiswaController) CreateMahasiswa(cntx *gin.Context) {
	var mahasiswaRequest request.CreateMahasiswaRequest

	var err = cntx.ShouldBindJSON(&mahasiswaRequest)
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

	mahasiswa, err := c.mahasiswaService.Create(mahasiswaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)

	cntx.JSON(http.StatusCreated, mahasiswaResponse)
}

func (c *mahasiswaController) UpdateMahasiswa(cntx *gin.Context) {
	var mahasiswaRequest request.UpdateMahasiswaRequest

	var err = cntx.ShouldBindJSON(&mahasiswaRequest)
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

	mahasiswa, err := c.mahasiswaService.Update(id, mahasiswaRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var mahasiswaResponse = helper.ConvertToMahasiswaResponse(mahasiswa)

	cntx.JSON(http.StatusOK, mahasiswaResponse)
}

func (c *mahasiswaController) DeleteMahasiswa(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.mahasiswaService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}