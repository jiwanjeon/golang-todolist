package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jiwanjeon/go-todolist/pkg/models"
	"github.com/jiwanjeon/go-todolist/pkg/routes"
	_ "github.com/lib/pq"
)

func main() {
	migrate := flag.Bool("fork", true, "bool")
	if *migrate == true {
		models.Migrate()
		router := mux.NewRouter()
		routes.Routes(router)
		http.Handle("/", router)
		log.Fatal(http.ListenAndServe(":9010", router))
	} else {
		fmt.Println("You should parsing argument with migrate")
	}

}