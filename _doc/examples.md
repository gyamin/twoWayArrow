# 利用例

## CSVファイルデータをDBに取り込む
CSVファイルのデータをデータベースに取り込む

1. NewFileReader にCSVファイルパスを指定して初期化
2. FileReader にCSVファイルのデータ設定を登録する   
    1列目の値が、DBテーブルのカラム"code"に該当する数値
    2列目の値が、DBテーブルのカラム"name"に該当する文字列
    ```go
        fr.AddDefinitions("code", 1, "int")
        fr.AddDefinitions("name", 2, "string")
        fr.AddDefinitions("market", 3, "string")
    ```
3. NewDataRegister にテーブル名とデータベーストランザクションポインタを渡して初期化
4. ImportCsvToDb に FileReader, DataRegister, DB登録前にメモリに保持する行数,進捗表示フラグ を指定して処理実行

#### コードサンプル
```go
package main

import (
	"github.com/gyamin/twoWayArrow/cmd"
	"github.com/gyamin/twoWayArrow/csv"
	"github.com/gyamin/twoWayArrow/db"
	testDb "github.com/gyamin/twoWayArrow/test/db"
	"log"
	"strconv"
)

func main() {
	log.Print("Start")
	csvFile := "./test/csv/data_j.csv"

	fr := csv.NewFileReader(csvFile, true)
	fr.AddDefinitions("code", 1, "int")
	fr.AddDefinitions("name", 2, "string")
	fr.AddDefinitions("market", 3, "string")

	connection := testDb.NewConnection()
	tx, err := connection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	dr := db.NewDataRegister("stock_codes", tx)

	affectedRows := dr.DeleteAll()
	log.Print(strconv.FormatInt(affectedRows, 10) + " Deleted")

	affectedRows = cmd.ImportCsvToDb(fr, dr, 1000, true)
	log.Print(strconv.FormatInt(affectedRows, 10) + " Created")

	tx.Commit()
	log.Print("End")
}
```