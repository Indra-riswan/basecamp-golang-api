package entity

type Brand struct {
	ID        uint        `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama      string      `gorm:"type:char(50)" json:"nama"`
	Marketing string      `gorm:"type:char(50)" json:"marketing"`
	Penjualan []Transaksi `gorm:"foreignKey:Brand;references:Nama;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"penjualan"`
}
