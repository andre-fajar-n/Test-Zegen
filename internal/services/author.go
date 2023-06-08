package services

import (
	"context"
	"time"
	"zegen/gen/models"
	"zegen/gen/restapi/operations/author"
	"zegen/internal/repositories"

	"github.com/go-openapi/strfmt"
)

func (s *service) CreateAuthor(ctx context.Context, form *author.CreateAuthorParams) (*uint64, error) {
	logger := s.rt.Logger.With().
		Interface("form", form.Data).
		Logger()

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	data := &models.Author{
		ModelTrackTime: models.ModelTrackTime{
			CreatedAt: &nowStrfmt,
		},
		AuthorData: models.AuthorData{
			Name:    *form.Data.Name,
			Country: *form.Data.Country,
		},
	}
	var err error

	data, err = s.repo.CreateAuthor(ctx, nil, data)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.CreateAuthor")
		return nil, err
	}

	return data.ID, nil
}

func (s *service) UpdateAuthor(ctx context.Context, form *author.UpdateAuthorParams) error {
	logger := s.rt.Logger.With().
		Uint64("authorID", form.AuthorID).
		Interface("form", form.Data).
		Logger()

	filter := []repositories.ColumnValue{
		{
			Column: "id",
			Value:  form.AuthorID,
		},
	}
	currData, err := s.repo.FindOneAuthorByFilter(ctx, filter, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneAuthorByFilter")
		return err
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	currData.Name = *form.Data.Name
	currData.Country = *form.Data.Country
	currData.UpdatedAt = &nowStrfmt

	err = s.repo.UpdateAuthor(ctx, nil, currData)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.UpdateAuthor")
		return err
	}

	return nil
}

func (s *service) SoftDeleteAuthor(ctx context.Context, form *author.SoftDeleteAuthorParams) error {
	logger := s.rt.Logger.With().
		Uint64("authorID", form.AuthorID).
		Logger()

	filter := []repositories.ColumnValue{
		{
			Column: "id",
			Value:  form.AuthorID,
		},
	}
	_, err := s.repo.FindOneAuthorByFilter(ctx, filter, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneAuthorByFilter")
		return err
	}

	now := time.Now().UTC()

	err = s.repo.SoftDeleteAuthor(ctx, nil, form.AuthorID, now)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.SoftDeleteAuthor")
		return err
	}

	return nil
}
