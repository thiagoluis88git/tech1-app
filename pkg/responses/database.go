package responses

import (
	"errors"

	"gorm.io/driver/postgres"
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
	var code int
	message := err.Error()
	code = DATABASE_ERROR

	var mySqlError *postgres.ErrMessage

	if errors.As(err, &mySqlError) {
		switch mySqlError.Code {
		case "1062":
			code = DATABASE_CONFLICT_ERROR
		default:
			code = DATABASE_CONSTRAINT_ERROR
		}
	}

	return &LocalError{
		Code:    code,
		Message: message,
	}
}
