package usecase

import (
	"context"

	"github.com/kupanglie/simple-go-app/service/user/internal/entity"
	"github.com/kupanglie/simple-go-app/service/user/internal/repository"
)

type user struct {
	userRP repository.User
}

func NewUserUC(userRP repository.User) *user {
	return &user{
		userRP: userRP,
	}
}

func (u *user) Add(ctx context.Context, name string) (entity.User, error) {
	return u.userRP.Add(ctx, name)
}

func (u *user) Edit(ctx context.Context, id, name string) (entity.User, error) {
	return u.userRP.Edit(ctx, id, name)
}

func (u *user) Find(ctx context.Context, id string) (entity.User, error) {
	return u.userRP.Find(ctx, id)
}

func (u *user) Delete(ctx context.Context, id string) (entity.User, error) {
	return u.userRP.Delete(ctx, id)
}
