package model

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

var ErrUpdateFail error = errors.New("Update corpus failed")

var ErrLoginFail error = errors.New("Login failed")

var ErrLoginNeeded error = errors.New("LoginNeeded")

var ErrCitedShouldNotEmpty = errors.New("cited should not empty")

var ErrNoCorpus error = errors.New("NoSuchCorpus")

var ErrNoUserTheme error = errors.New("NoSuchUserThemes")

var ErrActionNotAllow error = errors.New("ActionNotAllow")

var ErrAccessNotAllowed error = errors.New("AccessNotAllowed")

var ErrBadCorpusID error = errors.New("BadCorpusID")
