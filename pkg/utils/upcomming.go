package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/z9fr/ctftime-cli/pkg/models"
)

// Fetch upcomming ctf events and return as feed struct
func FetchUpcommingCTFRSS() gofeed.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://ctftime.org/event/list/upcoming/rss/")
	return *feed
}

func FetchUpcommingJSON() models.CTFevent {
	url := "https://ctftime.org/api/v1/events/"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		LogError(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		LogError(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		LogError(readErr)
	}

	evenz := models.CTFevent{}
	jsonErr := json.Unmarshal(body, &evenz)
	if jsonErr != nil {
		LogError(jsonErr)
	}

	return evenz
}
