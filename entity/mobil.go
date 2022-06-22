package entity

type Mobil struct {
	ID      uint   `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Unit    string `gorm:"type:char(50)" json:"unit"`
	Nopol   string `gorm:"type:char(20)" json:"nopol"`
	Gambar  string `gorm:"type:char(50)" json:"gambar"`
	Isready bool   `gorm:"type:bool;default:true" json:"is_ready"`
	Own     string `gorm:"not null" json:"-"`
	Mitra   Mitra  `gorm:"foreignKey:Own;references:Nama;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"mitra"`
}
