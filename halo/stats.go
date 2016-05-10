package halo

import (
	"fmt"
	"net/url"
	"time"
)

// Stats Structs
type EventsForMatchStruct struct {
	GameEvents []struct {
		EventName      string `json:"EventName"`
		RoundIndex     int    `json:"RoundIndex"`
		TimeSinceStart string `json:"TimeSinceStart"`
	} `json:"GameEvents"`
	IsCompleteSetOfEvents bool `json:"IsCompleteSetOfEvents"`
	Links                 struct {
		StatsMatchDetails struct {
			AcknowledgementTypeID                    int         `json:"AcknowledgementTypeId"`
			AuthenticationLifetimeExtensionSupported bool        `json:"AuthenticationLifetimeExtensionSupported"`
			AuthorityID                              string      `json:"AuthorityId"`
			Path                                     string      `json:"Path"`
			QueryString                              interface{} `json:"QueryString"`
			RetryPolicyID                            string      `json:"RetryPolicyId"`
			TopicName                                string      `json:"TopicName"`
		} `json:"StatsMatchDetails"`
		UgcFilmManifest struct {
			AcknowledgementTypeID                    int    `json:"AcknowledgementTypeId"`
			AuthenticationLifetimeExtensionSupported bool   `json:"AuthenticationLifetimeExtensionSupported"`
			AuthorityID                              string `json:"AuthorityId"`
			Path                                     string `json:"Path"`
			QueryString                              string `json:"QueryString"`
			RetryPolicyID                            string `json:"RetryPolicyId"`
			TopicName                                string `json:"TopicName"`
		} `json:"UgcFilmManifest"`
	} `json:"Links"`
}

type MatchesForPlayerStruct struct {
	Start       int `json:"Start"`
	Count       int `json:"Count"`
	ResultCount int `json:"ResultCount"`
	Results     []struct {
		Links struct {
			StatsMatchDetails struct {
				AuthorityID                              string      `json:"AuthorityId"`
				Path                                     string      `json:"Path"`
				QueryString                              interface{} `json:"QueryString"`
				RetryPolicyID                            string      `json:"RetryPolicyId"`
				TopicName                                string      `json:"TopicName"`
				AcknowledgementTypeID                    int         `json:"AcknowledgementTypeId"`
				AuthenticationLifetimeExtensionSupported bool        `json:"AuthenticationLifetimeExtensionSupported"`
			} `json:"StatsMatchDetails"`
			UgcFilmManifest struct {
				AuthorityID                              string `json:"AuthorityId"`
				Path                                     string `json:"Path"`
				QueryString                              string `json:"QueryString"`
				RetryPolicyID                            string `json:"RetryPolicyId"`
				TopicName                                string `json:"TopicName"`
				AcknowledgementTypeID                    int    `json:"AcknowledgementTypeId"`
				AuthenticationLifetimeExtensionSupported bool   `json:"AuthenticationLifetimeExtensionSupported"`
			} `json:"UgcFilmManifest"`
		} `json:"Links"`
		ID struct {
			MatchID  string `json:"MatchId"`
			GameMode int    `json:"GameMode"`
		} `json:"Id"`
		HopperID   string `json:"HopperId"`
		MapID      string `json:"MapId"`
		MapVariant struct {
			ResourceType int    `json:"ResourceType"`
			ResourceID   string `json:"ResourceId"`
			OwnerType    int    `json:"OwnerType"`
			Owner        string `json:"Owner"`
		} `json:"MapVariant"`
		GameBaseVariantID string `json:"GameBaseVariantId"`
		GameVariant       struct {
			ResourceType int    `json:"ResourceType"`
			ResourceID   string `json:"ResourceId"`
			OwnerType    int    `json:"OwnerType"`
			Owner        string `json:"Owner"`
		} `json:"GameVariant"`
		MatchDuration      string `json:"MatchDuration"`
		MatchCompletedDate struct {
			ISO8601Date time.Time `json:"ISO8601Date"`
		} `json:"MatchCompletedDate"`
		Teams []struct {
			ID    int `json:"Id"`
			Score int `json:"Score"`
			Rank  int `json:"Rank"`
		} `json:"Teams"`
		Players []struct {
			Player struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"Player"`
			TeamID           int         `json:"TeamId"`
			Rank             int         `json:"Rank"`
			Result           int         `json:"Result"`
			TotalKills       int         `json:"TotalKills"`
			TotalDeaths      int         `json:"TotalDeaths"`
			TotalAssists     int         `json:"TotalAssists"`
			PreMatchRatings  interface{} `json:"PreMatchRatings"`
			PostMatchRatings interface{} `json:"PostMatchRatings"`
		} `json:"Players"`
		IsTeamGame                 bool        `json:"IsTeamGame"`
		SeasonID                   interface{} `json:"SeasonId"`
		MatchCompletedDateFidelity int         `json:"MatchCompletedDateFidelity"`
	} `json:"Results"`
	Links struct {
		Self struct {
			AuthorityID                              string `json:"AuthorityId"`
			Path                                     string `json:"Path"`
			QueryString                              string `json:"QueryString"`
			RetryPolicyID                            string `json:"RetryPolicyId"`
			TopicName                                string `json:"TopicName"`
			AcknowledgementTypeID                    int    `json:"AcknowledgementTypeId"`
			AuthenticationLifetimeExtensionSupported bool   `json:"AuthenticationLifetimeExtensionSupported"`
		} `json:"Self"`
	} `json:"Links"`
}

func EventsForMatch(baseurl, title, matchid string) []byte {
	verifyValidID(matchid, "Match ID")
	url, err := url.Parse(fmt.Sprintf("%s/stats/%s/matches/%s/events", baseurl, title, matchid))
	checkErr(err)
	q := url.Query()
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func MatchesForPlayer(baseurl, title, player, modes string, start, count int) []byte {
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
