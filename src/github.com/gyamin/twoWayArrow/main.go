package main

import (
	"fmt"
	"github.com/gyamin/twoWayArrow/internal/csv"
	"github.com/gyamin/twoWayArrow/internal/db"
	"os"
)

func main() {
	p, _ := os.Getwd()
	fmt.Println(p)
	//csvFile := "./src/github.com/gyamin/twoWayArrow/test/csv/data_j.csv"
	csvFile := "./test/csv/data_j.csv"

	fr := csv.NewFileReader(csvFile)
	fr.AddDefinitions("code", 1, "string")
	fr.AddDefinitions("name", 2, "string")
	fr.AddDefinitions("market", 3, "string")

	for {
		csvData := fr.ConvertFileToMapArray(1000)
		if len(csvData) == 0 {
			break
		}
		insertSql := db.BuildInsertSql("stocks", csvData)
		println(len(insertSql))
	}
}
