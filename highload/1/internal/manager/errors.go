package manager

import (
	"errors"
)

var (
	NotFound             = errors.New("not found")
	BadRequest           = errors.New("bad request")
	IncorrectPassOrLogin = errors.New("логин или пароль некорректны")
)
