package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var print bool = true
var saveToFile bool = false
var importFromFile bool = true

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
	return gamesCompleted
}

func getAllMatchIDs(gamertag string, gameTypes []string, h *halo.Halo) [4][]string {
	matchIDs := [4][]string{}
	var ID string

	gamesCompleted := totalGamesCompleted(gamertag, gameTypes, h)
	//	fmt.Println("Games Completed: ", gamesCompleted)
	for k := 0; k < len(gameTypes); k++ { // Loop through game types
		for i := 0; i < gamesCompleted[k]; i = i + 25 { // Loop through game history 25 at a time
			//		for i := 0; i < 25; i = i + 25 { // Loop through game history 25 at a time
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
	return matchIDs
}

func getPlayersInMatch(ID, gameType string, h *halo.Halo) []string {
	playerList := []string{}
	if gameType == "arena" {
		a, err := h.CarnageReportArena(ID)
		if err != nil {
			log.Fatal(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID = a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	if gameType == "warzone" {
		a, err := h.CarnageReportWarzone(ID)
		if err != nil {
			log.Fatal(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID = a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	if gameType == "custom" {
		a, err := h.CarnageReportCustom(ID)
		if err != nil {
			log.Fatal(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID = a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	if gameType == "campaign" {
		a, err := h.CarnageReportCampaign(ID)
		if err != nil {
			log.Fatal(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID = a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	return playerList
}

func printNicely(playerMap []map[string][]string, gameTypes []string, gamertag, gamertag2 string) {
	gamesTogether := 0
	// Print the results nicely
	if print == true { // turn on printing easily
		for x := 0; x < len(playerMap); x++ {
			fmt.Printf("\n%v", gameTypes[x]) // Print Gametype
			for k, _ := range playerMap[x] {
				fmt.Printf("\n\t%v", k) // Print Key (Match ID)
				for _, v := range playerMap[x][k] {
					fmt.Printf("\n\t\t%v", v) // Print Values (Gamertag)
					if v == gamertag2 {       // count how many games gamertag2 has played with gamertag
						gamesTogether++
					}
				}
			}
		}
		fmt.Printf("\n\n%v played %v with Thomas\n\n", gamertag, gamesTogether)
	}
}

func writeToFile(playerMap []map[string][]string, file string) {

	if saveToFile == true {
		out, err := json.Marshal(playerMap)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(file, out, 0644)
		if err != nil {
			panic(err)
		}
	}

}

func openFile(playerMap []map[string][]string, file string) []map[string][]string {
	out, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(out, &playerMap)
	if err != nil {
		panic(err)
	}
	return playerMap
}

func main() {
	h := halo.NewHalo(baseurl, title, getAPIKey(), 200)
	gamertag := "smoke721"
	gamertag2 := "motta13"
	gameTypes := []string{"arena", "warzone", "custom", "campaign"}
	playerMap := make([]map[string][]string, 4)

	if importFromFile == false {
		matchIDs := getAllMatchIDs(gamertag, gameTypes, h) // Get Match IDs

		// Fill Match IDs and player lists into Playermap
		for i := 0; i < len(playerMap); i++ { // Loop through game types
			playerMap[i] = map[string][]string{}
			for j := 0; j < len(matchIDs[i]); j++ { // Loop through MatchIDs in each game type
				playerMap[i][matchIDs[i][j]] = getPlayersInMatch(matchIDs[i][j], gameTypes[i], h) // Add list of players to playerMap
			}
		}
	} else {
		playerMap = openFile(playerMap, "/tmp/smoke721")
	}
	printNicely(playerMap, gameTypes, gamertag, gamertag2) // Print Nicely
	writeToFile(playerMap, "/tmp/smoke721")

}
