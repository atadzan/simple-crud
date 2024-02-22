package database

type Config struct {
	Host    string
	Port    string
	DBName  string
	SSLMode string
}

func NewDatabasePoolConn()
