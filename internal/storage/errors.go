package storage

import "errors"

var (
	ErrStorageIsNotAvailable = errors.New("storage is not available")
	ErrUnknownError          = errors.New("unknown error")
	ErrStorageNotExists      = errors.New("storage not exists")
	ErrEmptyData             = errors.New("empty data")
	ErrInternalError         = errors.New("internal error")
)
