package manager

import (
	"errors"
)

var (
	NotFound             = errors.New("not found")
	IncorrectPassOrLogin = errors.New("логин или пароль некорректны")
)
