package utils

import (
	"github.com/mmcdole/gofeed"
)

// Fetch upcomming ctf events and return as feed struct
func FetchUpcommingCTF() gofeed.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://ctftime.org/event/list/upcoming/rss/")
	return *feed
}
