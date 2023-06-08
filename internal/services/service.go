package services

import (
	"context"
	"zegen/gen/restapi/operations/authentication"
	"zegen/gen/restapi/operations/author"
	"zegen/internal/repositories"
	"zegen/runtime"
)

type (
	IService interface {
		userService
		authorService
	}

	userService interface {
		Register(ctx context.Context, req authentication.RegisterParams) (*uint64, error)
		Login(ctx context.Context, req *authentication.LoginParams) (token, expiredAt *string, err error)
	}

	authorService interface {
		CreateAuthor(ctx context.Context, form *author.CreateAuthorParams) (*uint64, error)
		UpdateAuthor(ctx context.Context, form *author.UpdateAuthorParams) error
		SoftDeleteAuthor(ctx context.Context, form *author.SoftDeleteAuthorParams) error
	}

	service struct {
		rt   runtime.Runtime
		repo repositories.IRepository
	}
)

func New(
	rt runtime.Runtime,
	repo repositories.IRepository,
) IService {
	return &service{
		rt,
		repo,
	}
}