package apiobat

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoConnect(MONGOCONNSTRINGENV, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(MONGOCONNSTRINGENV),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

func Register(db *mongo.Database, col string, userdata User) error {
	cols := db.Collection(col)

	hash, _ := HashPassword(userdata.Password)
	user := bson.D{
		{Key: "username", Value: userdata.Username},
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

func InsertObat(db *mongo.Database, col string, obat Obat) (InsertedID primitive.ObjectID, status bool, err error) {
	cols := db.Collection(col)

	dataobat := bson.D{
		{Key: "jenis_obat", Value: obat.Jenis_Obat},
		{Key: "nama_obat", Value: obat.Nama_Obat},
		{Key: "deskripsi", Value: obat.Deskripsi},
	}

	result, err := cols.InsertOne(context.Background(), dataobat)
	if err != nil {
		return InsertedID, false, err
	}
	InsertedID = result.InsertedID.(primitive.ObjectID)

	return InsertedID, true, nil
}

func UpdateObat(db *mongo.Database, col string, _id primitive.ObjectID, obat Obat) (status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "jenis_obat", Value: obat.Jenis_Obat},
			{Key: "nama_obat", Value: obat.Nama_Obat},
			{Key: "deskripsi", Value: obat.Deskripsi},
		}},
	}

	result, err := cols.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return true, nil
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

func InsertPenyakit(db *mongo.Database, col string, penyakit Penyakit) (InsertedID primitive.ObjectID, status bool, err error) {
	cols := db.Collection(col)

	datapenyakit := bson.D{
		{Key: "jenis_penyakit", Value: penyakit.Jenis_Penyakit},
		{Key: "nama_penyakit", Value: penyakit.Nama_Penyakit},
		{Key: "deskripsi", Value: penyakit.Deskripsi},
		{Key: "obat", Value: bson.D{
			{Key: "nama_obat", Value: penyakit.Obat.Nama_Obat},
		}},
	}

	result, err := cols.InsertOne(context.Background(), datapenyakit)
	if err != nil {
		return InsertedID, false, err
	}
	InsertedID = result.InsertedID.(primitive.ObjectID)

	return InsertedID, true, nil
}

func UpdatePenyakit(db *mongo.Database, col string, _id primitive.ObjectID, penyakit Penyakit) (status bool, err error) {
	cols := db.Collection(col)

	filter := bson.M{"_id": _id}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "jenis_penyakit", Value: penyakit.Jenis_Penyakit},
			{Key: "nama_penyakit", Value: penyakit.Nama_Penyakit},
			{Key: "deskripsi", Value: penyakit.Deskripsi},
			{Key: "obat", Value: bson.D{
				{Key: "nama_obat", Value: penyakit.Obat.Nama_Obat},
			}},
		}},
	}

	result, err := cols.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, fmt.Errorf("data tidak di temukan dengan ID: %s", _id)
	}

	return true, nil
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
