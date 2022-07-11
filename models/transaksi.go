package models


type transaksi struct {
	no_transaksi      string `json:"no_transaksi"`
	tanggal_transaksi string `json:"tanggal_transaksi"`
	id_pelanggan      string `json:"id_pelanggan"`
	nama_pelanggan    string `json:"nama_pelanggan"`
	id_kasir          string `json:"id_kasir"`
	subtotal_makanan  int    `json:"subtotal_makanan"`
	subtotal_minuman  int    `json:"subtotal_minuman"`
	total_bayar       int    `json:"total_bayar"`
}

func (transaksi) tablename() string {

	return "transaksi"
}
