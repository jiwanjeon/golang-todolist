package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jiwanjeon/go-todolist/pkg/models"
	"github.com/jiwanjeon/go-todolist/pkg/utils"
	_ "github.com/lib/pq"
)

var NewTodo models.Todo

func GetTodo(w http.ResponseWriter, r *http.Request) {
	newTodos := models.GetAllTodos()
	res, _ := json.Marshal(newTodos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	// TODO 
	fmt.Println("string일까 int일까?: ", todoId)
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

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	CreateTodo := &models.Todo{}
	utils.ParseBody(r, CreateTodo)
	b := CreateTodo.CreateTodo()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

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

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodo = &models.Todo{}
	utils.ParseBody(r, updateTodo)
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	// TODO: == 아닌가?? 
	if err != nil {
		fmt.Println("error while parsing")
	}
	todoDetails, db := models.GetTodoById(ID)
	if updateTodo.Title != "" {
		todoDetails.Title = updateTodo.Title
	}
	if updateTodo.Description != "" {
		todoDetails.Description = updateTodo.Description
	}
	if updateTodo.Condition != bool(false) || bool(true) {
		todoDetails.Condition = updateTodo.Condition
	}
	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodo = &models.Todo{}
	utils.ParseBody(r, updateTodo)
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	updateTodo.Condition = true
	todoDetails, db := models.GetTodoById(ID)
	todoDetails.Condition = updateTodo.Condition

	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func InCompleteTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodo = &models.Todo{}
	utils.ParseBody(r, updateTodo)
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	updateTodo.Condition = false
	todoDetails, db := models.GetTodoById(ID)
	todoDetails.Condition = updateTodo.Condition

	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}