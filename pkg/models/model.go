package models

import "time"

type CTFevent []struct {
	Organizers []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"organizers"`
	Onsite        bool      `json:"onsite"`
	Finish        time.Time `json:"finish"`
	Description   string    `json:"description"`
	Weight        float64   `json:"weight"`
	Title         string    `json:"title"`
	URL           string    `json:"url"`
	IsVotableNow  bool      `json:"is_votable_now"`
	Restrictions  string    `json:"restrictions"`
	Format        string    `json:"format"`
	Start         time.Time `json:"start"`
	Participants  int       `json:"participants"`
	CtftimeURL    string    `json:"ctftime_url"`
	Location      string    `json:"location"`
	LiveFeed      string    `json:"live_feed"`
	PublicVotable bool      `json:"public_votable"`
	Duration      struct {
		Hours int `json:"hours"`
		Days  int `json:"days"`
	} `json:"duration"`
	Logo     string `json:"logo"`
	FormatID int    `json:"format_id"`
	ID       int    `json:"id"`
	CtfID    int    `json:"ctf_id"`
}

type CTFTEAM struct {
	Academic     bool   `json:"academic"`
	PrimaryAlias string `json:"primary_alias"`
	Name         string `json:"name"`
	Rating       struct {
		Num2011 struct {
			CountryPlace int `default:"0" json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2011"`
		Num2012 struct {
			CountryPlace int `default:"0" json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2012"`
		Num2013 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2013"`
		Num2014 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2014"`
		Num2015 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2015"`
		Num2016 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2016"`
		Num2017 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2017"`
		Num2018 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2018"`
		Num2019 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2019"`
		Num2020 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2020"`
		Num2021 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2021"`
		Num2022 struct {
			CountryPlace int `json:"country_place"`

			RatingPoints float64 `default:"0" json:"rating_points"`
		} `json:"2022"`
	} `json:"rating"`
	Logo    string   `json:"logo"`
	Country string   `json:"country"`
	ID      int      `json:"id"`
	Aliases []string `json:"aliases"`
}
