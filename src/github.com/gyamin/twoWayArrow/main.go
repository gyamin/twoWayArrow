package main

import (
	"fmt"
	"github.com/gyamin/twoWayArrow/internal/csv"
	"github.com/gyamin/twoWayArrow/internal/db"
	testdb "github.com/gyamin/twoWayArrow/test/db"
	"log"
	"os"
)

func main() {
	p, _ := os.Getwd()
	fmt.Println(p)
	//csvFile := "./src/github.com/gyamin/twoWayArrow/test/csv/data_j.csv"
	csvFile := "./test/csv/data_j.csv"

	fr := csv.NewFileReader(csvFile, true)
	fr.AddDefinitions("code", 1, "int")
	fr.AddDefinitions("name", 2, "string")
	fr.AddDefinitions("market", 3, "string")

	connection := testdb.NewConnection()
	tx, err := connection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	dr := db.NewDataRegister("stock_codes", tx)

	affectedRows := dr.DeleteAll()

	for {
		csvData := fr.ConvertFileToMapArray(1000)
		if len(csvData) == 0 {
			break
		}
		affectedRows = dr.CreateData(csvData)
		println(affectedRows)
	}
	tx.Commit()
}
