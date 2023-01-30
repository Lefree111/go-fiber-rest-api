package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Data struct {
	Id      string `json:"id"`
	User_id int    `json:"user_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type Data2 struct {
	Id      int    `json:"id"`
	User_id int    `json:"user_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type Read_Api_version_1 struct {
	Data []Data2 `json:"data"`
}

var db *gorm.DB
var err error

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func NewPostgreSQLClient() {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USER")
		dbname   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(Data{})
}

// /**
// * Read api version 1
//  */

func Read_version_1() {

	res, err := http.Get("https://gorest.co.in/public/v1/posts")
	CheckErr(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	CheckErr(err)

	var responce Read_Api_version_1

	json.Unmarshal(body, &responce)

	for _, p := range responce.Data {
		fmt.Println("id: ", p.Id)
		fmt.Println("user_id: ", p.User_id)
		fmt.Println("title: ", p.Title)
		fmt.Println("body: ", p.Body)
	}

}

// /**
// * Read api version 2
//  */

func Read_version_2() {
	res, err := http.Get("https://gorest.co.in/public/v2/posts")
	CheckErr(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	CheckErr(err)

	var responce []Data

	json.Unmarshal(body, &responce)

	for _, p := range responce {
		fmt.Println("id: ", p.Id)
		fmt.Println("user_id: ", p.User_id)
		fmt.Println("title: ", p.Title)
		fmt.Println("body: ", p.Body)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
