package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	file, err := xlsx.OpenFile(filepath.Join(rootPath, "data/data.xlsx"))
	must(err)
	sheet := findSheetTerm(file)

	rows := sheet.Rows
	headerRowAt, headerDesc := loadHeaderRow(rows)
	parsedCells := loadCellData(headerDesc, rows[headerRowAt+1:])

	var b bytes.Buffer
	b.WriteString(`export default `)
	jw := json.NewEncoder(&b)
	jw.SetIndent("", "  ")
	must(jw.Encode(parsedCells))

	outFile := "src/_tmp/data.js"
	err = ioutil.WriteFile(filepath.Join(rootPath, outFile), b.Bytes(), 0644)
	must(err)

	fmt.Printf("  num rows: %v\n", len(parsedCells))
	fmt.Printf("output to %v\n", outFile)
}

func findSheetTerm(file *xlsx.File) *xlsx.Sheet {
	for _, sh := range file.Sheets {
		if strings.Contains(sh.Name, "(term)") {
			return sh
		}
	}
	panic(`can not find sheet "term"`)
}
