package models

type Makanan struct {
	Id_makanan   string `gorm:"column:id_makanan" json:"id_makanan"`
	Nama_makanan string `gorm:"column:nama_makanan" json:"nama_makanan"`
	Jenis        string `gorm:"column:jenis" json:"jenis"`
	Harga        int    `gorm:"column:harga" json:"harga"`
}

type InputNewMakanan struct {
	NamaMakanan string `json:"nama_makanan"`
	Jenis       string `json:"jenis"`
	Harga       int    `json:"harga"`
}
