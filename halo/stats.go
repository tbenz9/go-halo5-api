package halo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

func (h *Halo) EventsForMatch(matchid string) EventsForMatchStruct {
	verifyValidID(matchid, "Match ID")
	var j EventsForMatchStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/matches/%s/events", h.baseurl, h.title, matchid))
	if err != nil {
		log.Fatal("EventsForMatch URL Parse Failed: ", err)
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("EventsForMatch Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) MatchesForPlayer(player, modes string, start, count int) MatchesForPlayerStruct {
	var j MatchesForPlayerStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/players/%s/matches", h.baseurl, h.title, player))
	if err != nil {
		log.Fatal("MatchesForPlayer URL Parse Failed: ", err)
	}
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
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("MatchesForPlayer Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func PlayerLeaderboard(baseurl, title, seasonid, playlistid string, count int) []byte {
	verifyValidID(playlistid, "Playlist ID")
	verifyValidID(seasonid, "Season ID")
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

func CarnageReportArena(baseurl, title, matchid string) []byte {
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/arena/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func CarnageReportCampaign(baseurl, title, matchid string) []byte {
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/campaign/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func CarnageReportCustom(baseurl, title, matchid string) []byte {
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/custom/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func CarnageReportWarzone(baseurl, title, matchid string) []byte {
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/warzone/matches/%s", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func ServiceRecordArena(baseurl, title, players, seasonid string) []byte {
	verifyValidID(seasonid, "Season ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/arena", baseurl, title))
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

func ServiceRecordCampaign(baseurl, title, players string) []byte {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/campaign", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func ServiceRecordCustom(baseurl, title, players string) []byte {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/custom", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func ServiceRecordWarzone(baseurl, title, players string) []byte {
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/warzone", baseurl, title))
	checkErr(err)
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}
