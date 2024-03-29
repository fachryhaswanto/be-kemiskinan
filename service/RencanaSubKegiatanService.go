package service

import (
	"errors"
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type RencanaSubKegiatanService interface {
	FindAll() ([]model.RencanaSubKegiatan, error)
	FindById(id int) (model.RencanaSubKegiatan, error)
	FindBySearch(whereClase map[string]interface{}) ([]model.RencanaSubKegiatan, error)
	SumAllPaguRencanaSubKegiatan(tahun, tipe string) int64
	SumPaguRencanaSubKegiatanByInstansi(tahun, tipe string, instansis []model.Instansi) []int64
	CountJumlahRencanaSubKegiatan(tahun string, tipe string) (int64, error)
	CountJumlahRencanaSubKegiatanByInstansi(tahun string, tipe string, instansis []model.Instansi) []int64
	Create(rencanaSubKegiatanRequest request.CreateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error)
	Update(id int, rencanaSubKegiatanRequest request.UpdateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error)
	Delete(id int) (model.RencanaSubKegiatan, error)
}

type rencanaSubKegiatanService struct {
	rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository
	rencanaKegiatanRepository    repository.RencanaKegiatanRepository
	fokusBelanjaRepository       repository.FokusBelanjaRepository
}

func NewRencanaSubKegiatanService(rencanaSubKegiatanRepository repository.RencanaSubKegiatanRepository, rencanaKegiatanRepository repository.RencanaKegiatanRepository, fokusBelanjaRepository repository.FokusBelanjaRepository) *rencanaSubKegiatanService {
	return &rencanaSubKegiatanService{rencanaSubKegiatanRepository, rencanaKegiatanRepository, fokusBelanjaRepository}
}

func (s *rencanaSubKegiatanService) FindAll() ([]model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatans, err = s.rencanaSubKegiatanRepository.FindAll()

	return rencanaSubKegiatans, err
}

func (s *rencanaSubKegiatanService) FindById(id int) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan, err = s.rencanaSubKegiatanRepository.FindById(id)

	return rencanaSubKegiatan, err
}

func (s *rencanaSubKegiatanService) FindBySearch(whereClause map[string]interface{}) ([]model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatans, err = s.rencanaSubKegiatanRepository.FindBySearch(whereClause)

	return rencanaSubKegiatans, err
}

func (s *rencanaSubKegiatanService) SumAllPaguRencanaSubKegiatan(tahun, tipe string) int64 {
	var totalPaguRencanaSubKegiatan = s.rencanaSubKegiatanRepository.SumAllPaguRencanaSubKegiatan(tahun, tipe)

	return totalPaguRencanaSubKegiatan
}

func (s *rencanaSubKegiatanService) SumPaguRencanaSubKegiatanByInstansi(tahun, tipe string, instansis []model.Instansi) []int64 {
	var totalPaguRencanaSubKegiatan = s.rencanaSubKegiatanRepository.SumPaguRencanaSubKegiatanByInstansi(tahun, tipe, instansis)

	return totalPaguRencanaSubKegiatan
}

func (s *rencanaSubKegiatanService) CountJumlahRencanaSubKegiatan(tahun string, tipe string) (int64, error) {
	var jumlahRencanaSubKegiatan, err = s.rencanaSubKegiatanRepository.CountJumlahRencanaSubKegiatan(tahun, tipe)

	return jumlahRencanaSubKegiatan, err
}

func (s *rencanaSubKegiatanService) CountJumlahRencanaSubKegiatanByInstansi(tahun string, tipe string, instansis []model.Instansi) []int64 {
	var jumlahRencanaSubKegiatan = s.rencanaSubKegiatanRepository.CountJumlahRencanaSubKegiatanByInstansi(tahun, tipe, instansis)

	return jumlahRencanaSubKegiatan
}

