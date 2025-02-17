package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	xlsx "github.com/tealeg/xlsx"
)

func textconv(filename string, w io.Writer) error {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if row == nil {
				continue
			}
			cels := make([]string, len(row.Cells))
			for i, cell := range row.Cells {
				var s string
				if cell.Type() == xlsx.CellTypeStringFormula {
					s = cell.Formula()
				} else {
					s = cell.String()
				}

				s = strings.Replace(s, "\\", "\\\\", -1)
				s = strings.Replace(s, "\n", "\\n", -1)
				s = strings.Replace(s, "\r", "\\r", -1)
				s = strings.Replace(s, "\t", "\\t", -1)

				cels[i] = s
			}
			fmt.Fprintf(w, "[%s] %s\n", sheet.Name, strings.Join(cels, "\t"))
		}
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: git-xlsx-textconv file.xlsx")
	}
	excelFileName := os.Args[1]

	err := textconv(excelFileName, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
