package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jiwanjeon/go_todolist/graph"
	"github.com/jiwanjeon/go_todolist/graph/generated"
	"github.com/jiwanjeon/go_todolist/pkg/models"
	"github.com/jiwanjeon/go_todolist/pkg/routes"
	_ "github.com/lib/pq"
)
var db *gorm.DB;
const defaultPort = "8080"
func main() {
	models.Migrate()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	migrate := flag.Bool("fork", true, "bool")
	if *migrate == true{
		models.Migrate()
		router := mux.NewRouter()
		routes.Routes(router)

		//router.Handle("/", router)
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":9010", router))
	} else {
		fmt.Println("You should parsing argument with migrate")
	}
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}

// func initDB() {
// 	var err error
// 	dialect := os.Getenv("DIALECT")
// 	host := os.Getenv("HOST")
// 	dbPort := os.Getenv("DBPORT")
// 	user := os.Getenv("USER")
// 	dbName := os.Getenv("NAME")
// 	password := os.Getenv("PASSWORD")
// 	// Database connection string
// 	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	
// 	// Opening connection to database
// 	db, err = gorm.Open(dialect, dbURI)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("Successfully connected to database!")
// 	}

// 	db.AutoMigrate(&model.Todo{})
// }

// func main() {
// 	migrate := flag.Bool("fork", true, "bool")
// 	if *migrate == true {
// 		models.Migrate()
// 		router := mux.NewRouter()
// 		routes.Routes(router)
// 		http.Handle("/", router)
// 		log.Fatal(http.ListenAndServe(":9010", router))
// 	} else {
// 		fmt.Println("You should parsing argument with migrate")
// 	}
// }