package db

import (
	"fmt"
)

func BuildBulkInsertSql(tableName string, data map[string]interface{}) (sql string) {

	for key, _ := range data {
		fmt.Print(key)
		fmt.Print(" ")
	}
	fmt.Println()

	sql = ""
	//sql = sql + "INSERT INTO " + tableName + " ( " + strings.Join(columns, ",") + ") VALUES ("
	//
	//for _ , line := range data {
	//	sql = sql + "VALUES ( " + strings.Join(line, ",")
	//}

	return sql
}
