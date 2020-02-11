package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func TestConvertFileToArray(t *testing.T) {
	p, _ := os.Getwd()
	fmt.Println(p)

	file, err := os.Open("./../../test/csv/data_j.xls")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	columns := []int{1, 2, 3}

	reader, data := ConvertFileToArray(reader, columns, 10)

	for elem := range data {
		println(elem)
	}
}
