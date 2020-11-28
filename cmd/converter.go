package cmd

import (
	"github.com/gyamin/twoWayArrow/csv"
	"github.com/gyamin/twoWayArrow/db"
	"log"
	"strconv"
)

func ImportCsvToDb(reader csv.FileReader, register db.DataRegister, i int, showLog bool) (importedRows int64) {

	for {
		csvData := reader.ConvertFileToMapArray(i)
		if len(csvData) == 0 {
			break
		}
		importedRows = importedRows + register.CreateData(csvData)
		if showLog {
			log.Print(strconv.FormatInt(importedRows, 10))
		}
	}
	return importedRows
}
