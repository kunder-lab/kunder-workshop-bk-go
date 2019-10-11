package dbConnections

import (
	"database/sql"
	"fmt"
	"kunder-workshop-bk-go/settings"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	dbInfo := settings.V.DB
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbInfo.User, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.Table))
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
