package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jiwanjeon/go_todolist/pkg/config"
)

var db *gorm.DB

// gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

type Person struct {
	gorm.Model

	Name string
	Email string
	Todos []Todo
}

type Todo struct {
	gorm.Model

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
	// TODO : https://gorm.io/docs/query.html, https://gorm.io/docs/advanced_query.html
	db := db.Where("ID = ?", Id).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(ID int64) Todo{
	var todo Todo
	db.Where("ID=?", ID).Delete(todo)
	return todo
}