package rest

import (
	"zegen/gen/models"
	"zegen/gen/restapi/operations"
	"zegen/internal/utils/jwt"
	"zegen/runtime"

	"github.com/go-openapi/strfmt"
)

func parseToken(rt *runtime.Runtime, token string) (*models.Principal, error) {
	secret := rt.Cfg.JwtSecret
	maker, err := jwt.NewJWTMaker(secret)
	if err != nil {
		return nil, err
	}

	payload, err := maker.VerifyToken(token)
	if err != nil {
		return nil, rt.SetError(401, "Unauthorized: invalid API key token: %v", err)
	}

	return &models.Principal{
		ExpiredAt: strfmt.DateTime(payload.ExpiredAt),
		UserID:    payload.UserID,
		Username:  payload.Username,
	}, nil
}

func Authorization(rt *runtime.Runtime, api *operations.ServerAPI) {
	api.AuthorizationAuth = func(token string) (*models.Principal, error) {
		return parseToken(rt, token)
	}
}
