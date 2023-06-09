package services

import (
	"context"
	"net/http"
	"time"
	"zegen/gen/models"
	"zegen/gen/restapi/operations/book"
	"zegen/internal/repositories"
	"zegen/internal/utils"

	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

func (s *service) CreateBook(ctx context.Context, form *book.CreateBookParams) (*uint64, error) {
	logger := s.rt.Logger.With().
		Interface("form", form.Data).
		Logger()

	authors, err := s.repo.FindAuthorsByIDs(ctx, form.Data.AuthorIds, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindAuthorsByIDs")
		return nil, err
	}

	if len(authors) != len(form.Data.AuthorIds) {
		return nil, s.rt.SetError(http.StatusBadRequest, utils.ERR_THERE_IS_DATA_DELETED)
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	bookAuthors := make([]*models.BookAuthor, len(authors))
	for i, v := range authors {
		bookAuthors[i] = &models.BookAuthor{
			BookAuthorData: models.BookAuthorData{
				AuthorID: *v.ID,
			},
			ModelTrackTime: models.ModelTrackTime{
				CreatedAt: &nowStrfmt,
			},
		}
	}

	data := &models.Book{
		ModelTrackTime: models.ModelTrackTime{
			CreatedAt: &nowStrfmt,
		},
		BookData: models.BookData{
			Title:         *form.Data.Title,
			PublishedYear: *form.Data.PublishedYear,
			Isbn:          *form.Data.Isbn,
		},
		BookForeignKey: models.BookForeignKey{
			Authors: bookAuthors,
		},
	}

	err = s.rt.Db.Transaction(func(tx *gorm.DB) error {
		data, err = s.repo.CreateBook(ctx, tx, data)
		if err != nil {
			logger.Error().Err(err).Msg("error repo.CreateBook")
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("error transaction")
		return nil, err
	}

	return data.ID, nil
}

func (s *service) UpdateBook(ctx context.Context, form *book.UpdateBookParams) error {
	logger := s.rt.Logger.With().
		Uint64("bookID", form.BookID).
		Interface("form", form.Data).
		Logger()

	filter := []repositories.ColumnValue{
		{
			Column: "id",
			Value:  form.BookID,
		},
	}
	currData, err := s.repo.FindOneBookByFilter(ctx, filter, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneBookByFilter")
		return err
	}

	authors, err := s.repo.FindAuthorsByIDs(ctx, form.Data.AuthorIds, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindAuthorsByIDs")
		return err
	}

	if len(authors) != len(form.Data.AuthorIds) {
		return s.rt.SetError(http.StatusBadRequest, utils.ERR_THERE_IS_DATA_DELETED)
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	bookAuthors := make([]*models.BookAuthor, 0)

	// map key: authorID
	// index location in authors
	mapCurrAuthorByAuthorID := make(map[uint64]int)
	for i, v := range currData.Authors {
		mapCurrAuthorByAuthorID[v.AuthorID] = i
		v.DeletedAt = &nowStrfmt
		bookAuthors = append(bookAuthors, v)
	}

	for _, v := range authors {
		val, ok := mapCurrAuthorByAuthorID[*v.ID]
		if ok {
			bookAuthors[val].DeletedAt = nil
		} else {
			bookAuthors = append(bookAuthors, &models.BookAuthor{
				BookAuthorData: models.BookAuthorData{
					AuthorID: *v.ID,
				},
				ModelTrackTime: models.ModelTrackTime{
					CreatedAt: &nowStrfmt,
				},
			})
		}
	}

	currData.Title = *form.Data.Title
	currData.PublishedYear = *form.Data.PublishedYear
	currData.Isbn = *form.Data.Isbn
	currData.Authors = bookAuthors
	currData.UpdatedAt = &nowStrfmt

	err = s.rt.Db.Transaction(func(tx *gorm.DB) error {
		err = s.repo.UpdateBook(ctx, tx, currData)
		if err != nil {
			logger.Error().Err(err).Msg("error repo.UpdateBook")
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("error transaction")
		return err
	}

	return nil
}

func (s *service) SoftDeleteBook(ctx context.Context, form *book.SoftDeleteBookParams) error {
	logger := s.rt.Logger.With().
		Uint64("bookID", form.BookID).
		Logger()

	filter := []repositories.ColumnValue{
		{
			Column: "id",
			Value:  form.BookID,
		},
	}
	_, err := s.repo.FindOneBookByFilter(ctx, filter, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneBookByFilter")
		return err
	}

	now := time.Now().UTC()

	err = s.repo.SoftDeleteBook(ctx, nil, form.BookID, now)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.SoftDeleteBook")
		return err
	}

	return nil
}
