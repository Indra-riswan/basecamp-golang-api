package dto

type BrandDto struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Website   string `json:"website" form:"password" binding:"required"`
	Marketing string `json:"marketing" form:"marketing" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Nohp      uint   `json:"nohp" form:"nohp" binding:"required,number"`
	Gambar    string `json:"gambar" form:"gambar" binding:"required"`
}

type BrandDtoUpdate struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Website   string `json:"website" form:"password" binding:"required"`
	Marketing string `json:"marketing" form:"marketing" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Nohp      uint   `json:"nohp" form:"nohp" binding:"required,number"`
	Gambar    string `json:"gambar" form:"gambar" binding:"required"`
}

type BrandDtoLogin struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
