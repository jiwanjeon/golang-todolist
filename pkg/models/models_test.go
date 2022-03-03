package models_test

import (
	"testing"

	"github.com/jiwanjeon/go_todolist/pkg/models"
	"github.com/stretchr/testify/require"
)


func TestCreateTodo(t *testing.T) {
	// 직접 데이타베이스에 접근해도 되고 전역 변수를 생성해서 해도 됨
	dbData := &models.Todo{
		Title: "test-title-console-check",
		Description: "test-description-console-check",
		Condition: true,
	}

	submittedTodo := dbData.CreateTodo()

	require.Equal(t, dbData.Title, submittedTodo.Title,"They should be equal")
	require.Equal(t, dbData.Description, submittedTodo.Description,"They should be equal")
	require.Equal(t, dbData.Condition, submittedTodo.Condition,"They should be equal")
	todoFromDb, _ := models.GetTodoById(int64(submittedTodo.ID))
	require.Equal(t, dbData.Title, todoFromDb.Title,"They should be equal")
	require.Equal(t, dbData.Description, todoFromDb.Description,"They should be equal")
	require.Equal(t, dbData.Condition, todoFromDb.Condition,"They should be equal")
}

func TestDeleteTodo(t *testing.T) {
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
	require.Equal(t, dbData.Title, submittedTodo.Title,"They should be equal")
	require.Equal(t, dbData.Description, submittedTodo.Description,"They should be equal")
	require.Equal(t, dbData.Condition, submittedTodo.Condition,"They should be equal")
}