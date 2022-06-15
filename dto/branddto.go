package dto

type BrandDto struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Marketing string `json:"marketing" form:"marketing" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Nohp      uint   `json:"nohp" form:"nohp" binding:"required,number"`
}

type BrandDtoUpdate struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Marketing string `json:"marketing" form:"marketing" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Nohp      uint   `json:"nohp" form:"nohp" binding:"required,number"`
}
