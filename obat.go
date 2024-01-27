package apiobat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aiteung/atdb"
	image "github.com/intern-monitoring/backend-intermoni"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var imageUrl string

func MongoConnect(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func InsertOneDoc(db *mongo.Database, col string, docs interface{}) (insertedID primitive.ObjectID, err error) {
	cols := db.Collection(col)
	result, err := cols.InsertOne(context.Background(), docs)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, err
}

func Register(db *mongo.Database, col string, userdata User) error {
	cols := db.Collection(col)

	hash, _ := HashPassword(userdata.Password)
	user := bson.D{
		{Key: "username", Value: userdata.Username},
		{Key: "phone_number", Value: userdata.Phone_Number},
		{Key: "password", Value: hash},
	}

	result, err := cols.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	userdata.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func Login(db *mongo.Database, col string, userdata User) (user User, status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"username": userdata.Username}

	err = cols.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		err = fmt.Errorf("username tidak ditemukan")
		return user, false, err
	}

	if !CheckPasswordHash(userdata.Password, user.Password) {
		err = fmt.Errorf("password salah")
		return user, false, err
	}

	return user, true, nil
}

func GetAllUser(db *mongo.Database, col string) (user []User, err error) {
	cols := db.Collection(col)

	cursor, err := cols.Find(context.Background(), bson.M{})
	if err != nil {
		return user, err
	}

	err = cursor.All(context.Background(), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByID(db *mongo.Database, col string, _id primitive.ObjectID) (user User, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	err = cols.FindOne(context.Background(), filter).Decode(&obat)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("data tidak di temukan dengan ID: ", _id)
		} else {
			fmt.Println("error retrieving data for ID", _id, ":", err.Error())
		}
	}

	return user, nil
}

// obat
func GetAllObat(db *mongo.Database, col string) (obat []Obat, err error) {
	cols := db.Collection(col)

	cursor, err := cols.Find(context.Background(), bson.M{})
	if err != nil {
		return obat, err
	}

	err = cursor.All(context.Background(), &obat)
	if err != nil {
		return obat, err
	}

	return obat, nil
}

func GetObatByID(db *mongo.Database, col string, _id primitive.ObjectID) (obat Obat, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	err = cols.FindOne(context.Background(), filter).Decode(&obat)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("data tidak di temukan dengan ID: ", _id)
		} else {
			fmt.Println("error retrieving data for ID", _id, ":", err.Error())
		}
	}

	return obat, nil
}

func InsertObat(db *mongo.Database, col string, r *http.Request) (docs Obat, err error) {
	jenisobat := r.FormValue("jenis_obat")
	namaobat := r.FormValue("nama_obat")
	deskripsi := r.FormValue("deskripsi")

	imageUrl, err := image.SaveFileToGithub("Febriand1", "fax.mp4@gmail.com", "Image", "pemrog", r)
	if err != nil {
		return docs, fmt.Errorf("error save file: %s", err)
	}

	objectID := primitive.NewObjectID()

	dataobat := bson.D{
		{Key: "_id", Value: objectID},
		{Key: "jenis_obat", Value: jenisobat},
		{Key: "nama_obat", Value: namaobat},
		{Key: "deskripsi", Value: deskripsi},
		{Key: "gambar", Value: imageUrl},
	}

	InsertedID, err := InsertOneDoc(db, col, dataobat)
	if err != nil {
		fmt.Printf("InsertObat: %v\n", err)
		return docs, err
	}

	docs.ID = InsertedID

	return docs, nil
}

func UpdateObat(db *mongo.Database, col string, _id primitive.ObjectID, r *http.Request) (docs Obat, err error) {
	cols := db.Collection(col)

	jenisobat := r.FormValue("jenis_obat")
	namaobat := r.FormValue("nama_obat")
	deskripsi := r.FormValue("deskripsi")
	gambar := r.FormValue("file")

	if gambar != "" {
		imageUrl = gambar
	} else {
		imageUrl, err := image.SaveFileToGithub("Febriand1", "fax.mp4@gmail.com", "Image", "pemrog", r)
		if err != nil {
			return docs, fmt.Errorf("error save file: %s", err)
		}
		gambar = imageUrl
	}

	filter := bson.M{"_id": _id}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "jenis_obat", Value: jenisobat},
			{Key: "nama_obat", Value: namaobat},
			{Key: "deskripsi", Value: deskripsi},
			{Key: "gambar", Value: gambar},
		}},
	}

	result, err := cols.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return docs, err
	}

	if result.MatchedCount == 0 {
		return docs, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return docs, nil
}

func DeleteObat(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	result, err := cols.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return true, nil
}

// penyakit
func GetAllPenyakit(db *mongo.Database, col string) (penyakit []Penyakit, err error) {
	cols := db.Collection(col)

	cursor, err := cols.Find(context.Background(), bson.M{})
	if err != nil {
		return penyakit, err
	}

	err = cursor.All(context.Background(), &penyakit)
	if err != nil {
		return penyakit, err
	}

	return penyakit, nil
}

func GetPenyakitByID(db *mongo.Database, col string, _id primitive.ObjectID) (penyakit Penyakit, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	err = cols.FindOne(context.Background(), filter).Decode(&penyakit)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("data tidak di temukan dengan ID: ", _id)
		} else {
			fmt.Println("error retrieving data for ID", _id, ":", err.Error())
		}
	}

	return penyakit, nil
}

