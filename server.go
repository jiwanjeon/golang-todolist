package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jiwanjeon/go_todolist/graph"
	"github.com/jiwanjeon/go_todolist/graph/generated"
	"github.com/jiwanjeon/go_todolist/graph/model"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"
var db *gorm.DB;
func main() {
	initDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/todo"))
	http.Handle("/todo", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB() {
	var err error
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// Opening connection to database
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
		fmt.Println("This is server.go file")
	}
	db.AutoMigrate(&model.Todo{})
}

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 			port = defaultPort
// 	}

// 	initDB()
// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", handler.GraphQL(go_todolist.NewExecutableSchema(go_todolist.Config{Resolvers: &go_todolist.Resolver{
// 		DB: db,
// }})))

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }