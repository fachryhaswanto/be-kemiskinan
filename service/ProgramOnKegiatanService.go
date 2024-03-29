package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type ProgramOnKegiatanService interface {
	FindAll() ([]model.ProgramOnKegiatan, error)
	FindById(id int) (model.ProgramOnKegiatan, error)
	FindByProgramId(programId int) ([]model.ProgramOnKegiatan, error)
	FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.ProgramOnKegiatan, error)
	CountJumlahKegiatanAllInstansi(tahun string, instansis []model.Instansi) []int64
	CountJumlahKegiatanByInstansiId(tahun string, instansi model.Instansi) int64
	Create(programOnKegiatanRequest request.CreateProgramOnKegiatanRequest) (model.ProgramOnKegiatan, error)
	Update(id int, programOnKegiatanRequest request.UpdateProgramOnKegiatanRequest) (model.ProgramOnKegiatan, error)
	Delete(id int) (model.ProgramOnKegiatan, error)
}

type programOnKegiatanService struct {
	programOnKegiatanRepository repository.ProgramOnKegiatanRepository
}

func NewProgramOnKegiatanService(programOnKegiatanRepository repository.ProgramOnKegiatanRepository) *programOnKegiatanService {
	return &programOnKegiatanService{programOnKegiatanRepository}
}

func (s *programOnKegiatanService) FindAll() ([]model.ProgramOnKegiatan, error) {
	var programOnKegiatans, err = s.programOnKegiatanRepository.FindAll()

	return programOnKegiatans, err
}

func (s *programOnKegiatanService) FindById(id int) (model.ProgramOnKegiatan, error) {
	var programOnKegiatan, err = s.programOnKegiatanRepository.FindById(id)

	return programOnKegiatan, err
}

func (s *programOnKegiatanService) FindByProgramId(programId int) ([]model.ProgramOnKegiatan, error) {
	var programOnKegiatans, err = s.programOnKegiatanRepository.FindByProgramId(programId)

	return programOnKegiatans, err
}

func (s *programOnKegiatanService) FindBySearch(whereClause map[string]interface{}, tahun string) ([]model.ProgramOnKegiatan, error) {
	var programOnKegiatans, err = s.programOnKegiatanRepository.FindBySearch(whereClause, tahun)

	return programOnKegiatans, err
}

func (s *programOnKegiatanService) CountJumlahKegiatanAllInstansi(tahun string, instansis []model.Instansi) []int64 {
	var jumlahKegiatanAllInstansi = s.programOnKegiatanRepository.CountJumlahKegiatanAllIntansi(tahun, instansis)

	return jumlahKegiatanAllInstansi
}

func (s *programOnKegiatanService) CountJumlahKegiatanByInstansiId(tahun string, instansi model.Instansi) int64 {
	var jumlahKegiatanByInstansi = s.programOnKegiatanRepository.CountJumlahKegiatanByInstansiId(tahun, instansi)

	return jumlahKegiatanByInstansi
}

func (s *programOnKegiatanService) Create(programOnKegiatanRequest request.CreateProgramOnKegiatanRequest) (model.ProgramOnKegiatan, error) {
	var programOnKegiatan = model.ProgramOnKegiatan{
		ProgramId:  programOnKegiatanRequest.ProgramId,
		KegiatanId: programOnKegiatanRequest.KegiatanId,
	}

	newProgramOnKegiatan, err := s.programOnKegiatanRepository.Create(programOnKegiatan)

	return newProgramOnKegiatan, err
}

func (s *programOnKegiatanService) Update(id int, programOnKegiatanRequest request.UpdateProgramOnKegiatanRequest) (model.ProgramOnKegiatan, error) {
	var programOnKegiatan, err = s.programOnKegiatanRepository.FindById(id)

	programOnKegiatan.ProgramId = programOnKegiatanRequest.ProgramId
	programOnKegiatan.KegiatanId = programOnKegiatanRequest.KegiatanId

	updatedProgramOnKegiatan, err := s.programOnKegiatanRepository.Update(programOnKegiatan)

	return updatedProgramOnKegiatan, err
}

func (s *programOnKegiatanService) Delete(id int) (model.ProgramOnKegiatan, error) {
	var programOnKegiatan, err = s.programOnKegiatanRepository.FindById(id)

	deletedProgramOnKegiatan, err := s.programOnKegiatanRepository.Delete(programOnKegiatan)

	return deletedProgramOnKegiatan, err
}
