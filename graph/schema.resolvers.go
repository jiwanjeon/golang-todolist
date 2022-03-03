package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jiwanjeon/go_todolist/graph/generated"
	"github.com/jiwanjeon/go_todolist/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	todo := &model.Todo{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Condition:   input.Condition,
	}
	// err := r.DB.Create(&todo).Error
	// if err != nil {
	// 	return nil, err
	// }
	r.TodosList = append(r.TodosList, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.TodosList, nil
	// return []*model.Todo{
	// 	{
	// 		ID:          "123",
	// 		Title:       "Jiwan",
	// 		Description: "desc1",
	// 		Condition:   true,
	// 	},
	// }, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
