package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jiwanjeon/go-todolist/pkg/config"
)

var db *gorm.DB

//TODO: Interface 화
// type Manager interface {
// 	GetAll()
// 	Get()
// 	Create()
// }

type Person struct {
	gorm.Model

	Name string
	Email string
	Todos []Todo
}

type Todo struct {
	gorm.Model
	
	// todo
	// Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Condition bool `json:"conditions"`
}

//TODO: Interface 화
// type TodoManger struct {
// 	Manager
// }

var (
	person = &Person{Name: "Jiwan", Email: "jeonjiwan94@gmail.com"}
	todos = []Todo{
		{ Title: "Golang Study", Description: "Richard Templar", Condition: false},
		{ Title: "Todo2", Description: "Description 2", Condition: false},
		{ Title: "Todo3", Description: "Description 3", Condition: false},
	}
)

func Migrate() {
	db.AutoMigrate(&Todo{})
	db.AutoMigrate(&Person{})
}

func init() {
	config.Connet()
	db = config.GetDB()

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

//TODO: Interface 화
// func (m *TodoManger) GetAll() []Todo {
// 	var x []Todo
// 	db.Find(&x)
// 	return x
// }

func GetAllTodos() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}

func GetTodoById(Id int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	// TODO : https://gorm.io/docs/query.html, https://gorm.io/docs/advanced_query.html
	db := db.Where("ID = ?", Id).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(ID int64) Todo{
	var todo Todo
	db.Where("ID=?", ID).Delete(todo)
	return todo
}