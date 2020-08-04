package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var database = initDatabase()

func initDatabase() *sql.DB {
	DB_USERNAME := "root"
	DB_PASSWORD := "password"
	DB_SERVER := "localhost"
	DB_PORT := 3306
	DATABASE_NAME := "mysql"
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DB_USERNAME, DB_PASSWORD, DB_SERVER, DB_PORT, DATABASE_NAME)
	database, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return nil
	}
	database.SetConnMaxLifetime(100 * time.Second) //timeout
	database.SetMaxOpenConns(5)                    //max_conn
	return database
}

func GetTableStatus(tableName string, time int64) int64 {
	//SELECT
	//	table_schema,
	//  table_name,
	//  update_time
	//FROM
	//  information_schema.tables
	//WHERE
	//  table_name = 'subscriptions' and
	//  update_time > (NOW() - INTERVAL 10 MINUTE);
	var count int64
	row := database.QueryRow(`
		SELECT 
			count(*) as count 
		FROM 
			information_schema.tables 
		WHERE 
			table_name = ? and 
			update_time > (NOW() - INTERVAL ? MINUTE)`, tableName, time)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return 0
	}
	fmt.Println(count)
	return count
}
