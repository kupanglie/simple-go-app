package user

import (
	"github.com/kupanglie/simple-go-app/gql/graph/model"
	"golang.org/x/net/context"
)

type UserHandler interface {
	UserAdd(ctx context.Context, name string) (*model.UserData, error)
	UserEdit(ctx context.Context, userID string, name string) (*model.UserData, error)
	UserDelete(ctx context.Context, userID string) (*model.UserData, error)
	UserFind(ctx context.Context, userID string) (*model.UserData, error)
}
