package user_token

import (
	"myweb/infrastructure/goerror"
)

func TokenExpiredError() error {
	return &goerror.GoError{
		Code:      "1003",
		Message:   "token 已过期",
		Component: goerror.ErrInfrastructure,
	}
}

func TokenCreateError() error {
	return &goerror.GoError{
		Code:      "1001",
		Message:   "创建 token 失败",
		Component: goerror.ErrInfrastructure,
	}
}

func TokenHandleError(message string) error {
	return &goerror.GoError{
		Code:      "1002",
		Message:   message,
		Component: goerror.ErrInfrastructure,
	}
}
