package db

import (
	"database/sql"
	"log"
	"strings"
)

type DataRegister struct {
	tx        *sql.Tx
	tableName string
	keys      []string
}

func NewDataRegister(tableName string, tx *sql.Tx) (databaseRegister DataRegister) {
	databaseRegister.tx = tx
	databaseRegister.tableName = tableName
	return databaseRegister
}

func (dr *DataRegister) DeleteAll() (affectedRows int64) {
	sql := "DELETE FROM " + dr.tableName
	result, err := dr.tx.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, _ = result.RowsAffected()
	return affectedRows
}

func (dr *DataRegister) CreateData(data []map[string]interface{}) (affectedRows int64) {
	dr.keys = []string{}
	for key := range data[0] {
		dr.keys = append(dr.keys, key)
	}

	// insert文作成
	insertSql := dr.buildInsertSql()

	for _, elem := range data {
		stmt, err := dr.tx.Prepare(insertSql)
		if err != nil {
			log.Fatal(err)
		}

		// データ抽出
		values := dr.extractValues(elem)
		// Insert実行
		_, err = stmt.Exec(values...)
		if err != nil {
			log.Fatal(err)
		}
		affectedRows++
	}
	return affectedRows
}

// mapから値を取り出し、配列で返す
// {"id":10, "name":"太郎"} → [10, "太郎"] を返す
func (dr *DataRegister) extractValues(elem map[string]interface{}) (values []interface{}) {
	for _, key := range dr.keys {
		values = append(values, elem[key])
		//switch data[key].(type) {
		//case string:
		//	values = append(values, data[key].(string))
		//case int:
		//	values = append(values, strconv.Itoa(data[key].(int)))
		//}
	}
	return values
}

// Insert文を生成する
func (dr *DataRegister) buildInsertSql() (leadSql string) {
	sql := "INSERT INTO " + dr.tableName + " ("

	for i := 0; i < len(dr.keys); i++ {
		sql = sql + dr.keys[i] + ","
	}

	sql = strings.TrimRight(sql, ",")
	sql = sql + ") VALUES ("

	for i := 0; i < len(dr.keys); i++ {
		sql = sql + "?,"
	}

	sql = strings.TrimRight(sql, ",")
	sql = sql + ")"

	return sql
}
