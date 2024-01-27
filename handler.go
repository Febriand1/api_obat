package apiobat

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	user       User
	obat       Obat
	penyakit   Penyakit
	rs         RumahSakit
	credential Credential
	response   Response
)

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// user
func HandlerRegister(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	credential.Status = 400

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		credential.Message = "error parsing application/json: " + err.Error()
	}

	err = Register(mconn, collectionname, user)
	if err != nil {
		credential.Message = err.Error()
		return GCFReturnStruct(credential)
	}

	credential.Status = 200
	credential.Message = "Registrasi Berhasil"

	return GCFReturnStruct(credential)
}

func HandlerLogin(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	credential.Status = 400

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		credential.Message = "error parsing application/json: " + err.Error()
	}

	users, _, err := Login(mconn, collectionname, user)
	if err != nil {
		credential.Message = err.Error()
		return GCFReturnStruct(credential)
	}

	credential.Status = 200
	credential.Message = "Selamat Datang " + users.Username

	return GCFReturnStruct(credential)
}

func HandlerGetAllUser(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)

	data, err := GetAllUser(mconn, collectionname)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get All User Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerGetUserByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	data, err := GetUserByID(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get User By ID Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

// obat
func HandlerGetAllObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	data, err := GetAllObat(mconn, collectionname)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get All Obat Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerGetObatByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	data, err := GetObatByID(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get Obat By ID Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerInsertObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	// err := json.NewDecoder(r.Body).Decode(&obat)
	// if err != nil {
	// 	response.Message = "error parsing application/json: " + err.Error()
	// 	return GCFReturnStruct(response)
	// }

	data, err := InsertObat(mconn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Insert Obat Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerUpdateObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	// err = json.NewDecoder(r.Body).Decode(&obat)
	// if err != nil {
	// 	response.Message = "error parsing application/json: " + err.Error()
	// 	return GCFReturnStruct(response)
	// }

	data, err := UpdateObat(mconn, collectionname, ID, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Update Obat Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerDeleteObat(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	_, err = DeleteObat(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Delete Obat Success " + obat.Nama_Obat

	return GCFReturnStruct(response)
}

// penyakit
func HandlerGetAllPenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	data, err := GetAllPenyakit(mconn, collectionname)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get Penyakit Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerGetPenyakitByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	data, err := GetPenyakitByID(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get Penyakit By ID Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerInsertPenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	err := json.NewDecoder(r.Body).Decode(&penyakit)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	data, err := InsertPenyakit(mconn, collectionname, penyakit)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Insert Penyakit Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerUpdatePenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	err = json.NewDecoder(r.Body).Decode(&penyakit)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}

	data, err := UpdatePenyakit(mconn, collectionname, ID, penyakit)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Update Penyakit Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerDeletePenyakit(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	_, err = DeletePenyakit(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Delete Penyakit Success " + penyakit.Nama_Penyakit

	return GCFReturnStruct(response)
}

// rumah sakit
func HandlerGetAllRS(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	data, err := GetAllRS(mconn, collectionname)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get RS Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerGetRSByID(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	data, err := GetRSByID(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Get RS By ID Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerInsertRS(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	// err := json.NewDecoder(r.Body).Decode(&rs)
	// if err != nil {
	// 	response.Message = "error parsing application/json: " + err.Error()
	// 	return GCFReturnStruct(response)
	// }

	data, err := InsertRS(mconn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Insert RS Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerUpdateRS(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	// err = json.NewDecoder(r.Body).Decode(&rs)
	// if err != nil {
	// 	response.Message = "error parsing application/json: " + err.Error()
	// 	return GCFReturnStruct(response)
	// }

	data, err := UpdateRS(mconn, collectionname, ID, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Update RS Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}

	return GCFReturnStruct(responData)
}

func HandlerDeleteRS(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400

	id := r.URL.Query().Get("_id")
	if id == "" {
		response.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(response)
	}

	_, err = DeleteRS(mconn, collectionname, ID)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}

	response.Status = 200
	response.Message = "Delete RS Success " + rs.Nama_RS

	return GCFReturnStruct(response)
}
