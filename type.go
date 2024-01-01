package apiobat

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

type Obat struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis_Obat string             `bson:"jenis_obat,omitempty" json:"jenis_obat,omitempty"`
	Nama_Obat  string             `bson:"nama_obat,omitempty" json:"nama_obat,omitempty"`
	Deskripsi  string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
}

type Penyakit struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis_Penyakit string             `bson:"jenis_penyakit,omitempty" json:"jenis_penyakit,omitempty"`
	Nama_Penyakit  string             `bson:"nama_penyakit,omitempty" json:"nama_penyakit,omitempty"`
	Deskripsi      string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Obat           Obat               `bson:"obat,omitempty" json:"obat,omitempty"`
}

type ResponseUser struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ResponseObat struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Obat `json:"data"`
}

type ResponsePenyakit struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Penyakit `json:"data"`
}
