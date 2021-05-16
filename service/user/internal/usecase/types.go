package usecase

import (
	"context"

	"github.com/kupanglie/simple-go-app/service/user/internal/entity"
)

type User interface {
	Add(ctx context.Context, name string) (entity.User, error)
	Find(ctx context.Context, id string) (entity.User, error)
	Edit(ctx context.Context, id, name string) (entity.User, error)
	Delete(ctx context.Context, id string) (entity.User, error)
}
