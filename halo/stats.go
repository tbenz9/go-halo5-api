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

func (h *Halo) PlayerLeaderboard(seasonid, playlistid string, count int) PlayerLeaderboardStruct {
	var j PlayerLeaderboardStruct
	verifyValidID(playlistid, "Playlist ID")
	verifyValidID(seasonid, "Season ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/player-leaderboards/csr/%s/%s", h.baseurl, h.title, seasonid, playlistid))
	if err != nil {
		log.Fatal("PlayerLeaderboard URL Parse Failed: ", err)
	}
	q := url.Query()

	if count != 0 {
		q.Set("count", string(count))
	}
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("PlayerLeaderboard Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) CarnageReportArena(matchid string) CarnageReportArenaStruct {
	var j CarnageReportArenaStruct
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/arena/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		log.Fatal("CarnageReportArena URL Parse Failed: ", err)
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("CarnageReportArena Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) CarnageReportCampaign(matchid string) CarnageReportCampaignStruct {
	var j CarnageReportCampaignStruct
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/campaign/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		log.Fatal("CarnageReportCampaign URL Parse Failed: ", err)
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("CarnageReportCampaign Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) CarnageReportCustom(matchid string) CarnageReportCustomStruct {
	var j CarnageReportCustomStruct
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/custom/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		log.Fatal("CarnageReportCustom URL Parse Failed: ", err)
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("CarnageReportCustom Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) CarnageReportWarzone(matchid string) CarnageReportWarzoneStruct {
	var j CarnageReportWarzoneStruct
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/warzone/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		log.Fatal("CarnageReportWarzone URL Parse Failed: ", err)
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("CarnageReportWarzone Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) ServiceRecordArena(players, seasonid string) ServiceRecordArenaStruct {
	var j ServiceRecordArenaStruct
	verifyValidID(seasonid, "Season ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/arena", h.baseurl, h.title))
	if err != nil {
		log.Fatal("CarnageReportWarzone URL Parse Failed: ", err)
	}
	q := url.Query()
	q.Set("players", players)
	if seasonid != "" {
		q.Set("seasonId", seasonid)
	}
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("ServiceRecordArena Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) ServiceRecordCampaign(players string) ServiceRecordCampaignStruct {
	var j ServiceRecordCampaignStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/campaign", h.baseurl, h.title))
	if err != nil {
		log.Fatal("ServiceRecordCampaign URL Parse Failed: ", err)
	}
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("ServiceRecordCampaign Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) ServiceRecordCustom(players string) ServiceRecordCustomStruct {
	var j ServiceRecordCustomStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/custom", h.baseurl, h.title))
	if err != nil {
		log.Fatal("ServiceRecordCampaign URL Parse Failed: ", err)
	}
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("ServiceRecordCustom Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) ServiceRecordWarzone(players string) ServiceRecordWarzoneStruct {
	var j ServiceRecordWarzoneStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/warzone", h.baseurl, h.title))
	if err != nil {
		log.Fatal("ServiceRecordCampaign URL Parse Failed: ", err)
	}
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		log.Fatal("ServiceRecordWarzone Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}
