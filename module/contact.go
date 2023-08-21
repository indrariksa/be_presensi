package module

import (
	"context"
	"fmt"
	"github.com/indrariksa/be_presensi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllContact(db *mongo.Database, col string) (data []model.Kontak) {
	kontak := db.Collection(col)
	filter := bson.M{}
	cursor, err := kontak.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertKontak(db *mongo.Database, col string, nmkontak string, nmrkontak string, almt string, ktrngn string) (insertedID primitive.ObjectID, err error) {
	presensi := bson.M{
		"namakontak": nmkontak,
		"nomorhp":    nmrkontak,
		"alamat":     almt,
		"keterangan": ktrngn,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), presensi)
	if err != nil {
		fmt.Printf("Insert Kontak: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}
