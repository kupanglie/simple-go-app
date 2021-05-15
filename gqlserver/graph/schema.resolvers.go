package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kupanglie/simple-go-app/gql/graph/generated"
	"github.com/kupanglie/simple-go-app/gql/graph/model"
)

func (r *mutationResolver) UserAdd(ctx context.Context, name string) (*model.UserData, error) {
	panic(fmt.Errorf("asd"))
}

func (r *mutationResolver) UserEdit(ctx context.Context, userID string, name string) (*model.UserData, error) {
	panic(fmt.Errorf("dsa"))
}

func (r *mutationResolver) UserDelete(ctx context.Context, userID string) (*model.UserData, error) {
	panic(fmt.Errorf("sda"))
}

func (r *queryResolver) UserFind(ctx context.Context, userID string) (*model.UserData, error) {
	panic(fmt.Errorf("dassw"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
