package mysql

import (
	"app/internal/models"
	"app/internal/product"
	"app/internal/storage"
	"database/sql"
	"fmt"
)

type ProductSchema struct {
	Id        sql.NullInt64
	Name      sql.NullString
	Code      sql.NullString
	Amount    sql.NullInt64
	Size      sql.NullString
	Reserved  sql.NullInt64
	StorageId sql.NullInt64
}

type ProductRepository struct {
	connPool *sql.DB

	storageRepo storage.Repository
	// mutex       sync.Mutex
}

func NewProductRepository(conn *sql.DB, storageRepo storage.Repository) *ProductRepository {
	return &ProductRepository{
		connPool:    conn,
		storageRepo: storageRepo,
	}
}

func (repo *ProductRepository) ReserveProducts(productCodes map[string]int) error {
	transaction, err := repo.connPool.Begin()
	if err != nil {
		transaction.Rollback()
		return product.ErrInternalError
	}
	for code, amountToReserve := range productCodes {

		reserveQuery := fmt.Sprintf(
			"UPDATE products SET reserved = reserved + %v WHERE code = '%v'",

			amountToReserve,
			code,
		)
		_, err = transaction.Exec(reserveQuery)
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}
		// first reserved amount and is storage available
		checkQuery := fmt.Sprintf(`SELECT amount, reserved, storage_id FROM products where code='%v'`, code)
		rows, err := transaction.Query(checkQuery)
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}
		if !rows.Next() {
			transaction.Rollback()
			return product.ErrProductCodeNotFound
		}
		var amount, reserved, storageId int64
		err = rows.Scan(&amount, &reserved, &storageId)
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}

		isStorageAvailable, err := repo.storageRepo.IsStorageAvailable(fmt.Sprint(storageId))
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}
		if !isStorageAvailable {
			transaction.Rollback()
			return storage.ErrStorageIsNotAvailable
		}

		if amount-reserved < 0 {
			transaction.Rollback()
			return product.ErrNotEnoughAmount
		}

	}
	transaction.Commit()
	return nil
}

func (repo *ProductRepository) ReleaseProducts(productCodes map[string]int) error {
	transaction, err := repo.connPool.Begin()
	if err != nil {
		transaction.Rollback()
		return product.ErrInternalError
	}
	for code, amountToRelease := range productCodes {
		// release products
		releaseQuery := fmt.Sprintf(
			"UPDATE products SET reserved = reserved - %v WHERE code = '%v'",

			amountToRelease,
			code,
		)
		_, err = transaction.Exec(releaseQuery)
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}
		// check reserved amount and is storage available
		checkQuery := fmt.Sprintf(`SELECT reserved, storage_id FROM products where code='%v'`, code)
		rows, err := transaction.Query(checkQuery)
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}
		if !rows.Next() {
			transaction.Rollback()
			return product.ErrProductCodeNotFound
		}
		var reserved, storageId int64
		err = rows.Scan(&reserved, &storageId)
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}

		isStorageAvailable, err := repo.storageRepo.IsStorageAvailable(fmt.Sprint(storageId))
		if err != nil {
			transaction.Rollback()
			return product.ErrInternalError
		}
		if !isStorageAvailable {
			transaction.Rollback()
			return storage.ErrStorageIsNotAvailable
		}

		if reserved < 0 {
			transaction.Rollback()

			return product.ErrTooMuchProductToRelease
		}
	}

	transaction.Commit()
	return nil
}

func (repo *ProductRepository) AvailableProducts(storageId string) ([]models.Product, error) {

	isStorageAvailable, err := repo.storageRepo.IsStorageAvailable(storageId)
	if err != nil {
		return nil, product.ErrInternalError
	}
	if !isStorageAvailable {
		return nil, storage.ErrStorageIsNotAvailable
	}
	productsQuery := fmt.Sprintf(`
SELECT pr.name, pr.size, pr.code, pr.amount, pr.reserved
FROM storage as st
JOIN products as pr
	ON st.id = pr.storage_id
WHERE storage_id = %v`,
		storageId,
	)

	rows, err := repo.connPool.Query(productsQuery)
	if err != nil {
		return nil, product.ErrInternalError
	}
	availableProducts := make([]models.Product, 0)
	for rows.Next() {
		var result ProductSchema
		err = rows.Scan(&result.Name, &result.Size, &result.Code, &result.Amount, &result.Reserved)
		if err != nil {
			return nil, product.ErrInternalError
		}
		if result.Name.Valid && result.Code.Valid && result.Amount.Valid && result.Reserved.Valid {
			availableProducts = append(availableProducts,
				models.Product{
					Name:   result.Name.String,
					Size:   result.Size.String,
					Code:   result.Code.String,
					Amount: result.Amount.Int64 - result.Reserved.Int64,
				},
			)
		}

	}
	return availableProducts, nil
}
