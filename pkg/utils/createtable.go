package utils

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func CustamizeAndCreateTable(tbl table.Table) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	return tbl
}
