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

	var keys []string
	for key := range data[0] {
		keys = append(keys, key)
		sql = sql + key + ","
	}
	sql = strings.TrimRight(sql, ",")
	sql = sql + ") VALUES "

	for _, elem := range data {
		sql = sql + "("

		for _, key := range keys {
			switch elem[key].(type) {
			case string:
				sql = sql + "'" + elem[key].(string) + "',"
			case int:
				sql = sql + strconv.Itoa(elem[key].(int)) + ","
			}
		}
		sql = strings.TrimRight(sql, ",")
		sql = sql + "),"
	}
	sql = strings.TrimRight(sql, ",")

	return sql
}
