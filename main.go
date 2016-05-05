package main

import (
	"fmt"

	"github.com/tbenz9/go-halo5-api/halo"
)

var baseurl string = "https://www.haloapi.com"
var title string = "h5"

// Sample values for testing
var sampleGamertag string = "motta13"
var sampleMode string = "warzone"
var sampleMatchID string = "c35a35f8-f450-4836-a4c2-65100a7acb79"
var sampleSeasonID string = "b46c2095-4ca6-4f4b-a565-4702d7cfe586"   //February 2016 Season
var samplePlaylistID string = "2323b76a-db98-4e03-aa37-e171cfbdd1a4" //SWAT gametype 2016 Season
var samplePlayers string = "motta13,smoke721"

func main() {
	// Uncomment any of the below for sample output.
	fmt.Println(halo.MatchesForPlayer(baseurl, title, sampleGamertag, "warzone", 0, 0))
	//fmt.Println(events_for_match(baseurl, title, sampleMatchID))
	//fmt.Println(player_leaderboard(baseurl, title, sampleSeasonID, samplePlaylistID, 0))
	//fmt.Println(service_record_arena(baseurl, title, samplePlayers, sampleSeasonID))
	//fmt.Println(vehicles(baseurl, title))
	//fmt.Println(spartan_image(baseurl, title, gamertag, 0, ""))
}
