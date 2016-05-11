package main

import (
	"fmt"
	"os"

	"github.com/tbenz9/go-halo5-api/halo"
)

var baseurl string = "https://www.haloapi.com"
var title string = "h5"

// Sample values for testing
var sampleGamertag string = "motta13"
var sampleMode string = "warzone"
var sampleArenaMatchID string = "f15986a8-1132-48d7-9194-e23388ec6084"
var sampleCampaignMatchID string = "f9d5a884-68a5-4e01-a9cc-92239787559f"
var sampleCustomMatchID string = "5e0985de-309c-4031-8133-fea03500fd1b"
var sampleWarzoneMatchID string = "c35a35f8-f450-4836-a4c2-65100a7acb79"
var sampleSeasonID string = "b46c2095-4ca6-4f4b-a565-4702d7cfe586"      //February 2016 Season
var samplePlaylistID string = "2323b76a-db98-4e03-aa37-e171cfbdd1a4"    //SWAT gametype 2016 Season
var sampleGameVariantID string = "963ca478-369a-4a37-97e3-432fa13035e1" //Slayer
var sampleMapVariantsID string = "a44373ee-9f63-4733-befd-5cd8fbb1b44a" //Truth
var sampleRequisitionPacksID string = "d10141cb-68a5-4c6b-af38-4e4935f973f7"
var sampleRequisitionID string = "e4f549b2-90af-4dab-b2bc-11a46ea44103"
var samplePlayers string = "motta13,smoke721"

func getAPIKey() string {
	apikey := os.Getenv("HALO_API_KEY")
	if len(apikey) != 32 {
		fmt.Println("Invalid API Key")
	}
	return apikey
}

