package module
import (
	"context"
	// "crypto/rand"
	// "encoding/hex"
	"errors"
	"fmt"
	"os"
	// "strings"

	// "github.com/badoux/checkmail"
	"github.com/bimbingankonseling/kee/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/crypto/argon2"
)

func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// crud
func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server : insert")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

//login
func Login(db *mongo.Database, col string, insertedDoc model.User) (user model.User, Status bool, err error) {
	if insertedDoc.Username == "" || insertedDoc.Password == "" {
		return user, false, fmt.Errorf("mohon untuk melengkapi data")
	}

	// Periksa apakah pengguna dengan username tertentu ada
	userExists, _ := GetUserFromUsername(db, col, insertedDoc.Username)
	if userExists.Username == "" {
		err = fmt.Errorf("Username tidak ditemukan")
		return user, false, err
	}
	// Periksa apakah kata sandi benar
	if !CheckPasswordHash(insertedDoc.Password, userExists.Password) {
		err = fmt.Errorf("Password salah")
		return user, false, err
	}

	return userExists, true, nil
}

func GetUserFromUsername(db *mongo.Database, col string, username string) (user model.User, err error) {
	cols := db.Collection(col)
	filter := bson.M{"username": username}

	err = cols.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err := fmt.Errorf("no data found for username %s", username)
			return user, err
		}

		err := fmt.Errorf("error retrieving data for username %s: %s", username, err.Error())
		return user, err
	}

	return user, nil
}

//reservasi
func InsertReservasi(db *mongo.Database, col string, nama string, no_telp string, ttl string, status string, keluhan string) (insertedID primitive.ObjectID, err error) {
	reservasi := bson.M{
		"nama"		: nama,
		"no_telp"	: no_telp,
		"ttl"		: ttl,
		"status"	: status,
		"keluhan"	: keluhan,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), reservasi)
	if err != nil {
		fmt.Printf("InsertReservasi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllReservasi(db *mongo.Database, col string) (reservasi []model.Reservasi, err error) {
	cols := db.Collection(col)
	filter := bson.M{}

	cursor, err := cols.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("Error GetReservasi in colection", col, ":", err)
		return nil, err
	}

	err = cursor.All(context.Background(), &reservasi)
	if err != nil {
		fmt.Println(err)
	}

	return reservasi, nil
}

func UpdateReservasi(db *mongo.Database, doc model.Reservasi) (err error) {
	filter := bson.M{"_id": doc.ID}
	result, err := db.Collection("reservasi").UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		fmt.Printf("UpdateReservasi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified id")
		return
	}
	return nil
}

func DeleteReservasi(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
	cols := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := cols.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		err = fmt.Errorf("Data tidak berhasil dihapus")
		return false, err
	}

	return true, nil
}

//registrasi
func InsertRegistrasi(db *mongo.Database, col string, nama_lengkap string, no_telp string, ttl string, nim string, alamat string) (insertedID primitive.ObjectID, err error) {
	registrasi := bson.M{
		"nama_lengkap"	: nama_lengkap,
		"no_telp"		: no_telp,
		"ttl"			: ttl,
		"nim"			: nim,
		"alamat"		: alamat,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), registrasi)
	if err != nil {
		fmt.Printf("InsertRegistrasi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllRegistrasi(db *mongo.Database, col string) (registrasi []model.Registrasi, err error) {
	cols := db.Collection(col)
	filter := bson.M{}

	cursor, err := cols.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("Error GetRegistrasi in colection", col, ":", err)
		return nil, err
	}

	err = cursor.All(context.Background(), &registrasi)
	if err != nil {
		fmt.Println(err)
	}

	return registrasi, nil
}

func UpdateRegistrasi(db *mongo.Database, doc model.Registrasi) (err error) {
	filter := bson.M{"_id": doc.ID}
	result, err := db.Collection("registrasi").UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		fmt.Printf("UpdateRegistrasi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified id")
		return
	}
	return nil
}

func DeleteRegistrasi(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
	cols := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := cols.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount == 0 {
		err = fmt.Errorf("Data tidak berhasil dihapus")
		return false, err
	}

	return true, nil
}

// func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	datagedung := GetAllBangunanLineString(mconn, collectionname)
// 	return GCFReturnStruct(datagedung)
// }

// func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	var Response Credential
// 	Response.Status = false
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var datauser User
// 	err := json.NewDecoder(r.Body).Decode(&datauser)
// 	if err != nil {
// 		Response.Message = "error parsing application/json: " + err.Error()
// 	} else {
// 		if IsPasswordValid(mconn, collectionname, datauser) {
// 			Response.Status = true
// 			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
// 			if err != nil {
// 				Response.Message = "Gagal Encode Token : " + err.Error()
// 			} else {
// 				Response.Message = "Selamat Datang"
// 				Response.Token = tokenstring
// 			}
// 		} else {
// 			Response.Message = "Password Salah"
// 		}
// 	}

// 	return GCFReturnStruct(Response)
// }

// func GCFReturnStruct(DataStuct any) string {
// 	jsondata, _ := json.Marshal(DataStuct)
// 	return string(jsondata)
// }

// func InsertUser(db *mongo.Database, collection string, userdata User) string {
// 	hash, _ := HashPassword(userdata.Password)
// 	userdata.Password = hash
// 	atdb.InsertOneDoc(db, collection, userdata)
// 	return "Ini username : " + userdata.Username + "ini password : " + userdata.Password
// }