func InsertPenyakit(db *mongo.Database, col string, r *http.Request) (docs Penyakit, err error) {
	jenispenyakit := r.FormValue("jenis_penyakit")
	namapenyakit := r.FormValue("nama_penyakit")
	deskripsi := r.FormValue("deskripsi")
	namaobat := r.FormValue("nama_obat")

	obat := Obat{
		Nama_Obat: namaobat,
	}

	objectID := primitive.NewObjectID()

	datapenyakit := bson.D{
		{Key: "_id", Value: objectID},
		{Key: "jenis_penyakit", Value: jenispenyakit},
		{Key: "nama_penyakit", Value: namapenyakit},
		{Key: "deskripsi", Value: deskripsi},
		{Key: "obat", Value: bson.D{
			{Key: "nama_obat", Value: obat},
		}},
	}

	InsertedID, err := InsertOneDoc(db, col, datapenyakit)
	if err != nil {
		fmt.Printf("InsertPenyakit: %v\n", err)
		return docs, err
	}

	docs.ID = InsertedID

	return docs, nil
}

func UpdatePenyakit(db *mongo.Database, col string, _id primitive.ObjectID, r *http.Request) (docs Penyakit, err error) {
	cols := db.Collection(col)

	jenispenyakit := r.FormValue("jenis_penyakit")
	namapenyakit := r.FormValue("nama_penyakit")
	deskripsi := r.FormValue("deskripsi")
	namaobat := r.FormValue("nama_obat")

	obat := Obat{
		Nama_Obat: namaobat,
	}

	filter := bson.M{"_id": _id}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "jenis_penyakit", Value: jenispenyakit},
			{Key: "nama_penyakit", Value: namapenyakit},
			{Key: "deskripsi", Value: deskripsi},
			{Key: "obat", Value: bson.D{
				{Key: "nama_obat", Value: obat},
			}},
		}},
	}

	result, err := cols.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return docs, err
	}

	if result.MatchedCount == 0 {
		return docs, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return docs, nil
}

func DeletePenyakit(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	result, err := cols.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return true, nil
}

// rumah sakit
func GetAllRS(db *mongo.Database, col string) (rs []RumahSakit, err error) {
	cols := db.Collection(col)

	cursor, err := cols.Find(context.Background(), bson.M{})
	if err != nil {
		return rs, err
	}

	err = cursor.All(context.Background(), &rs)
	if err != nil {
		return rs, err
	}

	return rs, nil
}

func GetRSByID(db *mongo.Database, col string, _id primitive.ObjectID) (rs RumahSakit, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	err = cols.FindOne(context.Background(), filter).Decode(&rs)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("data tidak di temukan dengan ID: ", _id)
		} else {
			fmt.Println("error retrieving data for ID", _id, ":", err.Error())
		}
	}

	return rs, nil
}

func InsertRS(db *mongo.Database, col string, r *http.Request) (docs RumahSakit, err error) {
	namars := r.FormValue("nama_rs")
	notelp := r.FormValue("no_telp")
	alamat := r.FormValue("alamat")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	namaobat := r.FormValue("nama_obat")

	obat := Obat{
		Nama_Obat: namaobat,
	}

	objectID := primitive.NewObjectID()

	imageUrl, err := image.SaveFileToGithub("Febriand1", "fax.mp4@gmail.com", "Image", "pemrog", r)
	if err != nil {
		return docs, fmt.Errorf("error save file: %s", err)
	}

	datapenyakit := bson.D{
		{Key: "_id", Value: objectID},
		{Key: "nama_rs", Value: namars},
		{Key: "no_telp", Value: notelp},
		{Key: "alamat", Value: alamat},
		{Key: "latitude", Value: latitude},
		{Key: "longitude", Value: longitude},
		{Key: "gambar", Value: imageUrl},
		{Key: "obat", Value: bson.D{
			{Key: "nama_obat", Value: obat},
		}},
	}

	InsertedID, err := InsertOneDoc(db, col, datapenyakit)
	if err != nil {
		fmt.Printf("InsertPenyakit: %v\n", err)
		return docs, err
	}

	docs.ID = InsertedID

	return docs, nil
}

func UpdateRS(db *mongo.Database, col string, _id primitive.ObjectID, r *http.Request) (docs RumahSakit, err error) {
	cols := db.Collection(col)

	namars := r.FormValue("nama_rs")
	notelp := r.FormValue("no_telp")
	alamat := r.FormValue("alamat")
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	gambar := r.FormValue("file")
	namaobat := r.FormValue("nama_obat")

	obat := Obat{
		Nama_Obat: namaobat,
	}

	if gambar != "" {
		imageUrl = gambar
	} else {
		imageUrl, err := image.SaveFileToGithub("Febriand1", "fax.mp4@gmail.com", "Image", "pemrog", r)
		if err != nil {
			return docs, fmt.Errorf("error save file: %s", err)
		}
		gambar = imageUrl
	}

	filter := bson.M{"_id": _id}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "nama_rs", Value: namars},
			{Key: "no_telp", Value: notelp},
			{Key: "alamat", Value: alamat},
			{Key: "latitude", Value: latitude},
			{Key: "longitude", Value: longitude},
			{Key: "gambar", Value: gambar},
			{Key: "obat", Value: bson.D{
				{Key: "nama_obat", Value: obat},
			}},
		}},
	}

	result, err := cols.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return docs, err
	}

	if result.MatchedCount == 0 {
		return docs, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return docs, nil
}

func DeleteRS(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	result, err := cols.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		return false, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return true, nil
}
