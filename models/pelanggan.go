package models


type pelanggan struct {
	Id_pelanggan   string `json:"id_pelanggan"`
	Nama_pelanggan string `json:"nama_pelanggan"`
	Alamat         string `json:"alamat"`
	No_telepon     string `json:"no_telepon"`
}

func (pelanggan) tablename() string {

	return "pelanggan"

}
