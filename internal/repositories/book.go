package repositories

import (
	"context"
	"net/http"
	"zegen/gen/models"
	"zegen/internal/utils"

	"gorm.io/gorm"
)

func (r *repository) CreateBook(ctx context.Context, tx *gorm.DB, data *models.Book) (*models.Book, error) {
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

func (r *repository) FindOneBookByFilter(ctx context.Context, filter []ColumnValue, isDeleted bool) (*models.Book, error) {
	logger := r.rt.Logger.With().
		Interface("filter", filter).
		Bool("isDeleted", isDeleted).
		Logger()

	bookModel := models.Book{}
	db := r.rt.Db.Model(&bookModel)

	for _, v := range filter {
		db = db.Where(v.Column, v.Value)
	}

	if !isDeleted {
		db = db.Where("deleted_at IS NULL")
	}

	err := db.Preload("Authors").First(&bookModel).Error
	if err == gorm.ErrRecordNotFound {
		return nil, r.rt.SetError(http.StatusNotFound, utils.ERR_BOOK_NOT_FOUND)
	}
	if err != nil {
		logger.Error().Err(err).Msg("error query")
		return nil, err
	}

	return &bookModel, nil
}

func (r *repository) UpdateBook(ctx context.Context, tx *gorm.DB, data *models.Book) error {
	logger := r.rt.Logger.With().
		Interface("data", data).
		Logger()

	if tx == nil {
		tx = r.rt.Db
	}

	tx = tx.Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&data)
	if err := tx.Error; err != nil {
		logger.Error().Err(err).Msg("error query")
		return err
	}

	return nil
}

func (r *repository) SoftDeleteBook(ctx context.Context, tx *gorm.DB, data *models.Book) error {
	logger := r.rt.Logger.With().
		Interface("data", data).
		Logger()

	if tx == nil {
		tx = r.rt.Db
	}

	tx = tx.Select("deleted_at", "Authors").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&data)
	if err := tx.Error; err != nil {
		logger.Error().Err(err).Msg("error query")
		return err
	}

	return nil
}
