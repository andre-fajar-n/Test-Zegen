package rest

import (
	"context"
	"zegen/gen/models"
	"zegen/gen/restapi/operations"
	"zegen/gen/restapi/operations/authentication"
	"zegen/gen/restapi/operations/author"
	"zegen/gen/restapi/operations/book"
	"zegen/gen/restapi/operations/health"
	"zegen/internal/services"
	"zegen/internal/utils"
	"zegen/runtime"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *runtime.Runtime, api *operations.ServerAPI, apiService services.IService) {
	//  health
	api.HealthHealthHandler = health.HealthHandlerFunc(func(hp health.HealthParams) middleware.Responder {
		return health.NewHealthOK().WithPayload(&models.Success{
			Message: "Server up",
		})
	})

	// authentication
	{
		api.AuthenticationRegisterHandler = authentication.RegisterHandlerFunc(func(rp authentication.RegisterParams) middleware.Responder {
			userID, err := apiService.Register(context.Background(), rp)
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
			token, expiredAt, err := apiService.Login(context.Background(), &lp)
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
			authorID, err := apiService.CreateAuthor(context.Background(), &cap)
			if err != nil {
				errRes := rt.GetError(err)
				return author.NewCreateAuthorDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return author.NewCreateAuthorCreated().WithPayload(&models.SuccessCreateAuthor{
				Success: models.Success{
					Message: utils.SUCCESS_CREATE,
				},
				SuccessCreateAuthorAllOf1: models.SuccessCreateAuthorAllOf1{
					AuthorID: *authorID,
				},
			})
		})

		api.AuthorUpdateAuthorHandler = author.UpdateAuthorHandlerFunc(func(uap author.UpdateAuthorParams, p *models.Principal) middleware.Responder {
			err := apiService.UpdateAuthor(context.Background(), &uap)
			if err != nil {
				errRes := rt.GetError(err)
				return author.NewUpdateAuthorDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return author.NewUpdateAuthorOK().WithPayload(&models.Success{
				Message: utils.SUCCESS_UPDATE,
			})
		})

		api.AuthorSoftDeleteAuthorHandler = author.SoftDeleteAuthorHandlerFunc(func(sdap author.SoftDeleteAuthorParams, p *models.Principal) middleware.Responder {
			err := apiService.SoftDeleteAuthor(context.Background(), &sdap)
			if err != nil {
				errRes := rt.GetError(err)
				return author.NewSoftDeleteAuthorDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return author.NewSoftDeleteAuthorOK().WithPayload(&models.Success{
				Message: utils.SUCCESS_DELETE,
			})
		})

		api.AuthorFindOneAuthorHandler = author.FindOneAuthorHandlerFunc(func(foap author.FindOneAuthorParams) middleware.Responder {
			data, err := apiService.FindOneAuthor(context.Background(), &foap)
			if err != nil {
				errRes := rt.GetError(err)
				return author.NewFindOneAuthorDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return author.NewFindOneAuthorOK().WithPayload(&models.SuccessFindOneAuthor{
				Success: models.Success{
					Message: utils.SUCCESS_CREATE,
				},
				SuccessFindOneAuthorAllOf1: *data,
			})
		})
	}

	// book
	{
		api.BookCreateBookHandler = book.CreateBookHandlerFunc(func(cap book.CreateBookParams, p *models.Principal) middleware.Responder {
			bookID, err := apiService.CreateBook(context.Background(), &cap)
			if err != nil {
				errRes := rt.GetError(err)
				return book.NewCreateBookDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return book.NewCreateBookCreated().WithPayload(&models.SuccessCreateBook{
				Success: models.Success{
					Message: utils.SUCCESS_CREATE,
				},
				SuccessCreateBookAllOf1: models.SuccessCreateBookAllOf1{
					BookID: *bookID,
				},
			})
		})

		api.BookUpdateBookHandler = book.UpdateBookHandlerFunc(func(uap book.UpdateBookParams, p *models.Principal) middleware.Responder {
			err := apiService.UpdateBook(context.Background(), &uap)
			if err != nil {
				errRes := rt.GetError(err)
				return book.NewUpdateBookDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return book.NewUpdateBookOK().WithPayload(&models.Success{
				Message: utils.SUCCESS_UPDATE,
			})
		})

		api.BookSoftDeleteBookHandler = book.SoftDeleteBookHandlerFunc(func(sdap book.SoftDeleteBookParams, p *models.Principal) middleware.Responder {
			err := apiService.SoftDeleteBook(context.Background(), &sdap)
			if err != nil {
				errRes := rt.GetError(err)
				return book.NewSoftDeleteBookDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return book.NewSoftDeleteBookOK().WithPayload(&models.Success{
				Message: utils.SUCCESS_DELETE,
			})
		})

		api.BookFindOneBookHandler = book.FindOneBookHandlerFunc(func(fobp book.FindOneBookParams) middleware.Responder {
			data, err := apiService.FindOneBook(context.Background(), &fobp)
			if err != nil {
				errRes := rt.GetError(err)
				return book.NewFindOneBookDefault(int(errRes.Code())).WithPayload(&models.Error{
					Code:    int64(errRes.Code()),
					Message: errRes.Error(),
				})
			}
			return book.NewFindOneBookOK().WithPayload(&models.SuccessFindOneBook{
				Success: models.Success{
					Message: utils.SUCCESS_CREATE,
				},
				SuccessFindOneBookAllOf1: *data,
			})
		})
	}
}
