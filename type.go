package apiobat

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username         string             `bson:"username," json:"username,"`
	Password         string             `bson:"password," json:"password,"`
	Confirm_Password string             `bson:"confirm_password," json:"confirm_password,"`
}

// tambah gambar
type Obat struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis_Obat string             `bson:"jenis_obat,omitempty" json:"jenis_obat,omitempty"`
	Nama_Obat  string             `bson:"nama_obat,omitempty" json:"nama_obat,omitempty"`
	Deskripsi  string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Gambar     string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
}

type Penyakit struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis_Penyakit string             `bson:"jenis_penyakit,omitempty" json:"jenis_penyakit,omitempty"`
	Nama_Penyakit  string             `bson:"nama_penyakit,omitempty" json:"nama_penyakit,omitempty"`
	Deskripsi      string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Obat           Obat               `bson:"obat,omitempty" json:"obat,omitempty"`
}

type RumahSakit struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_RS   string             `bson:"nama_rs,omitempty" json:"nama_rs,omitempty"`
	No_Telp   string             `bson:"no_telp,omitempty" json:"no_telp,omitempty"`
	Alamat    string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Latitude  string             `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude string             `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Gambar    string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
}

type Credential struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
