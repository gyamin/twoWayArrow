package csv

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestConvertFileToArray(t *testing.T) {

	file, err := os.Open("./../../test/csv/data_j.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	columns := []int{1, 2, 3}

	total := 0

	for {
		data := ConvertFileToArray(reader, columns, 10)
		if len(data) == 0 {
			break
		}

		// データがある間は取得できる要素数は10件(totalが0の初回だけで確認)
		if total == 0 {
			if len(data) != 10 {
				t.Errorf("Return len of array is unexpected")
			}
		}

		total = total + len(data)
	}

	if total != 4017 {
		t.Errorf("Tha total row count is unexpected")
	}
}

func TestConvertFileToMapArray(t *testing.T) {

	file, err := os.Open("./../../test/csv/data_j.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var definitions []map[string]string
	definitions = append(definitions, map[string]string{"key": "code", "pos": "1", "type": "string"})
	definitions = append(definitions, map[string]string{"key": "name", "pos": "2", "type": "string"})
	definitions = append(definitions, map[string]string{"key": "marcket", "pos": "3", "type": "string"})

	total := 0

	for {
		data := ConvertFileToMapArray(*reader, definitions, 10)
		if len(data) == 0 {
			break
		}

		// データがある間は取得できる要素数は10件(totalが0の初回だけで確認)
		if total == 0 {
			if len(data) != 10 {
				t.Errorf("Return len of array is unexpected")
			}
		}

		total = total + len(data)
	}

}
