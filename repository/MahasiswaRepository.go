package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	FindAll() ([]model.Mahasiswa, error)
	FindById(id int) (model.Mahasiswa, error)
	FindByUserId(userId int) (model.Mahasiswa, error)
	FindAllRelation() ([]model.Mahasiswa, error)
	CountVerifiedKeluarga(id int, kelurahanId string) (int64, error)
	CountVerifiedIndividu(id int, kelurahanId string) (int64, error)
	DistinctKelurahan() ([]model.DistinctKelurahan, error)
	Create(mahasiswa model.Mahasiswa) (model.Mahasiswa, error)
	CreateBatch(mahasiswas []model.Mahasiswa) ([]model.Mahasiswa, error)
	Update(mahasiswa model.Mahasiswa) (model.Mahasiswa, error)
	Delete(mahasiswa model.Mahasiswa) (model.Mahasiswa, error)
}

type mahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) *mahasiswaRepository {
	return &mahasiswaRepository{db}
}

func (r *mahasiswaRepository) FindAll() ([]model.Mahasiswa, error) {
	var mahasiswas []model.Mahasiswa

	var err = r.db.Find(&mahasiswas).Error

	return mahasiswas, err
}

func (r *mahasiswaRepository) FindById(id int) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	var err = r.db.Take(&mahasiswa, id).Error

	return mahasiswa, err
}

func (r *mahasiswaRepository) FindByUserId(userId int) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	var err = r.db.Where("userId = ?", userId).Take(&mahasiswa).Error

	return mahasiswa, err
}

func (r *mahasiswaRepository) FindAllRelation() ([]model.Mahasiswa, error) {
	var mahasiswas []model.Mahasiswa

	var err = r.db.Model(&mahasiswas).Preload("User").Preload("KabupatenKota").Preload("Kecamatan").Preload("Kelurahan").Find(&mahasiswas).Error

	return mahasiswas, err
}

func (r *mahasiswaRepository) CountVerifiedIndividu(id int, kelurahanId string) (int64, error) {
	var count int64

	var err = r.db.Where("mahasiswaId = ? and kelurahanId = ?", id, kelurahanId).Table("individus").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *mahasiswaRepository) CountVerifiedKeluarga(id int, kelurahanId string) (int64, error) {
	var count int64

	var err = r.db.Where("mahasiswaId = ? and kelurahanId = ?", id, kelurahanId).Table("keluargas").Select("count(*)").Count(&count).Error

	return count, err
}

func (r *mahasiswaRepository) DistinctKelurahan() ([]model.DistinctKelurahan, error) {
	var distinctsKelurahan []model.DistinctKelurahan

	var err = r.db.Distinct("m.kelurahanId, kl.nama").Table("mahasiswas as m").Joins("inner join kelurahans as kl on kl.id = m.kelurahanId").Scan(&distinctsKelurahan).Error

	return distinctsKelurahan, err
}

func (r *mahasiswaRepository) Create(mahasiswa model.Mahasiswa) (model.Mahasiswa, error) {
	var err = r.db.Create(&mahasiswa).Error

	return mahasiswa, err
}

func (r *mahasiswaRepository) CreateBatch(mahasiswas []model.Mahasiswa) ([]model.Mahasiswa, error) {
	var err = r.db.Create(&mahasiswas).Error

	return mahasiswas, err
}

func (r *mahasiswaRepository) Update(mahasiswa model.Mahasiswa) (model.Mahasiswa, error) {
	var err = r.db.Model(&mahasiswa).Updates(model.Mahasiswa{
		UserId:          mahasiswa.UserId,
		NamaLengkap:     mahasiswa.NamaLengkap,
		Universitas:     mahasiswa.Universitas,
		JenisKelamin:    mahasiswa.JenisKelamin,
		TanggalLahir:    mahasiswa.TanggalLahir,
		KabupatenKotaId: mahasiswa.KabupatenKotaId,
		KecamatanId:     mahasiswa.KecamatanId,
		KelurahanId:     mahasiswa.KelurahanId,
		UrlFotoProfil:   mahasiswa.UrlFotoProfil,
	}).Error

	return mahasiswa, err
}

func (r *mahasiswaRepository) Delete(mahasiswa model.Mahasiswa) (model.Mahasiswa, error) {
	var err = r.db.Delete(&mahasiswa).Error

	return mahasiswa, err
}
