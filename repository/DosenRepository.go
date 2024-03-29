package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type DosenRepository interface {
	FindAll() ([]model.Dosen, error)
	FindById(id int) (model.Dosen, error)
	FindByUserId(userId int) (model.Dosen, error)
	FindAllRelation() ([]model.Dosen, error)
	DistinctDosen(kelurahanId string) ([]model.LokasiDosen, error)
	DistinctLokasiDosen(dosenId int) ([]model.LokasiDosen, error)
	FindMahasiswa(kelurahanId string) ([]model.Mahasiswa, error)
	Create(dosen model.Dosen) (model.Dosen, error)
	Update(dosen model.Dosen) (model.Dosen, error)
	Delete(dosen model.Dosen) (model.Dosen, error)
}

type dosenRepository struct {
	db *gorm.DB
}

func NewDosenRepository(db *gorm.DB) *dosenRepository {
	return &dosenRepository{db}
}

func (r *dosenRepository) FindAll() ([]model.Dosen, error) {
	var dosens []model.Dosen

	var err = r.db.Find(&dosens).Error

	return dosens, err
}

func (r *dosenRepository) FindById(id int) (model.Dosen, error) {
	var dosen model.Dosen

	var err = r.db.Take(&dosen, id).Error

	return dosen, err
}

func (r *dosenRepository) FindByUserId(userId int) (model.Dosen, error) {
	var dosen model.Dosen

	var err = r.db.Where("userId = ? ", userId).Take(&dosen).Error

	return dosen, err
}

func (r *dosenRepository) FindAllRelation() ([]model.Dosen, error) {
	var dosens []model.Dosen

	var err = r.db.Model(&dosens).Preload("User").Find(&dosens).Error

	return dosens, err
}

func (r *dosenRepository) DistinctDosen(kelurahanId string) ([]model.LokasiDosen, error) {
	var lokasiDosens []model.LokasiDosen

	var err = r.db.Distinct("dosenId").Table("lokasi_dosens").Where("kelurahanId = ?", kelurahanId).Find(&lokasiDosens).Error

	return lokasiDosens, err
}

func (r *dosenRepository) DistinctLokasiDosen(dosenId int) ([]model.LokasiDosen, error) {
	var lokasiDosens []model.LokasiDosen
	var err error

	if dosenId == 0 {
		err = r.db.Distinct("kelurahanId").Table("lokasi_dosens").Find(&lokasiDosens).Error
	} else {
		err = r.db.Distinct("dosenId, kelurahanId").Table("lokasi_dosens").Where("dosenId = ?", dosenId).Find(&lokasiDosens).Error
	}

	return lokasiDosens, err
}

func (r *dosenRepository) FindMahasiswa(kelurahanId string) ([]model.Mahasiswa, error) {
	var mahasiswas []model.Mahasiswa

	var err = r.db.Where("kelurahanId = ?", kelurahanId).Model(&mahasiswas).Preload("User").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&mahasiswas).Error

	return mahasiswas, err
}

func (r *dosenRepository) Create(dosen model.Dosen) (model.Dosen, error) {
	var err = r.db.Create(&dosen).Error

	return dosen, err
}

func (r *dosenRepository) Update(dosen model.Dosen) (model.Dosen, error) {
	var err = r.db.Model(&dosen).Updates(model.Dosen{
		UserId:        dosen.UserId,
		NamaLengkap:   dosen.NamaLengkap,
		Universitas:   dosen.Universitas,
		UrlFotoProfil: dosen.UrlFotoProfil,
	}).Error

	return dosen, err
}

func (r *dosenRepository) Delete(dosen model.Dosen) (model.Dosen, error) {
	var err = r.db.Delete(&dosen).Error

	return dosen, err
}
