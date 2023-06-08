package handlers

import (
	"context"
	"zegen/gen/restapi/operations/authentication"
	"zegen/gen/restapi/operations/author"
	"zegen/internal/repositories"
	"zegen/runtime"
)

type (
	IHandler interface {
		userHandler
		authorHandler
	}

	userHandler interface {
		Register(ctx context.Context, req authentication.RegisterParams) (*uint64, error)
		Login(ctx context.Context, req *authentication.LoginParams) (token, expiredAt *string, err error)
	}

	authorHandler interface {
		CreateAuthor(ctx context.Context, form *author.CreateAuthorParams) (*uint64, error)
	}

	handler struct {
		rt   runtime.Runtime
		repo repositories.IRepository
	}
)

func NewHandler(
	rt runtime.Runtime,
	repo repositories.IRepository,
) IHandler {
	return &handler{
		rt,
		repo,
	}
}
