package halo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func EventsForMatch(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/matches/%s/events", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func MatchesForPlayer(baseurl, title, player, modes string, start, count int) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/players/%s/matches", baseurl, title, player))
	checkErr(err)
	q := url.Query()

	if modes != "" {
		q.Set("modes", modes)
	}
	if start != 0 {
		q.Set("start", string(start))
	}
	if count != 0 {
		q.Set("count", string(count))
	}
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func PlayerLeaderboard(baseurl, title, seasonid, playlistid string, count int) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/player-leaderboards/csr/%s/%s", baseurl, title, seasonid, playlistid))
	checkErr(err)
	q := url.Query()

	if count != 0 {
		q.Set("count", string(count))
	}
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func CarnageReportArena(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/arena/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func CarnageReportCampaign(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/campaign/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func CarnageReportCustom(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/custom/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func CarnageReportWarzone(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/warzone/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func ServiceRecordArena(baseurl, title, players, seasonid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/arena/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	if seasonid != "" {
		q.Set("seasonId", seasonid)
	}
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func ServiceRecordCampaign(baseurl, title, players string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/campaign/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func ServiceRecordCustom(baseurl, title, players string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/custom/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func ServiceRecordWarzone(baseurl, title, players string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/warzone/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

// Begin profile section
func EmblemImage(baseurl, title, player string, size int) string {
	url, err := url.Parse(fmt.Sprintf("%s/profile/%s/profiles/%s/", baseurl, title, player))
	checkErr(err)
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", string(size))
	}
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func SpartanImage(baseurl, title, player string, size int, crop string) string {
	url, err := url.Parse(fmt.Sprintf("%s/profile/%s/profiles/%s/spartan/", baseurl, title, player))
	checkErr(err)
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", string(size))
	}
	if (strings.ToLower(crop) == "full") || (strings.ToLower(crop) == "portrait") {
		q.Set("crop", crop)
	}
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

// Begin metadata section
func CampaignMissions(baseurl, title string) string {
	return metadataRequest(baseurl, title, "campaign-missions", "")
}

func Commendations(baseurl, title string) string {
	return metadataRequest(baseurl, title, "commendations", "")
}

func CsrDesignations(baseurl, title string) string {
	return metadataRequest(baseurl, title, "csr-designations", "")
}

func Enemies(baseurl, title string) string {
	return metadataRequest(baseurl, title, "enemies", "")
}

func FlexibleStats(baseurl, title string) string {
	return metadataRequest(baseurl, title, "flexible-stats", "")
}

func GameBaseVariants(baseurl, title string) string {
	return metadataRequest(baseurl, title, "game-base-variants", "")
}

func GameVariants(baseurl, title, id string) string {
	return metadataRequest(baseurl, title, "game-variants", id)
}

func Impulses(baseurl, title string) string {
	return metadataRequest(baseurl, title, "impulses", "")
}

func MapVariants(baseurl, title, id string) string {
	return metadataRequest(baseurl, title, "map-variants", id)
}

func Maps(baseurl, title string) string {
	return metadataRequest(baseurl, title, "maps", "")
}

func Medals(baseurl, title string) string {
	return metadataRequest(baseurl, title, "medals", "")
}

func Playlists(baseurl, title string) string {
	return metadataRequest(baseurl, title, "playlists", "")
}

func RequisitionPacks(baseurl, title, id string) string {
	return metadataRequest(baseurl, title, "requisition-packs", id)
}

func Requisitions(baseurl, title, id string) string {
	return metadataRequest(baseurl, title, "requisitions", id)
}

func Seasons(baseurl, title string) string {
	return metadataRequest(baseurl, title, "seasons", "")
}

func Skulls(baseurl, title string) string {
	return metadataRequest(baseurl, title, "skulls", "")
}

func SpartanRanks(baseurl, title string) string {
	return metadataRequest(baseurl, title, "spartan-ranks", "")
}

func TeamColors(baseurl, title string) string {
	return metadataRequest(baseurl, title, "team-colors", "")
}

func Vehicles(baseurl, title string) string {
	return metadataRequest(baseurl, title, "vehicles", "")
}

func Weapons(baseurl, title string) string {
	return metadataRequest(baseurl, title, "weapons", "")
}

func metadataRequest(baseurl, title, datatype, id string) string {
	url, err := url.Parse(fmt.Sprintf("%s/metadata/%s/metadata/%s/%s", baseurl, title, datatype, id))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func sendRequest(url string) string {
	request, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	apikey := os.Getenv("HALO_API_KEY")
	request.Header.Set("Ocp-Apim-Subscription-Key", apikey)

	response, err := http.DefaultClient.Do(request)
	checkErr(err)
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	return string(contents)
}

// Sample values for testing
var sampleGamertag string = "motta13"
var sampleMode string = "warzone"
var sampleMatchID string = "c35a35f8-f450-4836-a4c2-65100a7acb79"
var sampleSeasonID string = "b46c2095-4ca6-4f4b-a565-4702d7cfe586"   //February 2016 Season
var samplePlaylistID string = "2323b76a-db98-4e03-aa37-e171cfbdd1a4" //SWAT gametype 2016 Season
var samplePlayers string = "motta13,smoke721"

func main() {
	// Uncomment any of the below for sample output.
	//fmt.Println(matches_for_player(baseurl, title, SampleGamertag, "warzone", 0, 0))
	//fmt.Println(events_for_match(baseurl, title, sampleMatchID))
	//fmt.Println(player_leaderboard(baseurl, title, sampleSeasonID, samplePlaylistID, 0))
	//fmt.Println(service_record_arena(baseurl, title, samplePlayers, sampleSeasonID))
	//fmt.Println(vehicles(baseurl, title))
	//fmt.Println(spartan_image(baseurl, title, gamertag, 0, ""))
}
