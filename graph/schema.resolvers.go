package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jiwanjeon/go-todolist/graph/generated"
	"github.com/jiwanjeon/go-todolist/graph/model"
	"github.com/jiwanjeon/go-todolist/pkg/models"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Condition:   input.Condition,
	}

	err := r.DB.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	// r.TodosList = append(r.TodosList, todo)
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID int) (bool, error) {

	r.DB.Where("ID = ?", todoID).Delete(&models.Todo{})
	return true, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// return r.TodosList, nil

	// Adding a todo-list
	return []*model.Todo{
		{
			ID:          "123",
			Title:       "Jiwan",
			Description: "desc1",
			Condition:   true,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
