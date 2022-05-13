package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/guptarohit/asciigraph"
	log "github.com/sirupsen/logrus"
	"github.com/z9fr/ctftime-cli/pkg/models"
)

func createGraph(url string) {
	//	url := "https://ctftime.org/api/v1/teams/1005/"

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	ctfteam := models.CTFTEAM{}
	jsonErr := json.Unmarshal(body, &ctfteam)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	a := ctfteam.Rating
	data := []float64{float64(a.Num2011.RatingPoints), float64(a.Num2012.RatingPoints), float64(a.Num2013.RatingPoints),
		float64(a.Num2014.RatingPoints), float64(a.Num2015.RatingPoints), float64(a.Num2016.RatingPoints), float64(a.Num2017.RatingPoints),
		float64(a.Num2018.RatingPoints), float64(a.Num2019.RatingPoints), float64(a.Num2020.RatingPoints), float64(a.Num2021.RatingPoints),
		float64(a.Num2022.RatingPoints)}

	graph := asciigraph.Plot(data, asciigraph.Width(50), asciigraph.Height(10), asciigraph.Caption("2011 2012 2013 2014 2015 2016 2017 2018 2019 2020 2021\n\n\t\t\tRating Points"), asciigraph.Precision(0))
	fmt.Println(graph)

}
