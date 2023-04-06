package goerror

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

const (
	ErrApplication    ErrComponent = "application"
	ErrDomain         ErrComponent = "domain  "
	ErrInfrastructure ErrComponent = "infrastructure"
	ErrInterfaces     ErrComponent = "interfaces"
)

const (
	BadRequest    ResponseErrType = "BadRequest"
	Forbidden     ResponseErrType = "Forbidden"
	NotFound      ResponseErrType = "NotFound"
	AlreadyExists ResponseErrType = "AlreadyExists"
	Unknown       ResponseErrType = "Unknown"
)

type GoError struct {
	error
	Code         string
	Message      string
	Data         map[string]interface{}
	Causes       []error
	Component    ErrComponent
	ResponseType ResponseErrType
}
