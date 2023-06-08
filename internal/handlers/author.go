package handlers

import (
	"context"
	"time"
	"zegen/gen/models"
	"zegen/gen/restapi/operations/author"
	"zegen/internal/repositories"

	"github.com/go-openapi/strfmt"
)

func (h *handler) CreateAuthor(ctx context.Context, form *author.CreateAuthorParams) (*uint64, error) {
	logger := h.rt.Logger.With().
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

	data, err = h.repo.CreateAuthor(ctx, nil, data)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.CreateAuthor")
		return nil, err
	}

	return data.ID, nil
}

func (h *handler) UpdateAuthor(ctx context.Context, form *author.UpdateAuthorParams) error {
	logger := h.rt.Logger.With().
		Uint64("authorID", form.AuthorID).
		Interface("form", form.Data).
		Logger()

	filter := []repositories.ColumnValue{
		{
			Column: "id",
			Value:  form.AuthorID,
		},
	}
	currData, err := h.repo.FindOneAuthorByFilter(ctx, filter, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindOneAuthorByFilter")
		return err
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)

	currData.Name = *form.Data.Name
	currData.Country = *form.Data.Country
	currData.UpdatedAt = &nowStrfmt

	err = h.repo.UpdateAuthor(ctx, nil, currData)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.UpdateAuthor")
		return err
	}

	return nil
}
