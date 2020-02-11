package csv

import (
	"encoding/csv"
)

// reader から columns で指定された要素を cnt 行分配列で返す
func ConvertFileToArray(reader *csv.Reader, columns []int, cnt int) (*csv.Reader, [][]string) {

	var data [][]string      // 戻り値で返すオブジェクト配列
	var readLine []string    // csvファイルの行 , 区切りで配列になる
	var convertLine []string // columnsの設定で抽出された配列
	var err error

	for {
		readLine, err = reader.Read()
		if err != nil {
			break
		}

		for elem := range columns {
			convertLine = append(convertLine, readLine[elem])
		}

		data = append(data, convertLine)
	}

	return reader, data
}
