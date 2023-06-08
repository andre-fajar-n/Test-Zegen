package rest

import (
	"context"
	"zegen/gen/models"
	"zegen/gen/restapi/operations"
	"zegen/gen/restapi/operations/authentication"
	"zegen/gen/restapi/operations/author"
	"zegen/gen/restapi/operations/health"
	"zegen/internal/handlers"
	"zegen/runtime"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *runtime.Runtime, api *operations.ServerAPI, apiHandler handlers.IHandler) {
	//  health
	api.HealthHealthHandler = health.HealthHandlerFunc(func(hp health.HealthParams) middleware.Responder {
		return health.NewHealthOK().WithPayload(&models.Success{
			Message: "Server up",
		})
	})

	// authentication
	{
		api.AuthenticationRegisterHandler = authentication.RegisterHandlerFunc(func(rp authentication.RegisterParams) middleware.Responder {
			userID, err := apiHandler.Register(context.Background(), rp)
			if err != nil {
				errRes := rt.GetError(err)
				return authentication.NewRegisterDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return authentication.NewRegisterCreated().WithPayload(&models.SuccessRegister{
				Success: models.Success{
					Message: "success register",
				},
				SuccessRegisterAllOf1: models.SuccessRegisterAllOf1{
					UserID: *userID,
				},
			})
		})

		api.AuthenticationLoginHandler = authentication.LoginHandlerFunc(func(lp authentication.LoginParams) middleware.Responder {
			token, expiredAt, err := apiHandler.Login(context.Background(), &lp)
			if err != nil {
				errRes := rt.GetError(err)
				return authentication.NewLoginDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}

			return authentication.NewLoginOK().WithPayload(&models.SuccessLogin{
				Success: models.Success{
					Message: "success login",
				},
				SuccessLoginAllOf1: models.SuccessLoginAllOf1{
					ExpiredAt: *expiredAt,
				},
			}).WithToken(*token)
		})
	}

	// author
	{
		api.AuthorCreateAuthorHandler = author.CreateAuthorHandlerFunc(func(cap author.CreateAuthorParams, p *models.Principal) middleware.Responder {
			authorID, err := apiHandler.CreateAuthor(context.Background(), &cap)
			if err != nil {
				errRes := rt.GetError(err)
				return author.NewCreateAuthorDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return author.NewCreateAuthorCreated().WithPayload(&models.SuccessCreateAuthor{
				Success: models.Success{
					Message: "success create",
				},
				SuccessCreateAuthorAllOf1: models.SuccessCreateAuthorAllOf1{
					AuthorID: *authorID,
				},
			})
		})
	}
}
