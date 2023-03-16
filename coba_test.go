package NPM

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	var jamKerja1 = JamKerja{
		Durasi:     8,
		Jam_masuk:  "08:00",
		Jam_keluar: "16:00",
		Gmt:        7,
		Hari:       []string{"Senin", "Rabu", "Kamis"},
		Shift:      1,
		Piket_tim:  "Piket A",
	}
	var jamKerja2 = JamKerja{
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
	lokasi := "Hawai"
	phonenumber := "6811110023231"
	checkin := "masuk"
	biodata := Karyawan{
		Nama:        "Parhan",
		PhoneNumber: "6284564562",
		Jabatan:     "DPR",
		Jam_kerja:   []JamKerja{jamKerja1, jamKerja2},
		Hari_kerja:  []string{"Senin", "Selasa"},
	}
	hasil := InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
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

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "628123456789"
	biodata := GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

func TestGetKaryawanFromStatus(t *testing.T) {
	checkin := "masuk"
	biodata := GetKaryawanFromStatus(checkin)
	fmt.Println(biodata)
}

//func TestGetKaryawanFromName(t *testing.T) {
//	namaewa := "DPR"
//	biodata := GetKaryawanFromName(namaewa)
//	fmt.Println(biodata)
//}

//func TestGetPresensi(t *testing.T) {
//	ps := GetPresensiCurrentMonth(Ulbimongoconn)
//	fmt.Print(ps)
//}
