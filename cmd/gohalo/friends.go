package main

import (
	"encoding/json"
	"fmt"
	"log"
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
var badGameVariantID string = "9aaaaaaa-369a-4a37-97e3-432fa13035e1"    //Slayer
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

func structToString(s interface{}) (string, error) {
	json, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func totalGamesCompleted(gamertag string, gameTypes []string, h *halo.Halo) [4]int {
	var gamesCompleted = [4]int{}
	a, err := h.ServiceRecordArena(gamertag, "")
	if err != nil {
		fmt.Println("a")
		log.Fatal(err)
	}
	w, err := h.ServiceRecordWarzone(gamertag)
	if err != nil {
		fmt.Println("w")
		log.Fatal(err)
	}
	c, err := h.ServiceRecordCustom(gamertag)
	if err != nil {
		fmt.Println("c")
		log.Fatal(err)
	}
	ca, err := h.ServiceRecordCampaign(gamertag)
	if err != nil {
		fmt.Println("ca")
		log.Fatal(err)
	}
	gamesCompleted[0] = a.Results[0].Result.ArenaStats.TotalGamesCompleted
	gamesCompleted[1] = w.Results[0].Result.WarzoneStat.TotalGamesCompleted
	gamesCompleted[2] = c.Results[0].Result.CustomStats.TotalGamesCompleted
	gamesCompleted[3] = ca.Results[0].Result.CampaignStat.TotalGamesCompleted
	//	time.Sleep(4 * time.Second)
	return gamesCompleted
}

func main() {
	var gamertag = "motta13"
	gameTypes := []string{"arena", "warzone", "custom", "campaign"}
	matchIDs := [4][]string{}
	var ID string
	h := halo.NewHalo(baseurl, title, getAPIKey(), 200)

	gamesCompleted := totalGamesCompleted(gamertag, gameTypes, h)
	fmt.Println(gamesCompleted)
	for k := 0; k < len(gameTypes); k++ { // Loop through game types
		for i := 0; i < gamesCompleted[k]; i = i + 25 { // Loop through game history 25 at a time
			//	for i := 0; i < 200; i = i + 25 { // Loop through game history 25 at a time
			a1, err := h.MatchesForPlayer(gamertag, gameTypes[k], i, i+25)
			if err != nil {
				fmt.Println("a")
				log.Fatal(err)
			}
			for j := 0; j < len(a1.Results); j++ { // Loop through results
				ID = a1.Results[j].ID.MatchID
				matchIDs[k] = append(matchIDs[k], ID)
			}
			//			time.Sleep(1 * time.Second)
		}
	}
	//fmt.Printf("%+v", matchIDs)
}
