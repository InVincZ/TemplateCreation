package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func toChar(i int) string {
	return string('A' - 1 + i)
}

func readexcel() {
	excelFileName := "/Users/Micha/Desktop/FillWordTemplate/ReadFile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		sheet_name := sheet.Name
		for row_number, row := range sheet.Rows {
			for cell_number, cell := range row.Cells {
				cell_string := toChar(cell_number + 1)
				val, _ := cell.String()
				if val != "" {
					fmt.Printf("Sheet: %s -- Cell:%s%d: %s\n", sheet_name, cell_string, row_number+1, val)
				}
			}
		}
	}
}
