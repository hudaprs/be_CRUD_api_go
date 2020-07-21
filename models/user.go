package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
)

// User Struct
type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
}

var err error
var db *gorm.DB

// Migration GORM Auto Migration
func Migration() {
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")

	db, err := gorm.Open("postgres", DBURI)

	if err != nil {
		fmt.Println("Failed connect to database")
	}
	defer db.Close()

	db.Debug().AutoMigrate(&User{})
}

// GetUsers Get All Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")

	db, err := gorm.Open("postgres", DBURI)

	if err != nil {
		fmt.Println("Failed connect to database")
	}
	defer db.Close()

	var users []User

	db.Table("users").Find(&users)

	json.NewEncoder(w).Encode(&users)
}

// CreateUser creating a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")

	db, err := gorm.Open("postgres", DBURI)

	if err != nil {
		fmt.Println("Failed connect to database")
	}
	defer db.Close()

	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)
	newUser := db.Save(&user)

	json.NewEncoder(w).Encode(newUser)
}

// GetUserByID get user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")

	db, err := gorm.Open("postgres", DBURI)

	if err != nil {
		fmt.Println("Failed connect to database")
	}
	defer db.Close()

	var user User

	id := mux.Vars(r)["id"]

	db.Table("users").Where("id = ?", id).Find(&user)

	json.NewEncoder(w).Encode(&user)
}

// UpdateUser updating existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")

	db, err := gorm.Open("postgres", DBURI)

	if err != nil {
		fmt.Println("Failed connect to database")
	}
	defer db.Close()

	id := mux.Vars(r)["id"]

	var user User

	db.Table("users").Where("id = ?", id).Find(&user)
	_ = json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

// DeleteUser delete existing user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")

	db, err := gorm.Open("postgres", DBURI)

	if err != nil {
		fmt.Println("Failed connect to database")
	}
	defer db.Close()

	var user User

	id := mux.Vars(r)["id"]

	db.Table("users").Where("id = ?", id).Find(&user)
	db.Unscoped().Delete(&user)
	json.NewEncoder(w).Encode(&user)
}
