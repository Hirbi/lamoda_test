package product

import "errors"

var (
	ErrStorageIdIsNull         = errors.New("storage id is null")
	ErrNotEnoughAmount         = errors.New("not enough amount")
	ErrEmptyData               = errors.New("empty data")
	ErrNoCodes                 = errors.New("no codes get")
	ErrTooMuchProductToRelease = errors.New("too much products to release")
	ErrProductCodeNotFound     = errors.New("product code not found")
	ErrInternalError           = errors.New("internal error")
)
