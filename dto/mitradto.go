package dto

type MitraDto struct {
	Nama  string `json:"nama" form:"nama" binding:"required"`
	Unit  string `json:"unit" form:"unit" binding:"required"`
	Nopol string `json:"nopol" form:"nopol" binding:"required"`
}

type MitraDtoUpdate struct {
	Nama  string `json:"nama" form:"nama" binding:"required"`
	Unit  string `json:"unit" form:"unit" binding:"required"`
	Nopol string `json:"nopol" form:"nopol" binding:"required"`
}