func main() {

	h := halo.NewHalo(baseurl, title, getAPIKey())
	//fmt.Println(h.Enemies())
	//fmt.Println(h.FlexibleStats())
	//fmt.Println(h.GameBaseVariants())
	//fmt.Println(h.Impulses())
	//fmt.Println(h.Maps())
	//fmt.Println(h.Medals())
	//fmt.Println(h.Playlists())
	//fmt.Println(h.Seasons())
	//fmt.Println(h.Skulls())
	//fmt.Println(h.SpartanRanks())
	//time.Sleep(time.Second * 10)
	//fmt.Println(h.TeamColors())
	//fmt.Println(h.Vehicles())
	//fmt.Println(h.Weapons())

	//fmt.Println(h.GameVariants(sampleGameVariantID))
	//fmt.Println(h.MapVariants(sampleMapVariantsID))
	//fmt.Println(h.Requisitions(sampleRequisitionID))
	//fmt.Println(h.MatchesForPlayer(sampleGamertag, "", 0, 0))
	//fmt.Println(h.PlayerLeaderboard(sampleSeasonID, samplePlaylistID, 0))
	fmt.Println(h.CarnageReportArena(sampleArenaMatchID))
	// Uncomment any of the below for sample output.

	// Metadata
	//fmt.Println(halo.CampaignMissions(baseurl, title))
	//fmt.Println(halo.Commendations(baseurl, title))
	//fmt.Println(halo.CsrDesignations(baseurl, title))
	//fmt.Println(string(halo.Enemies(baseurl, title)))
	//fmt.Println(string(halo.FlexibleStats(baseurl, title)))
	//fmt.Println(string(halo.GameBaseVariants(baseurl, title)))
	//fmt.Println(string(halo.GameVariants(baseurl, title)))
	//fmt.Println(string(halo.Impulses(baseurl, title)))
	//fmt.Println(string(halo.MapVariants(baseurl, title)))
	//fmt.Println(string(halo.Maps(baseurl, title)))
	//fmt.Println(string(halo.Medals(baseurl, title)))
	//fmt.Println(string(halo.Playlists(baseurl, title)))
	//fmt.Println(string(halo.RequisitionPacks(baseurl, title)))
	//fmt.Println(string(halo.Requisitions(baseurl, title)))
	//fmt.Println(string(halo.Seasons(baseurl, title)))
	//fmt.Println(string(halo.Sculls(baseurl, title)))
	//fmt.Println(string(halo.SpartanRanks(baseurl, title)))
	//fmt.Println(string(halo.TeamColors(baseurl, title)))
	//fmt.Println(string(halo.Vehicles(baseurl, title)))
	//fmt.Println(string(halo.Weapons(baseurl, title)))

	//ioutil.WriteFile("tmp/Enemies.txt", halo.Enemies(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/FlexibleStats.txt", halo.FlexibleStats(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/GameBaseVariants.txt", halo.GameBaseVariants(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/GameVariants.txt", halo.GameVariants(baseurl, title, sampleGameVariantID), 0755)
	//ioutil.WriteFile("tmp/Impulses.txt", halo.Impulses(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/MapVariants.txt", halo.MapVariants(baseurl, title, sampleMapVariantsID), 0755)
	//ioutil.WriteFile("tmp/Maps.txt", halo.Maps(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/Medals.txt", halo.Medals(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/Playlists.txt", halo.Playlists(baseurl, title), 0755)
	//fmt.Println("Sleeping")
	//time.Sleep(10 * time.Second)
	//ioutil.WriteFile("tmp/RequisitionPacks.txt", halo.RequisitionPacks(baseurl, title, sampleRequisitionPacksID), 0755)
	//ioutil.WriteFile("tmp/Requisitions.txt", halo.Requisitions(baseurl, title, sampleRequisitionID), 0755)
	//ioutil.WriteFile("tmp/Seasons.txt", halo.Seasons(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/Skulls.txt", halo.Skulls(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/SpartanRanks.txt", halo.SpartanRanks(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/TeamColors.txt", halo.TeamColors(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/Vehicles.txt", halo.Vehicles(baseurl, title), 0755)
	//ioutil.WriteFile("tmp/Weapons.txt", halo.Weapons(baseurl, title), 0755)

	// Matches For Player
	//fmt.Println(string(halo.MatchesForPlayer(baseurl, title, sampleGamertag, "warzone", 0, 0)))

	// Events For Match
	//fmt.Println(string(halo.EventsForMatch(baseurl, title, sampleArenaMatchID)))

	// Player Leaderboards
	//fmt.Println(string(halo.PlayerLeaderboard(baseurl, title, sampleSeasonID, samplePlaylistID, 0)))

	// Service Record: Arena
	//fmt.Println(string(halo.ServiceRecordArena(baseurl, title, samplePlayers, sampleSeasonID)))

	// Spartan Image
	//fmt.Println(string(halo.SpartanImage(baseurl, title, sampleGamertag, 0, "")))

	//	ioutil.WriteFile("tmp/EventsForMatch.txt", halo.EventsForMatch(baseurl, title, sampleArenaMatchID), 0755)
	//	ioutil.WriteFile("tmp/MatchesForPlayer.txt", halo.MatchesForPlayer(baseurl, title, sampleGamertag, "", 0, 0), 0755)
	//	ioutil.WriteFile("tmp/PlayerLeaderboard.txt", halo.PlayerLeaderboard(baseurl, title, sampleSeasonID, samplePlaylistID, 0), 0755)
	//	ioutil.WriteFile("tmp/CarnageReportArena.txt", halo.CarnageReportArena(baseurl, title, sampleArenaMatchID), 0755)
	//	ioutil.WriteFile("tmp/CarnageReportCampaign.txt", halo.CarnageReportCampaign(baseurl, title, sampleCampaignMatchID), 0755)
	//	ioutil.WriteFile("tmp/CarnageReportCustom.txt", halo.CarnageReportCustom(baseurl, title, sampleCustomMatchID), 0755)
	//	ioutil.WriteFile("tmp/CarnageReportWarzone.txt", halo.CarnageReportWarzone(baseurl, title, sampleWarzoneMatchID), 0755)
	//	fmt.Println("Sleeping")
	//	time.Sleep(10 * time.Second)
	//	ioutil.WriteFile("tmp/ServiceRecordArena.txt", halo.ServiceRecordArena(baseurl, title, samplePlayers, sampleSeasonID), 0755)
	//	ioutil.WriteFile("tmp/ServiceRecordCampaign.txt", halo.ServiceRecordCampaign(baseurl, title, samplePlayers), 0755)
	//	ioutil.WriteFile("tmp/ServiceRecordCustom.txt", halo.ServiceRecordCustom(baseurl, title, samplePlayers), 0755)
	//	ioutil.WriteFile("tmp/ServiceRecordWarzone.txt", halo.ServiceRecordWarzone(baseurl, title, samplePlayers), 0755)

}
