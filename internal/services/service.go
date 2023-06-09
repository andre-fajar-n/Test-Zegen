package services

import (
	"context"
	"zegen/gen/models"
	"zegen/gen/restapi/operations/authentication"
	"zegen/gen/restapi/operations/author"
	"zegen/gen/restapi/operations/book"
	"zegen/internal/repositories"
	"zegen/runtime"
)

type (
	IService interface {
		userService
		authorService
		bookService
	}

	userService interface {
		Register(ctx context.Context, req authentication.RegisterParams) (*uint64, error)
		Login(ctx context.Context, req *authentication.LoginParams) (token, expiredAt *string, err error)
	}

	authorService interface {
		CreateAuthor(ctx context.Context, form *author.CreateAuthorParams) (*uint64, error)
		UpdateAuthor(ctx context.Context, form *author.UpdateAuthorParams) error
		SoftDeleteAuthor(ctx context.Context, form *author.SoftDeleteAuthorParams) error
		FindOneAuthor(ctx context.Context, form *author.FindOneAuthorParams) (*models.SuccessFindOneAuthorAllOf1, error)
	}

	bookService interface {
		CreateBook(ctx context.Context, form *book.CreateBookParams) (*uint64, error)
		UpdateBook(ctx context.Context, form *book.UpdateBookParams) error
		SoftDeleteBook(ctx context.Context, form *book.SoftDeleteBookParams) error
		FindOneBook(ctx context.Context, form *book.FindOneBookParams) (*models.SuccessFindOneBookAllOf1, error)
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
