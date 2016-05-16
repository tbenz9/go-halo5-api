package halo

import "encoding/json"

type CarnageReportArenaStruct struct {
	PlayerStats []struct {
		XpInfo struct {
			PrevSpartanRank              int     `json:"PrevSpartanRank"`
			SpartanRank                  int     `json:"SpartanRank"`
			PrevTotalXP                  int     `json:"PrevTotalXP"`
			TotalXP                      int     `json:"TotalXP"`
			SpartanRankMatchXPScalar     float64 `json:"SpartanRankMatchXPScalar"`
			PlayerTimePerformanceXPAward int     `json:"PlayerTimePerformanceXPAward"`
			PerformanceXP                int     `json:"PerformanceXP"`
			PlayerRankXPAward            int     `json:"PlayerRankXPAward"`
			BoostAmount                  int     `json:"BoostAmount"`
			MatchSpeedWinAmount          int     `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount    int     `json:"ObjectivesCompletedAmount"`
		} `json:"XpInfo"`
		PreviousCsr struct {
			Tier              int         `json:"Tier"`
			DesignationID     int         `json:"DesignationId"`
			Csr               int         `json:"Csr"`
			PercentToNextTier int         `json:"PercentToNextTier"`
			Rank              interface{} `json:"Rank"`
		} `json:"PreviousCsr"`
		CurrentCsr struct {
			Tier              int         `json:"Tier"`
			DesignationID     int         `json:"DesignationId"`
			Csr               int         `json:"Csr"`
			PercentToNextTier int         `json:"PercentToNextTier"`
			Rank              interface{} `json:"Rank"`
		} `json:"CurrentCsr"`
		MeasurementMatchesLeft int           `json:"MeasurementMatchesLeft"`
		RewardSets             []interface{} `json:"RewardSets"`
		KilledOpponentDetails  []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledOpponentDetails"`
		KilledByOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledByOpponentDetails"`
		FlexibleStats struct {
			MedalStatCounts   []interface{} `json:"MedalStatCounts"`
			ImpulseStatCounts []struct {
				ID    string `json:"Id"`
				Count int    `json:"Count"`
			} `json:"ImpulseStatCounts"`
			MedalTimelapses   []interface{} `json:"MedalTimelapses"`
			ImpulseTimelapses []struct {
				ID        string `json:"Id"`
				Timelapse string `json:"Timelapse"`
			} `json:"ImpulseTimelapses"`
		} `json:"FlexibleStats"`
		CreditsEarned struct {
			Result                    int     `json:"Result"`
			TotalCreditsEarned        int     `json:"TotalCreditsEarned"`
			SpartanRankModifier       float64 `json:"SpartanRankModifier"`
			PlayerRankAmount          int     `json:"PlayerRankAmount"`
			TimePlayedAmount          float64 `json:"TimePlayedAmount"`
			BoostAmount               int     `json:"BoostAmount"`
			MatchSpeedWinAmount       int     `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount int     `json:"ObjectivesCompletedAmount"`
		} `json:"CreditsEarned"`
		MetaCommendationDeltas        []interface{} `json:"MetaCommendationDeltas"`
		ProgressiveCommendationDeltas []struct {
			ID               string `json:"Id"`
			PreviousProgress int    `json:"PreviousProgress"`
			Progress         int    `json:"Progress"`
		} `json:"ProgressiveCommendationDeltas"`
		BoostInfo struct {
			DefinitionID string `json:"DefinitionId"`
			CardConsumed bool   `json:"CardConsumed"`
		} `json:"BoostInfo"`
		Player struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		TeamID              int         `json:"TeamId"`
		Rank                int         `json:"Rank"`
		DNF                 bool        `json:"DNF"`
		AvgLifeTimeOfPlayer string      `json:"AvgLifeTimeOfPlayer"`
		PreMatchRatings     interface{} `json:"PreMatchRatings"`
		PostMatchRatings    interface{} `json:"PostMatchRatings"`
		PlayerScore         int         `json:"PlayerScore"`
		TotalKills          int         `json:"TotalKills"`
		TotalHeadshots      int         `json:"TotalHeadshots"`
		TotalWeaponDamage   float64     `json:"TotalWeaponDamage"`
		TotalShotsFired     int         `json:"TotalShotsFired"`
		TotalShotsLanded    int         `json:"TotalShotsLanded"`
		WeaponWithMostKills struct {
			WeaponID struct {
				StockID     int           `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponWithMostKills"`
		TotalMeleeKills                int     `json:"TotalMeleeKills"`
		TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
		TotalAssassinations            int     `json:"TotalAssassinations"`
		TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
		TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
		TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
		TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
		TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
		TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
		TotalDeaths                    int     `json:"TotalDeaths"`
		TotalAssists                   int     `json:"TotalAssists"`
		TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
		TotalGamesWon                  int     `json:"TotalGamesWon"`
		TotalGamesLost                 int     `json:"TotalGamesLost"`
		TotalGamesTied                 int     `json:"TotalGamesTied"`
		TotalTimePlayed                string  `json:"TotalTimePlayed"`
		TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
		MedalAwards                    []struct {
			MedalID int64 `json:"MedalId"`
			Count   int   `json:"Count"`
		} `json:"MedalAwards"`
		DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills             []interface{} `json:"EnemyKills"`
		WeaponStats            []struct {
			WeaponID struct {
				StockID     int           `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponStats"`
		Impulses          []interface{} `json:"Impulses"`
		TotalSpartanKills int           `json:"TotalSpartanKills"`
		FastestMatchWin   interface{}   `json:"FastestMatchWin"`
	} `json:"PlayerStats"`
	TeamStats []struct {
		TeamID     int `json:"TeamId"`
		Score      int `json:"Score"`
		Rank       int `json:"Rank"`
		RoundStats []struct {
			RoundNumber int `json:"RoundNumber"`
			Rank        int `json:"Rank"`
			Score       int `json:"Score"`
		} `json:"RoundStats"`
	} `json:"TeamStats"`
	IsMatchOver           bool   `json:"IsMatchOver"`
	TotalDuration         string `json:"TotalDuration"`
	MapVariantID          string `json:"MapVariantId"`
	GameVariantID         string `json:"GameVariantId"`
	PlaylistID            string `json:"PlaylistId"`
	MapID                 string `json:"MapId"`
	GameBaseVariantID     string `json:"GameBaseVariantId"`
	IsTeamGame            bool   `json:"IsTeamGame"`
	SeasonID              string `json:"SeasonId"`
	GameVariantResourceID struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"GameVariantResourceId"`
	MapVariantResourceID struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"MapVariantResourceId"`
}

type CarnageReportCampaignStruct struct {
	PlayerStats []struct {
		BiggestKillScore int `json:"BiggestKillScore"`
		FlexibleStats    struct {
			MedalStatCounts   []interface{} `json:"MedalStatCounts"`
			ImpulseStatCounts []struct {
				ID    string `json:"Id"`
				Count int    `json:"Count"`
			} `json:"ImpulseStatCounts"`
			MedalTimelapses   []interface{} `json:"MedalTimelapses"`
			ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
		} `json:"FlexibleStats"`
		Score          int         `json:"Score"`
		CharacterIndex interface{} `json:"CharacterIndex"`
		Player         struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		TeamID              int         `json:"TeamId"`
		Rank                int         `json:"Rank"`
		DNF                 bool        `json:"DNF"`
		AvgLifeTimeOfPlayer string      `json:"AvgLifeTimeOfPlayer"`
		PreMatchRatings     interface{} `json:"PreMatchRatings"`
		PostMatchRatings    interface{} `json:"PostMatchRatings"`
		PlayerScore         interface{} `json:"PlayerScore"`
		TotalKills          int         `json:"TotalKills"`
		TotalHeadshots      int         `json:"TotalHeadshots"`
		TotalWeaponDamage   float64     `json:"TotalWeaponDamage"`
		TotalShotsFired     int         `json:"TotalShotsFired"`
		TotalShotsLanded    int         `json:"TotalShotsLanded"`
		WeaponWithMostKills struct {
			WeaponID struct {
				StockID     int           `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponWithMostKills"`
		TotalMeleeKills                int           `json:"TotalMeleeKills"`
		TotalMeleeDamage               float64       `json:"TotalMeleeDamage"`
		TotalAssassinations            int           `json:"TotalAssassinations"`
		TotalGroundPoundKills          int           `json:"TotalGroundPoundKills"`
		TotalGroundPoundDamage         float64       `json:"TotalGroundPoundDamage"`
		TotalShoulderBashKills         int           `json:"TotalShoulderBashKills"`
		TotalShoulderBashDamage        float64       `json:"TotalShoulderBashDamage"`
		TotalGrenadeDamage             float64       `json:"TotalGrenadeDamage"`
		TotalPowerWeaponKills          int           `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponDamage         float64       `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int           `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponPossessionTime string        `json:"TotalPowerWeaponPossessionTime"`
		TotalDeaths                    int           `json:"TotalDeaths"`
		TotalAssists                   int           `json:"TotalAssists"`
		TotalGamesCompleted            int           `json:"TotalGamesCompleted"`
		TotalGamesWon                  int           `json:"TotalGamesWon"`
		TotalGamesLost                 int           `json:"TotalGamesLost"`
		TotalGamesTied                 int           `json:"TotalGamesTied"`
		TotalTimePlayed                string        `json:"TotalTimePlayed"`
		TotalGrenadeKills              int           `json:"TotalGrenadeKills"`
		MedalAwards                    []interface{} `json:"MedalAwards"`
		DestroyedEnemyVehicles         []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills                     []struct {
			Enemy struct {
				BaseID      int64         `json:"BaseId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"Enemy"`
			TotalKills int `json:"TotalKills"`
		} `json:"EnemyKills"`
		WeaponStats []struct {
			WeaponID struct {
				StockID     int64         `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponStats"`
		Impulses          []interface{} `json:"Impulses"`
		TotalSpartanKills int           `json:"TotalSpartanKills"`
	} `json:"PlayerStats"`
	TotalMissionPlaythroughTime string        `json:"TotalMissionPlaythroughTime"`
	Difficulty                  int           `json:"Difficulty"`
	Skulls                      []interface{} `json:"Skulls"`
	MissionCompleted            bool          `json:"MissionCompleted"`
	IsMatchOver                 bool          `json:"IsMatchOver"`
	TotalDuration               string        `json:"TotalDuration"`
	MapVariantID                string        `json:"MapVariantId"`
	GameVariantID               string        `json:"GameVariantId"`
	PlaylistID                  string        `json:"PlaylistId"`
	MapID                       string        `json:"MapId"`
	GameBaseVariantID           string        `json:"GameBaseVariantId"`
	IsTeamGame                  bool          `json:"IsTeamGame"`
	SeasonID                    interface{}   `json:"SeasonId"`
	GameVariantResourceID       struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"GameVariantResourceId"`
	MapVariantResourceID struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"MapVariantResourceId"`
}

type CarnageReportCustomStruct struct {
	PlayerStats []struct {
		KilledOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledOpponentDetails"`
		KilledByOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledByOpponentDetails"`
		FlexibleStats struct {
			MedalStatCounts   []interface{} `json:"MedalStatCounts"`
			ImpulseStatCounts []struct {
				ID    string `json:"Id"`
				Count int    `json:"Count"`
			} `json:"ImpulseStatCounts"`
			MedalTimelapses   []interface{} `json:"MedalTimelapses"`
			ImpulseTimelapses []struct {
				ID        string `json:"Id"`
				Timelapse string `json:"Timelapse"`
			} `json:"ImpulseTimelapses"`
		} `json:"FlexibleStats"`
		Player struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		TeamID              int         `json:"TeamId"`
		Rank                int         `json:"Rank"`
		DNF                 bool        `json:"DNF"`
		AvgLifeTimeOfPlayer string      `json:"AvgLifeTimeOfPlayer"`
		PreMatchRatings     interface{} `json:"PreMatchRatings"`
		PostMatchRatings    interface{} `json:"PostMatchRatings"`
		PlayerScore         int         `json:"PlayerScore"`
		TotalKills          int         `json:"TotalKills"`
		TotalHeadshots      int         `json:"TotalHeadshots"`
		TotalWeaponDamage   float64     `json:"TotalWeaponDamage"`
		TotalShotsFired     int         `json:"TotalShotsFired"`
		TotalShotsLanded    int         `json:"TotalShotsLanded"`
		WeaponWithMostKills struct {
			WeaponID struct {
				StockID     int           `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponWithMostKills"`
		TotalMeleeKills                int     `json:"TotalMeleeKills"`
		TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
		TotalAssassinations            int     `json:"TotalAssassinations"`
		TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
		TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
		TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
		TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
		TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
		TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
		TotalDeaths                    int     `json:"TotalDeaths"`
		TotalAssists                   int     `json:"TotalAssists"`
		TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
		TotalGamesWon                  int     `json:"TotalGamesWon"`
		TotalGamesLost                 int     `json:"TotalGamesLost"`
		TotalGamesTied                 int     `json:"TotalGamesTied"`
		TotalTimePlayed                string  `json:"TotalTimePlayed"`
		TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
		MedalAwards                    []struct {
			MedalID int64 `json:"MedalId"`
			Count   int   `json:"Count"`
		} `json:"MedalAwards"`
		DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills             []interface{} `json:"EnemyKills"`
		WeaponStats            []struct {
			WeaponID struct {
				StockID     int           `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponStats"`
		Impulses []struct {
			ID    int64 `json:"Id"`
			Count int   `json:"Count"`
		} `json:"Impulses"`
		TotalSpartanKills int `json:"TotalSpartanKills"`
	} `json:"PlayerStats"`
	TeamStats []struct {
		TeamID     int           `json:"TeamId"`
		Score      int           `json:"Score"`
		Rank       int           `json:"Rank"`
		RoundStats []interface{} `json:"RoundStats"`
	} `json:"TeamStats"`
	IsMatchOver           bool        `json:"IsMatchOver"`
	TotalDuration         string      `json:"TotalDuration"`
	MapVariantID          string      `json:"MapVariantId"`
	GameVariantID         string      `json:"GameVariantId"`
	PlaylistID            string      `json:"PlaylistId"`
	MapID                 string      `json:"MapId"`
	GameBaseVariantID     string      `json:"GameBaseVariantId"`
	IsTeamGame            bool        `json:"IsTeamGame"`
	SeasonID              interface{} `json:"SeasonId"`
	GameVariantResourceID struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"GameVariantResourceId"`
	MapVariantResourceID struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"MapVariantResourceId"`
}

type CarnageReportWarzoneStruct struct {
	PlayerStats []struct {
		XpInfo struct {
			PrevSpartanRank              int     `json:"PrevSpartanRank"`
			SpartanRank                  int     `json:"SpartanRank"`
			PrevTotalXP                  int     `json:"PrevTotalXP"`
			TotalXP                      int     `json:"TotalXP"`
			SpartanRankMatchXPScalar     float64 `json:"SpartanRankMatchXPScalar"`
			PlayerTimePerformanceXPAward int     `json:"PlayerTimePerformanceXPAward"`
			PerformanceXP                int     `json:"PerformanceXP"`
			PlayerRankXPAward            int     `json:"PlayerRankXPAward"`
			BoostAmount                  int     `json:"BoostAmount"`
			MatchSpeedWinAmount          int     `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount    int     `json:"ObjectivesCompletedAmount"`
		} `json:"XpInfo"`
		KilledOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledOpponentDetails"`
		KilledByOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledByOpponentDetails"`
		WarzoneLevel    int `json:"WarzoneLevel"`
		TotalPiesEarned int `json:"TotalPiesEarned"`
		FlexibleStats   struct {
			MedalStatCounts []struct {
				ID    string `json:"Id"`
				Count int    `json:"Count"`
			} `json:"MedalStatCounts"`
			ImpulseStatCounts []struct {
				ID    string `json:"Id"`
				Count int    `json:"Count"`
			} `json:"ImpulseStatCounts"`
			MedalTimelapses   []interface{} `json:"MedalTimelapses"`
			ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
		} `json:"FlexibleStats"`
		RewardSets []struct {
			RewardSet           string      `json:"RewardSet"`
			RewardSourceType    int         `json:"RewardSourceType"`
			SpartanRankSource   interface{} `json:"SpartanRankSource"`
			CommendationLevelID string      `json:"CommendationLevelId"`
			CommendationSource  string      `json:"CommendationSource"`
		} `json:"RewardSets"`
		CreditsEarned struct {
			Result                    int     `json:"Result"`
			TotalCreditsEarned        int     `json:"TotalCreditsEarned"`
			SpartanRankModifier       float64 `json:"SpartanRankModifier"`
			PlayerRankAmount          int     `json:"PlayerRankAmount"`
			TimePlayedAmount          float64 `json:"TimePlayedAmount"`
			BoostAmount               int     `json:"BoostAmount"`
			MatchSpeedWinAmount       int     `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount int     `json:"ObjectivesCompletedAmount"`
		} `json:"CreditsEarned"`
		MetaCommendationDeltas        []interface{} `json:"MetaCommendationDeltas"`
		ProgressiveCommendationDeltas []struct {
			ID               string `json:"Id"`
			PreviousProgress int    `json:"PreviousProgress"`
			Progress         int    `json:"Progress"`
		} `json:"ProgressiveCommendationDeltas"`
		Player struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		TeamID              int         `json:"TeamId"`
		Rank                int         `json:"Rank"`
		DNF                 bool        `json:"DNF"`
		AvgLifeTimeOfPlayer string      `json:"AvgLifeTimeOfPlayer"`
		PreMatchRatings     interface{} `json:"PreMatchRatings"`
		PostMatchRatings    interface{} `json:"PostMatchRatings"`
		PlayerScore         int         `json:"PlayerScore"`
		TotalKills          int         `json:"TotalKills"`
		TotalHeadshots      int         `json:"TotalHeadshots"`
		TotalWeaponDamage   float64     `json:"TotalWeaponDamage"`
		TotalShotsFired     int         `json:"TotalShotsFired"`
		TotalShotsLanded    int         `json:"TotalShotsLanded"`
		WeaponWithMostKills struct {
			WeaponID struct {
				StockID     int           `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponWithMostKills"`
		TotalMeleeKills                int     `json:"TotalMeleeKills"`
		TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
		TotalAssassinations            int     `json:"TotalAssassinations"`
		TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
		TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
		TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
		TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
		TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
		TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
		TotalDeaths                    int     `json:"TotalDeaths"`
		TotalAssists                   int     `json:"TotalAssists"`
		TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
		TotalGamesWon                  int     `json:"TotalGamesWon"`
		TotalGamesLost                 int     `json:"TotalGamesLost"`
		TotalGamesTied                 int     `json:"TotalGamesTied"`
		TotalTimePlayed                string  `json:"TotalTimePlayed"`
		TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
		MedalAwards                    []struct {
			MedalID int64 `json:"MedalId"`
			Count   int   `json:"Count"`
		} `json:"MedalAwards"`
		DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills             []struct {
			Enemy struct {
				BaseID      int64         `json:"BaseId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"Enemy"`
			TotalKills int `json:"TotalKills"`
		} `json:"EnemyKills"`
		WeaponStats []struct {
			WeaponID struct {
				StockID     int64         `json:"StockId"`
				Attachments []interface{} `json:"Attachments"`
			} `json:"WeaponId"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
		} `json:"WeaponStats"`
		Impulses []struct {
			ID    int `json:"Id"`
			Count int `json:"Count"`
		} `json:"Impulses"`
		TotalSpartanKills int `json:"TotalSpartanKills"`
	} `json:"PlayerStats"`
	TeamStats []struct {
		TeamID     int `json:"TeamId"`
		Score      int `json:"Score"`
		Rank       int `json:"Rank"`
		RoundStats []struct {
			RoundNumber int `json:"RoundNumber"`
			Rank        int `json:"Rank"`
			Score       int `json:"Score"`
		} `json:"RoundStats"`
	} `json:"TeamStats"`
	IsMatchOver           bool   `json:"IsMatchOver"`
	TotalDuration         string `json:"TotalDuration"`
	MapVariantID          string `json:"MapVariantId"`
	GameVariantID         string `json:"GameVariantId"`
	PlaylistID            string `json:"PlaylistId"`
	MapID                 string `json:"MapId"`
	GameBaseVariantID     string `json:"GameBaseVariantId"`
	IsTeamGame            bool   `json:"IsTeamGame"`
	SeasonID              string `json:"SeasonId"`
	GameVariantResourceID struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"GameVariantResourceId"`
	MapVariantResourceID struct {
		ResourceType int    `json:"ResourceType"`
		ResourceID   string `json:"ResourceId"`
		OwnerType    int    `json:"OwnerType"`
		Owner        string `json:"Owner"`
	} `json:"MapVariantResourceId"`
}

type CampaignMissionsStruct []struct {
	MissionNumber json.Number `json:"missionNumber"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	ImageURL      string      `json:"imageUrl"`
	Type          string      `json:"type"`
	ID            string      `json:"id"`
	ContentID     string      `json:"contentId"`
}

type CommendationsStruct []struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IconImageURL string `json:"iconImageUrl"`
	Levels       []struct {
		Reward struct {
			Xp               int `json:"xp"`
			RequisitionPacks []struct {
				Name                         string        `json:"name"`
				Description                  string        `json:"description"`
				LargeImageURL                string        `json:"largeImageUrl"`
				IsStack                      bool          `json:"isStack"`
				IsFeatured                   bool          `json:"isFeatured"`
				IsNew                        bool          `json:"isNew"`
				CreditPrice                  int           `json:"creditPrice"`
				IsPurchasableWithCredits     bool          `json:"isPurchasableWithCredits"`
				IsPurchasableFromMarketplace bool          `json:"isPurchasableFromMarketplace"`
				XboxMarketplaceProductID     interface{}   `json:"xboxMarketplaceProductId"`
				XboxMarketplaceProductURL    interface{}   `json:"xboxMarketplaceProductUrl"`
				MerchandisingOrder           int           `json:"merchandisingOrder"`
				Flair                        interface{}   `json:"flair"`
				StackedRequisitionPacks      []interface{} `json:"stackedRequisitionPacks"`
				ID                           string        `json:"id"`
				ContentID                    string        `json:"contentId"`
			} `json:"requisitionPacks"`
			ID        string `json:"id"`
			ContentID string `json:"contentId"`
		} `json:"reward"`
		Threshold int    `json:"threshold"`
		ID        string `json:"id"`
		ContentID string `json:"contentId"`
	} `json:"levels"`
	RequiredLevels []interface{} `json:"requiredLevels"`
	Reward         interface{}   `json:"reward"`
	Category       struct {
		Name         string `json:"name"`
		IconImageURL string `json:"iconImageUrl"`
		Order        int    `json:"order"`
		ID           string `json:"id"`
		ContentID    string `json:"contentId"`
	} `json:"category"`
	ID        string `json:"id"`
	ContentID string `json:"contentId"`
}

type CsrDesignationsStruct []struct {
	Name           string `json:"name"`
	BannerImageURL string `json:"bannerImageUrl"`
	Tiers          []struct {
		IconImageURL string `json:"iconImageUrl"`
		ID           string `json:"id"`
		ContentID    string `json:"contentId"`
	} `json:"tiers"`
	ID        string `json:"id"`
	ContentID string `json:"contentId"`
}

type EnemiesStruct []struct {
	ContentID         string      `json:"contentId"`
	Description       interface{} `json:"description"`
	Faction           string      `json:"faction"`
	ID                string      `json:"id"`
	LargeIconImageURL string      `json:"largeIconImageUrl"`
	Name              string      `json:"name"`
	SmallIconImageURL string      `json:"smallIconImageUrl"`
}

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

type FlexibleStatsStruct []struct {
	ContentID string `json:"contentId"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type GameBaseVariantsStruct []struct {
	ContentID          string   `json:"contentId"`
	IconURL            string   `json:"iconUrl"`
	ID                 string   `json:"id"`
	InternalName       string   `json:"internalName"`
	Name               string   `json:"name"`
	SupportedGameModes []string `json:"supportedGameModes"`
}

type GameVariantsStruct struct {
	ContentID         string `json:"contentId"`
	Description       string `json:"description"`
	GameBaseVariantID string `json:"gameBaseVariantId"`
	IconURL           string `json:"iconUrl"`
	ID                string `json:"id"`
	Name              string `json:"name"`
}

type ImpulsesStruct []struct {
	ContentID    string `json:"contentId"`
	ID           string `json:"id"`
	InternalName string `json:"internalName"`
}

type MapsStruct []struct {
	ContentID          string   `json:"contentId"`
	Description        string   `json:"description"`
	ID                 string   `json:"id"`
	ImageURL           string   `json:"imageUrl"`
	Name               string   `json:"name"`
	SupportedGameModes []string `json:"supportedGameModes"`
}

type MapVariantsStruct struct {
	ContentID   string `json:"contentId"`
	Description string `json:"description"`
	ID          string `json:"id"`
	MapID       string `json:"mapId"`
	MapImageURL string `json:"mapImageUrl"`
	Name        string `json:"name"`
}

type MatchesForPlayerStruct struct {
	Count int `json:"Count"`
	Links struct {
		Self struct {
			AcknowledgementTypeID                    int    `json:"AcknowledgementTypeId"`
			AuthenticationLifetimeExtensionSupported bool   `json:"AuthenticationLifetimeExtensionSupported"`
			AuthorityID                              string `json:"AuthorityId"`
			Path                                     string `json:"Path"`
			QueryString                              string `json:"QueryString"`
			RetryPolicyID                            string `json:"RetryPolicyId"`
			TopicName                                string `json:"TopicName"`
		} `json:"Self"`
	} `json:"Links"`
	ResultCount int `json:"ResultCount"`
	Results     []struct {
		GameBaseVariantID string `json:"GameBaseVariantId"`
		GameVariant       struct {
			Owner        string `json:"Owner"`
			OwnerType    int    `json:"OwnerType"`
			ResourceID   string `json:"ResourceId"`
			ResourceType int    `json:"ResourceType"`
		} `json:"GameVariant"`
		HopperID string `json:"HopperId"`
		ID       struct {
			GameMode int    `json:"GameMode"`
			MatchID  string `json:"MatchId"`
		} `json:"Id"`
		IsTeamGame bool `json:"IsTeamGame"`
		Links      struct {
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
		MapID      string `json:"MapId"`
		MapVariant struct {
			Owner        string `json:"Owner"`
			OwnerType    int    `json:"OwnerType"`
			ResourceID   string `json:"ResourceId"`
			ResourceType int    `json:"ResourceType"`
		} `json:"MapVariant"`
		MatchCompletedDate struct {
			ISO8601Date string `json:"ISO8601Date"`
		} `json:"MatchCompletedDate"`
		MatchCompletedDateFidelity int    `json:"MatchCompletedDateFidelity"`
		MatchDuration              string `json:"MatchDuration"`
		Players                    []struct {
			Player struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"Player"`
			PostMatchRatings interface{} `json:"PostMatchRatings"`
			PreMatchRatings  interface{} `json:"PreMatchRatings"`
			Rank             int         `json:"Rank"`
			Result           int         `json:"Result"`
			TeamID           int         `json:"TeamId"`
			TotalAssists     int         `json:"TotalAssists"`
			TotalDeaths      int         `json:"TotalDeaths"`
			TotalKills       int         `json:"TotalKills"`
		} `json:"Players"`
		SeasonID string `json:"SeasonId"`
		Teams    []struct {
			ID    int `json:"Id"`
			Rank  int `json:"Rank"`
			Score int `json:"Score"`
		} `json:"Teams"`
	} `json:"Results"`
	Start int `json:"Start"`
}

type MedalsStruct []struct {
	Classification string `json:"classification"`
	ContentID      string `json:"contentId"`
	Description    string `json:"description"`
	Difficulty     int    `json:"difficulty"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	SpriteLocation struct {
		Height         int    `json:"height"`
		Left           int    `json:"left"`
		SpriteHeight   int    `json:"spriteHeight"`
		SpriteSheetURI string `json:"spriteSheetUri"`
		SpriteWidth    int    `json:"spriteWidth"`
		Top            int    `json:"top"`
		Width          int    `json:"width"`
	} `json:"spriteLocation"`
}

type PlayerLeaderboardStruct struct {
	Count int `json:"Count"`
	Links struct {
		Self struct {
			AcknowledgementTypeID                    int    `json:"AcknowledgementTypeId"`
			AuthenticationLifetimeExtensionSupported bool   `json:"AuthenticationLifetimeExtensionSupported"`
			AuthorityID                              string `json:"AuthorityId"`
			Path                                     string `json:"Path"`
			QueryString                              string `json:"QueryString"`
			RetryPolicyID                            string `json:"RetryPolicyId"`
			TopicName                                string `json:"TopicName"`
		} `json:"Self"`
	} `json:"Links"`
	ResultCount int `json:"ResultCount"`
	Results     []struct {
		Player struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		Rank  int `json:"Rank"`
		Score struct {
			Csr               int `json:"Csr"`
			DesignationID     int `json:"DesignationId"`
			PercentToNextTier int `json:"PercentToNextTier"`
			Rank              int `json:"Rank"`
			Tier              int `json:"Tier"`
		} `json:"Score"`
	} `json:"Results"`
	Start int `json:"Start"`
}

type PlaylistsStruct []struct {
	ContentID   string      `json:"contentId"`
	Description string      `json:"description"`
	GameMode    string      `json:"gameMode"`
	ID          string      `json:"id"`
	ImageURL    interface{} `json:"imageUrl"`
	IsActive    bool        `json:"isActive"`
	IsRanked    bool        `json:"isRanked"`
	Name        string      `json:"name"`
}

type RequisitionPacksStruct struct {
	ContentID                    string        `json:"contentId"`
	CreditPrice                  int           `json:"creditPrice"`
	Description                  string        `json:"description"`
	Flair                        interface{}   `json:"flair"`
	ID                           string        `json:"id"`
	IsFeatured                   bool          `json:"isFeatured"`
	IsNew                        bool          `json:"isNew"`
	IsPurchasableFromMarketplace bool          `json:"isPurchasableFromMarketplace"`
	IsPurchasableWithCredits     bool          `json:"isPurchasableWithCredits"`
	IsStack                      bool          `json:"isStack"`
	LargeImageURL                string        `json:"largeImageUrl"`
	MerchandisingOrder           int           `json:"merchandisingOrder"`
	Name                         string        `json:"name"`
	StackedRequisitionPacks      []interface{} `json:"stackedRequisitionPacks"`
	XboxMarketplaceProductID     interface{}   `json:"xboxMarketplaceProductId"`
	XboxMarketplaceProductURL    interface{}   `json:"xboxMarketplaceProductUrl"`
}

type RequisitionsStruct struct {
	CategoryName               interface{} `json:"categoryName"`
	CertificationRequisitionID interface{} `json:"certificationRequisitionId"`
	ContentID                  string      `json:"contentId"`
	Description                string      `json:"description"`
	HideIfNotAcquired          bool        `json:"hideIfNotAcquired"`
	ID                         string      `json:"id"`
	InternalCategoryName       interface{} `json:"internalCategoryName"`
	InternalSubcategoryName    interface{} `json:"internalSubcategoryName"`
	IsCertification            bool        `json:"isCertification"`
	IsMythic                   bool        `json:"isMythic"`
	IsWearable                 bool        `json:"isWearable"`
	LargeImageURL              string      `json:"largeImageUrl"`
	LevelRequirement           int         `json:"levelRequirement"`
	Name                       string      `json:"name"`
	Rarity                     string      `json:"rarity"`
	RarityType                 string      `json:"rarityType"`
	SellPrice                  int         `json:"sellPrice"`
	SubcategoryName            interface{} `json:"subcategoryName"`
	SubcategoryOrder           int         `json:"subcategoryOrder"`
	SupportedGameModes         []string    `json:"supportedGameModes"`
	UseType                    string      `json:"useType"`
}

type SeasonsStruct []struct {
	ContentID string      `json:"contentId"`
	EndDate   string      `json:"endDate"`
	IconURL   interface{} `json:"iconUrl"`
	ID        string      `json:"id"`
	IsActive  bool        `json:"isActive"`
	Name      string      `json:"name"`
	Playlists []struct {
		ContentID   string      `json:"contentId"`
		Description string      `json:"description"`
		GameMode    string      `json:"gameMode"`
		ID          string      `json:"id"`
		ImageURL    interface{} `json:"imageUrl"`
		IsActive    bool        `json:"isActive"`
		IsRanked    bool        `json:"isRanked"`
		Name        string      `json:"name"`
	} `json:"playlists"`
	StartDate string `json:"startDate"`
}

type ServiceRecordWarzoneStruct struct {
	Results []struct {
		ID         string `json:"Id"`
		ResultCode int    `json:"ResultCode"`
		Result     struct {
			WarzoneStat struct {
				TotalPiesEarned int `json:"TotalPiesEarned"`
				ScenarioStats   []struct {
					TotalPiesEarned int `json:"TotalPiesEarned"`
					FlexibleStats   struct {
						MedalStatCounts []struct {
							ID    string `json:"Id"`
							Count int    `json:"Count"`
						} `json:"MedalStatCounts"`
						ImpulseStatCounts []struct {
							ID    string `json:"Id"`
							Count int    `json:"Count"`
						} `json:"ImpulseStatCounts"`
						MedalTimelapses   []interface{} `json:"MedalTimelapses"`
						ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
					} `json:"FlexibleStats"`
					MapID               string  `json:"MapId"`
					GameBaseVariantID   string  `json:"GameBaseVariantId"`
					TotalKills          int     `json:"TotalKills"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalWeaponDamage   float64 `json:"TotalWeaponDamage"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponWithMostKills struct {
						WeaponID struct {
							StockID     int           `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponWithMostKills"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
					TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					MedalAwards                    []struct {
						MedalID int `json:"MedalId"`
						Count   int `json:"Count"`
					} `json:"MedalAwards"`
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []struct {
						Enemy struct {
							BaseID      int           `json:"BaseId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"Enemy"`
						TotalKills int `json:"TotalKills"`
					} `json:"EnemyKills"`
					WeaponStats []struct {
						WeaponID struct {
							StockID     int           `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponStats"`
					Impulses          []interface{} `json:"Impulses"`
					TotalSpartanKills int           `json:"TotalSpartanKills"`
					FastestMatchWin   interface{}   `json:"FastestMatchWin"`
				} `json:"ScenarioStats"`
				TotalKills          int     `json:"TotalKills"`
				TotalHeadshots      int     `json:"TotalHeadshots"`
				TotalWeaponDamage   float64 `json:"TotalWeaponDamage"`
				TotalShotsFired     int     `json:"TotalShotsFired"`
				TotalShotsLanded    int     `json:"TotalShotsLanded"`
				WeaponWithMostKills struct {
					WeaponID struct {
						StockID     int           `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponWithMostKills"`
				TotalMeleeKills                int     `json:"TotalMeleeKills"`
				TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
				TotalAssassinations            int     `json:"TotalAssassinations"`
				TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
				TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
				TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
				TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
				TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
				TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
				TotalDeaths                    int     `json:"TotalDeaths"`
				TotalAssists                   int     `json:"TotalAssists"`
				TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
				TotalGamesWon                  int     `json:"TotalGamesWon"`
				TotalGamesLost                 int     `json:"TotalGamesLost"`
				TotalGamesTied                 int     `json:"TotalGamesTied"`
				TotalTimePlayed                string  `json:"TotalTimePlayed"`
				TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
				MedalAwards                    []struct {
					MedalID int `json:"MedalId"`
					Count   int `json:"Count"`
				} `json:"MedalAwards"`
				DestroyedEnemyVehicles []struct {
					Enemy struct {
						BaseID      int           `json:"BaseId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"DestroyedEnemyVehicles"`
				EnemyKills []struct {
					Enemy struct {
						BaseID      int           `json:"BaseId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"EnemyKills"`
				WeaponStats []struct {
					WeaponID struct {
						StockID     int           `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponStats"`
				Impulses []struct {
					ID    int64 `json:"Id"`
					Count int   `json:"Count"`
				} `json:"Impulses"`
				TotalSpartanKills int         `json:"TotalSpartanKills"`
				FastestMatchWin   interface{} `json:"FastestMatchWin"`
			} `json:"WarzoneStat"`
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			Xp          int `json:"Xp"`
		} `json:"Result"`
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

type ServiceRecordCustomStruct struct {
	Results []struct {
		ID         string `json:"Id"`
		ResultCode int    `json:"ResultCode"`
		Result     struct {
			CustomStats struct {
				CustomGameBaseVariantStats []struct {
					FlexibleStats struct {
						MedalStatCounts   []interface{} `json:"MedalStatCounts"`
						ImpulseStatCounts []struct {
							ID    string `json:"Id"`
							Count int    `json:"Count"`
						} `json:"ImpulseStatCounts"`
						MedalTimelapses   []interface{} `json:"MedalTimelapses"`
						ImpulseTimelapses []struct {
							ID        string `json:"Id"`
							Timelapse string `json:"Timelapse"`
						} `json:"ImpulseTimelapses"`
					} `json:"FlexibleStats"`
					GameBaseVariantID   string  `json:"GameBaseVariantId"`
					TotalKills          int     `json:"TotalKills"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalWeaponDamage   float64 `json:"TotalWeaponDamage"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponWithMostKills struct {
						WeaponID struct {
							StockID     int64         `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponWithMostKills"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
					TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					MedalAwards                    []struct {
						MedalID int64 `json:"MedalId"`
						Count   int   `json:"Count"`
					} `json:"MedalAwards"`
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []interface{} `json:"EnemyKills"`
					WeaponStats            []struct {
						WeaponID struct {
							StockID     int           `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponStats"`
					Impulses []struct {
						ID    int `json:"Id"`
						Count int `json:"Count"`
					} `json:"Impulses"`
					TotalSpartanKills int         `json:"TotalSpartanKills"`
					FastestMatchWin   interface{} `json:"FastestMatchWin"`
				} `json:"CustomGameBaseVariantStats"`
				TopGameBaseVariants []struct {
					GameBaseVariantRank      int    `json:"GameBaseVariantRank"`
					NumberOfMatchesCompleted int    `json:"NumberOfMatchesCompleted"`
					GameBaseVariantID        string `json:"GameBaseVariantId"`
					NumberOfMatchesWon       int    `json:"NumberOfMatchesWon"`
				} `json:"TopGameBaseVariants"`
				TotalKills          int     `json:"TotalKills"`
				TotalHeadshots      int     `json:"TotalHeadshots"`
				TotalWeaponDamage   float64 `json:"TotalWeaponDamage"`
				TotalShotsFired     int     `json:"TotalShotsFired"`
				TotalShotsLanded    int     `json:"TotalShotsLanded"`
				WeaponWithMostKills struct {
					WeaponID struct {
						StockID     int64         `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponWithMostKills"`
				TotalMeleeKills                int     `json:"TotalMeleeKills"`
				TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
				TotalAssassinations            int     `json:"TotalAssassinations"`
				TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
				TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
				TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
				TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
				TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
				TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
				TotalDeaths                    int     `json:"TotalDeaths"`
				TotalAssists                   int     `json:"TotalAssists"`
				TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
				TotalGamesWon                  int     `json:"TotalGamesWon"`
				TotalGamesLost                 int     `json:"TotalGamesLost"`
				TotalGamesTied                 int     `json:"TotalGamesTied"`
				TotalTimePlayed                string  `json:"TotalTimePlayed"`
				TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
				MedalAwards                    []struct {
					MedalID int64 `json:"MedalId"`
					Count   int   `json:"Count"`
				} `json:"MedalAwards"`
				DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
				EnemyKills             []interface{} `json:"EnemyKills"`
				WeaponStats            []struct {
					WeaponID struct {
						StockID     int           `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponStats"`
				Impulses []struct {
					ID    int `json:"Id"`
					Count int `json:"Count"`
				} `json:"Impulses"`
				TotalSpartanKills int         `json:"TotalSpartanKills"`
				FastestMatchWin   interface{} `json:"FastestMatchWin"`
			} `json:"CustomStats"`
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			Xp          int `json:"Xp"`
		} `json:"Result"`
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

type ServiceRecordCampaignStruct struct {
	Results []struct {
		ID         string `json:"Id"`
		ResultCode int    `json:"ResultCode"`
		Result     struct {
			CampaignStat struct {
				CampaignMissionStats []struct {
					FlexibleStats struct {
						MedalStatCounts   []interface{} `json:"MedalStatCounts"`
						ImpulseStatCounts []struct {
							ID    string `json:"Id"`
							Count int    `json:"Count"`
						} `json:"ImpulseStatCounts"`
						MedalTimelapses   []interface{} `json:"MedalTimelapses"`
						ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
					} `json:"FlexibleStats"`
					CoopStats struct {
						Num3 struct {
							HighestScore          int           `json:"HighestScore"`
							FastestCompletionTime string        `json:"FastestCompletionTime"`
							Skulls                []interface{} `json:"Skulls"`
							TotalTimesCompleted   int           `json:"TotalTimesCompleted"`
							AllSkullsOn           bool          `json:"AllSkullsOn"`
						} `json:"3"`
					} `json:"CoopStats"`
					SoloStats struct {
						Num2 struct {
							HighestScore          int           `json:"HighestScore"`
							FastestCompletionTime string        `json:"FastestCompletionTime"`
							Skulls                []interface{} `json:"Skulls"`
							TotalTimesCompleted   int           `json:"TotalTimesCompleted"`
							AllSkullsOn           bool          `json:"AllSkullsOn"`
						} `json:"2"`
					} `json:"SoloStats"`
					MissionID           string  `json:"MissionId"`
					TotalKills          int     `json:"TotalKills"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalWeaponDamage   float64 `json:"TotalWeaponDamage"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponWithMostKills struct {
						WeaponID struct {
							StockID     int64         `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponWithMostKills"`
					TotalMeleeKills                int           `json:"TotalMeleeKills"`
					TotalMeleeDamage               float64       `json:"TotalMeleeDamage"`
					TotalAssassinations            int           `json:"TotalAssassinations"`
					TotalGroundPoundKills          int           `json:"TotalGroundPoundKills"`
					TotalGroundPoundDamage         float64       `json:"TotalGroundPoundDamage"`
					TotalShoulderBashKills         int           `json:"TotalShoulderBashKills"`
					TotalShoulderBashDamage        float64       `json:"TotalShoulderBashDamage"`
					TotalGrenadeDamage             float64       `json:"TotalGrenadeDamage"`
					TotalPowerWeaponKills          int           `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponDamage         float64       `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int           `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponPossessionTime string        `json:"TotalPowerWeaponPossessionTime"`
					TotalDeaths                    int           `json:"TotalDeaths"`
					TotalAssists                   int           `json:"TotalAssists"`
					TotalGamesCompleted            int           `json:"TotalGamesCompleted"`
					TotalGamesWon                  int           `json:"TotalGamesWon"`
					TotalGamesLost                 int           `json:"TotalGamesLost"`
					TotalGamesTied                 int           `json:"TotalGamesTied"`
					TotalTimePlayed                string        `json:"TotalTimePlayed"`
					TotalGrenadeKills              int           `json:"TotalGrenadeKills"`
					MedalAwards                    []interface{} `json:"MedalAwards"`
					DestroyedEnemyVehicles         []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills                     []struct {
						Enemy struct {
							BaseID      int64         `json:"BaseId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"Enemy"`
						TotalKills int `json:"TotalKills"`
					} `json:"EnemyKills"`
					WeaponStats []struct {
						WeaponID struct {
							StockID     int           `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponStats"`
					Impulses []struct {
						ID    int `json:"Id"`
						Count int `json:"Count"`
					} `json:"Impulses"`
					TotalSpartanKills int         `json:"TotalSpartanKills"`
					FastestMatchWin   interface{} `json:"FastestMatchWin"`
				} `json:"CampaignMissionStats"`
				TotalKills          int     `json:"TotalKills"`
				TotalHeadshots      int     `json:"TotalHeadshots"`
				TotalWeaponDamage   float64 `json:"TotalWeaponDamage"`
				TotalShotsFired     int     `json:"TotalShotsFired"`
				TotalShotsLanded    int     `json:"TotalShotsLanded"`
				WeaponWithMostKills struct {
					WeaponID struct {
						StockID     int           `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponWithMostKills"`
				TotalMeleeKills                int           `json:"TotalMeleeKills"`
				TotalMeleeDamage               float64       `json:"TotalMeleeDamage"`
				TotalAssassinations            int           `json:"TotalAssassinations"`
				TotalGroundPoundKills          int           `json:"TotalGroundPoundKills"`
				TotalGroundPoundDamage         float64       `json:"TotalGroundPoundDamage"`
				TotalShoulderBashKills         int           `json:"TotalShoulderBashKills"`
				TotalShoulderBashDamage        float64       `json:"TotalShoulderBashDamage"`
				TotalGrenadeDamage             float64       `json:"TotalGrenadeDamage"`
				TotalPowerWeaponKills          int           `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponDamage         float64       `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int           `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponPossessionTime string        `json:"TotalPowerWeaponPossessionTime"`
				TotalDeaths                    int           `json:"TotalDeaths"`
				TotalAssists                   int           `json:"TotalAssists"`
				TotalGamesCompleted            int           `json:"TotalGamesCompleted"`
				TotalGamesWon                  int           `json:"TotalGamesWon"`
				TotalGamesLost                 int           `json:"TotalGamesLost"`
				TotalGamesTied                 int           `json:"TotalGamesTied"`
				TotalTimePlayed                string        `json:"TotalTimePlayed"`
				TotalGrenadeKills              int           `json:"TotalGrenadeKills"`
				MedalAwards                    []interface{} `json:"MedalAwards"`
				DestroyedEnemyVehicles         []struct {
					Enemy struct {
						BaseID      int           `json:"BaseId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"DestroyedEnemyVehicles"`
				EnemyKills []struct {
					Enemy struct {
						BaseID      int64         `json:"BaseId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"EnemyKills"`
				WeaponStats []struct {
					WeaponID struct {
						StockID     int           `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponStats"`
				Impulses []struct {
					ID    int `json:"Id"`
					Count int `json:"Count"`
				} `json:"Impulses"`
				TotalSpartanKills int         `json:"TotalSpartanKills"`
				FastestMatchWin   interface{} `json:"FastestMatchWin"`
			} `json:"CampaignStat"`
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			Xp          int `json:"Xp"`
		} `json:"Result"`
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

type ServiceRecordArenaStruct struct {
	Results []struct {
		ID         string `json:"Id"`
		ResultCode int    `json:"ResultCode"`
		Result     struct {
			ArenaStats struct {
				ArenaPlaylistStats []struct {
					PlaylistID             string      `json:"PlaylistId"`
					MeasurementMatchesLeft int         `json:"MeasurementMatchesLeft"`
					HighestCsr             interface{} `json:"HighestCsr"`
					Csr                    interface{} `json:"Csr"`
					CsrPercentile          interface{} `json:"CsrPercentile"`
					TotalKills             int         `json:"TotalKills"`
					TotalHeadshots         int         `json:"TotalHeadshots"`
					TotalWeaponDamage      float64     `json:"TotalWeaponDamage"`
					TotalShotsFired        int         `json:"TotalShotsFired"`
					TotalShotsLanded       int         `json:"TotalShotsLanded"`
					WeaponWithMostKills    struct {
						WeaponID struct {
							StockID     int           `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponWithMostKills"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
					TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					MedalAwards                    []struct {
						MedalID int64 `json:"MedalId"`
						Count   int   `json:"Count"`
					} `json:"MedalAwards"`
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []interface{} `json:"EnemyKills"`
					WeaponStats            []struct {
						WeaponID struct {
							StockID     int           `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponStats"`
					Impulses          []interface{} `json:"Impulses"`
					TotalSpartanKills int           `json:"TotalSpartanKills"`
					FastestMatchWin   interface{}   `json:"FastestMatchWin"`
				} `json:"ArenaPlaylistStats"`
				HighestCsrAttained struct {
					Tier              int         `json:"Tier"`
					DesignationID     int         `json:"DesignationId"`
					Csr               int         `json:"Csr"`
					PercentToNextTier int         `json:"PercentToNextTier"`
					Rank              interface{} `json:"Rank"`
				} `json:"HighestCsrAttained"`
				ArenaGameBaseVariantStats []struct {
					FlexibleStats struct {
						MedalStatCounts []struct {
							ID    string `json:"Id"`
							Count int    `json:"Count"`
						} `json:"MedalStatCounts"`
						ImpulseStatCounts []struct {
							ID    string `json:"Id"`
							Count int    `json:"Count"`
						} `json:"ImpulseStatCounts"`
						MedalTimelapses   []interface{} `json:"MedalTimelapses"`
						ImpulseTimelapses []struct {
							ID        string `json:"Id"`
							Timelapse string `json:"Timelapse"`
						} `json:"ImpulseTimelapses"`
					} `json:"FlexibleStats"`
					GameBaseVariantID   string  `json:"GameBaseVariantId"`
					TotalKills          int     `json:"TotalKills"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalWeaponDamage   float64 `json:"TotalWeaponDamage"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponWithMostKills struct {
						WeaponID struct {
							StockID     int           `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponWithMostKills"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
					TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					MedalAwards                    []struct {
						MedalID int `json:"MedalId"`
						Count   int `json:"Count"`
					} `json:"MedalAwards"`
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []interface{} `json:"EnemyKills"`
					WeaponStats            []struct {
						WeaponID struct {
							StockID     int64         `json:"StockId"`
							Attachments []interface{} `json:"Attachments"`
						} `json:"WeaponId"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
					} `json:"WeaponStats"`
					Impulses []struct {
						ID    int64 `json:"Id"`
						Count int   `json:"Count"`
					} `json:"Impulses"`
					TotalSpartanKills int    `json:"TotalSpartanKills"`
					FastestMatchWin   string `json:"FastestMatchWin"`
				} `json:"ArenaGameBaseVariantStats"`
				TopGameBaseVariants []struct {
					GameBaseVariantRank      int    `json:"GameBaseVariantRank"`
					NumberOfMatchesCompleted int    `json:"NumberOfMatchesCompleted"`
					GameBaseVariantID        string `json:"GameBaseVariantId"`
					NumberOfMatchesWon       int    `json:"NumberOfMatchesWon"`
				} `json:"TopGameBaseVariants"`
				HighestCsrPlaylistID       string      `json:"HighestCsrPlaylistId"`
				HighestCsrSeasonID         interface{} `json:"HighestCsrSeasonId"`
				ArenaPlaylistStatsSeasonID string      `json:"ArenaPlaylistStatsSeasonId"`
				TotalKills                 int         `json:"TotalKills"`
				TotalHeadshots             int         `json:"TotalHeadshots"`
				TotalWeaponDamage          float64     `json:"TotalWeaponDamage"`
				TotalShotsFired            int         `json:"TotalShotsFired"`
				TotalShotsLanded           int         `json:"TotalShotsLanded"`
				WeaponWithMostKills        struct {
					WeaponID struct {
						StockID     int64         `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponWithMostKills"`
				TotalMeleeKills                int     `json:"TotalMeleeKills"`
				TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
				TotalAssassinations            int     `json:"TotalAssassinations"`
				TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
				TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
				TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
				TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
				TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
				TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
				TotalDeaths                    int     `json:"TotalDeaths"`
				TotalAssists                   int     `json:"TotalAssists"`
				TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
				TotalGamesWon                  int     `json:"TotalGamesWon"`
				TotalGamesLost                 int     `json:"TotalGamesLost"`
				TotalGamesTied                 int     `json:"TotalGamesTied"`
				TotalTimePlayed                string  `json:"TotalTimePlayed"`
				TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
				MedalAwards                    []struct {
					MedalID int `json:"MedalId"`
					Count   int `json:"Count"`
				} `json:"MedalAwards"`
				DestroyedEnemyVehicles []struct {
					Enemy struct {
						BaseID      int64         `json:"BaseId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"DestroyedEnemyVehicles"`
				EnemyKills  []interface{} `json:"EnemyKills"`
				WeaponStats []struct {
					WeaponID struct {
						StockID     int64         `json:"StockId"`
						Attachments []interface{} `json:"Attachments"`
					} `json:"WeaponId"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
				} `json:"WeaponStats"`
				Impulses []struct {
					ID    int64 `json:"Id"`
					Count int   `json:"Count"`
				} `json:"Impulses"`
				TotalSpartanKills int    `json:"TotalSpartanKills"`
				FastestMatchWin   string `json:"FastestMatchWin"`
			} `json:"ArenaStats"`
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			Xp          int `json:"Xp"`
		} `json:"Result"`
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
type SkullsStruct []struct {
	ContentID   string `json:"contentId"`
	Description string `json:"description"`
	ID          string `json:"id"`
	ImageURL    string `json:"imageUrl"`
	MissionID   string `json:"missionId"`
	Name        string `json:"name"`
}

type SpartanRanksStruct []struct {
	ContentID string `json:"contentId"`
	ID        string `json:"id"`
	Reward    struct {
		ContentID        string        `json:"contentId"`
		ID               string        `json:"id"`
		RequisitionPacks []interface{} `json:"requisitionPacks"`
		Xp               int           `json:"xp"`
	} `json:"reward"`
	StartXp int `json:"startXp"`
}

type TeamColorsStruct []struct {
	Color       string      `json:"color"`
	ContentID   string      `json:"contentId"`
	Description interface{} `json:"description"`
	IconURL     string      `json:"iconUrl"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
}

type VehiclesStruct []struct {
	ContentID         string `json:"contentId"`
	Description       string `json:"description"`
	ID                string `json:"id"`
	IsUsableByPlayer  bool   `json:"isUsableByPlayer"`
	LargeIconImageURL string `json:"largeIconImageUrl"`
	Name              string `json:"name"`
	SmallIconImageURL string `json:"smallIconImageUrl"`
}

type WeaponsStruct []struct {
	ContentID         string      `json:"contentId"`
	Description       interface{} `json:"description"`
	ID                string      `json:"id"`
	IsUsableByPlayer  bool        `json:"isUsableByPlayer"`
	LargeIconImageURL string      `json:"largeIconImageUrl"`
	Name              string      `json:"name"`
	SmallIconImageURL string      `json:"smallIconImageUrl"`
	Type              string      `json:"type"`
}
