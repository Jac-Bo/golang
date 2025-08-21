package db

import (
	"os"

	_ "github.com/lib/pq"
)

var (
	UNAMEDB string = os.Getenv("UNAMEDB")
	PASSDB  string = os.Getenv("PASSDB")
	HOSTDB  string = os.Getenv("HOSTDB")
	DBNAME  string = os.Getenv("DBNAME")
)
