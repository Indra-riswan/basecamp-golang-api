package dto

type TransaksiDto struct {
	Brand     string `json:"brand" form:"brand" binding:"required"`
	Nopolunit string `json:"nopolunit" form:"nopolunit" binding:"required"`
	User      string `json:"user" form:"user" binding:"required"`
	Checker   string `json:"checker" form:"checker" binding:"required"`
	Status    string `json:"status" form:"status" binding:"required"`
	Driver    string `json:"driver,omitempty" form:"driver,omitempty"`
	Out       string `json:"out" form:"out" binding:"required"`
	In        string `json:"in" form:"in" binding:"required"`
	Day       uint   `json:"day" form:"day" binding:"required"`
	Perhari   uint   `json:"perhari" form:"perhari" binding:"required"`
	Op        uint   `json:"op" form:"op" binding:"required"`
}

type TransaksiDtoUpdate struct {
	Brand     string `json:"brand" form:"brand" binding:"required"`
	Nopolunit string `json:"nopolunit" form:"nopolunit" binding:"required"`
	User      string `json:"user" form:"user" binding:"required"`
	Checker   string `json:"checker" form:"checker" binding:"required"`
	Status    string `json:"status" form:"status" binding:"required"`
	Driver    string `json:"driver,omitempty" form:"driver,omitempty"`
	Out       string `json:"out" form:"out" binding:"required"`
	In        string `json:"in" form:"in" binding:"required"`
	Day       uint   `json:"day" form:"day" binding:"required"`
	Perhari   uint   `json:"perhari" form:"perhari" binding:"required"`
	Op        uint   `json:"op" form:"op" binding:"required"`
}
