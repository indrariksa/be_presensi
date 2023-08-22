package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/aiteung/atdb"
	"github.com/indrariksa/be_presensi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tes_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

//func InsertPresensi(db *mongo.Database, col string, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (InsertedID interface{}) {
//	var presensi model.Presensi
//	presensi.Latitude = long
//	presensi.Longitude = lat
//	presensi.Location = lokasi
//	presensi.Phone_number = phonenumber
//	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
//	presensi.Checkin = checkin
//	presensi.Biodata = biodata
//	return InsertOneDoc(db, col, presensi)
//}

func InsertPresensi(db *mongo.Database, col string, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (insertedID primitive.ObjectID, err error) {
	presensi := bson.M{
		"longitude":    long,
		"latitude":     lat,
		"location":     lokasi,
		"phone_number": phonenumber,
		"datetime":     primitive.NewDateTimeFromTime(time.Now().UTC()),
		"checkin":      checkin,
		"biodata":      biodata,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), presensi)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertKaryawan(db *mongo.Database, nama string, phone_number string, jabatan string, jam_kerja []model.JamKerja, hari_kerja []string) (InsertedID interface{}) {
	var karyawan model.Karyawan
	karyawan.Nama = nama
	karyawan.PhoneNumber = phone_number
	karyawan.Jabatan = jabatan
	karyawan.Jam_kerja = jam_kerja
	karyawan.Hari_kerja = hari_kerja
	return InsertOneDoc(db, "karyawan", karyawan)
}

func UpdatePresensi(db *mongo.Database, col string, id primitive.ObjectID, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"longitude":    long,
			"latitude":     lat,
			"location":     lokasi,
			"phone_number": phonenumber,
			"checkin":      checkin,
			"biodata":      biodata,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePresensi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func GetPresensiFromID(_id primitive.ObjectID, db *mongo.Database, col string) (staf model.Presensi, errs error) {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return staf, fmt.Errorf("no data found for ID %s", _id)
		}
		return staf, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return staf, nil
}

func GetKaryawanFromPhoneNumber(phone_number string, db *mongo.Database, col string) (staf model.Presensi, errs error) {
	karyawan := db.Collection(col)
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return staf, fmt.Errorf("no data found for phone number %s", phone_number)
		}
		return staf, fmt.Errorf("error retrieving data for phone number %s: %s", phone_number, err.Error())
	}
	return staf, nil
}

func GetKaryawanFromName(jabatan string, db *mongo.Database, col string) (staf model.Presensi) {
	nm_karyawan := db.Collection(col)
	filter := bson.M{"biodata.jabatan": jabatan}
	err := nm_karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromNama: %v\n", err)
	}
	return staf
}

func GetPresensiFromStatus(checkin string, db *mongo.Database, col string) (data model.Presensi) {
	karyawan := db.Collection(col)
	filter := bson.M{"checkin": checkin}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return data
}

func GetAllPresensiFromStatus(checkin string, db *mongo.Database, col string) (data []model.Presensi) {
	karyawan := db.Collection(col)
	filter := bson.M{"checkin": checkin}
	cursor, err := karyawan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllPresensi(db *mongo.Database, col string) (data []model.Presensi) {
	karyawan := db.Collection(col)
	filter := bson.M{}
	cursor, err := karyawan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeletePresensiByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := karyawan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with specific ID %s not found", _id)
	}

	return nil
}
