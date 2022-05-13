package smalltools

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Dic_to_excel(data map[string]string, sheetName string) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(sheetName)
	rowIndex := 1
	for k, v := range data {
		//设置单元格的值
		xlsx.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex), sheetName+"_"+k)
		xlsx.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex), v)
		rowIndex++
	}

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	err := xlsx.SaveAs("C:\\Users\\LingSi\\Desktop\\国际化\\脚本文件\\" + sheetName + ".xlsx")
	if err != nil {
		log.Fatal(err)
	}
}

func Rows_to_excel(rows []*row, sheetName string) {
	path := "C:\\Users\\LingSi\\Desktop\\国际化\\脚本文件\\ShadowBot.xlsx"
	if !isExist(path) {
		os.Create(path)
	}
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := xlsx.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index := xlsx.NewSheet(sheetName)
	rowIndex := 1
	for _, row := range rows {
		//设置单元格的值
		xlsx.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex), row.key_en)
		xlsx.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex), row.key_zh)
		xlsx.SetCellValue(sheetName, "C"+strconv.Itoa(rowIndex), row.prefix)
		xlsx.SetCellValue(sheetName, "D"+strconv.Itoa(rowIndex), row.file_full_path)
		xlsx.SetCellValue(sheetName, "E"+strconv.Itoa(rowIndex), row.mark)
		rowIndex++
	}
	xlsx.SetActiveSheet(index)
	xlsx.Save()
	if err != nil {
		log.Fatal(err)
	}
}

func ReadExcel(sheetName string) []*row {
	path := "C:\\Users\\LingSi\\Desktop\\国际化\\脚本文件\\UIAutomation.xlsx"
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	ret := []*row{}
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, row_data := range rows {
		data := &row{}
		data.key_en = row_data[0]
		data.key_zh = row_data[1]
		data.prefix = row_data[2]
		data.file_full_path = row_data[3]
		if len(row_data) == 4 {
			data.mark = ""
		} else {
			data.mark = row_data[4]
		}

		ret = append(ret, data)
	}
	return ret
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true

}
