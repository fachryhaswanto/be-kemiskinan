package controller

import (
	"kemiskinan/helper"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type keluargaController struct {
	keluargaService service.KeluargaService
}

func NewKeluargaController(keluargaService service.KeluargaService) *keluargaController {
	return &keluargaController{keluargaService}
}

func (c *keluargaController) GetKeluargas(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")

	var keluargas, err = c.keluargaService.FindAll(kabupatenKotaId)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var keluargasResponse []responses.KeluargaResponse

	for _, keluarga := range keluargas {
		var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)
		keluargasResponse = append(keluargasResponse, keluargaResponse)
	}

	cntx.JSON(http.StatusOK, keluargasResponse)
}

func (c *keluargaController) GetKeluargaById(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var keluarga, err = c.keluargaService.FindById(kabupatenKotaId, id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

func (c *keluargaController) GetIdKeluargaByKabupatenKota(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenKotaId")
	var idKeluarga = cntx.Param("idkeluarga")

	var keluarga, err = c.keluargaService.FindByIdKeluargaByKabupatenKota(kabupatenKotaId, idKeluarga)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	var keluargaResponse = helper.ConvertToKeluargaResponse(keluarga)

	cntx.JSON(http.StatusOK, keluargaResponse)
}

func (c *keluargaController) CountPenerimaByKabupatenKota(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")
	var penerimaParameter = cntx.Param("penerimaparameter")
	var nilai = cntx.Param("nilai")

	var jumlah, err = c.keluargaService.CountPenerimaByKabupatenKota(kabupatenKotaId, penerimaParameter, nilai)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountPenerimaByKecamatan(cntx *gin.Context) {
	var kecamatanId = cntx.Param("kecamatanid")
	var penerimaParameter = cntx.Param("penerimaparameter")
	var nilai = cntx.Param("nilai")

	var jumlah, err = c.keluargaService.CountPenerimaByKecamatan(kecamatanId, penerimaParameter, nilai)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)

}

func (c *keluargaController) CountPenerimaByKelurahan(cntx *gin.Context) {
	var kelurahanId = cntx.Param("kelurahanid")
	var penerimaParameter = cntx.Param("penerimaparameter")
	var nilai = cntx.Param("nilai")

	var jumlah, err = c.keluargaService.CountPenerimaByKelurahan(kelurahanId, penerimaParameter, nilai)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountDesilByKabupatenKota(cntx *gin.Context) {
	var kabupatenKotaId = cntx.Param("kabupatenkotaid")
	var nilaiDesil = cntx.Param("nilaidesil")

	var jumlah, err = c.keluargaService.CountDesilByKabupatenKota(kabupatenKotaId, nilaiDesil)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountDesilByKecamatan(cntx *gin.Context) {
	var kecamatanId = cntx.Param("kecamatanid")
	var nilaiDesil = cntx.Param("nilaidesil")

	var jumlah, err = c.keluargaService.CountDesilByKecamatan(kecamatanId, nilaiDesil)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}

func (c *keluargaController) CountDesilByKelurahan(cntx *gin.Context) {
	var kelurahanId = cntx.Param("kelurahanid")
	var nilaiDesil = cntx.Param("nilaidesil")

	var jumlah, err = c.keluargaService.CountDesilByKelurahan(kelurahanId, nilaiDesil)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var jumlahResponse = helper.ConvertToJumlahResponse(jumlah)

	cntx.JSON(http.StatusOK, jumlahResponse)
}