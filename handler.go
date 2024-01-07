package apiobat

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	user             User
	obat             Obat
	penyakit         Penyakit
	responseUser     ResponseUser
	responseObat     ResponseObat
	responsePenyakit ResponsePenyakit
)

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func HandlerLogin(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responseUser.Status = 400

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseUser.Message = "error parsing application/json: " + err.Error()
	}

	users, _, err := Login(mconn, collectionname, user)
	if err != nil {
		responseUser.Message = err.Error()
		return GCFReturnStruct(responseUser)
	}

	responseUser.Status = 200
	responseUser.Message = "Selamat Datang " + users.Username

	return GCFReturnStruct(responseUser)
}

// obat
func HandlerGetAllObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responseObat.Status = 400

	_, err := GetAllObat(mconn, collectionname)
	if err != nil {
		responseObat.Message = err.Error()
		return GCFReturnStruct(responseObat)
	}

	responseObat.Status = 200
	responseObat.Message = "Get Obat Success"

	return GCFReturnStruct(responseObat)
}

func HandlerGetObatByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responseObat.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		responseObat.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(responseObat)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseObat.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(responseObat)
	}

	obats, err := GetObatByID(mconn, collectionname, ID)
	if err != nil {
		responseObat.Message = err.Error()
		return GCFReturnStruct(responseObat)
	}

	responseObat.Status = 200
	responseObat.Message = "Get Obat By ID Success  " + obats.Nama_Obat

	return GCFReturnStruct(responseObat)
}

func HandlerInsertObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responseObat.Status = 400

	err := json.NewDecoder(r.Body).Decode(&obat)
	if err != nil {
		responseObat.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(responseObat)
	}

	_, err = InsertObat(mconn, collectionname, obat)
	if err != nil {
		responseObat.Message = err.Error()
		return GCFReturnStruct(responseObat)
	}

	responseObat.Status = 200
	responseObat.Message = "Insert Obat Success " + obat.Nama_Obat

	return GCFReturnStruct(responseObat)
}

func HandlerUpdateObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responseObat.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		responseObat.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(responseObat)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseObat.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(responseObat)
	}

	err = json.NewDecoder(r.Body).Decode(&obat)
	if err != nil {
		responseObat.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(responseObat)
	}

	_, err = UpdateObat(mconn, collectionname, ID, obat)
	if err != nil {
		responseObat.Message = err.Error()
		return GCFReturnStruct(responseObat)
	}

	responseObat.Status = 200
	responseObat.Message = "Update Obat Success " + obat.Nama_Obat

	return GCFReturnStruct(responseObat)
}

func HandlerDeleteObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responseObat.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		responseObat.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(responseObat)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responseObat.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(responseObat)
	}

	_, err = DeleteObat(mconn, collectionname, ID)
	if err != nil {
		responseObat.Message = err.Error()
		return GCFReturnStruct(responseObat)
	}

	responseObat.Status = 200
	responseObat.Message = "Delete Obat Success " + obat.Nama_Obat

	return GCFReturnStruct(responseObat)
}

// penyakit
func HandlerGetAllPenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responsePenyakit.Status = 400

	_, err := GetAllPenyakit(mconn, collectionname)
	if err != nil {
		responsePenyakit.Message = err.Error()
		return GCFReturnStruct(responsePenyakit)
	}

	responsePenyakit.Status = 200
	responsePenyakit.Message = "Get Penyakit Success"

	return GCFReturnStruct(responsePenyakit)
}

func HandlerGetPenyakitByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responsePenyakit.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		responsePenyakit.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(responsePenyakit)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responsePenyakit.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(responsePenyakit)
	}

	penyakits, err := GetPenyakitByID(mconn, collectionname, ID)
	if err != nil {
		responsePenyakit.Message = err.Error()
		return GCFReturnStruct(responsePenyakit)
	}

	responsePenyakit.Status = 200
	responsePenyakit.Message = "Get Penyakit By ID Success " + penyakits.Nama_Penyakit

	return GCFReturnStruct(responsePenyakit)
}

func HandlerInsertPenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responsePenyakit.Status = 400

	err := json.NewDecoder(r.Body).Decode(&penyakit)
	if err != nil {
		responsePenyakit.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(responsePenyakit)
	}

	_, err = InsertPenyakit(mconn, collectionname, penyakit)
	if err != nil {
		responsePenyakit.Message = err.Error()
		return GCFReturnStruct(responsePenyakit)
	}

	responsePenyakit.Status = 200
	responsePenyakit.Message = "Insert Penyakit Success " + penyakit.Nama_Penyakit

	return GCFReturnStruct(responsePenyakit)
}

func HandlerUpdatePenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responsePenyakit.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		responsePenyakit.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(responsePenyakit)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responsePenyakit.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(responsePenyakit)
	}

	err = json.NewDecoder(r.Body).Decode(&penyakit)
	if err != nil {
		responsePenyakit.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(responsePenyakit)
	}

	_, err = UpdatePenyakit(mconn, collectionname, ID, penyakit)
	if err != nil {
		responsePenyakit.Message = err.Error()
		return GCFReturnStruct(responsePenyakit)
	}

	responsePenyakit.Status = 200
	responsePenyakit.Message = "Update Penyakit Success " + penyakit.Nama_Penyakit

	return GCFReturnStruct(responsePenyakit)
}

func HandlerDeletePenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	responsePenyakit.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		responsePenyakit.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(responsePenyakit)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responsePenyakit.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(responsePenyakit)
	}

	_, err = DeletePenyakit(mconn, collectionname, ID)
	if err != nil {
		responsePenyakit.Message = err.Error()
		return GCFReturnStruct(responsePenyakit)
	}

	responsePenyakit.Status = 200
	responsePenyakit.Message = "Delete Penyakit Success " + penyakit.Nama_Penyakit

	return GCFReturnStruct(responsePenyakit)
}
