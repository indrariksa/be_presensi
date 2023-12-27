package NPM

import (
	"context"
	"fmt"
	"github.com/indrariksa/be_presensi/model"
	"github.com/indrariksa/be_presensi/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

//func TestInsertPresensi(t *testing.T) {
//	var jamKerja1 = model.JamKerja{
//		Durasi:     8,
//		Jam_masuk:  "08:00",
//		Jam_keluar: "16:00",
//		Gmt:        7,
//		Hari:       []string{"Senin", "Rabu", "Kamis"},
//		Shift:      1,
//		Piket_tim:  "Piket A",
//	}
//	var jamKerja2 = model.JamKerja{
//		Durasi:     8,
//		Jam_masuk:  "09:00",
//		Jam_keluar: "17:00",
//		Gmt:        7,
//		Hari:       []string{"Sabtu"},
//		Shift:      2,
//		Piket_tim:  "",
//	}
//
//	long := 98.345345
//	lat := 123.561651
//	lokasi := "New York"
//	phonenumber := "6811110023231"
//	checkin := "masuk"
//	biodata := model.Karyawan{
//		Nama:        "George Best",
//		PhoneNumber: "6284564562",
//		Jabatan:     "Rakyat",
//		Jam_kerja:   []model.JamKerja{jamKerja1, jamKerja2},
//		Hari_kerja:  []string{"Senin", "Selasa"},
//	}
//	hasil := module.InsertPresensi(module.MongoConn, "presensi", long, lat, lokasi, phonenumber, checkin, biodata)
//	fmt.Println(hasil)
//}

func TestInsertPresensi(t *testing.T) {
	var jamKerja1 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "08:00",
		Jam_keluar: "16:00",
		Gmt:        7,
		Hari:       []string{"Senin", "Rabu", "Kamis"},
		Shift:      1,
		Piket_tim:  "Piket A",
	}
	var jamKerja2 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "09:00",
		Jam_keluar: "17:00",
		Gmt:        7,
		Hari:       []string{"Sabtu"},
		Shift:      2,
		Piket_tim:  "",
	}

	long := 98.345345
	lat := 123.561651
	lokasi := "New York"
	phonenumber := "6811110023231"
	checkin := "masuk"
	biodata := model.Karyawan{
		Nama:        "Kindi Herdiansyah",
		PhoneNumber: "6284564562",
		Jabatan:     "Rakyat",
		Jam_kerja:   []model.JamKerja{jamKerja1, jamKerja2},
		Hari_kerja:  []string{"Senin", "Selasa"},
	}
	insertedID, err := module.InsertPresensi(module.MongoConn, "presensi", long, lat, lokasi, phonenumber, checkin, biodata)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestUpdatePresensi(t *testing.T) {
	col := "presensi"

	var jamKerja1 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "08:00",
		Jam_keluar: "16:00",
		Gmt:        7,
		Hari:       []string{"Senin", "Rabu", "Kamis"},
		Shift:      1,
		Piket_tim:  "Piket A",
	}
	var jamKerja2 = model.JamKerja{
		Durasi:     8,
		Jam_masuk:  "09:00",
		Jam_keluar: "17:00",
		Gmt:        7,
		Hari:       []string{"Sabtu"},
		Shift:      2,
		Piket_tim:  "",
	}

	// Define a test document
	doc := model.Presensi{
		ID:           primitive.NewObjectID(),
		Longitude:    98.345345,
		Latitude:     123.561651,
		Location:     "New York",
		Phone_number: "6811110023231",
		Checkin:      "masuk",
		Biodata: model.Karyawan{
			Nama:        "George Best",
			PhoneNumber: "6284564562",
			Jabatan:     "Rakyat",
			Jam_kerja:   []model.JamKerja{jamKerja1, jamKerja2},
			Hari_kerja:  []string{"Senin", "Selasa"},
		},
	}

	// Insert the test document into the collection
	if _, err := module.MongoConn.Collection(col).InsertOne(context.Background(), doc); err != nil {
		t.Fatalf("Failed to insert test document: %v", err)
	}

	// Define the fields to update
	long := 99.123456
	lat := 123.789012
	lokasi := "Los Angeles"
	phonenumber := "6811110023232"
	checkin := "pulang"
	biodata := model.Karyawan{
		Nama:        "Diego Maradona",
		PhoneNumber: "6281234567",
		Jabatan:     "Legenda",
		Jam_kerja:   []model.JamKerja{jamKerja1},
		Hari_kerja:  []string{"Senin", "Jumat"},
	}

	// Call UpdatePresensi with the test document ID and updated fields
	if err := module.UpdatePresensi(module.MongoConn, col, doc.ID, long, lat, lokasi, phonenumber, checkin, biodata); err != nil {
		t.Fatalf("UpdatePresensi failed: %v", err)
	}

	// Retrieve the updated document from the collection
	var updatedDoc model.Presensi
	if err := module.MongoConn.Collection(col).FindOne(context.Background(), bson.M{"_id": doc.ID}).Decode(&updatedDoc); err != nil {
		t.Fatalf("Failed to retrieve updated document: %v", err)
	}

	// Verify that the document was updated as expected
	if updatedDoc.Longitude != long || updatedDoc.Latitude != lat || updatedDoc.Location != lokasi || updatedDoc.Phone_number != phonenumber || updatedDoc.Checkin != checkin {
		t.Fatalf("Document was not updated as expected")
	}
}

