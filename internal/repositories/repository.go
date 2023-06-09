package repositories

import (
	"context"
	"time"
	"zegen/gen/models"
	"zegen/runtime"

	"gorm.io/gorm"
)

type (
	IRepository interface {
		userRepository
		authorRepository
		bookRepository
	}

	authorRepository interface {
		CreateAuthor(ctx context.Context, tx *gorm.DB, data *models.Author) (*models.Author, error)
		FindOneAuthorByFilter(ctx context.Context, filter []ColumnValue, isDeleted bool) (*models.Author, error)
		UpdateAuthor(ctx context.Context, tx *gorm.DB, data *models.Author) error
		SoftDeleteAuthor(ctx context.Context, tx *gorm.DB, authorID uint64, deletedAt time.Time) error
		FindAuthorsByIDs(ctx context.Context, IDs []uint64, isDeleted bool) ([]models.Author, error)
	}

	bookRepository interface {
		CreateBook(ctx context.Context, tx *gorm.DB, data *models.Book) (*models.Book, error)
		FindOneBookByFilter(ctx context.Context, filter []ColumnValue, isDeleted bool) (*models.Book, error)
		UpdateBook(ctx context.Context, tx *gorm.DB, data *models.Book) error
		SoftDeleteBook(ctx context.Context, tx *gorm.DB, data *models.Book) error
	}

	userRepository interface {
		CreateUser(ctx context.Context, tx *gorm.DB, data *models.User) (*models.User, error)
		FindUserBySingleColumn(ctx context.Context, column string, value interface{}, isDeleted bool) (*models.User, error)
		UsernameExist(ctx context.Context, username string) (bool, error)
	}

	repository struct {
		rt runtime.Runtime
	}

	ColumnValue struct {
		Column string
		Value  interface{}
	}
)

func New(rt runtime.Runtime) IRepository {
	return &repository{
		rt,
	}
}
