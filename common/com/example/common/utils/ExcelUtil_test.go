package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"testing"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Outter struct {
}

func (it *Outter) Write(p []byte) (n int, err error) {
	var f, e = os.Create("a.xls")
	if e != nil {
		return 0, e
	}
	defer f.Close()
	f.Write(p)
	return 0, nil
}

func TestCreateExcelFile(t *testing.T) {
	var o = Outter{}
	CreateExcelFile(func(xlsx *excelize.File) {
		SetExcelCellValue(xlsx, 0, 0, "a")
	}, &o)
}

func TestExport(t *testing.T) {

	var datas = []Data{
		{
			Name: "xiao ming",
			Age:  1,
		},
	}

	var o = Outter{}
	Export(ExportDTO{
		Titles: []ExcelTitle{
			{
				Title: "名字",
				Field: "Name",
			},
			{
				Title: "年龄",
				Field: "Age",
				Format: func(arg interface{}) interface{} {
					var item = arg.(int)
					return float64(item) / float64(100)
				},
			},
		},
		DataArray: datas,
	}, &o)
}
