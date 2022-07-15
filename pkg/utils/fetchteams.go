package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/z9fr/ctftime-cli/pkg/models"
)

func FetchTeamInformation(url string) (models.CTFTEAM, error) {

	httpClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.CTFTEAM{}, err
	}

	req.Header.Set("User-Agent", "ctftime cli https://github.com/z9fr/ctftime-cli/")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return models.CTFTEAM{}, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return models.CTFTEAM{}, readErr
	}

	ctfteam := models.CTFTEAM{}
	jsonErr := json.Unmarshal(body, &ctfteam)
	if jsonErr != nil {
		return models.CTFTEAM{}, jsonErr
	}

	return ctfteam, nil
}
