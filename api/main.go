package main

import (
	"api/controllers"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func main() {
	jsonString, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	type Config struct {
		User string `json:"user"`
		Pass string `json:"pass"`
	}

	c := new(Config)

	err = json.Unmarshal(jsonString, c)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	pgURL, err := pq.ParseURL("postgres://" + c.User + ":" + c.Pass + "@localhost:5432/openchat?sslmode=disable")
	if err != nil {
		log.Println("e1")
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Println("e2")
		log.Fatal(err)
	}
	controller := controllers.Controller{}
	router := mux.NewRouter()
	router.HandleFunc("/api/groups", controller.GetGroups(db)).Methods("GET")
	router.HandleFunc("/api/groups", controller.AddGroup(db)).Methods("POST")
	log.Println("Server up on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
