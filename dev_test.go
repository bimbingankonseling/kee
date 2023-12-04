package kee

import (
	"fmt"
	"testing"

	model "github.com/bimbingankonseling/kee/model"
	module "github.com/bimbingankonseling/kee/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "kee")

// TEST SIGN IN
func TestLogin(t *testing.T) {
	var doc model.User
	doc.Username = "gabriella"
	doc.Password = "bella123"
	user, Status, err := module.Login(db, "user", doc)
	fmt.Println("Status :", Status)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Welcome back:", user)
	}
}

func TestInsertPemasukan(t *testing.T) {
	var doc model.Pemasukan
	doc.Tanggal_masuk = "26/02/2023"
	doc.Jumlah_masuk = 50000
	doc.Sumber = "Gaji"
	doc.Deskripsi = "dari kantor"
	// doc.User = model.User{Username: "Fedhira Syaila"}

	hasil, err := module.InsertPemasukan(db, "pemasukan", doc.Tanggal_masuk, doc.Jumlah_masuk, doc.Sumber, doc.Deskripsi)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Printf("Data berhasil disimpan dengan id %s\n", hasil.Hex())
	}
	fmt.Println(hasil)
}

// func TestGeneratePasswordHash(t *testing.T) {
// 	password := "bellaa"
// 	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
// 	fmt.Println("Password:", password)
// 	fmt.Println("Hash:    ", hash)

// 	match := CheckPasswordHash(password, hash)
// 	fmt.Println("Match:   ", match)
// }
// func TestGeneratePrivateKeyPaseto(t *testing.T) {
// 	privateKey, publicKey := watoken.GenerateKey()
// 	fmt.Println(privateKey)
// 	fmt.Println(publicKey)
// 	hasil, err := watoken.Encode("gabril", privateKey)
// 	fmt.Println(hasil, err)
// }

// func TestHashFunction(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "keekonseling")
// 	var userdata User
// 	userdata.Username = "gabril"
// 	userdata.Password = "bellaa"

// 	filter := bson.M{"username": userdata.Username}
// 	res := atdb.GetOneDoc[User](mconn, "user", filter)
// 	fmt.Println("Mongo User Result: ", res)
// 	hash, _ := HashPassword(userdata.Password)
// 	fmt.Println("Hash Password : ", hash)
// 	match := CheckPasswordHash(userdata.Password, res.Password)
// 	fmt.Println("Match:   ", match)

// }

// func TestIsPasswordValid(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "keekonseling")
// 	var userdata User
// 	userdata.Username = "gabril"
// 	userdata.Password = "bellaa"

// 	anu := IsPasswordValid(mconn, "user", userdata)
// 	fmt.Println(anu)
// }

// func TestInsertUser(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "keekonseling")
// 	var userdata User
// 	userdata.Username = "gabril"
// 	userdata.Password = "bellaa"

// 	nama := InsertUser(mconn, "user", userdata)
// 	fmt.Println(nama)
// }