package dto

type MitraDto struct {
	Nama string `json:"nama" form:"nama" binding:"required"`
}

type MitraDtoUpdate struct {
	Nama string `json:"nama" form:"nama" binding:"required"`
}
