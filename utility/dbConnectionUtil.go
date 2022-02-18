package utilities

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var (
		host      = os.Getenv("host")
		port, err = strconv.Atoi(os.Getenv("port"))
		user      = os.Getenv("user")
		password  = os.Getenv("password")
		dbname    = os.Getenv("dbname")
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}

}
