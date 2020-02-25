package csv

import (
	"encoding/csv"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestConvertFileToArray(t *testing.T) {
	csvFile := "./../../test/csv/data_j.csv"

	cmdLine := "cat " + csvFile + " | wc -l"
	out, err := exec.Command("sh", "-c", cmdLine).Output()
	fileRowNum, _ := strconv.Atoi(strings.TrimSpace(string(out)))

	file, err := os.Open(csvFile)
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

	if total != fileRowNum {
		t.Errorf("Tha total row count is unexpected")
	}
}

func TestConvertFileToMapArray(t *testing.T) {
	csvFile := "./../../test/csv/data_j.csv"

	cmdLine := "cat " + csvFile + " | wc -l"
	out, err := exec.Command("sh", "-c", cmdLine).Output()
	fileRowNum, _ := strconv.Atoi(strings.TrimSpace(string(out)))

	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// CSV読み込み設定
	var definitions []map[string]string
	definitions = append(definitions, map[string]string{"key": "code", "pos": "1", "type": "string"})
	definitions = append(definitions, map[string]string{"key": "name", "pos": "2", "type": "string"})
	definitions = append(definitions, map[string]string{"key": "market", "pos": "3", "type": "string"})

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

	if total != fileRowNum {
		t.Errorf("Tha total row count is unexpected")
	}

}
