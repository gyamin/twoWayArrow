package main

import (
	"github.com/gyamin/twoWayArrow/internal/csv"
	"github.com/gyamin/twoWayArrow/internal/db"
	testDb "github.com/gyamin/twoWayArrow/test/db"
	"log"
	"strconv"
)

func main() {
	log.Print("Start")
	//csvFile := "./src/github.com/gyamin/twoWayArrow/test/csv/data_j.csv"
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

	for {
		csvData := fr.ConvertFileToMapArray(1000)
		if len(csvData) == 0 {
			break
		}
		affectedRows = dr.CreateData(csvData)
		log.Print(strconv.FormatInt(affectedRows, 10) + " Created")
	}
	tx.Commit()
	log.Print("End")
}
