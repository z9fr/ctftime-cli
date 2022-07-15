package handler

import (
	"fmt"
	"github.com/fatih/color"

	"github.com/sirupsen/logrus"
	"github.com/z9fr/ctftime-cli/pkg/utils"
)

func TeamDetails(teamID string) {

	ctfTeamInfo, err := utils.FetchTeamInformation(fmt.Sprintf("https://ctftime.org/api/v1/teams/%s/", teamID))

	if err != nil {
		logrus.Fatal(err)
	}

	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.Underline)
	boldWhite.Println(fmt.Sprintf("\n%s | %s\n", ctfTeamInfo.Name, ctfTeamInfo.Country))

	fmt.Println("Aliases")
	for _, a := range ctfTeamInfo.Aliases {
		fmt.Printf("- %s\n", a)
	}
	fmt.Printf("\n%s \n\n", ctfTeamInfo.Logo)

	utils.CreateGraph(ctfTeamInfo)

}
