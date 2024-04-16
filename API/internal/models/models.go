package models

import "github.com/jmoiron/sqlx"

type Env struct {
	DB *sqlx.DB
}

var E *Env
