package repositories

import (
	"context"
	"zegen/gen/models"

	"gorm.io/gorm"
)

func (r *repository) CreateAuthor(ctx context.Context, tx *gorm.DB, data *models.Author) (*models.Author, error) {
	logger := r.rt.Logger.With().
		Interface("data", data).
		Logger()

	if tx == nil {
		tx = r.rt.Db
	}

	err := tx.Select("*").Create(&data).Error
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return data, nil
}
