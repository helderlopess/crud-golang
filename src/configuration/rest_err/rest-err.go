package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"` //foi percebido que o int64 aprsentou erro e foi avisado na legenda
	Causes  []Causes `json:"causes"`
}

//objeto coletor de erros
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

//funcao implemetnada a partir de resterr para satisfazer a interface de err
func (r *RestErr) Error() string { return r.Message }

//construtor do objeto
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

//*RestErr significa ponteiro
//metodo para cada tipo de erro
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

//caso seja passado algum erro invalido, ie, passado valor da variavel invalido

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

// criar para codigo 500 400 e 404 e 403 caso envie um token incorretos
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusInternalServerError,
	}
}
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
