package model

type Pasien struct {
	Id int `json:"id" gorm:"column:id;primaryKey;type:integer;not null"`
	Kode string `json:"kode_pasien" gorm:"column:kode_pasien;type:varchar(10);not null"`
	Nama string `json:"nama_pasien" gorm:"column:nama_pasien;type:varchar(100);not null"`
	Deskripsi string `json:"deskripsi" gorm:"column:deskripsi;type:text;"`
	JumlahKunjungan int `json:"jumlah_kunjungan" gorm:"column:jumlah_kunjungan;type:integer;not null"`
	TanggalDaftar string `json:"tanggal_daftar" gorm:"column:tanggal_daftar;type:date;not null"`
}

func (Pasien) TableName() string {
	return "pasien"
}