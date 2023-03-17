package module

import (
	"context"
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

func InsertPresensi(db *mongo.Database, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (InsertedID interface{}) {
	var presensi model.Presensi
	presensi.Latitude = long
	presensi.Longitude = lat
	presensi.Location = lokasi
	presensi.Phone_number = phonenumber
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	return InsertOneDoc(db, "presensi", presensi)
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

func GetKaryawanFromPhoneNumber(phone_number string, db *mongo.Database, col string) (staf model.Presensi) {
	karyawan := db.Collection(col)
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return staf
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
