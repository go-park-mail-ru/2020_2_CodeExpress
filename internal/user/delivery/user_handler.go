package delivery

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"

	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/consts"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/models"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/session"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/builder"
	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/error_response"
	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/responser"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/validator"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/user"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userUsecase    user.UserUsecase
	sessionUsecase session.SessionUsecase
}

func NewUserHandler(userUsecase user.UserUsecase, sessionUsecase session.SessionUsecase) *UserHandler {
	return &UserHandler{
		userUsecase:    userUsecase,
		sessionUsecase: sessionUsecase,
	}
}

func (uh *UserHandler) Configure(e *echo.Echo) {
	e.POST("/api/v1/user/register", uh.handlerRegisterUser())
	e.POST("/api/v1/user/login", uh.handlerLoginUser())
	e.GET("/api/v1/user/current", uh.handlerCurrentUserInfo())
	e.DELETE("/api/v1/user/logout", uh.handlerUserLogout())
	e.POST("/api/v1/user/change/profile", uh.handlerUpdateProfile())
	e.POST("/api/v1/user/change/password", uh.handlerUpdatePassword())
	e.POST("/api/v1/user/change/avatar", uh.handlerUpdateAvatar())
	e.Static("/avatars", "avatars")
}

func (uh *UserHandler) handlerRegisterUser() echo.HandlerFunc {
	type Request struct {
		Name             string `json:"username" validate:"required"`
		Email            string `json:"email" validate:"required,email"`
		Password         string `json:"password" validate:"required,gte=8"`
		RepeatedPassword string `json:"repeated_password" validate:"required,eqfield=Password"`
	}

	return func(ctx echo.Context) error {
		req := &Request{}
		if err := validator.NewRequestValidator(ctx).Validate(req); err != nil {
			if err.Error != nil { //TODO: Обрабатывать ошибки валидатора
				logrus.Info(err.Error)
				validator.GetValidationError(err)
			}
			return ctx.JSON(err.StatusCode, err.UserError)
		}

		user := &models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}

		if err := uh.userUsecase.CreateUser(user); err != nil {
			return RespondWithError(err, ctx)
		}

		session := models.NewSession(user.ID)
		if err := uh.sessionUsecase.CreateSession(session); err != nil {
			return RespondWithError(err, ctx)
		}

		cookie := builder.BuildCookie(session)
		ctx.SetCookie(cookie)

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) handlerLoginUser() echo.HandlerFunc {
	type Request struct {
		Name     string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required,gte=8"`
	}

	return func(ctx echo.Context) error {
		req := &Request{}
		if err := validator.NewRequestValidator(ctx).Validate(req); err != nil {
			if err.Error != nil { //TODO: Обрабатывать ошибки валидатора
				logrus.Info(err.Error)
				validator.GetValidationError(err)
			}
			return ctx.JSON(err.StatusCode, err.UserError)
		}

		user, err := uh.userUsecase.LoginUser(req.Name, req.Password)
		if err != nil {
			return RespondWithError(err, ctx)
		}

		session := models.NewSession(user.ID)
		if err := uh.sessionUsecase.CreateSession(session); err != nil {
			return RespondWithError(err, ctx)
		}

		cookie := builder.BuildCookie(session)
		ctx.SetCookie(cookie)

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) handlerCurrentUserInfo() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie(ConstSessionName)
		if err != nil {
			errResp := NewErrorResponse(ErrNotAuthorized, err)
			return RespondWithError(errResp, ctx)
		}

		session, errResp := uh.sessionUsecase.GetByID(cookie.Value)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		user, errResp := uh.userUsecase.GetByID(session.UserID)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) handlerUserLogout() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie(ConstSessionName)
		if err != nil {
			errResp := NewErrorResponse(ErrNotAuthorized, err)
			return RespondWithError(errResp, ctx)
		}

		session, errResp := uh.sessionUsecase.GetByID(cookie.Value)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		errResp = uh.sessionUsecase.DeleteSession(session)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		return ctx.JSON(http.StatusOK, OKRespose)
	}
}

