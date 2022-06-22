package dto

type MobilDto struct {
	Unit   string `json:"unit" form:"unit" binding:"required"`
	Nopol  string `json:"nopol" form:"nopol" binding:"required"`
	Gambar string `json:"gambar" form:"gambar" binding:"required"`
	Own    string `json:"own" form:"own" binding:"required"`
}

type MobilDtoUpdate struct {
	Gambar  string `json:"gambar" form:"gambar" binding:"required"`
	Isready bool   `json:"isready" form:"isready"`
}
