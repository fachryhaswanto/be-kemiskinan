package helper

import (
	"kemiskinan/model"
	"kemiskinan/responses"
)

func ConvertToBidangUrusanResponse(b model.BidangUrusan) responses.BidangUrusanResponse {
	return responses.BidangUrusanResponse{
		Id:               b.Id,
		KodeBidangUrusan: b.KodeBidangUrusan,
		NamaBidangUrusan: b.NamaBidangUrusan,
		CreatedAt:        b.CreatedAt,
		UpdatedAt:        b.UpdatedAt,
	}
}

func ConvertToInstansiResponse(i model.Instansi) responses.InstansiResponse {
	return responses.InstansiResponse{
		Id:           i.Id,
		KodeInstansi: i.KodeInstansi,
		NamaInstansi: i.NamaInstansi,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
	}
}

func ConvertToBidangUrusanOnInstansiResponse(b model.BidangUrusanOnInstansi) responses.BidangUrusanOnInstansiResponse {
	return responses.BidangUrusanOnInstansiResponse{
		Id:         b.Id,
		InstansiId: b.InstansiId,
		Instansi: responses.InstansiResponse{
			Id:           b.Instansi.Id,
			KodeInstansi: b.Instansi.KodeInstansi,
			NamaInstansi: b.Instansi.NamaInstansi,
			CreatedAt:    b.Instansi.CreatedAt,
			UpdatedAt:    b.Instansi.UpdatedAt,
		},
		BidangUrusanId: b.BidangUrusanId,
		BidangUrusan: responses.BidangUrusanResponse{
			Id:               b.BidangUrusan.Id,
			KodeBidangUrusan: b.BidangUrusan.KodeBidangUrusan,
			NamaBidangUrusan: b.BidangUrusan.NamaBidangUrusan,
			CreatedAt:        b.BidangUrusan.CreatedAt,
			UpdatedAt:        b.BidangUrusan.UpdatedAt,
		},
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func ConvertToProgramResponse(p model.Program) responses.ProgramResponse {
	return responses.ProgramResponse{
		Id:          p.Id,
		Tahun:       p.Tahun,
		KodeProgram: p.KodeProgram,
		NamaProgram: p.NamaProgram,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ConvertToRencanaProgramResponse(r model.RencanaProgram) responses.RencanaProgramResponse {
	return responses.RencanaProgramResponse{
		Id:         r.Id,
		InstansiId: r.InstansiId,
		Instansi: responses.InstansiResponse{
			Id:           r.Instansi.Id,
			KodeInstansi: r.Instansi.KodeInstansi,
			NamaInstansi: r.Instansi.NamaInstansi,
			CreatedAt:    r.Instansi.CreatedAt,
			UpdatedAt:    r.UpdatedAt,
		},
		ProgramId: r.ProgramId,
		Program: responses.ProgramResponse{
			Id:          r.Program.Id,
			Tahun:       r.Program.Tahun,
			KodeProgram: r.Program.KodeProgram,
			NamaProgram: r.Program.NamaProgram,
			CreatedAt:   r.Program.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		},
		PaguProgram: r.PaguProgram,
		Tipe:        r.Tipe,
		Tahun:       r.Tahun,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func ConvertToInstansiOnProgramResponse(i model.InstansiOnProgram) responses.InstansiOnProgramResponse {
	return responses.InstansiOnProgramResponse{
		Id:         i.Id,
		InstansiId: i.InstansiId,
		Instansi: responses.InstansiResponse{
			Id:           i.Instansi.Id,
			KodeInstansi: i.Instansi.KodeInstansi,
			NamaInstansi: i.Instansi.NamaInstansi,
			CreatedAt:    i.Instansi.CreatedAt,
			UpdatedAt:    i.Instansi.UpdatedAt,
		},
		ProgramId: i.ProgramId,
		Program: responses.ProgramResponse{
			Id:          i.Program.Id,
			Tahun:       i.Program.Tahun,
			KodeProgram: i.Program.KodeProgram,
			NamaProgram: i.Program.NamaProgram,
			CreatedAt:   i.Program.CreatedAt,
			UpdatedAt:   i.Program.UpdatedAt,
		},
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

func ConvertToIndikatorSasaranResponse(i model.IndikatorSasaran) responses.IndikatorSasaranResponse {
	return responses.IndikatorSasaranResponse{
		Id:               i.Id,
		RencanaProgramId: i.RencanaProgramId,
		RencanaProgram: responses.RencanaProgramResponse{
			Id:          i.RencanaProgram.Id,
			InstansiId:  i.RencanaProgram.InstansiId,
			ProgramId:   i.RencanaProgram.ProgramId,
			PaguProgram: i.RencanaProgram.PaguProgram,
			Tipe:        i.RencanaProgram.Tipe,
			Tahun:       i.RencanaProgram.Tahun,
			CreatedAt:   i.RencanaProgram.CreatedAt,
			UpdatedAt:   i.RencanaProgram.UpdatedAt,
		},
		NamaIndikatorSasaran: i.NamaIndikatorSasaran,
		CreatedAt:            i.CreatedAt,
		UpdatedAt:            i.UpdatedAt,
	}
}

func ConvertToKegiatanResponse(k model.Kegiatan) responses.KegiatanResponse {
	return responses.KegiatanResponse{
		Id:           k.Id,
		Tahun:        k.Tahun,
		KodeKegiatan: k.KodeKegiatan,
		NamaKegiatan: k.NamaKegiatan,
		CreatedAt:    k.CreatedAt,
		UpdatedAt:    k.UpdatedAt,
	}
}

func ConvertToRencanaKegiatanResponse(r model.RencanaKegiatan) responses.RencanaKegiatanResponse {
	return responses.RencanaKegiatanResponse{
		Id:               r.Id,
		RencanaProgramId: r.RencanaProgramId,
		RencanaProgram: responses.RencanaProgramResponse{
			Id:         r.RencanaProgram.Id,
			InstansiId: r.RencanaProgram.InstansiId,
			Instansi: responses.InstansiResponse{
				Id:           r.RencanaProgram.Instansi.Id,
				KodeInstansi: r.RencanaProgram.Instansi.KodeInstansi,
				NamaInstansi: r.RencanaProgram.Instansi.NamaInstansi,
				CreatedAt:    r.RencanaProgram.Instansi.CreatedAt,
				UpdatedAt:    r.RencanaProgram.Instansi.UpdatedAt,
			},
			ProgramId: r.RencanaProgram.ProgramId,
			Program: responses.ProgramResponse{
				Id:          r.RencanaProgram.Program.Id,
				Tahun:       r.RencanaProgram.Program.Tahun,
				KodeProgram: r.RencanaProgram.Program.KodeProgram,
				NamaProgram: r.RencanaProgram.Program.NamaProgram,
				CreatedAt:   r.RencanaProgram.Program.CreatedAt,
				UpdatedAt:   r.RencanaProgram.Program.UpdatedAt,
			},
			PaguProgram: r.RencanaProgram.PaguProgram,
			Tipe:        r.RencanaProgram.Tipe,
			Tahun:       r.RencanaProgram.Tahun,
			CreatedAt:   r.RencanaProgram.CreatedAt,
			UpdatedAt:   r.RencanaProgram.UpdatedAt,
		},
		KegiatanId: r.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           r.Kegiatan.Id,
			Tahun:        r.Kegiatan.Tahun,
			KodeKegiatan: r.Kegiatan.KodeKegiatan,
			NamaKegiatan: r.Kegiatan.NamaKegiatan,
			CreatedAt:    r.Kegiatan.CreatedAt,
			UpdatedAt:    r.Kegiatan.UpdatedAt,
		},
		PaguKegiatan: r.PaguKegiatan,
		Tipe:         r.Tipe,
		Tahun:        r.Tahun,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}

func ConvertToProgramOnKegiatanResponse(p model.ProgramOnKegiatan) responses.ProgramOnKegiatanResponse {
	return responses.ProgramOnKegiatanResponse{
		Id:        p.Id,
		ProgramId: p.ProgramId,
		Program: responses.ProgramResponse{
			Id:          p.Program.Id,
			Tahun:       p.Program.Tahun,
			KodeProgram: p.Program.KodeProgram,
			NamaProgram: p.Program.NamaProgram,
			CreatedAt:   p.Program.CreatedAt,
			UpdatedAt:   p.Program.UpdatedAt,
		},
		KegiatanId: p.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           p.Kegiatan.Id,
			Tahun:        p.Kegiatan.Tahun,
			KodeKegiatan: p.Kegiatan.KodeKegiatan,
			NamaKegiatan: p.Kegiatan.NamaKegiatan,
			CreatedAt:    p.Kegiatan.CreatedAt,
			UpdatedAt:    p.Kegiatan.UpdatedAt,
		},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ConvertToSasaranResponse(s model.Sasaran) responses.SasaranResponse {
	return responses.SasaranResponse{
		Id:               s.Id,
		RencanaProgramId: s.RencanaProgramId,
		RencanaProgram: responses.RencanaProgramResponse{
			Id:          s.RencanaProgram.Id,
			InstansiId:  s.RencanaProgram.InstansiId,
			ProgramId:   s.RencanaProgram.ProgramId,
			PaguProgram: s.RencanaProgram.PaguProgram,
			Tipe:        s.RencanaProgram.Tipe,
			Tahun:       s.RencanaProgram.Tahun,
			CreatedAt:   s.RencanaProgram.CreatedAt,
			UpdatedAt:   s.RencanaProgram.UpdatedAt,
		},
		NamaSasaran: s.NamaSasaran,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func ConvertToSubKegiatanResponse(s model.SubKegiatan) responses.SubKegiatanResponse {
	return responses.SubKegiatanResponse{
		Id:              s.Id,
		Tahun:           s.Tahun,
		KodeSubKegiatan: s.KodeSubKegiatan,
		NamaSubKegiatan: s.NamaSubKegiatan,
		CreatedAt:       s.CreatedAt,
		UpdatedAt:       s.UpdatedAt,
	}
}

func ConvertToRencanaSubKegiatanResponse(r model.RencanaSubKegiatan) responses.RencanaSubKegiatanResponse {
	return responses.RencanaSubKegiatanResponse{
		Id:                r.Id,
		RencanaKegiatanId: r.RencanaKegiatanId,
		RencanaKegiatan: responses.RencanaKegiatanResponse{
			Id:               r.RencanaKegiatan.Id,
			RencanaProgramId: r.RencanaKegiatan.RencanaProgramId,
			RencanaProgram: responses.RencanaProgramResponse{
				Id:         r.RencanaKegiatan.RencanaProgram.Id,
				InstansiId: r.RencanaKegiatan.RencanaProgram.InstansiId,
				Instansi: responses.InstansiResponse{
					Id:           r.RencanaKegiatan.RencanaProgram.Instansi.Id,
					KodeInstansi: r.RencanaKegiatan.RencanaProgram.Instansi.KodeInstansi,
					NamaInstansi: r.RencanaKegiatan.RencanaProgram.Instansi.NamaInstansi,
					CreatedAt:    r.RencanaKegiatan.RencanaProgram.Instansi.CreatedAt,
					UpdatedAt:    r.RencanaKegiatan.RencanaProgram.Instansi.UpdatedAt,
				},
				ProgramId: r.RencanaKegiatan.RencanaProgram.ProgramId,
				Program: responses.ProgramResponse{
					Id:          r.RencanaKegiatan.RencanaProgram.Program.Id,
					Tahun:       r.RencanaKegiatan.RencanaProgram.Program.Tahun,
					KodeProgram: r.RencanaKegiatan.RencanaProgram.Program.KodeProgram,
					NamaProgram: r.RencanaKegiatan.RencanaProgram.Program.NamaProgram,
					CreatedAt:   r.RencanaKegiatan.RencanaProgram.Program.CreatedAt,
					UpdatedAt:   r.RencanaKegiatan.RencanaProgram.Program.UpdatedAt,
				},
				PaguProgram: r.RencanaKegiatan.RencanaProgram.PaguProgram,
				Tipe:        r.RencanaKegiatan.RencanaProgram.Tipe,
				Tahun:       r.RencanaKegiatan.Tahun,
				CreatedAt:   r.RencanaKegiatan.RencanaProgram.CreatedAt,
				UpdatedAt:   r.RencanaKegiatan.RencanaProgram.UpdatedAt,
			},
			KegiatanId: r.RencanaKegiatan.KegiatanId,
			Kegiatan: responses.KegiatanResponse{
				Id:           r.RencanaKegiatan.Kegiatan.Id,
				Tahun:        r.RencanaKegiatan.Kegiatan.Tahun,
				KodeKegiatan: r.RencanaKegiatan.Kegiatan.KodeKegiatan,
				NamaKegiatan: r.RencanaKegiatan.Kegiatan.NamaKegiatan,
				CreatedAt:    r.RencanaKegiatan.Kegiatan.CreatedAt,
				UpdatedAt:    r.RencanaKegiatan.Kegiatan.UpdatedAt,
			},
			PaguKegiatan: r.RencanaKegiatan.PaguKegiatan,
			Tipe:         r.RencanaKegiatan.Tipe,
			CreatedAt:    r.RencanaKegiatan.CreatedAt,
			UpdatedAt:    r.RencanaKegiatan.UpdatedAt,
		},
		SubKegiatanId: r.SubKegiatanId,
		SubKegiatan: responses.SubKegiatanResponse{
			Id:              r.SubKegiatan.Id,
			Tahun:           r.SubKegiatan.Tahun,
			KodeSubKegiatan: r.SubKegiatan.KodeSubKegiatan,
			NamaSubKegiatan: r.SubKegiatan.NamaSubKegiatan,
			CreatedAt:       r.SubKegiatan.CreatedAt,
			UpdatedAt:       r.SubKegiatan.UpdatedAt,
		},
		Tipe:            r.Tipe,
		PaguSubKegiatan: r.PaguSubKegiatan,
		Tahun:           r.Tahun,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}
}

func ConvertToKebijakanResponse(k model.Kebijakan) responses.KebijakanResponse {
	return responses.KebijakanResponse{
		Id:               k.Id,
		RencanaProgramId: k.RencanaProgramId,
		RencanaProgram: responses.RencanaProgramResponse{
			Id:          k.RencanaProgram.Id,
			InstansiId:  k.RencanaProgram.InstansiId,
			ProgramId:   k.RencanaProgram.ProgramId,
			PaguProgram: k.RencanaProgram.PaguProgram,
			Tipe:        k.RencanaProgram.Tipe,
			Tahun:       k.RencanaProgram.Tahun,
			CreatedAt:   k.RencanaProgram.CreatedAt,
			UpdatedAt:   k.RencanaProgram.UpdatedAt,
		},
		NamaKebijakan: k.NamaKebijakan,
		CreatedAt:     k.CreatedAt,
		UpdatedAt:     k.UpdatedAt,
	}
}

func ConvertToKegiatanOnSubKegiatanResponse(k model.KegiatanOnSubKegiatan) responses.KegiatanOnSubKegiatanResponse {
	return responses.KegiatanOnSubKegiatanResponse{
		Id:         k.Id,
		KegiatanId: k.KegiatanId,
		Kegiatan: responses.KegiatanResponse{
			Id:           k.Kegiatan.Id,
			Tahun:        k.Kegiatan.Tahun,
			KodeKegiatan: k.Kegiatan.KodeKegiatan,
			NamaKegiatan: k.Kegiatan.NamaKegiatan,
			CreatedAt:    k.Kegiatan.CreatedAt,
			UpdatedAt:    k.Kegiatan.UpdatedAt,
		},
		SubKegiatanId: k.SubKegiatanId,
		SubKegiatan: responses.SubKegiatanResponse{
			Id:              k.SubKegiatan.Id,
			Tahun:           k.SubKegiatan.Tahun,
			KodeSubKegiatan: k.SubKegiatan.KodeSubKegiatan,
			NamaSubKegiatan: k.SubKegiatan.NamaSubKegiatan,
			CreatedAt:       k.SubKegiatan.CreatedAt,
			UpdatedAt:       k.SubKegiatan.UpdatedAt,
		},
		CreatedAt: k.CreatedAt,
		UpdatedAt: k.UpdatedAt,
	}
}

func ConvertToFokusBelanjaResponse(f model.FokusBelanja) responses.FokusBelanjaResponse {
	return responses.FokusBelanjaResponse{
		Id:                   f.Id,
		RencanaSubKegiatanId: f.RencanaSubKegiatanId,
		RencanaSubKegiatan: responses.RencanaSubKegiatanResponse{
			Id:                f.RencanaSubKegiatan.Id,
			RencanaKegiatanId: f.RencanaSubKegiatan.RencanaKegiatanId,
			SubKegiatanId:     f.RencanaSubKegiatanId,
			PaguSubKegiatan:   f.RencanaSubKegiatan.PaguSubKegiatan,
			Tipe:              f.RencanaSubKegiatan.Tipe,
			Tahun:             f.RencanaSubKegiatan.Tahun,
			CreatedAt:         f.RencanaSubKegiatan.CreatedAt,
			UpdatedAt:         f.RencanaSubKegiatan.UpdatedAt,
		},
		NamaFokusBelanja: f.NamaFokusBelanja,
		Indikator:        f.Indikator,
		Target:           f.Target,
		Satuan:           f.Satuan,
		PaguFokusBelanja: f.PaguFokusBelanja,
		Keterangan:       f.Keterangan,
		Tahun:            f.Tahun,
		CreatedAt:        f.CreatedAt,
		UpdatedAt:        f.UpdatedAt,
	}
}

func ConvertToRealisasiResponse(r model.Realisasi) responses.RealisasiResponse {
	return responses.RealisasiResponse{
		Id:             r.Id,
		FokusBelanjaId: r.FokusBelanjaId,
		FokusBelanja: responses.FokusBelanjaResponse{
			Id:               r.FokusBelanja.Id,
			NamaFokusBelanja: r.FokusBelanja.NamaFokusBelanja,
			Indikator:        r.FokusBelanja.Indikator,
			Target:           r.FokusBelanja.Target,
			Satuan:           r.FokusBelanja.Satuan,
			PaguFokusBelanja: r.FokusBelanja.PaguFokusBelanja,
			Keterangan:       r.FokusBelanja.Keterangan,
			Tahun:            r.FokusBelanja.Tahun,
			CreatedAt:        r.FokusBelanja.CreatedAt,
			UpdatedAt:        r.FokusBelanja.UpdatedAt,
		},
		Tanggal:         r.Tanggal,
		NomorSp2d:       r.NomorSp2d,
		RealisasiPagu:   r.RealisasiPagu,
		RealisasiTarget: r.RealisasiTarget,
		Bulan:           r.Bulan,
		Keterangan:      r.Keterangan,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}
}

func ConvertToDetailLokasiResponse(d model.Detail_Lokasi) responses.Detail_LokasiResponse {
	return responses.Detail_LokasiResponse{
		Id:             d.Id,
		FokusBelanjaid: d.FokusBelanjaId,
		FokusBelanja: responses.FokusBelanjaResponse{
			Id:                   d.FokusBelanja.Id,
			RencanaSubKegiatanId: d.FokusBelanja.RencanaSubKegiatanId,
			NamaFokusBelanja:     d.FokusBelanja.NamaFokusBelanja,
			Indikator:            d.FokusBelanja.Indikator,
			Target:               d.FokusBelanja.Target,
			Satuan:               d.FokusBelanja.Satuan,
			PaguFokusBelanja:     d.FokusBelanja.PaguFokusBelanja,
			Keterangan:           d.FokusBelanja.Keterangan,
			Tahun:                d.FokusBelanja.Tahun,
			CreatedAt:            d.FokusBelanja.CreatedAt,
			UpdatedAt:            d.FokusBelanja.UpdatedAt,
		},
		KabupatenKotaId: d.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         d.KabupatenKota.Id,
			ProvinsiId: d.KabupatenKota.ProvinsiId,
			Nama:       d.KabupatenKota.Nama,
		},
		KecamatanId: d.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              d.Kecamatan.Id,
			KabupatenKotaId: d.Kecamatan.KabupatenKotaId,
			Nama:            d.Kecamatan.Nama,
		},
		KelurahanId: d.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          d.Kelurahan.Id,
			KecamatanId: d.Kelurahan.KecamatanId,
			Nama:        d.Kelurahan.Nama,
		},
		TipePenerima: d.TipePenerima,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

func ConvertToProvinsiResponse(p model.Provinsi) responses.ProvinsiResponse {
	return responses.ProvinsiResponse{
		Id:   p.Id,
		Nama: p.Nama,
	}
}

func ConvertToKabupatenKotaResponse(k model.Kabupaten_Kota) responses.Kabupaten_KotaResponse {
	return responses.Kabupaten_KotaResponse{
		Id:         k.Id,
		ProvinsiId: k.ProvinsiId,
		Nama:       k.Nama,
	}
}

func ConvertToKecamatanResponse(k model.Kecamatan) responses.KecamatanResponse {
	return responses.KecamatanResponse{
		Id:              k.Id,
		KabupatenKotaId: k.KabupatenKotaId,
		Nama:            k.Nama,
	}
}

func ConvertToKelurahanResponse(k model.Kelurahan) responses.KelurahanResponse {
	return responses.KelurahanResponse{
		Id:          k.Id,
		KecamatanId: k.KecamatanId,
		Nama:        k.Nama,
	}
}

func ConvertToDistinctKelurahanResponse(d model.DistinctKelurahan) responses.DistinctKelurahan {
	return responses.DistinctKelurahan{
		KelurahanId: d.KelurahanId,
		Nama:        d.Nama,
	}
}

func ConvertToKeluargaResponse(k model.Keluarga) responses.KeluargaResponse {
	return responses.KeluargaResponse{
		Id:         k.Id,
		IdKeluarga: k.IdKeluarga,
		ProvinsiId: k.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   k.Provinsi.Id,
			Nama: k.Provinsi.Nama,
		},

		KabupatenKotaId: k.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         k.KabupatenKota.Id,
			ProvinsiId: k.KabupatenKota.ProvinsiId,
			Nama:       k.KabupatenKota.Nama,
		},

		KecamatanId: k.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              k.Kecamatan.Id,
			KabupatenKotaId: k.Kecamatan.KabupatenKotaId,
			Nama:            k.Kecamatan.Nama,
		},

		KelurahanId: k.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          k.Kelurahan.Id,
			KecamatanId: k.Kelurahan.KecamatanId,
			Nama:        k.Kelurahan.Nama,
		},

		DesilKesejahteraan:     k.DesilKesejahteraan,
		Alamat:                 k.Alamat,
		KepalaKeluarga:         k.KepalaKeluarga,
		Nik:                    k.Nik,
		PadanDukcapil:          k.PadanDukcapil,
		JenisKelamin:           k.JenisKelamin,
		TanggalLahir:           k.TanggalLahir,
		Pekerjaan:              k.Pekerjaan,
		Pendidikan:             k.Pendidikan,
		KepemilikanRumah:       k.KepemilikanRumah,
		Simpanan:               k.Simpanan,
		JenisAtap:              k.JenisAtap,
		JenisDinding:           k.JenisDinding,
		JenisLantai:            k.JenisLantai,
		SumberPenerangan:       k.SumberPenerangan,
		BahanBakarMemasak:      k.BahanBakarMemasak,
		SumberAirMinum:         k.SumberAirMinum,
		FasilitasBuangAirBesar: k.FasilitasBuangAirBesar,
		PenerimaBPNT:           k.PenerimaBPNT,
		PenerimaBPUM:           k.PenerimaBPUM,
		PenerimaBST:            k.PenerimaBST,
		PenerimaPKH:            k.PenerimaPKH,
		PenerimaSembako:        k.PenerimaSembako,

		UserId: k.UserId,
		User: responses.UserResponse{
			Id:        k.User.Id,
			Username:  k.User.Username,
			Password:  k.User.Password,
			Email:     k.User.Email,
			NoHp:      k.User.NoHp,
			Role:      k.User.Role,
			CreatedAt: k.User.CreatedAt,
			UpdatedAt: k.User.UpdatedAt,
		},

		MahasiswaId: k.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:           k.Mahasiswa.Id,
			UserId:       k.Mahasiswa.UserId,
			NamaLengkap:  k.Mahasiswa.NamaLengkap,
			Universitas:  k.Mahasiswa.Universitas,
			JenisKelamin: k.Mahasiswa.JenisKelamin,
			TanggalLahir: k.Mahasiswa.TanggalLahir,
			KelurahanId:  k.Mahasiswa.KelurahanId,
			CreatedAt:    k.Mahasiswa.CreatedAt,
			UpdatedAt:    k.Mahasiswa.UpdatedAt,
		},

		StatusVerifikasi: k.StatusVerifikasi,
	}
}

func ConvertToMonevResponse(m model.Monev) responses.MonevResponse {
	return responses.MonevResponse{
		Id:              m.Id,
		NamaPenerima:    m.NamaPenerima,
		KabupatenKotaId: m.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         m.KabupatenKota.Id,
			ProvinsiId: m.KabupatenKota.ProvinsiId,
			Nama:       m.KabupatenKota.Nama,
		},
		KecamatanId: m.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              m.Kecamatan.Id,
			KabupatenKotaId: m.Kecamatan.KabupatenKotaId,
			Nama:            m.Kecamatan.Nama,
		},
		KelurahanId: m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		JenisBantuan:  m.JenisBantuan,
		VolumeBantuan: m.VolumeBantuan,
		SatuanVolume:  m.SatuanVolume,
	}
}

func ConvertToJumlahResponse(jumlah int64) responses.JumlahResponse {
	return responses.JumlahResponse{
		Jumlah: jumlah,
	}
}

func ConvertToUserResponse(u model.User) responses.UserResponse {
	return responses.UserResponse{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		NoHp:      u.NoHp,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ConvertToAdminResponse(a model.Admin) responses.AdminResponse {
	return responses.AdminResponse{
		Id:            a.Id,
		UserId:        a.UserId,
		NamaLengkap:   a.NamaLengkap,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToAdminOpdResponse(a model.AdminOpd) responses.AdminOpdResponse {
	return responses.AdminOpdResponse{
		Id:     a.Id,
		UserId: a.UserId,
		User: responses.UserResponse{
			Id:        a.User.Id,
			Username:  a.User.Username,
			Password:  a.User.Password,
			Email:     a.User.Email,
			NoHp:      a.User.NoHp,
			Role:      a.User.Role,
			CreatedAt: a.User.CreatedAt,
			UpdatedAt: a.User.UpdatedAt,
		},
		NamaLengkap: a.NamaLengkap,
		InstansiId:  a.InstansiId,
		Instansi: responses.InstansiResponse{
			Id:           a.Instansi.Id,
			KodeInstansi: a.Instansi.KodeInstansi,
			NamaInstansi: a.Instansi.NamaInstansi,
			CreatedAt:    a.Instansi.CreatedAt,
			UpdatedAt:    a.Instansi.UpdatedAt,
		},
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToAdminWithRelationResponse(a model.Admin) responses.AdminWithRelationResponse {
	return responses.AdminWithRelationResponse{
		Id:     a.Id,
		UserId: a.UserId,
		User: responses.UserResponse{
			Id:        a.User.Id,
			Username:  a.User.Username,
			Password:  a.User.Password,
			Email:     a.User.Email,
			NoHp:      a.User.NoHp,
			Role:      a.User.Role,
			CreatedAt: a.User.CreatedAt,
			UpdatedAt: a.User.UpdatedAt,
		},
		NamaLengkap:   a.NamaLengkap,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToAnalisResponse(a model.Analis) responses.AnalisResponse {
	return responses.AnalisResponse{
		Id:            a.Id,
		UserId:        a.UserId,
		NamaLengkap:   a.NamaLengkap,
		Universitas:   a.Universitas,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToAnalisWithRelationResponse(a model.Analis) responses.AnalisWithRelationResponse {
	return responses.AnalisWithRelationResponse{
		Id:     a.Id,
		UserId: a.UserId,
		User: responses.UserResponse{
			Id:        a.User.Id,
			Username:  a.User.Username,
			Password:  a.User.Password,
			Email:     a.User.Email,
			NoHp:      a.User.NoHp,
			Role:      a.User.Role,
			CreatedAt: a.User.CreatedAt,
			UpdatedAt: a.User.UpdatedAt,
		},
		NamaLengkap:   a.NamaLengkap,
		Universitas:   a.Universitas,
		UrlFotoProfil: a.UrlFotoProfil,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func ConvertToPusbangResponse(p model.Pusbang) responses.PusbangResponse {
	return responses.PusbangResponse{
		Id:            p.Id,
		UserId:        p.UserId,
		NamaLengkap:   p.NamaLengkap,
		Universitas:   p.Universitas,
		UrlFotoProfil: p.UrlFotoProfil,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func ConvertToPusbangWithRelationResponse(p model.Pusbang) responses.PusbangWithRelationResponse {
	return responses.PusbangWithRelationResponse{
		Id:     p.Id,
		UserId: p.UserId,
		User: responses.UserResponse{
			Id:        p.User.Id,
			Username:  p.User.Username,
			Password:  p.User.Password,
			Email:     p.User.Email,
			NoHp:      p.User.NoHp,
			Role:      p.User.Role,
			CreatedAt: p.User.CreatedAt,
			UpdatedAt: p.User.UpdatedAt,
		},
		NamaLengkap:   p.NamaLengkap,
		Universitas:   p.Universitas,
		UrlFotoProfil: p.UrlFotoProfil,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func ConvertToDosenResponse(d model.Dosen) responses.DosenResponse {
	return responses.DosenResponse{
		Id:            d.Id,
		UserId:        d.UserId,
		NamaLengkap:   d.NamaLengkap,
		Universitas:   d.Universitas,
		UrlFotoProfil: d.UrlFotoProfil,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

func ConvertToDosenWithRelationResponse(d model.Dosen) responses.DosenWithRelationResponse {
	return responses.DosenWithRelationResponse{
		Id:     d.Id,
		UserId: d.UserId,
		User: responses.UserResponse{
			Id:        d.User.Id,
			Username:  d.User.Username,
			Password:  d.User.Password,
			Email:     d.User.Email,
			NoHp:      d.User.NoHp,
			Role:      d.User.Role,
			CreatedAt: d.User.CreatedAt,
			UpdatedAt: d.User.UpdatedAt,
		},
		NamaLengkap:   d.NamaLengkap,
		Universitas:   d.Universitas,
		UrlFotoProfil: d.UrlFotoProfil,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

func ConvertToMahasiswaResponse(m model.Mahasiswa) responses.MahasiswaResponse {
	return responses.MahasiswaResponse{
		Id:              m.Id,
		UserId:          m.UserId,
		NamaLengkap:     m.NamaLengkap,
		Universitas:     m.Universitas,
		JenisKelamin:    m.JenisKelamin,
		TanggalLahir:    m.TanggalLahir,
		KabupatenKotaId: m.KabupatenKotaId,
		KecamatanId:     m.KecamatanId,
		KelurahanId:     m.KelurahanId,
		UrlFotoProfil:   m.UrlFotoProfil,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func ConvertToMahasiswaWithRelationResponse(m model.Mahasiswa) responses.MahasiswaWithRelationResponse {
	return responses.MahasiswaWithRelationResponse{
		Id:     m.Id,
		UserId: m.UserId,
		User: responses.UserResponse{
			Id:        m.User.Id,
			Username:  m.User.Username,
			Password:  m.User.Password,
			Email:     m.User.Email,
			NoHp:      m.User.NoHp,
			Role:      m.User.Role,
			CreatedAt: m.User.CreatedAt,
			UpdatedAt: m.User.UpdatedAt,
		},
		NamaLengkap:     m.NamaLengkap,
		Universitas:     m.Universitas,
		JenisKelamin:    m.JenisKelamin,
		TanggalLahir:    m.TanggalLahir,
		KabupatenKotaId: m.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         m.KabupatenKota.Id,
			ProvinsiId: m.KabupatenKota.ProvinsiId,
			Nama:       m.KabupatenKota.Nama,
		},
		KecamatanId: m.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              m.Kecamatan.Id,
			KabupatenKotaId: m.Kecamatan.KabupatenKotaId,
			Nama:            m.KabupatenKota.Nama,
		},
		KelurahanId: m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		UrlFotoProfil: m.UrlFotoProfil,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

func ConvertToMahasiswaWithVerifiedCountResponse(m model.Mahasiswa, jumlahVerifiedIndividu int64, jumlahVerifiedKeluarga int64) responses.MahasiswaWithVerifiedCountResponse {
	return responses.MahasiswaWithVerifiedCountResponse{
		Id:     m.Id,
		UserId: m.UserId,
		User: responses.UserResponse{
			Id:        m.User.Id,
			Username:  m.User.Username,
			Password:  m.User.Password,
			Email:     m.User.Email,
			NoHp:      m.User.NoHp,
			Role:      m.User.Role,
			CreatedAt: m.User.CreatedAt,
			UpdatedAt: m.User.UpdatedAt,
		},
		NamaLengkap:     m.NamaLengkap,
		Universitas:     m.Universitas,
		JenisKelamin:    m.JenisKelamin,
		TanggalLahir:    m.TanggalLahir,
		KabupatenKotaId: m.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         m.KabupatenKota.Id,
			ProvinsiId: m.KabupatenKota.ProvinsiId,
			Nama:       m.KabupatenKota.Nama,
		},
		KecamatanId: m.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              m.Kecamatan.Id,
			KabupatenKotaId: m.Kecamatan.KabupatenKotaId,
			Nama:            m.KabupatenKota.Nama,
		},
		KelurahanId: m.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          m.Kelurahan.Id,
			KecamatanId: m.Kelurahan.KecamatanId,
			Nama:        m.Kelurahan.Nama,
		},
		JumlahIndividuVerified: jumlahVerifiedIndividu,
		JumlahKeluargaVerified: jumlahVerifiedKeluarga,
		UrlFotoProfil:          m.UrlFotoProfil,
		CreatedAt:              m.CreatedAt,
		UpdatedAt:              m.UpdatedAt,
	}
}

func ConvertToLokasiDosenResponse(l model.LokasiDosen) responses.LokasiDosenResponse {
	return responses.LokasiDosenResponse{
		Id:              l.Id,
		DosenId:         l.DosenId,
		KabupatenKotaId: l.KabupatenKotaId,
		KecamatanId:     l.KecamatanId,
		KelurahanId:     l.KelurahanId,
		CreatedAt:       l.CreatedAt,
		UpdatedAt:       l.UpdatedAt,
	}
}

func ConvertToLokasiDosenWithRelationResponse(l model.LokasiDosen) responses.LokasiDosenWithRelationResponse {
	return responses.LokasiDosenWithRelationResponse{
		Id:      l.Id,
		DosenId: l.DosenId,
		Dosen: responses.DosenResponse{
			Id:          l.Dosen.Id,
			UserId:      l.Dosen.UserId,
			NamaLengkap: l.Dosen.NamaLengkap,
			Universitas: l.Dosen.Universitas,
			CreatedAt:   l.Dosen.CreatedAt,
			UpdatedAt:   l.Dosen.UpdatedAt,
		},
		KabupatenKotaId: l.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         l.KabupatenKota.Id,
			ProvinsiId: l.KabupatenKota.ProvinsiId,
			Nama:       l.KabupatenKota.Nama,
		},
		KecamatanId: l.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              l.Kecamatan.Id,
			KabupatenKotaId: l.Kecamatan.KabupatenKotaId,
			Nama:            l.Kecamatan.Nama,
		},
		KelurahanId: l.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          l.Kelurahan.Id,
			KecamatanId: l.Kelurahan.KecamatanId,
			Nama:        l.Kelurahan.Nama,
		},
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}

func ConvertToKeluargaVerifikasiResponse(k model.KeluargaVerifikasi) responses.KeluargaVerifikasiResponse {
	return responses.KeluargaVerifikasiResponse{
		Id:         k.Id,
		IdKeluarga: k.IdKeluarga,
		ProvinsiId: k.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   k.Provinsi.Id,
			Nama: k.Provinsi.Nama,
		},

		KabupatenKotaId: k.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         k.KabupatenKota.Id,
			ProvinsiId: k.KabupatenKota.ProvinsiId,
			Nama:       k.KabupatenKota.Nama,
		},

		KecamatanId: k.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              k.Kecamatan.Id,
			KabupatenKotaId: k.Kecamatan.KabupatenKotaId,
			Nama:            k.Kecamatan.Nama,
		},

		KelurahanId:            k.KelurahanId,
		DesilKesejahteraan:     k.DesilKesejahteraan,
		Alamat:                 k.Alamat,
		KepalaKeluarga:         k.KepalaKeluarga,
		Nik:                    k.Nik,
		PadanDukcapil:          k.PadanDukcapil,
		JenisKelamin:           k.JenisKelamin,
		TanggalLahir:           k.TanggalLahir,
		Pekerjaan:              k.Pekerjaan,
		Pendidikan:             k.Pendidikan,
		KepemilikanRumah:       k.KepemilikanRumah,
		Simpanan:               k.Simpanan,
		JenisAtap:              k.JenisAtap,
		JenisDinding:           k.JenisDinding,
		JenisLantai:            k.JenisLantai,
		SumberPenerangan:       k.SumberPenerangan,
		BahanBakarMemasak:      k.BahanBakarMemasak,
		SumberAirMinum:         k.SumberAirMinum,
		FasilitasBuangAirBesar: k.FasilitasBuangAirBesar,
		PenerimaBPNT:           k.PenerimaBPNT,
		PenerimaBPUM:           k.PenerimaBPUM,
		PenerimaBST:            k.PenerimaBST,
		PenerimaPKH:            k.PenerimaPKH,
		PenerimaSembako:        k.PenerimaSembako,
		PenerimaLainnya:        k.PenerimaLainnya,
		StatusResponden:        k.StatusResponden,

		UserId: k.UserId,
		User: responses.UserResponse{
			Id:        k.User.Id,
			Username:  k.User.Username,
			Password:  k.User.Password,
			Email:     k.User.Email,
			NoHp:      k.User.NoHp,
			Role:      k.User.Role,
			CreatedAt: k.User.CreatedAt,
			UpdatedAt: k.User.UpdatedAt,
		},

		MahasiswaId: k.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:           k.Mahasiswa.Id,
			UserId:       k.UserId,
			NamaLengkap:  k.Mahasiswa.NamaLengkap,
			Universitas:  k.Mahasiswa.Universitas,
			JenisKelamin: k.JenisKelamin,
			TanggalLahir: k.TanggalLahir,
			KelurahanId:  k.KelurahanId,
		},

		Kelurahan: responses.KelurahanResponse{
			Id:          k.Kelurahan.Id,
			KecamatanId: k.Kelurahan.KecamatanId,
			Nama:        k.Kelurahan.Nama,
		},

		UrlBukti: k.UrlBukti,

		CreatedAt: k.CreatedAt,
		UpdatedAt: k.UpdatedAt,
	}
}

func ConvertToIndividuResponse(i model.Individu) responses.IndividuResponse {
	return responses.IndividuResponse{
		Id:         i.Id,
		IdKeluarga: i.IdKeluarga,

		ProvinsiId: i.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   i.Provinsi.Id,
			Nama: i.Provinsi.Nama,
		},

		KabupatenKotaId: i.KabupatenKotaId,
		Kabupaten_Kota: responses.Kabupaten_KotaResponse{
			Id:         i.KabupatenKota.Id,
			ProvinsiId: i.KabupatenKota.ProvinsiId,
			Nama:       i.KabupatenKota.Nama,
		},

		KecamatanId: i.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              i.Kecamatan.Id,
			KabupatenKotaId: i.Kecamatan.KabupatenKotaId,
			Nama:            i.Kecamatan.Nama,
		},

		KelurahanId: i.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          i.Kelurahan.Id,
			KecamatanId: i.Kelurahan.KecamatanId,
			Nama:        i.Kelurahan.Nama,
		},

		DesilKesejahteraan: i.DesilKesejahteraan,
		Alamat:             i.Alamat,
		IdIndividu:         i.IdIndividu,
		Nama:               i.Nama,
		Nik:                i.Nik,
		PadanDukcapil:      i.PadanDukcapil,
		JenisKelamin:       i.JenisKelamin,
		Hubungan:           i.Hubungan,
		TanggalLahir:       i.TanggalLahir,
		StatusKawin:        i.StatusKawin,
		Pekerjaan:          i.Pekerjaan,
		Pendidikan:         i.Pendidikan,
		KategoriUsia:       i.KategoriUsia,
		PenerimaBPNT:       i.PenerimaBPNT,
		PenerimaBPUM:       i.PenerimaBPUM,
		PenerimaBST:        i.PenerimaBST,
		PenerimaPKH:        i.PenerimaPKH,
		PenerimaSembako:    i.PenerimaSembako,

		UserId: i.UserId,
		User: responses.UserResponse{
			Id:        i.User.Id,
			Username:  i.User.Username,
			Password:  i.User.Password,
			Email:     i.User.Email,
			NoHp:      i.User.NoHp,
			Role:      i.User.Role,
			CreatedAt: i.User.CreatedAt,
			UpdatedAt: i.User.UpdatedAt,
		},

		MahasiswaId: i.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:              i.Mahasiswa.Id,
			UserId:          i.Mahasiswa.UserId,
			NamaLengkap:     i.Mahasiswa.NamaLengkap,
			Universitas:     i.Mahasiswa.Universitas,
			JenisKelamin:    i.Mahasiswa.JenisKelamin,
			TanggalLahir:    i.Mahasiswa.TanggalLahir,
			KabupatenKotaId: i.Mahasiswa.KabupatenKotaId,
			KecamatanId:     i.Mahasiswa.KecamatanId,
			KelurahanId:     i.Mahasiswa.KelurahanId,
			CreatedAt:       i.Mahasiswa.CreatedAt,
			UpdatedAt:       i.Mahasiswa.UpdatedAt,
		},

		StatusVerifikasi: i.StatusVerifikasi,
	}
}

func ConvertToIndividuVerifikasiResponse(i model.IndividuVerifikasi) responses.IndividuVerifikasiResponse {
	return responses.IndividuVerifikasiResponse{
		Id:         i.Id,
		IdKeluarga: i.IdKeluarga,

		ProvinsiId: i.ProvinsiId,
		Provinsi: responses.ProvinsiResponse{
			Id:   i.Provinsi.Id,
			Nama: i.Provinsi.Nama,
		},

		KabupatenKotaId: i.KabupatenKotaId,
		KabupatenKota: responses.Kabupaten_KotaResponse{
			Id:         i.KabupatenKota.Id,
			ProvinsiId: i.KabupatenKota.ProvinsiId,
			Nama:       i.KabupatenKota.Nama,
		},

		KecamatanId: i.KecamatanId,
		Kecamatan: responses.KecamatanResponse{
			Id:              i.Kecamatan.Id,
			KabupatenKotaId: i.Kecamatan.KabupatenKotaId,
			Nama:            i.Kecamatan.Nama,
		},

		KelurahanId: i.KelurahanId,
		Kelurahan: responses.KelurahanResponse{
			Id:          i.Kelurahan.Id,
			KecamatanId: i.Kelurahan.KecamatanId,
			Nama:        i.Kelurahan.Nama,
		},

		DesilKesejahteraan: i.DesilKesejahteraan,
		Alamat:             i.Alamat,
		IdIndividu:         i.IdIndividu,
		Nama:               i.Nama,
		Nik:                i.Nik,
		PadanDukcapil:      i.PadanDukcapil,
		JenisKelamin:       i.JenisKelamin,
		Hubungan:           i.Hubungan,
		TanggalLahir:       i.TanggalLahir,
		StatusKawin:        i.StatusKawin,
		Pekerjaan:          i.Pekerjaan,
		Pendidikan:         i.Pendidikan,
		KategoriUsia:       i.KategoriUsia,
		PenerimaBPNT:       i.PenerimaBPNT,
		PenerimaBPUM:       i.PenerimaBPUM,
		PenerimaBST:        i.PenerimaBST,
		PenerimaPKH:        i.PenerimaPKH,
		PenerimaSembako:    i.PenerimaSembako,
		PenerimaLainnya:    i.PenerimaLainnya,
		StatusResponden:    i.StatusResponden,

		UserId: i.UserId,
		User: responses.UserResponse{
			Id:        i.User.Id,
			Username:  i.User.Username,
			Password:  i.User.Password,
			Email:     i.User.Email,
			NoHp:      i.User.NoHp,
			Role:      i.User.Role,
			CreatedAt: i.User.CreatedAt,
			UpdatedAt: i.User.UpdatedAt,
		},

		MahasiswaId: i.MahasiswaId,
		Mahasiswa: responses.MahasiswaResponse{
			Id:              i.Mahasiswa.Id,
			UserId:          i.Mahasiswa.UserId,
			NamaLengkap:     i.Mahasiswa.NamaLengkap,
			Universitas:     i.Mahasiswa.Universitas,
			JenisKelamin:    i.Mahasiswa.JenisKelamin,
			TanggalLahir:    i.Mahasiswa.TanggalLahir,
			KabupatenKotaId: i.Mahasiswa.KabupatenKotaId,
			KecamatanId:     i.Mahasiswa.KecamatanId,
			KelurahanId:     i.Mahasiswa.KelurahanId,
			CreatedAt:       i.Mahasiswa.CreatedAt,
			UpdatedAt:       i.Mahasiswa.UpdatedAt,
		},
		UrlBukti: i.UrlBukti,

		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

func ConvertToTagIndividuResponse(t model.TagIndividu) responses.TagIndividuResponse {
	return responses.TagIndividuResponse{
		Id:             t.Id,
		FokusBelanjaId: t.FokusBelanjaId,
		FokusBelanja: responses.FokusBelanjaResponse{
			Id:                   t.FokusBelanja.Id,
			RencanaSubKegiatanId: t.FokusBelanja.RencanaSubKegiatanId,
			NamaFokusBelanja:     t.FokusBelanja.NamaFokusBelanja,
			Indikator:            t.FokusBelanja.Indikator,
			Target:               t.FokusBelanja.Target,
			Satuan:               t.FokusBelanja.Satuan,
			PaguFokusBelanja:     t.FokusBelanja.PaguFokusBelanja,
			Keterangan:           t.FokusBelanja.Keterangan,
			Tahun:                t.FokusBelanja.Tahun,
			CreatedAt:            t.FokusBelanja.CreatedAt,
			UpdatedAt:            t.FokusBelanja.UpdatedAt,
		},
		IndividuId: t.IndividuId,
		Individu: responses.IndividuResponse{
			Id:         t.Individu.Id,
			IdKeluarga: t.Individu.IdKeluarga,

			ProvinsiId: t.Individu.ProvinsiId,
			Provinsi: responses.ProvinsiResponse{
				Id:   t.Individu.Provinsi.Id,
				Nama: t.Individu.Provinsi.Nama,
			},

			KabupatenKotaId: t.Individu.KabupatenKotaId,
			Kabupaten_Kota: responses.Kabupaten_KotaResponse{
				Id:         t.Individu.KabupatenKota.Id,
				ProvinsiId: t.Individu.KabupatenKota.ProvinsiId,
				Nama:       t.Individu.KabupatenKota.Nama,
			},

			KecamatanId: t.Individu.KecamatanId,
			Kecamatan: responses.KecamatanResponse{
				Id:              t.Individu.Kecamatan.Id,
				KabupatenKotaId: t.Individu.Kecamatan.KabupatenKotaId,
				Nama:            t.Individu.Kecamatan.Nama,
			},

			KelurahanId: t.Individu.KelurahanId,
			Kelurahan: responses.KelurahanResponse{
				Id:          t.Individu.Kelurahan.Id,
				KecamatanId: t.Individu.Kelurahan.KecamatanId,
				Nama:        t.Individu.Kelurahan.Nama,
			},

			DesilKesejahteraan: t.Individu.DesilKesejahteraan,
			Alamat:             t.Individu.Alamat,
			IdIndividu:         t.Individu.IdIndividu,
			Nama:               t.Individu.Nama,
			Nik:                t.Individu.Nik,
			PadanDukcapil:      t.Individu.PadanDukcapil,
			JenisKelamin:       t.Individu.JenisKelamin,
			Hubungan:           t.Individu.Hubungan,
			TanggalLahir:       t.Individu.TanggalLahir,
			StatusKawin:        t.Individu.StatusKawin,
			Pekerjaan:          t.Individu.Pekerjaan,
			Pendidikan:         t.Individu.Pendidikan,
			KategoriUsia:       t.Individu.KategoriUsia,
			PenerimaBPNT:       t.Individu.PenerimaBPNT,
			PenerimaBPUM:       t.Individu.PenerimaBPUM,
			PenerimaBST:        t.Individu.PenerimaBST,
			PenerimaPKH:        t.Individu.PenerimaPKH,
			PenerimaSembako:    t.Individu.PenerimaSembako,

			UserId: t.Individu.UserId,
			User: responses.UserResponse{
				Id:        t.Individu.User.Id,
				Username:  t.Individu.User.Username,
				Password:  t.Individu.User.Password,
				Email:     t.Individu.User.Email,
				NoHp:      t.Individu.User.NoHp,
				Role:      t.Individu.User.Role,
				CreatedAt: t.Individu.User.CreatedAt,
				UpdatedAt: t.Individu.User.UpdatedAt,
			},

			MahasiswaId: t.Individu.MahasiswaId,
			Mahasiswa: responses.MahasiswaResponse{
				Id:              t.Individu.Mahasiswa.Id,
				UserId:          t.Individu.Mahasiswa.UserId,
				NamaLengkap:     t.Individu.Mahasiswa.NamaLengkap,
				Universitas:     t.Individu.Mahasiswa.Universitas,
				JenisKelamin:    t.Individu.Mahasiswa.JenisKelamin,
				TanggalLahir:    t.Individu.Mahasiswa.TanggalLahir,
				KabupatenKotaId: t.Individu.Mahasiswa.KabupatenKotaId,
				KecamatanId:     t.Individu.Mahasiswa.KecamatanId,
				KelurahanId:     t.Individu.Mahasiswa.KelurahanId,
				CreatedAt:       t.Individu.Mahasiswa.CreatedAt,
				UpdatedAt:       t.Individu.Mahasiswa.UpdatedAt,
			},

			StatusVerifikasi: t.Individu.StatusVerifikasi,
		},

		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func ConvertToTagKeluargaResponse(t model.TagKeluarga) responses.TagKeluargaResponse {
	return responses.TagKeluargaResponse{
		Id:             t.Id,
		FokusBelanjaId: t.FokusBelanjaId,
		FokusBelanja: responses.FokusBelanjaResponse{
			Id:                   t.FokusBelanja.Id,
			RencanaSubKegiatanId: t.FokusBelanja.RencanaSubKegiatanId,
			NamaFokusBelanja:     t.FokusBelanja.NamaFokusBelanja,
			Indikator:            t.FokusBelanja.Indikator,
			Target:               t.FokusBelanja.Target,
			Satuan:               t.FokusBelanja.Satuan,
			PaguFokusBelanja:     t.FokusBelanja.PaguFokusBelanja,
			Keterangan:           t.FokusBelanja.Keterangan,
			Tahun:                t.FokusBelanja.Tahun,
			CreatedAt:            t.FokusBelanja.CreatedAt,
			UpdatedAt:            t.FokusBelanja.UpdatedAt,
		},
		KeluargaId: t.KeluargaId,
		Keluarga: responses.KeluargaResponse{
			Id:         t.Keluarga.Id,
			IdKeluarga: t.Keluarga.IdKeluarga,
			ProvinsiId: t.Keluarga.ProvinsiId,
			Provinsi: responses.ProvinsiResponse{
				Id:   t.Keluarga.Provinsi.Id,
				Nama: t.Keluarga.Provinsi.Nama,
			},

			KabupatenKotaId: t.Keluarga.KabupatenKotaId,
			KabupatenKota: responses.Kabupaten_KotaResponse{
				Id:         t.Keluarga.KabupatenKota.Id,
				ProvinsiId: t.Keluarga.KabupatenKota.ProvinsiId,
				Nama:       t.Keluarga.KabupatenKota.Nama,
			},

			KecamatanId: t.Keluarga.KecamatanId,
			Kecamatan: responses.KecamatanResponse{
				Id:              t.Keluarga.Kecamatan.Id,
				KabupatenKotaId: t.Keluarga.Kecamatan.KabupatenKotaId,
				Nama:            t.Keluarga.Kecamatan.Nama,
			},

			KelurahanId: t.Keluarga.KelurahanId,
			Kelurahan: responses.KelurahanResponse{
				Id:          t.Keluarga.Kelurahan.Id,
				KecamatanId: t.Keluarga.Kelurahan.KecamatanId,
				Nama:        t.Keluarga.Kelurahan.Nama,
			},

			DesilKesejahteraan:     t.Keluarga.DesilKesejahteraan,
			Alamat:                 t.Keluarga.Alamat,
			KepalaKeluarga:         t.Keluarga.KepalaKeluarga,
			Nik:                    t.Keluarga.Nik,
			PadanDukcapil:          t.Keluarga.PadanDukcapil,
			JenisKelamin:           t.Keluarga.JenisKelamin,
			TanggalLahir:           t.Keluarga.TanggalLahir,
			Pekerjaan:              t.Keluarga.Pekerjaan,
			Pendidikan:             t.Keluarga.Pendidikan,
			KepemilikanRumah:       t.Keluarga.KepemilikanRumah,
			Simpanan:               t.Keluarga.Simpanan,
			JenisAtap:              t.Keluarga.JenisAtap,
			JenisDinding:           t.Keluarga.JenisDinding,
			JenisLantai:            t.Keluarga.JenisLantai,
			SumberPenerangan:       t.Keluarga.SumberPenerangan,
			BahanBakarMemasak:      t.Keluarga.BahanBakarMemasak,
			SumberAirMinum:         t.Keluarga.SumberAirMinum,
			FasilitasBuangAirBesar: t.Keluarga.FasilitasBuangAirBesar,
			PenerimaBPNT:           t.Keluarga.PenerimaBPNT,
			PenerimaBPUM:           t.Keluarga.PenerimaBPUM,
			PenerimaBST:            t.Keluarga.PenerimaBST,
			PenerimaPKH:            t.Keluarga.PenerimaPKH,
			PenerimaSembako:        t.Keluarga.PenerimaSembako,

			UserId: t.Keluarga.UserId,
			User: responses.UserResponse{
				Id:        t.Keluarga.User.Id,
				Username:  t.Keluarga.User.Username,
				Password:  t.Keluarga.User.Password,
				Email:     t.Keluarga.User.Email,
				NoHp:      t.Keluarga.User.NoHp,
				Role:      t.Keluarga.User.Role,
				CreatedAt: t.Keluarga.User.CreatedAt,
				UpdatedAt: t.Keluarga.User.UpdatedAt,
			},

			MahasiswaId: t.Keluarga.MahasiswaId,
			Mahasiswa: responses.MahasiswaResponse{
				Id:           t.Keluarga.Mahasiswa.Id,
				UserId:       t.Keluarga.Mahasiswa.UserId,
				NamaLengkap:  t.Keluarga.Mahasiswa.NamaLengkap,
				Universitas:  t.Keluarga.Mahasiswa.Universitas,
				JenisKelamin: t.Keluarga.Mahasiswa.JenisKelamin,
				TanggalLahir: t.Keluarga.Mahasiswa.TanggalLahir,
				KelurahanId:  t.Keluarga.Mahasiswa.KelurahanId,
				CreatedAt:    t.Keluarga.Mahasiswa.CreatedAt,
				UpdatedAt:    t.Keluarga.Mahasiswa.UpdatedAt,
			},

			StatusVerifikasi: t.Keluarga.StatusVerifikasi,
		},

		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func ConvertToStatistikProgramInstansiResponse(i model.Instansi, hasil int64) responses.StatistikProgramInstansiResponse {
	return responses.StatistikProgramInstansiResponse{
		InstansiId:    i.Id,
		NamaInstansi:  i.NamaInstansi,
		JumlahProgram: hasil,
	}
}

func ConvertToStatistikKegiatanInstansiResponse(i model.Instansi, hasil int64) responses.StatistikKegiatanInstansiResponse {
	return responses.StatistikKegiatanInstansiResponse{
		InstansiId:     i.Id,
		NamaInstansi:   i.NamaInstansi,
		JumlahKegiatan: hasil,
	}
}

func ConvertToStatistikSubKegiatanInstansiResponse(i model.Instansi, hasil int64) responses.StatistikSubKegiatanInstansiResponse {
	return responses.StatistikSubKegiatanInstansiResponse{
		InstansiId:        i.Id,
		NamaInstansi:      i.NamaInstansi,
		JumlahSubKegiatan: hasil,
	}
}

func ConvertToStatistikRencanaProgramInstansiResponse(i model.Instansi, hasil int64) responses.StatistikRencanaProgramInstansiResponse {
	return responses.StatistikRencanaProgramInstansiResponse{
		InstansiId:           i.Id,
		NamaInstansi:         i.NamaInstansi,
		JumlahRencanaProgram: hasil,
	}
}

func ConvertToStatistikRencanaKegiatanInstansiResponse(i model.Instansi, hasil int64) responses.StatistikRencanaKegiatanInstansiResponse {
	return responses.StatistikRencanaKegiatanInstansiResponse{
		InstansiId:            i.Id,
		NamaInstansi:          i.NamaInstansi,
		JumlahRencanaKegiatan: hasil,
	}
}

func ConvertToStatistikRencanaSubKegiatanInstansiResponse(i model.Instansi, hasil int64) responses.StatistikRencanaSubKegiatanInstansiResponse {
	return responses.StatistikRencanaSubKegiatanInstansiResponse{
		InstansiId:               i.Id,
		NamaInstansi:             i.NamaInstansi,
		JumlahRencanaSubKegiatan: hasil,
	}
}

func ConvertToStatistikFokusBelanjaInstansiResponse(i model.Instansi, hasil int64) responses.StatistikFokusBelanjaInstansiResponse {
	return responses.StatistikFokusBelanjaInstansiResponse{
		InstansiId:         i.Id,
		NamaInstansi:       i.NamaInstansi,
		JumlahFokusBelanja: hasil,
	}
}

func ConvertToStatistikSumPaguRencanaProgramResponse(i model.Instansi, jumlahRencanaProgram, totalPaguRencanaProgram int64) responses.StatistikSumPaguRencanaProgramResponse {
	return responses.StatistikSumPaguRencanaProgramResponse{
		InstansiId:              i.Id,
		NamaInstansi:            i.NamaInstansi,
		JumlahRencanaProgram:    jumlahRencanaProgram,
		TotalPaguRencanaProgram: totalPaguRencanaProgram,
	}
}

func ConvertToStatistikSumPaguRencanaKegiatanResponse(i model.Instansi, jumlahRencanaKegiatan, totalPaguRencanaKegiatan int64) responses.StatistikSumPaguRencanaKegiatanResponse {
	return responses.StatistikSumPaguRencanaKegiatanResponse{
		InstansiId:               i.Id,
		NamaInstansi:             i.NamaInstansi,
		JumlahRencanaKegiatan:    jumlahRencanaKegiatan,
		TotalPaguRencanaKegiatan: totalPaguRencanaKegiatan,
	}
}

func ConvertToStatistikSumPaguRencanaSubKegiatanResponse(i model.Instansi, jumlahRencanaSubKegiatan, totalPaguRencanaSubKegiatan int64) responses.StatistikSumPaguRencanaSubKegiatanResponse {
	return responses.StatistikSumPaguRencanaSubKegiatanResponse{
		InstansiId:                  i.Id,
		NamaInstansi:                i.NamaInstansi,
		JumlahRencanaSubKegiatan:    jumlahRencanaSubKegiatan,
		TotalPaguRencanaSubKegiatan: totalPaguRencanaSubKegiatan,
	}
}

func ConvertToStatistikSumPaguFokusBelanjaResponse(i model.Instansi, jumlahFokusBelanja, totalPaguFokusBelanja int64) responses.StatistikSumPaguFokusBelanjaResponse {
	return responses.StatistikSumPaguFokusBelanjaResponse{
		InstansiId:            i.Id,
		NamaInstansi:          i.NamaInstansi,
		JumlahFokusBelanja:    jumlahFokusBelanja,
		TotalPaguFokusBelanja: totalPaguFokusBelanja,
	}
}
