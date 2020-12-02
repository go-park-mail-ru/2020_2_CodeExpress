package mwares

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/mwares/monitoring"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/session"
	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/error_response"
	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/responser"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/user"

	. "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/consts"
	"github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/csrf"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type MiddlewareManager struct {
	sessionUsecase session.SessionUsecase
	userUsecase    user.UserUsecase
	monitoring     *monitoring.Monitoring
}

func NewMiddlewareManager(sessionUsecase session.SessionUsecase, userUsecase user.UserUsecase,
	monitoring *monitoring.Monitoring) *MiddlewareManager {
	return &MiddlewareManager{
		sessionUsecase: sessionUsecase,
		userUsecase:    userUsecase,
		monitoring:     monitoring,
	}
}

func (mm *MiddlewareManager) PanicRecovering(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				logrus.Warn(err)
			}
		}()
		defer func() error {
			if err := recover(); err != nil {
				status := strconv.Itoa(ctx.Response().Status)
				path := ctx.Request().URL.Path
				method := ctx.Request().Method

				mm.monitoring.Hits.WithLabelValues(status, path, method).Inc()
				mm.monitoring.Duration.WithLabelValues(status, path, method).Observe(0)

				logrus.Warn(err)
				return ctx.JSON(http.StatusInternalServerError, nil)
			}
			return nil
		}()

		return next(ctx)
	}
}

func (mm *MiddlewareManager) AccessLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		logrus.Info(ctx.Request().RemoteAddr, " ", ctx.Request().Method, " ", ctx.Request().URL)

		start := time.Now()
		err := next(ctx)
		end := time.Now()

		workTime := end.Sub(start)

		status := strconv.Itoa(ctx.Response().Status)
		path := ctx.Request().URL.Path
		method := ctx.Request().Method

		mm.monitoring.Hits.WithLabelValues(status, path, method).Inc()
		mm.monitoring.Duration.WithLabelValues(status, path, method).Observe(workTime.Seconds())

		logrus.Info("Status: ", ctx.Response().Status, " Work time: ", workTime) //end.Sub(start)
		logrus.Println()

		return err
	}
}

func (mm *MiddlewareManager) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowHeaders:     ConstAllowedHeaders,
		AllowMethods:     ConstAllowedMethods,
		AllowOrigins:     ConstAllowedOrigins,
		ExposeHeaders:    ConstAllowedExpose,
	})
}

func (mm *MiddlewareManager) CheckCSRF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		csrfErr := errors.New("Bad csrf token received")
		sessionCookie, err := ctx.Cookie(ConstSessionName)

		if err != nil {
			return RespondWithError(NewErrorResponse(ErrNotAuthorized, err), ctx)
		}

		session, errResp := mm.sessionUsecase.GetByID(sessionCookie.Value)
		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		token := ctx.Request().Header.Get(ConstCSRFTokenName)

		if token == "" {
			return RespondWithError(NewErrorResponse(ErrBadRequest, csrfErr), ctx)
		}

		errResp = csrf.ValidateCSRFToken(session, token)

		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		return next(ctx)
	}
}

func (mm *MiddlewareManager) XSS() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection: "1; mode=block",
	})
}

func (mm *MiddlewareManager) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie(ConstSessionName)

		if err != nil {
			return RespondWithError(NewErrorResponse(ErrNotAuthorized, err), ctx)
		}

		session, errResp := mm.sessionUsecase.GetByID(cookie.Value)

		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		_, errResp = mm.userUsecase.GetById(session.UserID)

		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		ctx.Set(ConstAuthedUserParam, session.UserID)

		return next(ctx)
	}
}

func (mm *MiddlewareManager) CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userID := ctx.Get(ConstAuthedUserParam).(uint64)

		isAdmin, errResp := mm.userUsecase.CheckAdmin(userID)

		if errResp != nil {
			return RespondWithError(errResp, ctx)
		}

		if !isAdmin {
			return RespondWithError(NewErrorResponse(ErrNotAdmin, nil), ctx)
		}

		return next(ctx)
	}
}
