package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

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

func (r *mutationResolver) UpdateTodo(ctx context.Context, todoID int, input model.TodoInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdateOrder(ctx context.Context, todoID int, input model.TodoUpdateInput) (*model.Todo, error) {
	updatedTodo := model.Todo{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Condition:   input.Condition,
	}
	r.DB.Save(&updatedTodo)
	return &updatedTodo, nil
}
