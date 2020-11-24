package csv

import (
	"encoding/csv"
	"os"
	"strconv"
)

type FileReader struct {
	filePath    string
	reader      *csv.Reader
	definitions []map[string]interface{}
}

func NewFileReader(filePath string) (fileReader FileReader) {
	fileReader = FileReader{}
	fileReader.filePath = filePath
	file, err := os.Open(fileReader.filePath)
	if err != nil {
		panic(err)
	}
	fileReader.reader = csv.NewReader(file)
	return fileReader
}

func (fr *FileReader) AddDefinitions(key string, position int, kata string) {
	def := make(map[string]interface{})
	def["key"] = key
	def["position"] = position
	def["kata"] = kata
	fr.definitions = append(fr.definitions, def)
}

// reader から definition の設定にしたがってmapを作成し rowNum 行分配列で返す
// 	definition が [{"pos":0, "key":"id", "type":"int"}, {"pos":2, "key":"name", "type":"string"} で、
//	csvの行データが10,Tokyo,太郎の場合、{"id":10, "name":"太郎"} を作成する
func (fr FileReader) ConvertFileToMapArray(rowNum int, rowHeader bool) []map[string]interface{} {

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

		readLine, err = fr.reader.Read()
		if err != nil {
			break
		}

		// 行ヘッダーが有る場合、1行目スキップ
		if rowHeader && i == 1 {
			continue
		}

		// definition を走査
		mapLine := make(map[string]interface{})
		for _, def := range fr.definitions {
			pos := def["position"]
			key := def["key"]
			kata := def["kata"]

			switch kata {
			case "string":
				mapLine[key.(string)] = readLine[pos.(int)]
			case "int":
				mapLine[key.(string)], _ = strconv.Atoi(readLine[pos.(int)])
			}
		}

		data = append(data, mapLine)
	}

	return data
}
