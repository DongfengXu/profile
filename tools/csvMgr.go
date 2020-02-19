package tools

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

type CsvTable struct {
	FileName string
	Records  []CsvRecord
}

type CsvRecord struct {
	Record map[string]string
}

func (c *CsvRecord) GetInt(field string) int {
	var r int
	var err error
	if r, err = strconv.Atoi(c.Record[field]); err != nil {
		beego.Error(err)
		panic(err)
	}
	return r
}

func (c *CsvRecord) GetString(field string) string {
	data, ok := c.Record[field]
	if ok {
		return data
	} else {
		beego.Warning("Get fileld failed! fileld:", field)
		return ""
	}
}

func LoadCsvCfg(filename string, row int) *CsvTable {
	file, err := os.Open(filename)
	if err != nil {
		beego.Error(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if reader == nil {
		beego.Error("NewReader return nil, file:", file)
		return nil
	}
	records, err := reader.ReadAll()
	if err != nil {
		beego.Error(err)
		return nil
	}
	if len(records) < row {
		beego.Warning(filename, " is empty")
		return nil
	}
	colNum := len(records[0])
	recordNum := len(records)
	var allRecords []CsvRecord
	for i := row; i < recordNum; i++ {
		record := &CsvRecord{make(map[string]string)}
		for k := 0; k < colNum; k++ {
			record.Record[records[0][k]] = records[i][k]
		}
		allRecords = append(allRecords, *record)
	}
	var result = &CsvTable{
		filename,
		allRecords,
	}
	return result
}


//生成csv文件
func CreateCSV(txtname string,title map[string] *Profile_CIO)  {
	f , err := os.Create(txtname)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")// 写入UTF-8 BOM
	w:=csv.NewWriter(f)

	for _, record := range title {
		if err := w.Write([]string{record.APP_CODE, record.CONEXT,record.CIO}); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()
}