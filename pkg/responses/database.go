package responses

import (
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	DATABASE_ERROR            = 1
	DATABASE_CONSTRAINT_ERROR = 2
	DATABASE_CONFLICT_ERROR   = 3
	MALFORMED_DATA_ERROR      = 4
	LOGIC_ERROR               = 5
)

type LocalError struct {
	Code    int
	Message string
}

func (er LocalError) Error() string {
	return er.Message
}

func GetDatabaseError(err error) *LocalError {
	var localError *pgconn.PgError
	code := DATABASE_ERROR
	message := "Unknown error"

	if errors.As(err, &localError) {
		iCode, err := strconv.Atoi(localError.Code)
		if err != nil {
			iCode = DATABASE_ERROR
		}

		code = iCode
		message = localError.Message
	}

	return &LocalError{
		Message: message,
		Code:    code,
	}
}
