package csv

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestConvertFileToMapArray(t *testing.T) {
	csvFile := "./../../test/csv/data_j.csv"

	cmdLine := "cat " + csvFile + " | wc -l"
	out, _ := exec.Command("sh", "-c", cmdLine).Output()
	fileRowNum, _ := strconv.Atoi(strings.TrimSpace(string(out)))

	fr := NewFileReader(csvFile)
	fr.AddDefinitions("code", 1, "string")
	fr.AddDefinitions("name", 2, "string")
	fr.AddDefinitions("market", 3, "string")

	total := 0

	for {
		data := fr.ConvertFileToMapArray(10)
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
