package mysql

import (
	"app/internal/storage"
	"database/sql"
	"fmt"
)

type StorageRepository struct {
	connPool *sql.DB
}

func NewProductRepository(conn *sql.DB) *StorageRepository {
	return &StorageRepository{
		connPool: conn,
	}
}

type StorageSchema struct {
	id          sql.NullInt64
	name        sql.NullString
	isAvailable sql.NullBool
}

func (repo *StorageRepository) IsStorageAvailable(storageId string) (bool, error) {
	query := fmt.Sprintf("SELECT id, is_available FROM storage WHERE id = %v", storageId)
	queryRows, err := repo.connPool.Query(query)
	if err != nil {
		return false, storage.ErrInternalError
	}
	hasStorageInfo := queryRows.Next()
	if !hasStorageInfo {
		return false, storage.ErrStorageNotExists
	}
	var result StorageSchema
	err = queryRows.Scan(&result.id, &result.isAvailable)
	if result.isAvailable.Valid && err == nil {
		return result.isAvailable.Bool, nil
	}
	if err != nil {
		return false, storage.ErrInternalError
	}
	return false, nil
}
