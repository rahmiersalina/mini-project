package models

type detail_transaksi_minuman struct {
	id_minuman       string `json:"id_miniman"`
	no_transaksi     string `json:"no_transaksi"`
	quantity_minuman int    `jsom:"quantity_minuman"`
	jumlah_minuman   int    `json:"jumlah_minuman"`
}

func (detail_transaksi_minuman) tablename() string {

	return "detail_transaksi_minuman"
}
