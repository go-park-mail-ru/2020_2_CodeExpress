package consts

import (
	"errors"
	"net/http"
)

const (
	ErrInternal = iota
	ErrBadRequest
	ErrEmailAlreadyExist
	ErrNameAlreadyExist
	ErrIncorrectLoginOrPassword
	ErrNotAuthorized
	ErrNoEmail
	ErrNoUsername
	ErrNoAvatar
	ErrWrongOldPassword
	ErrNewPasswordIsOld
	ErrArtistNotExist
	ErrTrackNotExist
	ErrAlbumNotExist
	ErrTitleAlreadyExist
	ErrNoFavoritesTracks
	ErrEmptySearchQuery
)

var Errors = map[int]error{
	ErrInternal:                 errors.New("Internal server error"),
	ErrBadRequest:               errors.New("Bad request received"),
	ErrEmailAlreadyExist:        errors.New("Email already exists"),
	ErrNameAlreadyExist:         errors.New("Name already exists"),
	ErrIncorrectLoginOrPassword: errors.New("Incorrect login or password"),
	ErrNotAuthorized:            errors.New("Not authorized"),
	ErrNoEmail:                  errors.New("No email field"),
	ErrNoUsername:               errors.New("No username field"),
	ErrNoAvatar:                 errors.New("No avatar field"),
	ErrWrongOldPassword:         errors.New("Wrong old password"),
	ErrNewPasswordIsOld:         errors.New("New password matches old"),
	ErrArtistNotExist:           errors.New("Artist not found"),
	ErrTrackNotExist:            errors.New("Track not found"),
	ErrAlbumNotExist:            errors.New("Album not found"),
	ErrTitleAlreadyExist:        errors.New("Title already exists"),
	ErrNoFavoritesTracks:        errors.New("User has no favorite tracks"),
	ErrEmptySearchQuery:         errors.New("Empty search query"),
}

var StatusCodes = map[int]int{
	ErrInternal:                 http.StatusInternalServerError,
	ErrBadRequest:               http.StatusBadRequest,
	ErrEmailAlreadyExist:        http.StatusForbidden,
	ErrNameAlreadyExist:         http.StatusForbidden,
	ErrIncorrectLoginOrPassword: http.StatusNotFound,
	ErrNotAuthorized:            http.StatusNotFound,
	ErrNoEmail:                  http.StatusBadRequest,
	ErrNoUsername:               http.StatusBadRequest,
	ErrNoAvatar:                 http.StatusBadRequest,
	ErrWrongOldPassword:         http.StatusBadRequest,
	ErrNewPasswordIsOld:         http.StatusBadRequest,
	ErrArtistNotExist:           http.StatusNotFound,
	ErrTrackNotExist:            http.StatusNotFound,
	ErrAlbumNotExist:            http.StatusNotFound,
	ErrTitleAlreadyExist:        http.StatusForbidden,
	ErrNoFavoritesTracks:        http.StatusNotFound,
	ErrEmptySearchQuery:         http.StatusBadRequest,
}
