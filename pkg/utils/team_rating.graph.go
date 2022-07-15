package utils

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/z9fr/ctftime-cli/pkg/models"
)

func CreateGraph(ctfteam models.CTFTEAM) {
	a := ctfteam.Rating
	data := []float64{float64(a.Num2011.RatingPoints), float64(a.Num2012.RatingPoints), float64(a.Num2013.RatingPoints),
		float64(a.Num2014.RatingPoints), float64(a.Num2015.RatingPoints), float64(a.Num2016.RatingPoints), float64(a.Num2017.RatingPoints),
		float64(a.Num2018.RatingPoints), float64(a.Num2019.RatingPoints), float64(a.Num2020.RatingPoints), float64(a.Num2021.RatingPoints),
		float64(a.Num2022.RatingPoints)}

	graph := asciigraph.Plot(data, asciigraph.Width(50), asciigraph.Height(10), asciigraph.Caption("2011 2012 2013 2014 2015 2016 2017 2018 2019 2020 2021\n\n\t\t\tRating Points"), asciigraph.Precision(0))
	fmt.Println(graph)

}
