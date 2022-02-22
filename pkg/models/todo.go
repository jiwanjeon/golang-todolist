package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jiwanjeon/go-todolist/pkg/config"
)

var db *gorm.DB

type Person struct {
	gorm.Model

	Name string
	Email string
	Todos []Todo
}

type Todo struct {
	gorm.Model
	// Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Condition bool `json:"conditions"`
}

var (
	person = &Person{Name: "Jiwan", Email: "jeonjiwan94@gmail.com"}
	todos = []Todo{
		{ Title: "Golang Study", Description: "Richard Templar", Condition: false},
		{ Title: "Todo2", Description: "Description 2", Condition: false},
		{ Title: "Todo3", Description: "Description 3", Condition: false},
	}
)

func init() {
	config.Connet()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&Person{})

	// DB에 데이터 테이블 생성하기 위한 코드
	// db.Create(person)
	// for idx := range todos {
	// 	db.Create(&todos[idx])
	// }
}

func (b *Todo) CreateTodo() *Todo{
	// TODO
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTodos() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}

func GetTodoById(Id int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	// TODO
	db := db.Where("ID = ?", Id).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(ID int64) Todo{
	var todo Todo
	db.Where("ID=?", ID).Delete(todo)
	return todo
}

func CompleteTodo(Condition bool) (*Todo, *gorm.DB) {
	var todo Todo
	db := db.Where("Condition = ?", Condition).Find(&todo)
	return &todo, db
}

func InCompleteTodo(Condition bool) (*Todo, *gorm.DB) {
	var todo Todo
	db := db.Where("Condition = ?", Condition).Find(&todo)
	return &todo, db
}