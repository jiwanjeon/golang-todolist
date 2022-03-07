package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jiwanjeon/go_todolist/graph/generated"
	"github.com/jiwanjeon/go_todolist/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	// todo := &model.Todo{
	// 	ID:          input.ID,
	// 	Title:       input.Title,
	// 	Description: input.Description,
	// 	Condition:   input.Condition,
	// }
	// r.TodosList = append(r.TodosList, todo)
	// return todo, nil

	todo := model.Todo{
		// ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Condition:   input.Condition,
	}

	r.DB.Create(&todo)
	return &todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID int) (bool, error) {
	r.DB.Where("ID = ?", todoID).Delete(&model.Todo{})
	return true, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, todoID int, input model.TodoInput) (*model.Todo, error) {
	// helper := int64(todoID)
	// ID, _ := strconv.ParseInt(todoID, 0, 0)
	updatedTodo := model.Todo{
		ID:          todoID,
		Title:       input.Title,
		Description: input.Description,
		Condition:   input.Condition,
	}
	r.DB.Save(&updatedTodo)
	return &updatedTodo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// return r.TodosList, nil
	// return []*model.Todo{
	// 	{
	// 		ID:          "123",
	// 		Title:       "Jiwan",
	// 		Description: "desc1",
	// 		Condition:   true,
	// 	},
	// }, nil
	var todos []*model.Todo
	r.DB.Find(&todos)
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
