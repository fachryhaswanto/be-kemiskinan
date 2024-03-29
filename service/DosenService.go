package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type DosenService interface {
	FindAll() ([]model.Dosen, error)
	FindById(id int) (model.Dosen, error)
	FindByUserId(userId int) (model.Dosen, error)
	FindAllRelation() ([]model.Dosen, error)
	DistinctDosen(kelurahanId string) ([]model.LokasiDosen, error)
	DistinctLokasiDosen(dosenId int) ([]model.LokasiDosen, error)
	FindMahasiswa(kelurahanId string) ([]model.Mahasiswa, error)
	Create(dosenRequest request.CreateDosenRequest) (model.Dosen, error)
	Update(id int, dosenRequest request.UpdateDosenRequest) (model.Dosen, error)
	Delete(id int) (model.Dosen, error)
}

type dosenService struct {
	dosenRepository repository.DosenRepository
}

func NewDosenService(dosenRepository repository.DosenRepository) *dosenService {
	return &dosenService{dosenRepository}
}

func (s *dosenService) FindAll() ([]model.Dosen, error) {
	var dosens, err = s.dosenRepository.FindAll()

	return dosens, err
}

func (s *dosenService) FindById(id int) (model.Dosen, error) {
	var dosen, err = s.dosenRepository.FindById(id)

	return dosen, err
}

func (s *dosenService) FindByUserId(userId int) (model.Dosen, error) {
	var dosen, err = s.dosenRepository.FindByUserId(userId)

	return dosen, err
}

func (s *dosenService) FindAllRelation() ([]model.Dosen, error) {
	var dosens, err = s.dosenRepository.FindAllRelation()

	return dosens, err
}

func (s *dosenService) DistinctDosen(kelurahanId string) ([]model.LokasiDosen, error) {
	var lokasiDosens, err = s.dosenRepository.DistinctDosen(kelurahanId)

	return lokasiDosens, err
}

func (s *dosenService) DistinctLokasiDosen(dosenId int) ([]model.LokasiDosen, error) {
	var lokasiDosens, err = s.dosenRepository.DistinctLokasiDosen(dosenId)

	return lokasiDosens, err
}

func (s *dosenService) FindMahasiswa(kelurahanId string) ([]model.Mahasiswa, error) {
	var mahasiswas, err = s.dosenRepository.FindMahasiswa(kelurahanId)

	return mahasiswas, err
}

func (s *dosenService) Create(dosenRequest request.CreateDosenRequest) (model.Dosen, error) {
	var dosen = model.Dosen{
		UserId:        dosenRequest.UserId,
		NamaLengkap:   dosenRequest.NamaLengkap,
		Universitas:   dosenRequest.Universitas,
		UrlFotoProfil: dosenRequest.UrlFotoProfil,
	}

	newDosen, err := s.dosenRepository.Create(dosen)

	return newDosen, err
}

func (s *dosenService) Update(id int, dosenRequest request.UpdateDosenRequest) (model.Dosen, error) {
	var dosen, err = s.dosenRepository.FindById(id)

	dosen.UserId = dosenRequest.UserId
	dosen.NamaLengkap = dosenRequest.NamaLengkap
	dosen.Universitas = dosenRequest.Universitas
	dosen.UrlFotoProfil = dosenRequest.UrlFotoProfil

	updatedDosen, err := s.dosenRepository.Update(dosen)

	return updatedDosen, err
}

func (s *dosenService) Delete(id int) (model.Dosen, error) {
	var dosen, err = s.dosenRepository.FindById(id)

	deletedInstansi, err := s.dosenRepository.Delete(dosen)

	return deletedInstansi, err
}
