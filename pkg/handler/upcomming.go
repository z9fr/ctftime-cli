package handler

import (
	"fmt"

	"github.com/z9fr/ctftime-cli/pkg/utils"
)

func DisplayUpcommingCTF(padding int) {
	events := utils.FetchUpcommingCTF()

	if padding > len(events.Items) {
		fmt.Println("not enough events") // handle errors properly later
	}

	for i := 0; i <= padding; i++ {
		fmt.Println(events.Items[0])
	}

	tbl := utils.CreateTable("ID", "Name", "Score", "Added")
	tbl.AddRow(1, "test", "score", "added")
	tbl.Print()
}
