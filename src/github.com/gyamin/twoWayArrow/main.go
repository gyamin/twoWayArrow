package main

import (
	"fmt"
	"github.com/gyamin/twoWayArrow/internal/csv"
	"github.com/gyamin/twoWayArrow/internal/db"
	testdb "github.com/gyamin/twoWayArrow/test/db"
	"os"
)

func main() {
	p, _ := os.Getwd()
	fmt.Println(p)
	//csvFile := "./src/github.com/gyamin/twoWayArrow/test/csv/data_j.csv"
	csvFile := "./test/csv/data_j.csv"

	fr := csv.NewFileReader(csvFile)
	fr.AddDefinitions("code", 1, "int")
	fr.AddDefinitions("name", 2, "string")
	fr.AddDefinitions("market", 3, "string")

	for {
		csvData := fr.ConvertFileToMapArray(1000, true)
		if len(csvData) == 0 {
			break
		}

		connection := testdb.NewConnection()
		dr := db.NewDataRegister("stock_codes", csvData, connection)
		affectedRow := dr.CreateData()
		println(affectedRow)
		//insertSql := db.BuildInsertSql("stocks", csvData)
		//println(len(insertSql))
	}
}
