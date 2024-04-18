package responses

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type BusinessResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"msgError"`
}

func (br BusinessResponse) Error() string {
	return br.Message
}

/*
*

	Regras de validação de erros

	1) Verificação de erros de rede (BFF chamando serviços externos)
		Nesse caso o status code já vem do próprio serviço externo. Então só
		precisamos adaptar a mensagem de erro e usar o próprio status code para
		retornar pro usuário do BFF

	2) Verificação de erros de banco de dados (BFF chamando o banco do Microserviço)
		Nesse caso a mensagem de erro já vem do banco de dados (Ex: Duplicate keys). Então
		precisamos adaptar o status code e usar a própria mensagem de erro para
		retornar para o usuário do BFF

	3) Verificação de erros de outro Use Case (Use Case com dependência de outro Use Case)
		Nesse caso o status e a mensagem já estão prontas, é só repassar

	4) Default
		Caso não seja NetworkError ou LocalError, retornará um statuso code 500
		para o usuário

*
*/
func GetResponseError(err error, service string, logicError string) error {
	var networkError *NetworkError
	var databaseError *LocalError
	var businessError *BusinessResponse

	statusCode := http.StatusInternalServerError
	message := "Unexpected internal error"

	if errors.As(err, &networkError) {
		statusCode = networkError.Code
		message = getBusinessMessageError(statusCode, service)
	} else if errors.As(err, &databaseError) {
		message = databaseError.Message
		statusCode = getBusinessStatusCodeError(databaseError.Code)
	} else if errors.As(err, &businessError) {
		statusCode = businessError.StatusCode
		message = businessError.Message
	}

	businessResponse := &BusinessResponse{
		StatusCode: statusCode,
		Message:    message,
	}

	if strings.TrimSpace(logicError) != "" {
		businessResponse.Message = fmt.Sprintf("%v: %v", businessResponse.Message, logicError)
	}

	return businessResponse
}

func getBusinessMessageError(statusCode int, service string) string {
	var message string

	switch statusCode {
	case http.StatusBadRequest:
		message = fmt.Sprintf("Bad request trying to execute %v", service)
	case http.StatusUnauthorized:
		message = fmt.Sprintf("Unauthorized error trying to execute %v", service)
	case http.StatusForbidden:
		message = fmt.Sprintf("Forbiden error trying to execute %v", service)
	case http.StatusNotFound:
		message = fmt.Sprintf("Not found trying to execute %v", service)
	case http.StatusConflict:
		message = fmt.Sprintf("Conflit with some data using the service %v", service)
	case http.StatusUnprocessableEntity:
		message = fmt.Sprintf("Logic error found in service %v", service)
	case http.StatusGone:
		message = fmt.Sprintf("Gone error found in service %v", service)
	case http.StatusPreconditionRequired:
		message = fmt.Sprintf("Precondition Required found in service %v", service)
	case http.StatusRequestedRangeNotSatisfiable:
		message = fmt.Sprintf("Range Not Satisfiable found in service %v", service)
	case http.StatusLengthRequired:
		message = fmt.Sprintf("Length Required found in service %v", service)
	case http.StatusRequestEntityTooLarge:
		message = fmt.Sprintf("Some data is too large to be accepted found in service %v", service)
	case http.StatusLocked:
		message = fmt.Sprintf("The resource is locked or not finished found in service %v", service)
	default:
		message = fmt.Sprintf("Unexpected internal error trying to execute service %v", service)
	}

	return message
}

func getBusinessStatusCodeError(databaseStatusCode int) int {
	var code int

	switch databaseStatusCode {
	case DATABASE_ERROR:
		code = http.StatusInternalServerError
	case DATABASE_CONSTRAINT_ERROR:
		code = http.StatusUnprocessableEntity
	case DATABASE_CONFLICT_ERROR:
		code = http.StatusConflict
	case MALFORMED_DATA_ERROR:
		code = http.StatusBadRequest
	}

	return code
}
