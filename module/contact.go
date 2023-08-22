package module

import (
	"context"
	"errors"
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
		"nama_kontak": nmkontak,
		"nomor_hp":    nmrkontak,
		"alamat":      almt,
		"keterangan":  ktrngn,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), presensi)
	if err != nil {
		fmt.Printf("Insert Kontak: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateKontak(db *mongo.Database, col string, id primitive.ObjectID, nmkontak string, nmrkontak string, almt string, ktrngn string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama_kontak": nmkontak,
			"nomor_hp":    nmrkontak,
			"alamat":      almt,
			"keterangan":  ktrngn,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateKontak: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteKontakByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	kontak := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := kontak.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
