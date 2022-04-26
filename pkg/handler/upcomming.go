package handler

import (
	"fmt"

	"github.com/rodaine/table"
	"github.com/z9fr/ctftime-cli/pkg/utils"
)

func DisplayUpcommingCTF(padding int) {
	events := utils.FetchUpcommingCTF()

	if padding > len(events.Items) {
		fmt.Println("not enough events") // handle errors properly later
	}

	tbl := utils.CustamizeAndCreateTable(table.New("ctf_id", "ctf_name", "start_date", "url"))

	for i := 0; i <= padding; i++ {
		tbl.AddRow(events.Items[i].Custom["ctf_id"], events.Items[i].Custom["ctf_name"], events.Items[i].Custom["start_date"], events.Items[i].Custom["url"])
	}

	tbl.Print()
}
