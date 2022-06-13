package entity

import "time"

type Transaksi struct {
	ID         uint   `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Brand      string `gorm:"type:char(50)" json:"brand"`
	Mitra      string `gorm:"type:char(20)" json:"mitra"`
	Unit       string `gorm:"type:char(20)" json:"unit"`
	Nopol      string `gorm:"type:char(50)" json:"nopol"`
	User       string `gorm:"type:char(50)" json:"user"`
	Checker    string `gorm:"type:char(50)" json:"checker"`
	Status     string `gorm:"type:char(20)" json:"status"`
	Driver     string `gorm:"type:char(20)" json:"driver"`
	Out        string `gorm:"type:char(10)" json:"out"`
	In         string `gorm:"type:char(10)" json:"in"`
	Day        uint   `gorm:"type:int" json:"day"`
	Perhari    uint   `gorm:"type:int" json:"perhari"`
	Op         uint   `gorm:"type:int" json:"op"`
	Total      uint   `gorm:"not null" json:"total"`
	TotalOp    uint   `gorm:"not null" json:"total_op"`
	SisaProfit uint   `gorm:"not null" json:"sisa_profit"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
