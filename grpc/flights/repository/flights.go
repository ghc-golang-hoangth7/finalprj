package repository

import "database/sql"

type FlightDb struct {
	db *sql.DB
}
