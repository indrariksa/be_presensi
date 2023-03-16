package NPM

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

//func MongoConnect2(mconn DBInfo) (db *mongo.Database) {
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
//	if err != nil {
//		fmt.Printf("AIteung Mongo, MongoConnect: %v\n", err)
//	}
//	return client.Database(mconn.DBName)
//}
//
//var DBUlbimongoinfo = DBInfo{
//	DBString: MongoString,
//	DBName:   "tes_db",
//}
//
//var Ulbimongoconn = MongoConnect2(DBUlbimongoinfo)

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPresensi(long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata Karyawan) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Latitude = long
	presensi.Longitude = lat
	presensi.Location = lokasi
	presensi.Phone_number = phonenumber
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	return InsertOneDoc("tes_db", "presensi", presensi)
}

func InsertKaryawan(nama string, phone_number string, jabatan string, jam_kerja []JamKerja, hari_kerja []string) (InsertedID interface{}) {
	var karyawan Karyawan
	karyawan.Nama = nama
	karyawan.PhoneNumber = phone_number
	karyawan.Jabatan = jabatan
	karyawan.Jam_kerja = jam_kerja
	karyawan.Hari_kerja = hari_kerja
	return InsertOneDoc("tes_db", "karyawan", karyawan)
}

func GetKaryawanFromPhoneNumber(phone_number string) (staf Presensi) {
	karyawan := MongoConnect("tes_db").Collection("presensi")
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return staf
}

func GetKaryawanFromStatus(checkin string) (staf Presensi) {
	karyawan := MongoConnect("tes_db").Collection("presensi")
	filter := bson.M{"checkin": checkin}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return staf
}

func GetKaryawanFromName(jabatan string) (staf Presensi) {
	nm_karyawan := MongoConnect("tes_db").Collection("presensi")
	filter := bson.M{"biodata.jabatan": jabatan}
	err := nm_karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromNama: %v\n", err)
	}
	return staf
}

//func GetPresensiCurrentMonth(mongoconn *mongo.Database) (allpresensi []Presensi) {
//	startdate, enddate := GetFirstLastDateCurrentMonth()
//	coll := mongoconn.Collection("presensi")
//	today := bson.M{
//		"$gte": primitive.NewDateTimeFromTime(startdate),
//		"$lte": primitive.NewDateTimeFromTime(enddate),
//	}
//	filter := bson.M{"datetime": today}
//	cursor, err := coll.Find(context.TODO(), filter)
//	if err != nil {
//		fmt.Printf("getPresensiTodayFromPhoneNumber: %v\n", err)
//	}
//	err = cursor.All(context.TODO(), &allpresensi)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	return
//}
