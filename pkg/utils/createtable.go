package utils

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func CreateTable(fields ...string) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New(fields)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	return tbl
}
