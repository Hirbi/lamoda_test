package server

import (
	"app/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func initConnPool() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Name,
	)
	connPool, err := sql.Open(config.Config.Database.Driver, psqlInfo)
	if err != nil {
		return nil, err
	}
	return connPool, err
}
