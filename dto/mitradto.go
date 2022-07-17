package dto

type MitraDto struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type MitraDtoUpdate struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type MitraDtoLogin struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
