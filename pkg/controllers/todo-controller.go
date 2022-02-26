package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/jiwanjeon/go-todolist/pkg/config"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jiwanjeon/go-todolist/pkg/models"
	"github.com/jiwanjeon/go-todolist/pkg/utils"
	_ "github.com/lib/pq"
)

var NewTodo models.Todo

// Receiving All Todo-list Data
func GetTodo(w http.ResponseWriter, r *http.Request) {
	newTodos := models.GetAllTodos()
	res, _ := json.Marshal(newTodos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res)
}

// Receiving Todo-list specific ID Data 
func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todoDetails, _ := models.GetTodoById(ID)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Creating Todo-list
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	CreateTodo := &models.Todo{}
	utils.ParseBody(r, CreateTodo)
	CreateTodoList := CreateTodo.CreateTodo()
	res, _ := json.Marshal(CreateTodoList)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Delete Todo-list
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todo := models.DeleteTodo(ID)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Update Todo-list information
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var parsedRequest = &models.Todo{}
	utils.ParseBody(r, parsedRequest)
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		fmt.Println("error while doing something")
	}

	convertIntTodoId := int64(todoId)
	todoDetails, _ := models.GetTodoById(convertIntTodoId)

	fmt.Println("(before)todoDetails: ", todoDetails)
	if parsedRequest.Title != "" {
		todoDetails.Title = parsedRequest.Title
	}
	if parsedRequest.Description != "" {
		todoDetails.Description = parsedRequest.Description
	}
	if parsedRequest.Condition != bool(false) || bool(true) {
		todoDetails.Condition = parsedRequest.Condition
	}
	DatabaseUpdate(&todoDetails, w)
}

func DatabaseUpdate(todoDetails interface{}, w http.ResponseWriter) {
	db := config.GetDB()
	db.Save(todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.Write(res)
}

// Change Todo-list Condition change into true
func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	var parsedRequest = &models.Todo{}
	utils.ParseBody(r, parsedRequest)
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	parsedRequest.Condition = true
	todoDetails, db := models.GetTodoById(ID)
	todoDetails.Condition = parsedRequest.Condition

	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Change Todo-list Condition change into false
func InCompleteTodo(w http.ResponseWriter, r *http.Request) {
	var parsedRequest = &models.Todo{}
	utils.ParseBody(r, parsedRequest)
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	parsedRequest.Condition = false
	todoDetails, db := models.GetTodoById(ID)
	todoDetails.Condition = parsedRequest.Condition

	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}