package models_test

import (
	"testing"
	"fmt"

	// Unit Test하는데 도움주는 Library
	"github.com/stretchr/testify/require"
	// Require : it terminate test if assertion is not successful.
	// "github.com/stretchr/testify/assert"
	// Assert: it returns boolean value. It returns true if assertion is successful
	// and false otherwise, allowing other assertions to be executed in spite of failures upstream.

	"github.com/jiwanjeon/go-todolist/pkg/models"
)

// func TestGetAllTodo(t *testing.T) {
// 	dbData := &models.Todo{
// 		Title: "Test-getAllTodo-Title",
// 		Description: "Test-getAllTodo-Description",
// 		Condition: false,
// 	}
	
// 	submittedTodo := models.GetAllTodos()
// 	fmt.Println(submittedTodo.Title)
// 	// require.Equal(t, dbData.ID, submittedTodo.ID,"They should be equal")
// 	// require.Equal(t, dbData.Title, submittedTodo.Title, "They Should be equal")
// 	// require.Equal(t, dbData.Description, submittedTodo.Description, "They Should be equal")
// 	// require.Equal(t, dbData.Condition, submittedTodo.Condition, "They Should be equal")
// }

func TestCreateTodo(t *testing.T) {
	// 직접 데이타베이스에 접근해도 되고 전역 변수를 생성해서 해도 됨
	dbData := &models.Todo{
		Title: "test-title-console-check",
		Description: "test-description-console-check",
		Condition: true,
	}

	submittedTodo := dbData.CreateTodo()

	// TODO: require 확인하기
	require.Equal(t, dbData.Title, submittedTodo.Title,"They should be equal")
	require.Equal(t, dbData.Description, submittedTodo.Description,"They should be equal")
	require.Equal(t, dbData.Condition, submittedTodo.Condition,"They should be equal")
	fmt.Println("DataBase: ", dbData.ID)
	todoFromDb, _ := models.GetTodoById(int64(submittedTodo.ID))
	fmt.Println("After Creation: ", todoFromDb.ID)
	require.Equal(t, dbData.Title, todoFromDb.Title,"They should be equal")
	require.Equal(t, dbData.Description, todoFromDb.Description,"They should be equal")
	require.Equal(t, dbData.Condition, todoFromDb.Condition,"They should be equal")
}

func TestDeleteTodo(t *testing.T) {
	// require.Fail(t, "not implemented")

	dbData := &models.Todo{
		Title: "delete-test-title",
		Description: "delete-test-description",
		Condition: true,
	}
	submittedTodo := dbData.CreateTodo()

	todoFromDb, _ := models.GetTodoById(int64(submittedTodo.ID))
	deleteTodoFromDb := models.DeleteTodo(int64(todoFromDb.ID))
	require.NotEqual(t, dbData.Title, deleteTodoFromDb.Title, "they should not be equal")
	require.NotEqual(t, dbData.Description, deleteTodoFromDb.Description, "they should not be equal")
	require.NotEqual(t, dbData.Condition, deleteTodoFromDb.Condition, "they should not be equal")
}

func TestGetTodoById(t *testing.T) {
	dbData := &models.Todo{
		Title: "GetId-test-title",
		Description: "GetId-test-description",
		Condition: true,
	}

	submittedTodo := dbData.CreateTodo()
	fmt.Println("dbId: ", dbData.ID)

	// require.Equal(t, dbData.ID, submittedTodo.ID,"They should be equal")
	require.Equal(t, dbData.Title, submittedTodo.Title,"They should be equal")
	require.Equal(t, dbData.Description, submittedTodo.Description,"They should be equal")
	require.Equal(t, dbData.Condition, submittedTodo.Condition,"They should be equal")
}