package halo

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (h *Halo) EventsForMatch(matchid string) (EventsForMatchStruct, error) {
	err := verifyValidID(matchid)
	var j EventsForMatchStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/matches/%s/events", h.baseurl, h.title, matchid))
	if err != nil {
		return EventsForMatchStruct{}, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return EventsForMatchStruct{}, fmt.Errorf(
			"EventsForMatch request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return EventsForMatchStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) MatchesForPlayer(player, modes string, start, count int) (MatchesForPlayerStruct, error) {
	var j MatchesForPlayerStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/players/%s/matches", h.baseurl, h.title, player))
	if err != nil {
		return MatchesForPlayerStruct{}, err
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
		return MatchesForPlayerStruct{}, fmt.Errorf(
			"MatchesForPlayer request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return MatchesForPlayerStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) PlayerLeaderboard(seasonid, playlistid string, count int) (PlayerLeaderboardStruct, error) {
	var j PlayerLeaderboardStruct
	err := verifyValidID(playlistid)
	err = verifyValidID(seasonid)
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/player-leaderboards/csr/%s/%s", h.baseurl, h.title, seasonid, playlistid))
	if err != nil {
		return PlayerLeaderboardStruct{}, err
	}
	q := url.Query()

	if count != 0 {
		q.Set("count", string(count))
	}
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return PlayerLeaderboardStruct{}, fmt.Errorf(
			"PlayerLeaderboard request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return PlayerLeaderboardStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) CarnageReportArena(matchid string) (CarnageReportArenaStruct, error) {
	var j CarnageReportArenaStruct
	err := verifyValidID(matchid)
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/arena/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		return CarnageReportArenaStruct{}, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return CarnageReportArenaStruct{}, fmt.Errorf(
			"CarnageReportArena request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return CarnageReportArenaStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) CarnageReportCampaign(matchid string) (CarnageReportCampaignStruct, error) {
	var j CarnageReportCampaignStruct
	err := verifyValidID(matchid)
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/campaign/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		return CarnageReportCampaignStruct{}, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return CarnageReportCampaignStruct{}, fmt.Errorf(
			"CarnageReportCampaign request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return CarnageReportCampaignStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) CarnageReportCustom(matchid string) (CarnageReportCustomStruct, error) {
	var j CarnageReportCustomStruct
	err := verifyValidID(matchid)
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/custom/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		return CarnageReportCustomStruct{}, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return CarnageReportCustomStruct{}, fmt.Errorf(
			"CarnageReportCustom request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return CarnageReportCustomStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) CarnageReportWarzone(matchid string) (CarnageReportWarzoneStruct, error) {
	var j CarnageReportWarzoneStruct
	err := verifyValidID(matchid)
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/warzone/matches/%s", h.baseurl, h.title, matchid))
	if err != nil {
		return CarnageReportWarzoneStruct{}, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return CarnageReportWarzoneStruct{}, fmt.Errorf(
			"CarnageReportWarzone request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return CarnageReportWarzoneStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) ServiceRecordArena(players, seasonid string) (ServiceRecordArenaStruct, error) {
	var j ServiceRecordArenaStruct
	err := verifyValidID(seasonid)
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/arena", h.baseurl, h.title))
	if err != nil {
		return ServiceRecordArenaStruct{}, err
	}
	q := url.Query()
	q.Set("players", players)
	if seasonid != "" {
		q.Set("seasonId", seasonid)
	}
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return ServiceRecordArenaStruct{}, fmt.Errorf(
			"ServiceRecordArena request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return ServiceRecordArenaStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) ServiceRecordCampaign(players string) (ServiceRecordCampaignStruct, error) {
	var j ServiceRecordCampaignStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/campaign", h.baseurl, h.title))
	if err != nil {
		return ServiceRecordCampaignStruct{}, err
	}
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return ServiceRecordCampaignStruct{}, fmt.Errorf(
			"ServiceRecordCampaign request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return ServiceRecordCampaignStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) ServiceRecordCustom(players string) (ServiceRecordCustomStruct, error) {
	var j ServiceRecordCustomStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/custom", h.baseurl, h.title))
	if err != nil {
		return ServiceRecordCustomStruct{}, err
	}
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return ServiceRecordCustomStruct{}, fmt.Errorf(
			"ServiceRecordCustom request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return ServiceRecordCustomStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) ServiceRecordWarzone(players string) (ServiceRecordWarzoneStruct, error) {
	var j ServiceRecordWarzoneStruct
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/servicerecords/warzone", h.baseurl, h.title))
	if err != nil {
		return ServiceRecordWarzoneStruct{}, err
	}
	q := url.Query()
	q.Set("players", players)
	url.RawQuery = q.Encode()
	jsonObject, err := h.sendRequest(url.String())
	if err != nil {
		return ServiceRecordWarzoneStruct{}, fmt.Errorf(
			"ServiceRecordWarzone request Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return ServiceRecordWarzoneStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}
