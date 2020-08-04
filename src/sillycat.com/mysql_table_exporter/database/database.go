package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sillycat.com/mysql_table_exporter/config"
	"time"
)

var database = initDatabase()

func initDatabase() *sql.DB {
	dbUserName := config.GetEnv("DB_USERNAME", "root")
	dbPassword := config.GetEnv("DB_PASSWORD", "password")
	dbServer := config.GetEnv("DB_SERVER", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	databaseName := config.GetEnv("DATABASE_NAME", "mysql")
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUserName, dbPassword, dbServer, dbPort, databaseName)
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
