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

	affectedRows = cmd.ImportCsvToDb(fr, dr, 1000, true)
	log.Print(strconv.FormatInt(affectedRows, 10) + " Created")

	tx.Commit()
	log.Print("End")
}
