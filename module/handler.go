package module

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bimbingankonseling/kee/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/whatsauth/watoken"
)

var (
	Responsed           model.Credential
	reservasiResponse   model.ReservasiResponse
	registrasiResponse 	model.RegistrasiResponse
	datauser            model.User
	reservasi           model.Reservasi
	registrasi 	        model.Registrasi
)

//login
func GCFHandlerLogin(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Responsed model.Credential
	Responsed.Status = false

	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Responsed.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(Responsed)
	}
	user, status1, err := Login(conn, collectionname, datauser)
	if err != nil {
		Responsed.Message = err.Error()
		return GCFReturnStruct(Responsed)
	}
	Responsed.Status = true
	tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		Responsed.Message = "Gagal Encode Token : " + err.Error()
	} else {
		Responsed.Message = "Selamat Datang " + user.Username + " di KeeKonseling" + strconv.FormatBool(status1)
		Responsed.Token = tokenstring
	}
	return GCFReturnStruct(Responsed)
}

func GCFHandlerGetAll(MONGOCONNSTRINGENV, dbname, col string, docs interface{}) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	data := GetAllDocs(conn, col, docs)
	return GCFReturnStruct(data)
}

// func GCFHandlerGetUserFromUsername(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
// 	Responsed.Status = false

// 	username := r.URL.Query().Get("username")
// 	if username == "" {
// 		Responsed.Message = "Missing 'username' parameter in the URL"
// 		return GCFReturnStruct(Responsed)
// 	}

// 	datauser.Username = username

// 	user, err := GetUserFromUsername(mconn, collectionname, username)
// 	if err != nil {
// 		Responsed.Message = "Error retrieving user data: " + err.Error()
// 		return GCFReturnStruct(Responsed)
// 	}

// 	Responsed.Status = true
// 	Responsed.Message = "Hello user"
// 	Responsed.Data = []model.User{user}

// 	return GCFReturnStruct(Responsed)
// }

//reservasi
func GCFHandlerInsertReservasi(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	reservasiResponse.Status = false
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		reservasiResponse.Message = "error parsing application/json1:"
		return GCFReturnStruct(reservasiResponse)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(reservasiResponse)
	}

	err = json.NewDecoder(r.Body).Decode(&reservasi)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(reservasiResponse)
	}

	_, err = InsertReservasi(mconn, collectionname, reservasi.Nama, reservasi.No_telp, reservasi.TTL, reservasi.Status, reservasi.Keluhan)
	if err != nil {
		reservasiResponse.Message = "error inserting Reservasi: " + err.Error()
		return GCFReturnStruct(reservasiResponse)
	}

	reservasiResponse.Status = true
	reservasiResponse.Message = "Insert Pemasukan success"
	reservasiResponse.Data = []model.Reservasi{reservasi}
	return GCFReturnStruct(reservasiResponse)
}

func GCFHandlerGetAllReservasi(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	reservasiResponse.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		reservasiResponse.Message = "error parsing application/json1:"
		return GCFReturnStruct(reservasiResponse)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(reservasiResponse)
	}

	reservasi, err := GetAllReservasi(mconn, collectionname)
	if err != nil {
		reservasiResponse.Message = err.Error()
		return GCFReturnStruct(reservasiResponse)
	}

	reservasiResponse.Status = true
	reservasiResponse.Message = "Get reservasi success"
	reservasiResponse.Data = reservasi

	return GCFReturnStruct(reservasiResponse)
}

func GCFHandlerUpdateReservasi(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)

	reservasiResponse.Status = false

	// get token from header
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		reservasiResponse.Message = "error parsing application/json1:"
		return GCFReturnStruct(reservasiResponse)
	}

	// decode token
	_, err1 := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	if err1 != nil {
		reservasiResponse.Message = "error parsing application/json2: " + err1.Error() + ";" + token
		return GCFReturnStruct(reservasiResponse)
	}

	err := json.NewDecoder(r.Body).Decode(&reservasi)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(reservasiResponse)
	}
	err = UpdateReservasi(conn, reservasi)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json4: " + err.Error()
		return GCFReturnStruct(reservasiResponse)
	}
	reservasiResponse.Status = true
	reservasiResponse.Message = "Reservasi berhasil diupdate"
	reservasiResponse.Data = []model.Reservasi{reservasi}
	return GCFReturnStruct(reservasiResponse)
}

func GCFHandlerDeleteReservasi(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	reservasiResponse.Status = false

	token := r.Header.Get("Authorization")
	if token == "" {
		reservasiResponse.Message = "error parsing application/json1:"
		return GCFReturnStruct(reservasiResponse)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)
	if err != nil {
		reservasiResponse.Message = "error parsing application/json2:" + err.Error() + ";" + token
		return GCFReturnStruct(reservasiResponse)
	}

	id := r.URL.Query().Get("_id")
	if id == "" {
		reservasiResponse.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(reservasiResponse)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		reservasiResponse.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(reservasiResponse)
	}

	_, err = DeleteReservasi(mconn, collectionname, ID)
	if err != nil {
		reservasiResponse.Message = err.Error()
		return GCFReturnStruct(reservasiResponse)
	}

	reservasiResponse.Status = true
	reservasiResponse.Message = "Delete todo success"

	return GCFReturnStruct(reservasiResponse)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

