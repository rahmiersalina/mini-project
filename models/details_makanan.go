package models

type detail_transaksi_makanan struct {
	id_makanan       string `json:"id_makanan"`
	no_transaksi     string `json:"no_transaksi"`
	quantity_makanan int    `jsom:"quantity_makanan"`
	jumlah_makanan   int    `json:"jumlah_makanan"`
}

func (detail_transaksi_makanan) tablename() string {

	return "detail_transaksi_makanan"
}