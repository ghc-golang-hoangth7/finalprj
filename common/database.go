package common

import (
	"database/sql"
	"fmt"
)

const DatabaseDriver = "postgres"

func InitDb(config *Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.PostgresHost, config.PostgresPort,
		config.PostgresUser, config.PostgresPassword,
		config.DbName)
	return sql.Open(DatabaseDriver, dataSource)
}
