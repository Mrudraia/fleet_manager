package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mrudraia/fleet_manager/aws"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	DatabaseConnection()
}

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "root"
	password := "root"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, dbUser, dbName, password)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database ...", err)
	}

	fmt.Println("Database connection successful ...")
}

func main() {
	fmt.Println("Fleet Manager")
	handleRequests()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside home page")
	fmt.Fprintf(w, "Welcome to the HomePage!")
	json.Marshal("dfdfd")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/v1/account", aws.GetAccount)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
