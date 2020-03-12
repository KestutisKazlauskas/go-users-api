package users_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"os"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_database = "mysql_users_database"
)

var (
	Clinet *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	database = os.Getenv(mysql_users_database)
)

func init() {
	var err error

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		database,
	)

	Clinet, err = sql.Open("mysql",  dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Clinet.Ping(); err != nil {
		panic(err)
	}

	log.Println("Succesfull connected to database")
}