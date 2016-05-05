package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var baseurl string = "https://www.haloapi.com"
var title string = "h5"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func events_for_match(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/matches/%s/events", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func matches_for_player(baseurl, title, player, modes string, start, count int) string {
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
	response := send_request(url.String())
	return response
}

func player_leaderboard(baseurl, title, seasonid, playlistid string, count int) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/player-leaderboards/csr/%s/%s", baseurl, title, seasonid, playlistid))
	checkErr(err)
	q := url.Query()

	if count != 0 {
		q.Set("count", string(count))
	}
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func carnage_report_arena(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/arena/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func carnage_report_campaign(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/campaign/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func carnage_report_custom(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/custom/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func carnage_report_warzone(baseurl, title, matchid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/warzone/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func service_record_arena(baseurl, title, players, seasonid string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/arena/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	if seasonid != "" {
		q.Set("seasonId", seasonid)
	}
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func service_record_campaign(baseurl, title, players string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/campaign/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func service_record_custom(baseurl, title, players string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/custom/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func service_record_warzone(baseurl, title, players string) string {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/warzone/", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

// Begin profile section
func emblem_image(baseurl, title, player string, size int) string {
	url, err := url.Parse(fmt.Sprintf("%s/profile/%s/profiles/%s/", baseurl, title, player))
	checkErr(err)
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", string(size))
	}
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func spartan_image(baseurl, title, player string, size int, crop string) string {
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
	response := send_request(url.String())
	return response
}

// Begin metadata section
func campaign_missions(baseurl, title string) string {
	return metadata_request(baseurl, title, "campaign-missions", "")
}

func commendations(baseurl, title string) string {
	return metadata_request(baseurl, title, "commendations", "")
}

func csr_designations(baseurl, title string) string {
	return metadata_request(baseurl, title, "csr-designations", "")
}

func enemies(baseurl, title string) string {
	return metadata_request(baseurl, title, "enemies", "")
}

func flexible_stats(baseurl, title string) string {
	return metadata_request(baseurl, title, "flexible-stats", "")
}

func game_base_variants(baseurl, title string) string {
	return metadata_request(baseurl, title, "game-base-variants", "")
}

func game_variants(baseurl, title, id string) string {
	return metadata_request(baseurl, title, "game-variants", id)
}

func impulses(baseurl, title string) string {
	return metadata_request(baseurl, title, "impulses", "")
}

func map_variants(baseurl, title, id string) string {
	return metadata_request(baseurl, title, "map-variants", id)
}

func maps(baseurl, title string) string {
	return metadata_request(baseurl, title, "maps", "")
}

func medals(baseurl, title string) string {
	return metadata_request(baseurl, title, "medals", "")
}

func playlists(baseurl, title string) string {
	return metadata_request(baseurl, title, "playlists", "")
}

func requisition_packs(baseurl, title, id string) string {
	return metadata_request(baseurl, title, "requisition-packs", id)
}

func requisitions(baseurl, title, id string) string {
	return metadata_request(baseurl, title, "requisitions", id)
}

func seasons(baseurl, title string) string {
	return metadata_request(baseurl, title, "seasons", "")
}

func skulls(baseurl, title string) string {
	return metadata_request(baseurl, title, "skulls", "")
}

func spartan_ranks(baseurl, title string) string {
	return metadata_request(baseurl, title, "spartan-ranks", "")
}

func team_colors(baseurl, title string) string {
	return metadata_request(baseurl, title, "team-colors", "")
}

func vehicles(baseurl, title string) string {
	return metadata_request(baseurl, title, "vehicles", "")
}

func weapons(baseurl, title string) string {
	return metadata_request(baseurl, title, "weapons", "")
}

func metadata_request(baseurl, title, datatype, id string) string {
	url, err := url.Parse(fmt.Sprintf("%s/metadata/%s/metadata/%s/%s", baseurl, title, datatype, id))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := send_request(url.String())
	return response
}

func send_request(url string) string {
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
