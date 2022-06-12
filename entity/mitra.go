package entity

type Mitra struct {
	ID        uint        `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama      string      `gorm:"type:char(50)" json:"nama"`
	Unit      string      `gorm:"type:char(10)" json:"unit"`
	Nopol     string      `gorm:"type:char(10)" json:"nopol"`
	Penjualan []Transaksi `gorm:"foreignkey:Nopolunit;references:Nopol;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"penjualan"`
}
