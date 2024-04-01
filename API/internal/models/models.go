package models

import "database/sql"

type Env struct {
	DB *sql.DB
}

var E *Env
