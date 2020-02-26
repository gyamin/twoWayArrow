package db

import (
	"log"
	"strconv"
	"strings"
)

func BuildInsertSql(tableName string, data []map[string]interface{}) (sql string) {

	if len(data) < 1 {
		log.Fatal("Empty data is given. Can not build insert SQL.")
	}

	sql = "INSERT INTO " + tableName + " ("

	for key, _ := range data[0] {
		sql = sql + key + ","
	}
	sql = strings.TrimRight(sql, ",")
	sql = sql + ") VALUES "

	for _, elem := range data {
		sql = sql + "("
		for _, val := range elem {
			switch val.(type) {
			case string:
				sql = sql + val.(string) + ","
			case int:
				sql = sql + strconv.Itoa(val.(int)) + ","
			}
		}
		sql = sql + ")"
	}

	return sql
}
