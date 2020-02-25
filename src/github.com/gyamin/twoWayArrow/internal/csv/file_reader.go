package csv

import (
	"encoding/csv"
	"strconv"
)

// reader から columns で指定された要素を rowNum 行分配列で返す
func ConvertFileToArray(reader *csv.Reader, columns []int, rowNum int) [][]string {

	var data [][]string      // 戻り値で返すオブジェクト配列
	var readLine []string    // csvファイルの行 , 区切りで配列になる
	var convertLine []string // columnsの設定で抽出された配列
	var err error

	var i int

	for {
		if i >= rowNum {
			break
		} else {
			i++
		}

		readLine, err = reader.Read()
		if err != nil {
			break
		}

		for elem := range columns {
			convertLine = append(convertLine, readLine[elem])
		}

		data = append(data, convertLine)
	}

	return data
}

// reader から definition の設定にしたがってmapを作成し rowNum 行分配列で返す
// 	definition が [{"pos":0, "key":"id", "type":"int"}, {"pos":2, "key":"name", "type":"string"} で、
//	csvの行データが10,Tokyo,太郎の場合、{"id":10, "name":"太郎"} を作成する
func ConvertFileToMapArray(reader csv.Reader, definitions []map[string]string, rowNum int) []map[string]interface{} {

	var data []map[string]interface{} // 戻り値で返すmap配列
	var readLine []string             // csvファイルの行 , 区切りで配列になる
	var err error

	var i int
	for {
		// rowNum行で読み込み止める
		if i >= rowNum {
			break
		} else {
			i++
		}

		readLine, err = reader.Read()
		if err != nil {
			break
		}

		// definition を走査
		mapLine := make(map[string]interface{})
		for _, def := range definitions {
			pos, _ := strconv.Atoi(def["pos"])
			key := def["key"]
			kata := def["type"]

			switch kata {
			case "string":
				mapLine[key] = readLine[pos]
			case "int":
				mapLine[key], _ = strconv.Atoi(readLine[pos])
			}
		}
		data = append(data, mapLine)
	}

	return data
}
