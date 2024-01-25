package apiobat

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mconn = MongoConnect("MONGOSTRING", "psikofarmaka")

func TestRegister(t *testing.T) {
	var data User
	data.Username = "fhulan"
	data.Password = "secret"

	err := Register(mconn, "user", data)
	if err != nil {
		t.Errorf("Error registering user: %v", err)
	} else {
		fmt.Println("Register success", data.Username)
	}
}

func TestLogIn(t *testing.T) {
	var data User
	data.Username = "fhulan"
	data.Password = "secret"

	user, status, err := Login(mconn, "user", data)
	fmt.Println("Status", status)
	if err != nil {
		t.Errorf("Error logging in user: %v", err)
	} else {
		fmt.Println("Login success", user)
	}
}

// obat
func TestGetAllObat(t *testing.T) {
	var data []Obat

	data, err := GetAllObat(mconn, "obat")
	if err != nil {
		t.Errorf("Error getting all obat: %v", err)
	} else {
		fmt.Println("Get all obat success", data)
	}
}

func TestGetObatByID(t *testing.T) {
	var data Obat

	id := "6592618179e4a43f5118d5ec"

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("get obat by id error")
		return
	} else {

		data, err = GetObatByID(mconn, "obat", ID)
		if err != nil {
			t.Errorf("Error getting obat by id: %v", err)
		} else {
			fmt.Println("Get obat by id success", data)
		}
	}
}

// func TestInsertObat(t *testing.T) {
// 	var data Obat
// 	data.Jenis_Obat = "Obat Anti-Psikosis"
// 	data.Nama_Obat = "chlorpromazine"
// 	data.Deskripsi = "Deskripsi Obat"

// 	_, err := InsertObat(mconn, "obat", r)
// 	if err != nil {
// 		t.Errorf("Error inserting obat: %v", err)
// 	} else {
// 		fmt.Println("Insert obat success", data)
// 	}
// }

// func TestUpdateObat(t *testing.T) {
// 	var data Obat
// 	data.Jenis_Obat = "Obat Anti-Psikosis"
// 	data.Nama_Obat = "Chlorpromazine"
// 	data.Deskripsi = "obat untuk menangani gejala psikosis, seperti halusinasi, dan pikiran tidak wajar. Obat ini bekerja dengan menyeimbangkan zat kimia di otak yang umumnya tidak normal pada penderita gangguan jiwa. Obat ini dapat mengurangi halusinasi, serta membantu pasien berpikir lebih jernih dan menjadi tidak agresif sehingga ia bisa melakukan aktivitas sehari-hari."

// 	id := "6592618179e4a43f5118d5ec"

// 	ID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		fmt.Printf("update obat error")
// 		return
// 	} else {

// 		status, err := UpdateObat(mconn, "obat", ID, data)
// 		fmt.Println("Status", status)
// 		if err != nil {
// 			t.Errorf("Error updating obat: %v", err)
// 		} else {
// 			fmt.Println("Update obat success", data)
// 		}
// 	}
// }

func TestDeleteObat(t *testing.T) {
	id := "659262931be188bed64b2600"

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("delete obat error")
		return
	} else {

		status, err := DeleteObat(mconn, "obat", ID)
		fmt.Println("Status", status)
		if err != nil {
			t.Errorf("Error deleting obat: %v", err)
		} else {
			fmt.Println("Delete obat success")
		}
	}
}

// penyakit
func TestGetAllPenyakit(t *testing.T) {
	var data []Penyakit

	data, err := GetAllPenyakit(mconn, "penyakit")
	if err != nil {
		t.Errorf("Error getting all penyakit: %v", err)
	} else {
		fmt.Println("Get all penyakit success", data)
	}
}

func TestGetPenyakitByID(t *testing.T) {
	var data Penyakit

	id := "6592618179e4a43f5118d5ec"

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("get penyakit by id error")
		return
	} else {

		data, err = GetPenyakitByID(mconn, "penyakit", ID)
		if err != nil {
			t.Errorf("Error getting penyakit by id: %v", err)
		} else {
			fmt.Println("Get penyakit by id success", data)
		}
	}
}

func TestInsertPenyakit(t *testing.T) {
	var data Penyakit
	data.Jenis_Penyakit = "Kejiwaan"
	data.Nama_Penyakit = "Skizofrenia"
	data.Deskripsi = "Deskripsi Penyakit"
	data.Obat.Nama_Obat = "Chlorpromazine"

	_, err := InsertPenyakit(mconn, "penyakit", data)
	if err != nil {
		t.Errorf("Error inserting penyakit: %v", err)
	} else {
		fmt.Println("Insert penyakit success", data)
	}
}

func TestUpdatePenyakit(t *testing.T) {
	var data Penyakit
	data.Jenis_Penyakit = "Kejiwaan"
	data.Nama_Penyakit = "Skizofrenia"
	data.Deskripsi = "gangguan mental berat yang dapat memengaruhi tingkah laku, emosi, dan komunikasi. Penderita skizofrenia bisa mengalami halusinasi, delusi, kekacauan berpikir, dan perubahan perilaku."
	data.Obat.Nama_Obat = "Chlorpromazine"

	id := "6592646ddff04a4ce4aa716e"

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("update penyakit error")
		return
	} else {

		status, err := UpdatePenyakit(mconn, "penyakit", ID, data)
		fmt.Println("Status", status)
		if err != nil {
			t.Errorf("Error updating penyakit: %v", err)
		} else {
			fmt.Println("Update penyakit success", data)
		}
	}
}

func TestDeletePenyakit(t *testing.T) {
	id := "659263df1db7d0aa75bc9b7d"

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Delete penyakit error")
		return
	} else {

		status, err := DeletePenyakit(mconn, "penyakit", ID)
		fmt.Println("Status", status)
		if err != nil {
			t.Errorf("Error deleting penyakit: %v", err)
		} else {
			fmt.Println("Delete penyakit success")
		}
	}
}
