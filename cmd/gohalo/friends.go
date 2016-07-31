package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/tbenz9/go-halo5-api/halo"
)

var baseurl string = "https://www.haloapi.com"
var title string = "h5"

var print bool = true
var saveToFile bool = false
var importFromFile bool = false

type gameData struct {
	ID           string
	Type         string
	Completed    string
	Players      []string
	Rank         int
	Result       int
	MapID        string
	BaseMap      string
	TeamID       int
	TotalAssists int
	TotalDeaths  int
	TotalKills   int
}

func getAPIKey() string {
	apikey := os.Getenv("HALO_API_KEY")
	if len(apikey) != 32 {
		fmt.Println("Invalid API Key")
	}
	return apikey
}

// I have to get my total games for each gameType so I know how many times
// to loop in the getAllMatchIDs function.
func totalGamesCompleted(gamertag string, gameTypes []string, h *halo.Halo) [4]int {
	//var wg sync.WaitGroup
	//wg.Add(4)
	var gamesCompleted = [4]int{}
	//go func() {
	a, err := h.ServiceRecordArena(gamertag, "")
	if err != nil {
		fmt.Println("a")
		log.Fatal(err)
	}
	gamesCompleted[0] = a.Results[0].Result.ArenaStats.TotalGamesCompleted
	//	wg.Done()
	//	}()
	//	go func() {
	w, err := h.ServiceRecordWarzone(gamertag)
	if err != nil {
		fmt.Println("w")
		log.Fatal(err)
	}
	gamesCompleted[1] = w.Results[0].Result.WarzoneStat.TotalGamesCompleted
	//		wg.Done()
	//	}()
	//	go func() {
	c, err := h.ServiceRecordCustom(gamertag)
	if err != nil {
		fmt.Println("c")
		log.Fatal(err)
	}
	gamesCompleted[2] = c.Results[0].Result.CustomStats.TotalGamesCompleted
	//		wg.Done()
	//	}()
	//	go func() {
	ca, err := h.ServiceRecordCampaign(gamertag)
	if err != nil {
		fmt.Println("ca")
		log.Fatal(err)
	}
	gamesCompleted[3] = ca.Results[0].Result.CampaignStat.TotalGamesCompleted
	//		wg.Done()
	//	}()
	//	fmt.Println("waiting in TotalGamesCompleted")
	//	wg.Wait()
	return gamesCompleted
}

func getAllMatchIDs(gamertag string, gameTypes []string, h *halo.Halo, allGames *[]*gameData) {
	//var wg sync.WaitGroup
	//var mutex = &sync.Mutex{}
	gamesCompleted := totalGamesCompleted(gamertag, gameTypes, h)
	fmt.Println("Games Completed: ", gamesCompleted)
	for k := range gameTypes { // Loop through game types
		//for i := 0; i < gamesCompleted[k]; i = i + 25 { // Loop through game history 25 at a time
		for i := 0; i < 25; i = i + 25 { // For Testing: Only do the first 25 games of each type.
			//		wg.Add(1)
			//		go func() {
			//			defer wg.Done()
			a1, err := h.MatchesForPlayer(gamertag, gameTypes[k], i, i+25)
			if err != nil {
				fmt.Println("a")
				log.Fatal(err)
			}
			for j := 0; j < len(a1.Results); j++ { // Loop through results
				game := new(gameData)
				game.ID = a1.Results[j].ID.MatchID
				game.Type = gameTypes[k]
				game.Completed = a1.Results[j].MatchCompletedDate.ISO8601Date
				game.BaseMap = a1.Results[j].MapID
				game.Rank = a1.Results[j].Players[0].Rank
				game.Result = a1.Results[j].Players[0].Result
				game.TeamID = a1.Results[j].Players[0].TeamID
				game.TotalAssists = a1.Results[j].Players[0].TotalAssists
				game.TotalDeaths = a1.Results[j].Players[0].TotalDeaths
				game.TotalKills = a1.Results[j].Players[0].TotalKills
				//			mutex.Lock()
				*allGames = append(*allGames, game)
				//			mutex.Unlock()
			}

			//		}()
			fmt.Println("Waiting in getAllMatchIDs")
			//		wg.Wait()
		}
	}
}

func getPlayersInMatch(h *halo.Halo, game *gameData) {
	playerList := []string{}
	if game.Type == "arena" {
		a, err := h.CarnageReportArena(game.ID)
		if err != nil {
			fmt.Println(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID := a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	if game.Type == "warzone" {
		a, err := h.CarnageReportWarzone(game.ID)
		if err != nil {
			fmt.Println(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID := a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	if game.Type == "custom" {
		a, err := h.CarnageReportCustom(game.ID)
		if err != nil {
			fmt.Println(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID := a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	if game.Type == "campaign" {
		a, err := h.CarnageReportCampaign(game.ID)
		if err != nil {
			fmt.Println(err)
		}
		for j := 0; j < len(a.PlayerStats); j++ { // Loop through results
			ID := a.PlayerStats[j].Player.Gamertag
			playerList = append(playerList, ID)
		}
	}
	game.Players = playerList
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
					fmt.Printf("\n\t\t%v", v)                             // Print Values (Gamertag)
					if strings.ToLower(v) == strings.ToLower(gamertag2) { // count how many games gamertag2 has played with gamertag
						gamesTogether++
					}
				}
			}
		}
		fmt.Printf("\n\n%v played %v with %v\n\n", gamertag, gamesTogether, gamertag2)
	}
}

func writeToFile(allGames *[]*gameData, file string) {

	if saveToFile == true {
		out, err := json.Marshal(allGames)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(file, out, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func openFile(allGames *[]*gameData, file string) []*gameData {
	tmp := []*gameData{}
	out, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(out, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func main() {
	//	var wg sync.WaitGroup
	h := halo.NewHalo(baseurl, title, getAPIKey(), 200)
	gamertag := "motta13"
	//gamertag2 := "casket velocity"
	gameTypes := []string{"arena", "warzone", "custom", "campaign"}
	allGames := []*gameData{}

	if importFromFile == true {
		allGames = openFile(&allGames, "/tmp/mt")
		for i := range allGames {
			fmt.Printf("\n%+v", allGames[i])
		}
		fmt.Printf("%+v", allGames)
	} else {
		getAllMatchIDs(gamertag, gameTypes, h, &allGames) // Get Match IDs
		for i := range allGames {
			//			wg.Add(1)
			//			go func() {
			//			defer wg.Done()
			getPlayersInMatch(h, allGames[i])
			fmt.Printf("\nGame: %+v", allGames[i])
			//			}()
		}
		//		wg.Wait()
	}

	if saveToFile == true {
		writeToFile(&allGames, "/tmp/mt")
	}
}
