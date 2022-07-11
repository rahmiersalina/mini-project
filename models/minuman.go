package models

type Minuman struct {
	Id_minuman   string `gorm:"column:id_minuman" json:"id_minuman"`
	Nama_minuman string `gorm:"column:nama_minuman" json:"nama_minuman"`
	Jenis        string `gorm:"column:jenis" json:"jenis"`
	Harga        int    `gorm:"column:jarga" json:"harga"`
}

type InputNewMinuman struct {
	NamaMinuman string `json:"nama_minuman"`
	Jenis       string `json:"jenis"`
	Harga       int    `json:"harga"`
}
