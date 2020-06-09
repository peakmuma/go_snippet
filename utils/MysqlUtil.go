package utils

import (
	"database/sql"
	"fmt"
	// the sql-driver
    _ "github.com/go-sql-driver/mysql"
)

func connectDB() {
	db, err := sql.Open("mysql", "root:root@127.0.0.1:3306/lock_test")
	query(db, "select id, name from t1 where id = 1")
}

func query(db *sql.DB, sql string) {
	row := db.QueryRow(sql)
	fmt.Println("the res is: ", row)
}