func (s *rencanaSubKegiatanService) Create(rencanaSubKegiatanRequest request.CreateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan = model.RencanaSubKegiatan{
		RencanaKegiatanId: rencanaSubKegiatanRequest.RencanaKegiatanId,
		SubKegiatanId:     rencanaSubKegiatanRequest.SubKegiatanId,
		PaguSubKegiatan:   rencanaSubKegiatanRequest.PaguSubKegiatan,
		Tipe:              rencanaSubKegiatanRequest.Tipe,
		Tahun:             rencanaSubKegiatanRequest.Tahun,
	}

	var rencanaKegiatan, _ = s.rencanaKegiatanRepository.FindById(rencanaSubKegiatan.RencanaKegiatanId)

	var totalPaguRencanaSubKegiatan, _ = s.rencanaSubKegiatanRepository.SumPaguRencanaSubKegiatan(rencanaSubKegiatan.RencanaKegiatanId)
	totalPaguRencanaSubKegiatan = totalPaguRencanaSubKegiatan + rencanaSubKegiatan.PaguSubKegiatan

	if rencanaKegiatan.PaguKegiatan >= totalPaguRencanaSubKegiatan {
		newRencanaSubKegiatan, err := s.rencanaSubKegiatanRepository.Create(rencanaSubKegiatan)
		return newRencanaSubKegiatan, err
	} else {
		return rencanaSubKegiatan, errors.New("pagu sub kegiatan melebihi pagu kegiatan")
	}
}

func (s *rencanaSubKegiatanService) Update(id int, rencanaSubKegiatanRequest request.UpdateRencanaSubKegiatanRequest) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan, _ = s.rencanaSubKegiatanRepository.FindById(id)

	var currentPagu = rencanaSubKegiatan.PaguSubKegiatan

	rencanaSubKegiatan.RencanaKegiatanId = rencanaSubKegiatanRequest.RencanaKegiatanId
	rencanaSubKegiatan.SubKegiatanId = rencanaSubKegiatanRequest.SubKegiatanId
	rencanaSubKegiatan.PaguSubKegiatan = rencanaSubKegiatanRequest.PaguSubKegiatan
	rencanaSubKegiatan.Tipe = rencanaSubKegiatanRequest.Tipe
	rencanaSubKegiatan.Tahun = rencanaSubKegiatanRequest.Tahun

	var rencanaKegiatan, _ = s.rencanaKegiatanRepository.FindById(rencanaSubKegiatan.RencanaKegiatanId)

	var totalPaguFokusBelanja, _ = s.fokusBelanjaRepository.SumPaguFokusBelanja(id)

	var totalPaguRencanaSubKegiatan, _ = s.rencanaSubKegiatanRepository.SumPaguRencanaSubKegiatan(rencanaSubKegiatan.RencanaKegiatanId)
	totalPaguRencanaSubKegiatan = totalPaguRencanaSubKegiatan - currentPagu + rencanaSubKegiatanRequest.PaguSubKegiatan

	if rencanaSubKegiatanRequest.PaguSubKegiatan >= totalPaguFokusBelanja {
		if rencanaKegiatan.PaguKegiatan >= totalPaguRencanaSubKegiatan {
			updatedRencanaSubKegiatan, err := s.rencanaSubKegiatanRepository.Update(rencanaSubKegiatan)
			return updatedRencanaSubKegiatan, err
		} else {
			return rencanaSubKegiatan, errors.New("pagu sub kegiatan melebihi pagu kegiatannya")
		}
	} else {
		return rencanaSubKegiatan, errors.New("pagu parent lebih kecil dari pagu child")
	}

}

func (s *rencanaSubKegiatanService) Delete(id int) (model.RencanaSubKegiatan, error) {
	var rencanaSubKegiatan, err = s.rencanaSubKegiatanRepository.FindById(id)

	deletedRencanaSubKegiatan, err := s.rencanaSubKegiatanRepository.Delete(rencanaSubKegiatan)

	return deletedRencanaSubKegiatan, err
}
