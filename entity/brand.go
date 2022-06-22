package entity

type Brand struct {
	ID        uint        `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama      string      `gorm:"type:char(50);uniqueIndex" json:"nama"`
	Password  string      `gorm:"->;<-;not null" json:"-"`
	Marketing string      `gorm:"type:char(50)" json:"marketing"`
	Email     string      `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Nohp      uint        `gorm:"type:int" json:"nohp"`
	Gambar    string      `gorm:"type:varchar(100)" json:"gambar"`
	Penjualan []Transaksi `gorm:"foreignKey:Brand;references:Nama;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"penjualan"`
}
