package goerror

type ResponseErrType string

type ErrComponent string

type Error interface {
	error
	Code() string
	Message() string
	Cause() error
	Causes() []error
	Data() map[string]interface{}
	String() string
	ResponseErrType() ResponseErrType
	SetResponseType(r ResponseErrType) Error
	Component() ErrComponent
	SetComponent(c ErrComponent) Error
	Retryable() bool
	SetRetryable() Error
}