func (uh *UserHandler) handlerUpdateProfile() echo.HandlerFunc {
	type Request struct {
		Name  string `json:"username" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	return func(ctx echo.Context) error {
		req := &Request{}
		if err := validator.NewRequestValidator(ctx).Validate(req); err != nil {
			if err.Error != nil { //TODO: Обрабатывать ошибки валидатора
				logrus.Info(err.Error)
				validator.GetValidationError(err)
			}
			return ctx.JSON(err.StatusCode, err.UserError)
		}

		cookie, err := ctx.Cookie(ConstSessionName)
		if err != nil {
			errResp := NewErrorResponse(ErrNotAuthorized, err)
			return RespondWithError(errResp, ctx)
		}

		session, errResp := uh.sessionUsecase.GetByID(cookie.Value)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		user, errResp := uh.userUsecase.GetByID(session.UserID)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		user.Name = req.Name
		user.Email = req.Email
		if errResp = uh.userUsecase.UpdateProfile(user); errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) handlerUpdatePassword() echo.HandlerFunc {
	type Request struct {
		Password         string `json:"password" validate:"required,gte=8"`
		RepeatedPassword string `json:"repeated_password" validate:"required,eqfield=Password"`
	}

	return func(ctx echo.Context) error {
		req := &Request{}
		if err := validator.NewRequestValidator(ctx).Validate(req); err != nil {
			if err.Error != nil { //TODO: Обрабатывать ошибки валидатора
				logrus.Info(err.Error)
				validator.GetValidationError(err)
			}
			return ctx.JSON(err.StatusCode, err.UserError)
		}

		cookie, err := ctx.Cookie(ConstSessionName)
		if err != nil {
			errResp := NewErrorResponse(ErrNotAuthorized, err)
			return RespondWithError(errResp, ctx)
		}

		session, errResp := uh.sessionUsecase.GetByID(cookie.Value)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		user, errResp := uh.userUsecase.GetByID(session.UserID)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		user.Password = req.Password
		if errResp = uh.userUsecase.UpdatePassword(user); errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		return ctx.JSON(http.StatusOK, user)
	}
}

func (uh *UserHandler) handlerUpdateAvatar() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie(ConstSessionName)
		if err != nil {
			errResp := NewErrorResponse(ErrNotAuthorized, err)
			return RespondWithError(errResp, ctx)
		}

		session, errResp := uh.sessionUsecase.GetByID(cookie.Value)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		user, errResp := uh.userUsecase.GetByID(session.UserID)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		formFile, err := ctx.FormFile("avatar")
		if err != nil {
			errResp := NewErrorResponse(ErrNoAvatar, err)
			return RespondWithError(errResp, ctx)
		}

		if formFile.Size > int64(10<<20) { //TODO: magic number
			errResp := NewErrorResponse(ErrNoAvatar, err) //TODO: другая ошибка будет
			return RespondWithError(errResp, ctx)
		}

		source, err := formFile.Open()
		if err != nil {
			errResp := NewErrorResponse(ErrNoAvatar, err)
			return RespondWithError(errResp, ctx)
		}
		defer func() {
			_ = source.Close()
		}()

		fileExtension := "png" //TODO: убрать в другое место; захордкоженное расширение
		randBytes := md5.Sum([]byte(fileExtension + user.Name))
		randString := hex.EncodeToString(randBytes[:])
		fileName := randString + "." + fileExtension
		pathToNewFile := "./avatars/" + fileName
		destination, err := os.OpenFile(pathToNewFile, os.O_WRONLY|os.O_CREATE, os.FileMode(int(0777)))
		if err != nil {
			return RespondWithError(NewErrorResponse(ErrInternal, err), ctx)
		}
		defer func() {
			_ = destination.Close()
		}()

		if _, err := io.Copy(destination, source); err != nil {
			return RespondWithError(NewErrorResponse(ErrInternal, err), ctx)
		}

		user.Avatar = pathToNewFile
		if errResp = uh.userUsecase.UpdateProfile(user); errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		return ctx.JSON(http.StatusOK, user)
	}
}