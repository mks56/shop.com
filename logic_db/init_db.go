package logic_db

import (
	"database/sql"
	"fmt"

	"shop.com/err"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "1234"
	dbname   = "shop_db"
)

func InitDb() *sql.DB {
	configData := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, result := sql.Open("postgres", configData)
	err.CheckError(result)

	result = db.Ping()
	err.CheckError(result)

	fmt.Println("Db connected")
	return db
}
