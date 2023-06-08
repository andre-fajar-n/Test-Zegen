package repositories

import (
	"context"
	"net/http"
	"zegen/gen/models"
	"zegen/internal/utils"

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

func (r *repository) FindOneAuthorByFilter(ctx context.Context, filter []ColumnValue, isDeleted bool) (*models.Author, error) {
	logger := r.rt.Logger.With().
		Interface("filter", filter).
		Bool("isDeleted", isDeleted).
		Logger()

	authorModel := models.Author{}
	db := r.rt.Db.Model(&authorModel)

	for _, v := range filter {
		db = db.Where(v.Column, v.Value)
	}

	if !isDeleted {
		db = db.Where("deleted_at IS NULL")
	}

	err := db.First(&authorModel).Error
	if err == gorm.ErrRecordNotFound {
		return nil, r.rt.SetError(http.StatusNotFound, utils.ERR_AUTHOR_NOT_FOUND)
	}
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return &authorModel, nil
}

func (r *repository) UpdateAuthor(ctx context.Context, tx *gorm.DB, data *models.Author) error {
	logger := r.rt.Logger.With().
		Interface("data", data).
		Logger()

	if tx == nil {
		tx = r.rt.Db
	}

	tx = tx.Select("*").Updates(&data)

	if err := tx.Error; err != nil {
		logger.Error().Err(err).Msg("error query")
		return err
	}

	return nil
}
