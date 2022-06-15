package entity

type Mitra struct {
	ID        uint        `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama      string      `gorm:"type:char(50);uniqueIndex" json:"nama"`
	Penjualan []Transaksi `gorm:"foreignKey:Mitra;references:Nama;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"penjualan"`
}
