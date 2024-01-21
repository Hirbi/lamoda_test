package storage

type Repository interface {
	IsStorageAvailable(storageId string) (bool, error)
}
