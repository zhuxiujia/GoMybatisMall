package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"reflect"
	"strconv"
)

func CreateExcelFile(Create func(xlsx *excelize.File), out io.Writer) {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet1")
	Create(xlsx)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.Write(out)
	if err != nil {
		fmt.Println(err)
	}
}

func SetExcelCellValue(xlsx *excelize.File, x int, y int, obj interface{}) {
	if y < 0 || x < 0 {
		panic("x and y must >= 0 !")
	}
	// Set value of a cell.
	xlsx.SetCellValue("Sheet1", ToAlphaString(x)+strconv.Itoa(y+1), obj)
}

func ToAlphaString(value int) string {
	if value < 0 {
		return ""
	}
	var ans string
	i := value + 1
	for i > 0 {
		ans = string((i-1)%26+65) + ans
		i = (i - 1) / 26
	}
	return ans
}

type ExcelTitle struct {
	Title  string
	Field  string
	Format func(arg interface{}) interface{}
}

type ExportDTO struct {
	Titles    []ExcelTitle
	DataArray interface{}
}

func Export(dto ExportDTO, out io.Writer) {
	CreateExcelFile(func(xlsx *excelize.File) {
		for x, title := range dto.Titles {
			SetExcelCellValue(xlsx, x, 0, title.Title)
		}
		var valuesValue = reflect.ValueOf(dto.DataArray)
		var valuesType = reflect.TypeOf(dto.DataArray)
		var itemKind = valuesType.Kind()
		if itemKind == reflect.Slice || itemKind == reflect.Array {
		} else {
			panic("Export only support [] data")
		}

		for y := 0; y < valuesValue.Len(); y++ {
			var value = valuesValue.Index(y)
			if (value.Kind() == reflect.Interface || value.Kind() == reflect.Ptr) && value.CanAddr() {
				value = value.Elem()
			}
			for x := 0; x < len(dto.Titles); x++ {
				var title = dto.Titles[x]
				println(value.String())
				if value.Kind() == reflect.Struct {
					var field = value.FieldByName(title.Field)
					if field.IsValid() {
						var v = field.Interface()
						if title.Format != nil {
							SetExcelCellValue(xlsx, x, 1+y, title.Format(v))
						} else {
							SetExcelCellValue(xlsx, x, 1+y, v)
						}
					} else {
						SetExcelCellValue(xlsx, x, 1+y, "")
					}
				} else if value.Kind() == reflect.Map {
					var field = value.MapIndex(reflect.ValueOf(title.Field))
					if field.IsValid() {
						var v = field.Interface()
						if title.Format != nil {
							SetExcelCellValue(xlsx, x, 1+y, title.Format(v))
						} else {
							SetExcelCellValue(xlsx, x, 1+y, v)
						}
					} else {
						SetExcelCellValue(xlsx, x, 1+y, "")
					}
				}

			}
		}

	}, out)
}
