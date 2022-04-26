package handler

import (
	"errors"
	"fmt"

	"github.com/rodaine/table"
	"github.com/z9fr/ctftime-cli/pkg/utils"
)

func DisplayUpcommingCTFRSS(padding int) {
	events := utils.FetchUpcommingCTFRSS()

	if padding > len(events.Items) {
		utils.LogError(errors.New("Not Enough Events to Display"))
	}

	tbl := utils.CustamizeAndCreateTable(table.New("ctf_title", "ctf_name", "weight", "start_date", "ctf_url", "ctftime_url"))

	for i := 0; i <= padding; i++ {
		tbl.AddRow(events.Items[i].Title, events.Items[i].Custom["ctf_name"], events.Items[i].Custom["weight"],
			events.Items[i].Custom["start_date"], events.Items[i].Custom["url"],
			events.Items[i].Link)
	}

	tbl.Print()
}

func DisplayUpcommingCTFJSON() {
	ctfevents := utils.FetchUpcommingJSON()

	tbl := utils.CustamizeAndCreateTable(table.New("ctf_id", "ctf_title", "weight", "start_date", "duration", "ctf_url", "ctftime_url"))

	for _, ctf := range ctfevents {
		duration := fmt.Sprintf("%d days %d hours", ctf.Duration.Days, ctf.Duration.Hours)

		tbl.AddRow(ctf.ID, ctf.Title, ctf.Weight, ctf.Start.Format("02-Jan-2006"), duration, ctf.URL, ctf.CtftimeURL)
	}

	tbl.Print()
}
