package models


type kasir struct {
	id_kasir   string `json:"id_kasir"`
	nama_kasir string `json:"nama_kasir"`
}

func tablename() string {

	return "kasir"
}