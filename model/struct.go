package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string 			`json:"username" bson:"username"`
	Password string 			`json:"password" bson:"password"`
}

type Reservasi struct {
	ID		primitive.ObjectID	`bson:"_id,omitempty" json:"_id,omitempty"`
	Nama	string				`json:"nama" bson:"nama"`
	No_telp string				`json:"no_telp" bson:"no_telp"`
	TTL		string				`json:"ttl" bson:"ttl"`
	Status	string				`json:"status" bson:"status"`
	Keluhan	string				`json:"keluhan" bson:"keluhan"`
}

type Registrasi struct {
	ID				primitive.ObjectID	`bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_lengkap	string				`json:"nama_lengkap" bson:"nama_lengkap"`
	No_telp			string				`json:"no_telp" bson:"no_telp"`
	TTL				string				`json:"ttl" bson:"ttl"`
	NIM				string				`json:"nim" bson:"nim"`
	Alamat			string				`json:"alamat" bson:"alamat"`
}

type Response struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type ReservasiResponse struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message,omitempty" bson:"message,omitempty"`
	Data    []Reservasi `json:"data" bson:"data"`
}

type RegistrasiResponse struct {
	Status  bool        	`json:"status" bson:"status"`
	Message string			`json:"message,omitempty" bson:"message,omitempty"`
	Data    []Registrasi 	`json:"data" bson:"data"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    []User `bson:"data,omitempty" json:"data,omitempty"`
}
