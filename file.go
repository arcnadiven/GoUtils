package utils

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

func ReadExcelToObject(filePath, sheet string, object interface{}) error {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()
	rows, err := file.GetRows(sheet)
	if err != nil {
		return errors.WithStack(err)
	}

	objectList := reflect.ValueOf(object).Elem()

	// 建立sheet第一行跟struct field对应关系，struct tag为excel
	rowIdxs := map[int]int{}
	for rowIdx, row := range rows[0] {
		for idx := 0; idx < objectList.Type().Elem().NumField(); idx++ {
			if objectList.Type().Elem().Field(idx).Tag.Get("excel") == row {
				rowIdxs[rowIdx] = idx
			}
		}
	}

	for _, row := range rows[1:] {

		elem := reflect.New(objectList.Type().Elem()).Elem()

		for rowIdx, colCell := range row {
			fieldIdx, ok := rowIdxs[rowIdx]
			if !ok {
				continue
			}
			switch elem.Field(fieldIdx).Type().String() {
			case "int":
				cell, err := strconv.ParseInt(colCell, 10, 64)
				if err != nil {
					elem.Field(fieldIdx).SetInt(0)
					continue
				}
				elem.Field(fieldIdx).SetInt(cell)
			case "float64":
				cell, err := strconv.ParseFloat(colCell, 64)
				if err != nil {
					return errors.WithStack(err)
				}
				elem.Field(fieldIdx).SetFloat(cell)
			default:
				elem.Field(fieldIdx).SetString(colCell)
			}
		}

		objectList.Set(reflect.Append(objectList, elem))
	}
	return nil
}

func ReadRowsInSheet(filePath, sheet string) ([]string, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()
	rows, err := file.GetRows(sheet)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret := make([]string, 0)
	for _, row := range rows {
		for _, cell := range row {
			ret = append(ret, strings.TrimSpace(cell))
		}
	}
	return ret, nil
}

func CountLines(filePath string) (map[string]int, error) {
	ret := make(map[string]int)

	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()
	euCount := 0
	//usCount := int(0)

	list := file.GetSheetList()
	for idx, sheet := range list {
		if idx < 3 {
			continue
		}
		if strings.HasSuffix(sheet, "US") || strings.HasSuffix(sheet, "CA") {
			continue
		}
		rows, err := file.GetRows(sheet)
		if err != nil {
			return nil, err
		}
		ret[sheet] = len(rows)
		//if sheet[len(sheet)-2:] != "US"  {
		//	euCount += len(rows)
		//}
		euCount += len(rows)
	}
	return ret, nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func NewExcelTemplate(filePath string, fields []string) error {
	file := excelize.NewFile()
	defer file.Close()
	_, err := file.NewSheet("Sheet1")
	if err != nil {
		return errors.WithStack(err)
	}
	cell, err := excelize.CoordinatesToCellName(1, 1)
	if err != nil {
		return errors.WithStack(err)
	}
	err = file.SetSheetRow("Sheet1", cell, &fields)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := file.SaveAs(filePath); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
