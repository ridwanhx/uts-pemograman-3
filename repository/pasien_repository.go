package repository

import (
	"errors"
	"go-uts-pasien-klinik/config"
	"go-uts-pasien-klinik/model"
)

func GetAllPasien() ([]model.Pasien, error) {
	var data []model.Pasien;
	result := config.GetDB().Find(&data)

	return data, result.Error;
}

func InsertPasien(mhs model.Pasien) error {
	result := config.GetDB().Create(&mhs)
	return result.Error
}

func GetPasienById(id string) ([]model.Pasien, error) {
	var data []model.Pasien
	result := config.GetDB().Where("id = ?", id).First(&data)

	return data, result.Error
}

func ReplacePasienById(id string, pasien model.Pasien) error {
	result := config.GetDB().Model(&model.Pasien{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"kode_pasien":      pasien.Kode,
			"nama_pasien":      pasien.Nama,
			"deskripsi":        pasien.Deskripsi,
			"jumlah_kunjungan": pasien.JumlahKunjungan,
			"tanggal_daftar":   pasien.TanggalDaftar,
		})

	// Jika tidak ada row yang diupdate → ID tidak ditemukan
	if result.RowsAffected == 0 {
		return errors.New("pasien tidak ditemukan")
	}

	return result.Error
}



func UpdatePasienById(id string, updateData map[string]interface{}) error {
	result := config.GetDB().Model(&model.Pasien{}).Where("id = ?", id).Updates(updateData)

	if result.RowsAffected == 0 {
		return errors.New("pasien tidak ditemukan")
	}

	return result.Error
}


func DeletePasienById(id string) error {
	result := config.GetDB().Where("id = ?", id).Delete(&model.Pasien{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("pasien tidak ditemukan")
	}

	return nil
}
