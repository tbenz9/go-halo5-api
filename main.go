package main

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

	// Matches For Player
	//fmt.Println(string(halo.MatchesForPlayer(baseurl, title, sampleGamertag, "warzone", 0, 0)))

	// Events For Match
	//fmt.Println(string(halo.EventsForMatch(baseurl, title, sampleMatchID)))

	// Player Leaderboards
	//fmt.Println(string(halo.PlayerLeaderboard(baseurl, title, sampleSeasonID, samplePlaylistID, 0)))

	// Service Record: Arena
	//fmt.Println(string(halo.ServiceRecordArena(baseurl, title, samplePlayers, sampleSeasonID)))

	// Vehicles
	//fmt.Println(string(halo.Vehicles(baseurl, title)))

	// Spartan Image
	//fmt.Println(string(halo.SpartanImage(baseurl, title, sampleGamertag, 0, "")))
}
