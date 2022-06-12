package dto

type BrandDto struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Marketing string `json:"marketing" form:"marketing" binding:"required"`
}

type BrandDtoUpdate struct {
	Nama      string `json:"nama" form:"nama" binding:"required"`
	Marketing string `json:"marketing" form:"marketing" binding:"required"`
}
