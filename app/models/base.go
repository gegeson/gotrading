package models

import (
	"database/sql"
	"fmt"
	"log"
	"main/config"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableNamesignalEvent = "signal_events"
)
var Dbconnection *sql.DB

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	Dbconnection, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatal(err)
	}
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(
			time DATETIME PRIMARY KEY NOT NULL,
			product_code STRING,
			side STRING,
			price FLOAT,
			size FLOAT
		)`, tableNamesignalEvent)
	Dbconnection.Exec(cmd)

	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s(
			time DATETIME PRIMARY KEY NOT NULL,
			open FLOAT,
			close FLOAT,
			high FLOAT,
			low open FLOAT,
			valume FLOAT
		)`, tableName)
		Dbconnection.Exec(c)
	}
}