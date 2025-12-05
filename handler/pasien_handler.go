package handler

import (
	"go-uts-pasien-klinik/model"
	"go-uts-pasien-klinik/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllPasien(c *fiber.Ctx) error {
	data, err := repository.GetAllPasien()

	if (err != nil) {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mengambil data.",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data pasien.",
		"data": data,
	})
}

func InsertPasien(c *fiber.Ctx) error {
	var pasien model.Pasien

	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data yang dikirim salah.",
			"error": err.Error(),
		})
	}

	if errs := validatePasienInput(pasien); len(errs) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validasi gagal.",
			"errors":  errs,
		})
	}

	if err := repository.InsertPasien(pasien); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menambahkan data baru.",
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Data pasien berhasil ditambahkan.",
		"data": pasien,
	})
}

func GetPasienById(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := repository.GetPasienById(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Pasien dengan ID tersebut tidak ditemukan",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data pasien berdasarkan ID.",
		"data": data,
	})
}

func ReplacePasienById(c *fiber.Ctx) error {
	id := c.Params("id")

	var pasien model.Pasien
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah.",
			"error":   err.Error(),
		})
	}

	if errs := validatePasienInput(pasien); len(errs) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validasi gagal.",
			"errors":  errs,
		})
	}

	if err := repository.ReplacePasienById(id, pasien); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Pasien tidak ditemukan.",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data pasien berhasil diganti.",
		"data":    pasien,
	})
}


func UpdatePasienById(c *fiber.Ctx) error {
	id := c.Params("id")

	var updateData map[string]interface{}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data yang dikirimkan salah.",
			"error":   err.Error(),
		})
	}

	if err := repository.UpdatePasienById(id, updateData); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Pasien dengan ID tersebut tidak ditemukan.",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data pasien berhasil diubah.",
		"data":    updateData,
	})
}


func DeletePasienById(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repository.DeletePasienById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Gagal menghapus data pasien.",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data pasien berhasil dihapus.",
	})
}



// extensions
func validatePasienInput(p model.Pasien) map[string]string {
	errors := make(map[string]string)

	if p.Kode == "" {
		errors["kode_pasien"] = "Kode pasien tidak boleh kosong."
	}
	if p.Nama == "" {
		errors["nama_pasien"] = "Nama pasien tidak boleh kosong."
	}
	if p.TanggalDaftar == "" {
		errors["tanggal_daftar"] = "Tanggal daftar tidak boleh kosong."
	}
	if p.JumlahKunjungan < 0 {
		errors["jumlah_kunjungan"] = "Jumlah kunjungan tidak boleh negatif."
	}

	return errors
}