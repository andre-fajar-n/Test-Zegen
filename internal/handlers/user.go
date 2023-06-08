package handlers

import (
	"context"
	"net/http"
	"time"
	"zegen/gen/models"
	"zegen/gen/restapi/operations/authentication"
	utilsHash "zegen/internal/utils/hash"
	utilsJwt "zegen/internal/utils/jwt"

	"github.com/go-openapi/strfmt"
)

func (h *handler) Register(ctx context.Context, req authentication.RegisterParams) (*uint64, error) {
	logger := h.rt.Logger.With().
		Interface("req", req.Data).
		Logger()

	existUsername, err := h.repo.UsernameExist(ctx, *req.Data.Username)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.UsernameExist")
		return nil, err
	}
	if existUsername {
		return nil, h.rt.SetError(http.StatusBadRequest, "username already exist")
	}

	now := time.Now().UTC()
	nowStrfmt := strfmt.DateTime(now)
	user, err := h.repo.CreateUser(ctx, nil, &models.User{
		ModelTrackTime: models.ModelTrackTime{
			CreatedAt: &nowStrfmt,
		},
		UserData: models.UserData{
			Username: *req.Data.Username,
			Password: utilsHash.HashSha256(*req.Data.Password),
		},
	})
	if err != nil {
		logger.Error().Err(err).Msg("error repo.Create")
		return nil, err
	}

	return user.ID, nil
}

func (h *handler) Login(ctx context.Context, req *authentication.LoginParams) (token, expiredAt *string, err error) {
	logger := h.rt.Logger.With().
		Interface("req", req.Data).
		Logger()

	user, err := h.repo.FindUserBySingleColumn(ctx, "username", req.Data.Username, false)
	if err != nil {
		logger.Error().Err(err).Msg("error repo.FindUserBySingleColumn")
		return
	}

	if user.Password != utilsHash.HashSha256(*req.Data.Password) {
		return nil, nil, h.rt.SetError(http.StatusForbidden, "wrong password")
	}

	secret := h.rt.Cfg.JwtSecret
	maker, err := utilsJwt.NewJWTMaker(secret)
	if err != nil {
		logger.Error().Err(err).Msg("error NewJWTMaker")
		return
	}

	tempToken, payload, err := maker.CreateToken(*user.ID, user.Username, time.Duration(h.rt.Cfg.JwtExp)*time.Second)
	if err != nil {
		logger.Error().Err(err).Msg("error CreateToken")
		return
	}
	token = &tempToken
	tempExp := payload.ExpiredAt.Format(time.RFC3339)
	expiredAt = &tempExp

	return
}
