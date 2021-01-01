package apperror

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func NewGorm(err error, msg string) appError {
	switch err.Error() {
	case gorm.ErrRecordNotFound.Error():
		return appError{
			code:    CodeNotFound,
			err:     err,
			message: msg + ": " + err.Error(),
		}
	default:
		return appError{
			code:    CodeError,
			err:     err,
			message: msg + ": " + err.Error(),
		}
	}
}

func NewGormFind(err error, mname string) Error {
	msg := fmt.Sprintf("Find record `%s` error", mname)
	return NewGorm(err, msg)
}

func NewGormDelete(err error, mname string) Error {
	msg := fmt.Sprintf("Delete record `%s` error", mname)
	return NewGorm(err, msg)
}

func NewGormCreate(err error, mname string) Error {
	msg := fmt.Sprintf("Create record `%s` error", mname)
	return NewGorm(err, msg)
}

func NewGormSave(err error, mname string) Error {
	msg := fmt.Sprintf("Update record `%s` error", mname)
	return NewGorm(err, msg)
}
