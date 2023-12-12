package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	numbersOfRows = 150_000
	sheetName     = "report"
)

func main() {

	start := time.Now()

	err := createReport()
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}

func createReport() error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create a new sheet.
	err := f.SetSheetName("Sheet1", sheetName)
	if err != nil {
		return err
	}

	err = addHeader(f)
	if err != nil {
		return err
	}

	fillData(f)

	// Save spreadsheet by the given path.
	if err := f.SaveAs("result.xlsx"); err != nil {
		return err
	}

	return nil
}

func addHeader(f *excelize.File) error {
	err := f.MergeCell(sheetName, "A1", "J1")
	if err != nil {
		return err
	}
	err = f.SetCellValue(sheetName, "A1", "my report")
	if err != nil {
		return err
	}

	err = f.MergeCell(sheetName, "B2", "I2")
	if err != nil {
		return err
	}
	err = f.SetCellValue(sheetName, "B2", "some random text")
	if err != nil {
		return err
	}

	return nil
}

func fillData(f *excelize.File) {

	for i := 3; i < numbersOfRows+3; i++ {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i), "aaaaaaaaaaa")
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i), "bbbbbbbbbbb")
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i), "ccccccccccc")
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i), "ddddddddddd")
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", i), "eeeeeeeeeee")
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", i), "fffffffffff")
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", i), "ggggggggggg")
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", i), "hhhhhhhhhhh")
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", i), "iiiiiiiiiii")
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", i), "jjjjjjjjjjj")
	}

}
