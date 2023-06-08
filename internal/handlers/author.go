package handlers

import (
	"context"
	"time"
	"zegen/gen/models"
	"zegen/gen/restapi/operations/author"

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
