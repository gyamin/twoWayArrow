package db

import (
	"log"
	"strings"
)

func BuildInsertSql(tableName string, data []map[string]interface{}) (sql string) {

	sql = "INSERT INTO " + tableName + " ("

	if len(data) < 1 {
		log.Fatal("Empty data is given. Can not build insert SQL.")
	}

	for key, _ := range data[0] {
		sql = sql + key + ","
	}
	sql = strings.TrimRight(sql, ",")
	sql = sql + " )"

	return sql
}
