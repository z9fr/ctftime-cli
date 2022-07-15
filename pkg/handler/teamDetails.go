package handler

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/z9fr/ctftime-cli/pkg/utils"
)

func TeamDetails(teamID string) {

	ctfTeamInfo, err := utils.FetchTeamInformation(fmt.Sprintf("https://ctftime.org/api/v1/teams/%s/", teamID))

	if err != nil {
		logrus.Fatal(err)
	}

	utils.CreateGraph(ctfTeamInfo)

}
