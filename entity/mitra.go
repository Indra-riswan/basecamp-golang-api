package entity

type Mitra struct {
	ID        uint        `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama      string      `gorm:"type:char(50)" json:"nama"`
	Penjualan []Transaksi `gorm:"foreignkey:Mitra;references:Nama;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"penjualan"`
}
