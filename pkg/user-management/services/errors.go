package services

import (
	"database/sql"
	"errors"

	"github.com/goshathebusiness/dombu/pkg/logging"
)

var (
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal error")
)

func handleError(logger *logging.Logger, err error) error {
	logger.Error(err)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	return ErrInternal
}