func TestUpdateKontak(t *testing.T) {
	col := "kontak"

	// Define a test document
	doc := model.Kontak{
		ID:         primitive.NewObjectID(),
		NamaKontak: "testing",
		NomorHp:    "6811110023231",
		Alamat:     "New York",
		Keterangan: "6811110023231",
	}

	// Insert the test document into the collection
	if _, err := module.MongoConn.Collection(col).InsertOne(context.Background(), doc); err != nil {
		t.Fatalf("Failed to insert test document: %v", err)
	}

	// Define the fields to update
	nama_kontak := "testing"
	nomor_hp := "6811110023231"
	alamat := "New York"
	keterangan := "6811110023222"

	// Call UpdatePresensi with the test document ID and updated fields
	if err := module.UpdateKontak(module.MongoConn, col, doc.ID, nama_kontak, nomor_hp, alamat, keterangan); err != nil {
		t.Fatalf("UpdateKontak failed: %v", err)
	}

	// Retrieve the updated document from the collection
	var updatedDoc model.Kontak
	if err := module.MongoConn.Collection(col).FindOne(context.Background(), bson.M{"_id": doc.ID}).Decode(&updatedDoc); err != nil {
		t.Fatalf("Failed to retrieve updated document: %v", err)
	}

	// Verify that the document was updated as expected
	if updatedDoc.NamaKontak != nama_kontak || updatedDoc.NomorHp != nomor_hp || updatedDoc.Alamat != alamat || updatedDoc.Keterangan != keterangan {
		t.Fatalf("Document was not updated as expected")
	}
}

//func TestInsertKaryawan(t *testing.T) {
//	var jamKerja1 = JamKerja{
//		Durasi:     8,
//		Jam_masuk:  "08:00",
//		Jam_keluar: "16:00",
//		Gmt:        7,
//		Hari:       []string{"Senin", "Selasa"},
//		Shift:      1,
//		Piket_tim:  "Piket A",
//	}
//	var jamKerja2 = JamKerja{
//		Durasi:     8,
//		Jam_masuk:  "09:00",
//		Jam_keluar: "17:00",
//		Gmt:        7,
//		Hari:       []string{"Sabtu"},
//		Shift:      2,
//		Piket_tim:  "",
//	}
//
//	nama := "Bulan"
//	phone_number := "08123456789"
//	jabatan := "DPR"
//	jam_kerja := []JamKerja{jamKerja1, jamKerja2}
//	hari_kerja := []string{"Senin", "Selasa"}
//	hasil := InsertKaryawan(nama, phone_number, jabatan, jam_kerja, hari_kerja)
//	fmt.Println(hasil)
//}

func TestGetKaryawanFromStatus(t *testing.T) {
	checkin := "masuk"
	biodata := module.GetPresensiFromStatus(checkin, module.MongoConn, "presensi")
	fmt.Println(biodata)
}

func TestGetAllKaryawanFromStatus(t *testing.T) {
	checkin := "masuk"
	data := module.GetAllPresensiFromStatus(checkin, module.MongoConn, "presensi")
	fmt.Println(data)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "628122221814"
	biodata, err := module.GetKaryawanFromPhoneNumber(phonenumber, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling GetKaryawanFromPhoneNumber: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetPresensiFromID(t *testing.T) {
	id := "6422e300f590e691c91082cb"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPresensiFromID(objectID, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllPresensi(module.MongoConn, "presensi")
	fmt.Println(data)
}

func TestDeletePresensiByID(t *testing.T) {
	id := "646611fecb7c6963b9e2baad" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePresensiByID(objectID, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetPresensiFromID(objectID, module.MongoConn, "presensi")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestGetAllKontak(t *testing.T) {
	data := module.GetAllContacts(module.MongoConn, "contacts")
	fmt.Println(data)
}

func TestInsertKontak(t *testing.T) {
	nama_kontak := "testing"
	nomor_hp := "6811110023231"
	alamat := "New York"
	keterangan := "6811110023231"

	insertedID, err := module.InsertKontak(module.MongoConn, "kontak", nama_kontak, nomor_hp, alamat, keterangan)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestCreateUser(t *testing.T) {
	username := "admin"    // Ganti dengan username yang diinginkan
	password := "admin123" // Ganti dengan password yang diinginkan

	insertedID, err := module.CreateUser(module.MongoConn, "users", username, password)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	// Verifikasi bahwa data pengguna telah berhasil dimasukkan
	if insertedID.IsZero() {
		t.Fatal("Invalid inserted user ID")
	}

	// Anda bisa menambahkan pengujian lainnya untuk memeriksa apakah pengguna berhasil disimpan di database
	// Misalnya, Anda dapat mencoba mendapatkan pengguna menggunakan fungsi GetUserByUsername
	// dan memeriksa apakah pengguna yang didapatkan sesuai dengan data yang dimasukkan sebelumnya.
}

func TestLogin_Success(t *testing.T) {
	// Data pengguna yang akan digunakan untuk pengujian
	username := "username" // Ganti dengan username yang ada di database
	password := "password" // Ganti dengan password yang sesuai dengan pengguna yang telah ada di database

	// Memanggil fungsi login
	loggedIn, err := module.Login(username, password, module.MongoConn, "users")
	if err != nil {
		t.Errorf("Error logging in: %v", err)
	}

	// Verifikasi bahwa login berhasil dilakukan
	if !loggedIn {
		t.Error("Login should be successful but it failed")
	}
}

func TestLogin_Failure(t *testing.T) {
	// Data pengguna yang akan digunakan untuk pengujian
	username := "username"      // Ganti dengan username yang mungkin tidak ada di database
	password := "wrongpassword" // Ganti dengan password yang salah untuk pengguna yang mungkin tidak ada di database

	// Memanggil fungsi login
	loggedIn, err := module.Login(username, password, module.MongoConn, "users")
	if err != nil {
		t.Errorf("Error logging in: %v", err)
	}

	// Verifikasi bahwa login gagal dilakukan
	if loggedIn {
		t.Error("Login should fail but it succeeded")
	}
}
