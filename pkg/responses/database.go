package responses

import (
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
	var postgresError *postgres.Dialector

	return &LocalError{
		Message: postgresError.Translate(err).Error(),
	}
}
