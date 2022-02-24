package models_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jiwanjeon/go-todolist/pkg/models"
)

/*
var (
	testPerson = &Person{Name: "Jiwan", Email: "jeonjiwan94@gmail.com"}
	testTodos = []Todo{
		{ Title: "Golang Study", Description: "Richard Templar", Condition: false},
		{ Title: "Todo2", Description: "Description 2", Condition: false},
		{ Title: "Todo3", Description: "Description 3", Condition: false},
		{ Title: Todo4, Description: "Description 4", Condition: false},
	}
)
*/

func TestCreateTodo(t *testing.T) {
	// 직접 데이타베이스에 접근해도 되고 전역 변수를 생성해서 해도 됨
	b := &models.Todo{
		Title: "test-title",
		Description: "test-description",
		Condition: true,
	}

	submittedTodo := b.CreateTodo()

	// TODO: requre 확인하기
	require.Equal(t, b.Title, submittedTodo.Title)
	require.Equal(t, b.Description, submittedTodo.Description)
	require.Equal(t, b.Condition, submittedTodo.Condition)

	todoFromDb, _ := models.GetTodoById(int64(submittedTodo.ID))
	require.Equal(t, b.Title, todoFromDb.Title)
	require.Equal(t, b.Description, todoFromDb.Description)
	require.Equal(t, b.Condition, todoFromDb.Condition)
}

func TestDeleteTodo(t *testing.T) {
	require.Fail(t, "not implemented")

	b := &models.Todo{
		Title: "test-title",
		Description: "test-description",
		Condition: true,
	}
	submittedTodo := b.CreateTodo()

	todoFromDb, _ := models.GetTodoById(submittedTodo.ID)
	require.Equal(t, b.Title, todoFromDb.Title)
	require.Equal(t, b.Description, todoFromDb.Description)
	require.Equal(t, b.Condition, todoFromDb.Condition)

	// delete

	// 실제로 데이터베이스 
}

func TestModels(t *testing.T) {
	
}