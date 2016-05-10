type CarnageReportArenaStruct struct {
	GameBaseVariantID     string `json:"GameBaseVariantId"`
	GameVariantID         string `json:"GameVariantId"`
	GameVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"GameVariantResourceId"`
	IsMatchOver          bool   `json:"IsMatchOver"`
	IsTeamGame           bool   `json:"IsTeamGame"`
	MapID                string `json:"MapId"`
	MapVariantID         string `json:"MapVariantId"`
	MapVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"MapVariantResourceId"`
	PlayerStats []struct {
		AvgLifeTimeOfPlayer string `json:"AvgLifeTimeOfPlayer"`
		CreditsEarned       struct {
			BoostAmount               int `json:"BoostAmount"`
			MatchSpeedWinAmount       int `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount int `json:"ObjectivesCompletedAmount"`
			PlayerRankAmount          int `json:"PlayerRankAmount"`
			Result                    int `json:"Result"`
			SpartanRankModifier       int `json:"SpartanRankModifier"`
			TimePlayedAmount          int `json:"TimePlayedAmount"`
			TotalCreditsEarned        int `json:"TotalCreditsEarned"`
		} `json:"CreditsEarned"`
		CurrentCsr struct {
			Csr               int         `json:"Csr"`
			DesignationID     int         `json:"DesignationId"`
			PercentToNextTier int         `json:"PercentToNextTier"`
			Rank              interface{} `json:"Rank"`
			Tier              int         `json:"Tier"`
		} `json:"CurrentCsr"`
		DNF                    bool          `json:"DNF"`
		DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills             []interface{} `json:"EnemyKills"`
		FlexibleStats          struct {
			ImpulseStatCounts []struct {
				Count int    `json:"Count"`
				ID    string `json:"Id"`
			} `json:"ImpulseStatCounts"`
			ImpulseTimelapses []struct {
				ID        string `json:"Id"`
				Timelapse string `json:"Timelapse"`
			} `json:"ImpulseTimelapses"`
			MedalStatCounts []interface{} `json:"MedalStatCounts"`
			MedalTimelapses []interface{} `json:"MedalTimelapses"`
		} `json:"FlexibleStats"`
		Impulses                []interface{} `json:"Impulses"`
		KilledByOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledByOpponentDetails"`
		KilledOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledOpponentDetails"`
		MeasurementMatchesLeft int `json:"MeasurementMatchesLeft"`
		MedalAwards            []struct {
			Count   int `json:"Count"`
			MedalID int `json:"MedalId"`
		} `json:"MedalAwards"`
		MetaCommendationDeltas []interface{} `json:"MetaCommendationDeltas"`
		Player                 struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		PlayerScore      int         `json:"PlayerScore"`
		PostMatchRatings interface{} `json:"PostMatchRatings"`
		PreMatchRatings  interface{} `json:"PreMatchRatings"`
		PreviousCsr      struct {
			Csr               int         `json:"Csr"`
			DesignationID     int         `json:"DesignationId"`
			PercentToNextTier int         `json:"PercentToNextTier"`
			Rank              interface{} `json:"Rank"`
			Tier              int         `json:"Tier"`
		} `json:"PreviousCsr"`
		ProgressiveCommendationDeltas []struct {
			ID               string `json:"Id"`
			PreviousProgress int    `json:"PreviousProgress"`
			Progress         int    `json:"Progress"`
		} `json:"ProgressiveCommendationDeltas"`
		Rank                           int           `json:"Rank"`
		RewardSets                     []interface{} `json:"RewardSets"`
		TeamID                         int           `json:"TeamId"`
		TotalAssassinations            int           `json:"TotalAssassinations"`
		TotalAssists                   int           `json:"TotalAssists"`
		TotalDeaths                    int           `json:"TotalDeaths"`
		TotalGamesCompleted            int           `json:"TotalGamesCompleted"`
		TotalGamesLost                 int           `json:"TotalGamesLost"`
		TotalGamesTied                 int           `json:"TotalGamesTied"`
		TotalGamesWon                  int           `json:"TotalGamesWon"`
		TotalGrenadeDamage             int           `json:"TotalGrenadeDamage"`
		TotalGrenadeKills              int           `json:"TotalGrenadeKills"`
		TotalGroundPoundDamage         int           `json:"TotalGroundPoundDamage"`
		TotalGroundPoundKills          int           `json:"TotalGroundPoundKills"`
		TotalHeadshots                 int           `json:"TotalHeadshots"`
		TotalKills                     int           `json:"TotalKills"`
		TotalMeleeDamage               int           `json:"TotalMeleeDamage"`
		TotalMeleeKills                int           `json:"TotalMeleeKills"`
		TotalPowerWeaponDamage         int           `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int           `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponKills          int           `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponPossessionTime string        `json:"TotalPowerWeaponPossessionTime"`
		TotalShotsFired                int           `json:"TotalShotsFired"`
		TotalShotsLanded               int           `json:"TotalShotsLanded"`
		TotalShoulderBashDamage        int           `json:"TotalShoulderBashDamage"`
		TotalShoulderBashKills         int           `json:"TotalShoulderBashKills"`
		TotalSpartanKills              int           `json:"TotalSpartanKills"`
		TotalTimePlayed                string        `json:"TotalTimePlayed"`
		TotalWeaponDamage              int           `json:"TotalWeaponDamage"`
		WeaponStats                    []struct {
			TotalDamageDealt    int    `json:"TotalDamageDealt"`
			TotalHeadshots      int    `json:"TotalHeadshots"`
			TotalKills          int    `json:"TotalKills"`
			TotalPossessionTime string `json:"TotalPossessionTime"`
			TotalShotsFired     int    `json:"TotalShotsFired"`
			TotalShotsLanded    int    `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponStats"`
		WeaponWithMostKills struct {
			TotalDamageDealt    int    `json:"TotalDamageDealt"`
			TotalHeadshots      int    `json:"TotalHeadshots"`
			TotalKills          int    `json:"TotalKills"`
			TotalPossessionTime string `json:"TotalPossessionTime"`
			TotalShotsFired     int    `json:"TotalShotsFired"`
			TotalShotsLanded    int    `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponWithMostKills"`
		XpInfo struct {
			BoostAmount                  int `json:"BoostAmount"`
			MatchSpeedWinAmount          int `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount    int `json:"ObjectivesCompletedAmount"`
			PerformanceXP                int `json:"PerformanceXP"`
			PlayerRankXPAward            int `json:"PlayerRankXPAward"`
			PlayerTimePerformanceXPAward int `json:"PlayerTimePerformanceXPAward"`
			PrevSpartanRank              int `json:"PrevSpartanRank"`
			PrevTotalXP                  int `json:"PrevTotalXP"`
			SpartanRank                  int `json:"SpartanRank"`
			SpartanRankMatchXPScalar     int `json:"SpartanRankMatchXPScalar"`
			TotalXP                      int `json:"TotalXP"`
		} `json:"XpInfo"`
	} `json:"PlayerStats"`
	PlaylistID string `json:"PlaylistId"`
	SeasonID   string `json:"SeasonId"`
	TeamStats  []struct {
		Rank       int `json:"Rank"`
		RoundStats []struct {
			Rank        int `json:"Rank"`
			RoundNumber int `json:"RoundNumber"`
			Score       int `json:"Score"`
		} `json:"RoundStats"`
		Score  int `json:"Score"`
		TeamID int `json:"TeamId"`
	} `json:"TeamStats"`
	TotalDuration string `json:"TotalDuration"`
}

type CarnageReportCampaignStruct struct {
	Difficulty            int    `json:"Difficulty"`
	GameBaseVariantID     string `json:"GameBaseVariantId"`
	GameVariantID         string `json:"GameVariantId"`
	GameVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"GameVariantResourceId"`
	IsMatchOver          bool   `json:"IsMatchOver"`
	IsTeamGame           bool   `json:"IsTeamGame"`
	MapID                string `json:"MapId"`
	MapVariantID         string `json:"MapVariantId"`
	MapVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"MapVariantResourceId"`
	MissionCompleted bool `json:"MissionCompleted"`
	PlayerStats      []struct {
		AvgLifeTimeOfPlayer    string        `json:"AvgLifeTimeOfPlayer"`
		BiggestKillScore       int           `json:"BiggestKillScore"`
		CharacterIndex         interface{}   `json:"CharacterIndex"`
		DNF                    bool          `json:"DNF"`
		DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills             []struct {
			Enemy struct {
				Attachments []interface{} `json:"Attachments"`
				BaseID      int           `json:"BaseId"`
			} `json:"Enemy"`
			TotalKills int `json:"TotalKills"`
		} `json:"EnemyKills"`
		FlexibleStats struct {
			ImpulseStatCounts []struct {
				Count int    `json:"Count"`
				ID    string `json:"Id"`
			} `json:"ImpulseStatCounts"`
			ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
			MedalStatCounts   []interface{} `json:"MedalStatCounts"`
			MedalTimelapses   []interface{} `json:"MedalTimelapses"`
		} `json:"FlexibleStats"`
		Impulses    []interface{} `json:"Impulses"`
		MedalAwards []interface{} `json:"MedalAwards"`
		Player      struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		PlayerScore                    interface{} `json:"PlayerScore"`
		PostMatchRatings               interface{} `json:"PostMatchRatings"`
		PreMatchRatings                interface{} `json:"PreMatchRatings"`
		Rank                           int         `json:"Rank"`
		Score                          int         `json:"Score"`
		TeamID                         int         `json:"TeamId"`
		TotalAssassinations            int         `json:"TotalAssassinations"`
		TotalAssists                   int         `json:"TotalAssists"`
		TotalDeaths                    int         `json:"TotalDeaths"`
		TotalGamesCompleted            int         `json:"TotalGamesCompleted"`
		TotalGamesLost                 int         `json:"TotalGamesLost"`
		TotalGamesTied                 int         `json:"TotalGamesTied"`
		TotalGamesWon                  int         `json:"TotalGamesWon"`
		TotalGrenadeDamage             int         `json:"TotalGrenadeDamage"`
		TotalGrenadeKills              int         `json:"TotalGrenadeKills"`
		TotalGroundPoundDamage         int         `json:"TotalGroundPoundDamage"`
		TotalGroundPoundKills          int         `json:"TotalGroundPoundKills"`
		TotalHeadshots                 int         `json:"TotalHeadshots"`
		TotalKills                     int         `json:"TotalKills"`
		TotalMeleeDamage               int         `json:"TotalMeleeDamage"`
		TotalMeleeKills                int         `json:"TotalMeleeKills"`
		TotalPowerWeaponDamage         int         `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int         `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponKills          int         `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponPossessionTime string      `json:"TotalPowerWeaponPossessionTime"`
		TotalShotsFired                int         `json:"TotalShotsFired"`
		TotalShotsLanded               int         `json:"TotalShotsLanded"`
		TotalShoulderBashDamage        int         `json:"TotalShoulderBashDamage"`
		TotalShoulderBashKills         int         `json:"TotalShoulderBashKills"`
		TotalSpartanKills              int         `json:"TotalSpartanKills"`
		TotalTimePlayed                string      `json:"TotalTimePlayed"`
		TotalWeaponDamage              int         `json:"TotalWeaponDamage"`
		WeaponStats                    []struct {
			TotalDamageDealt    int    `json:"TotalDamageDealt"`
			TotalHeadshots      int    `json:"TotalHeadshots"`
			TotalKills          int    `json:"TotalKills"`
			TotalPossessionTime string `json:"TotalPossessionTime"`
			TotalShotsFired     int    `json:"TotalShotsFired"`
			TotalShotsLanded    int    `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponStats"`
		WeaponWithMostKills struct {
			TotalDamageDealt    int    `json:"TotalDamageDealt"`
			TotalHeadshots      int    `json:"TotalHeadshots"`
			TotalKills          int    `json:"TotalKills"`
			TotalPossessionTime string `json:"TotalPossessionTime"`
			TotalShotsFired     int    `json:"TotalShotsFired"`
			TotalShotsLanded    int    `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponWithMostKills"`
	} `json:"PlayerStats"`
	PlaylistID                  string        `json:"PlaylistId"`
	SeasonID                    interface{}   `json:"SeasonId"`
	Skulls                      []interface{} `json:"Skulls"`
	TotalDuration               string        `json:"TotalDuration"`
	TotalMissionPlaythroughTime string        `json:"TotalMissionPlaythroughTime"`
}

type CarnageReportCustomStruct struct {
	GameBaseVariantID     string `json:"GameBaseVariantId"`
	GameVariantID         string `json:"GameVariantId"`
	GameVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"GameVariantResourceId"`
	IsMatchOver          bool   `json:"IsMatchOver"`
	IsTeamGame           bool   `json:"IsTeamGame"`
	MapID                string `json:"MapId"`
	MapVariantID         string `json:"MapVariantId"`
	MapVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"MapVariantResourceId"`
	PlayerStats []struct {
		AvgLifeTimeOfPlayer    string        `json:"AvgLifeTimeOfPlayer"`
		DNF                    bool          `json:"DNF"`
		DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills             []interface{} `json:"EnemyKills"`
		FlexibleStats          struct {
			ImpulseStatCounts []struct {
				Count int    `json:"Count"`
				ID    string `json:"Id"`
			} `json:"ImpulseStatCounts"`
			ImpulseTimelapses []struct {
				ID        string `json:"Id"`
				Timelapse string `json:"Timelapse"`
			} `json:"ImpulseTimelapses"`
			MedalStatCounts []interface{} `json:"MedalStatCounts"`
			MedalTimelapses []interface{} `json:"MedalTimelapses"`
		} `json:"FlexibleStats"`
		Impulses []struct {
			Count int `json:"Count"`
			ID    int `json:"Id"`
		} `json:"Impulses"`
		KilledByOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledByOpponentDetails"`
		KilledOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledOpponentDetails"`
		MedalAwards []struct {
			Count   int `json:"Count"`
			MedalID int `json:"MedalId"`
		} `json:"MedalAwards"`
		Player struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		PlayerScore                    int         `json:"PlayerScore"`
		PostMatchRatings               interface{} `json:"PostMatchRatings"`
		PreMatchRatings                interface{} `json:"PreMatchRatings"`
		Rank                           int         `json:"Rank"`
		TeamID                         int         `json:"TeamId"`
		TotalAssassinations            int         `json:"TotalAssassinations"`
		TotalAssists                   int         `json:"TotalAssists"`
		TotalDeaths                    int         `json:"TotalDeaths"`
		TotalGamesCompleted            int         `json:"TotalGamesCompleted"`
		TotalGamesLost                 int         `json:"TotalGamesLost"`
		TotalGamesTied                 int         `json:"TotalGamesTied"`
		TotalGamesWon                  int         `json:"TotalGamesWon"`
		TotalGrenadeDamage             int         `json:"TotalGrenadeDamage"`
		TotalGrenadeKills              int         `json:"TotalGrenadeKills"`
		TotalGroundPoundDamage         int         `json:"TotalGroundPoundDamage"`
		TotalGroundPoundKills          int         `json:"TotalGroundPoundKills"`
		TotalHeadshots                 int         `json:"TotalHeadshots"`
		TotalKills                     int         `json:"TotalKills"`
		TotalMeleeDamage               float64     `json:"TotalMeleeDamage"`
		TotalMeleeKills                int         `json:"TotalMeleeKills"`
		TotalPowerWeaponDamage         float64     `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int         `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponKills          int         `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponPossessionTime string      `json:"TotalPowerWeaponPossessionTime"`
		TotalShotsFired                int         `json:"TotalShotsFired"`
		TotalShotsLanded               int         `json:"TotalShotsLanded"`
		TotalShoulderBashDamage        int         `json:"TotalShoulderBashDamage"`
		TotalShoulderBashKills         int         `json:"TotalShoulderBashKills"`
		TotalSpartanKills              int         `json:"TotalSpartanKills"`
		TotalTimePlayed                string      `json:"TotalTimePlayed"`
		TotalWeaponDamage              float64     `json:"TotalWeaponDamage"`
		WeaponStats                    []struct {
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponStats"`
		WeaponWithMostKills struct {
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponWithMostKills"`
	} `json:"PlayerStats"`
	PlaylistID string      `json:"PlaylistId"`
	SeasonID   interface{} `json:"SeasonId"`
	TeamStats  []struct {
		Rank       int           `json:"Rank"`
		RoundStats []interface{} `json:"RoundStats"`
		Score      int           `json:"Score"`
		TeamID     int           `json:"TeamId"`
	} `json:"TeamStats"`
	TotalDuration string `json:"TotalDuration"`
}

type CarnageReportWarzoneStruct struct {
	GameBaseVariantID     string `json:"GameBaseVariantId"`
	GameVariantID         string `json:"GameVariantId"`
	GameVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"GameVariantResourceId"`
	IsMatchOver          bool   `json:"IsMatchOver"`
	IsTeamGame           bool   `json:"IsTeamGame"`
	MapID                string `json:"MapId"`
	MapVariantID         string `json:"MapVariantId"`
	MapVariantResourceID struct {
		Owner        string `json:"Owner"`
		OwnerType    int    `json:"OwnerType"`
		ResourceID   string `json:"ResourceId"`
		ResourceType int    `json:"ResourceType"`
	} `json:"MapVariantResourceId"`
	PlayerStats []struct {
		AvgLifeTimeOfPlayer string `json:"AvgLifeTimeOfPlayer"`
		CreditsEarned       struct {
			BoostAmount               int `json:"BoostAmount"`
			MatchSpeedWinAmount       int `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount int `json:"ObjectivesCompletedAmount"`
			PlayerRankAmount          int `json:"PlayerRankAmount"`
			Result                    int `json:"Result"`
			SpartanRankModifier       int `json:"SpartanRankModifier"`
			TimePlayedAmount          int `json:"TimePlayedAmount"`
			TotalCreditsEarned        int `json:"TotalCreditsEarned"`
		} `json:"CreditsEarned"`
		DNF                    bool          `json:"DNF"`
		DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
		EnemyKills             []struct {
			Enemy struct {
				Attachments []interface{} `json:"Attachments"`
				BaseID      int           `json:"BaseId"`
			} `json:"Enemy"`
			TotalKills int `json:"TotalKills"`
		} `json:"EnemyKills"`
		FlexibleStats struct {
			ImpulseStatCounts []struct {
				Count int    `json:"Count"`
				ID    string `json:"Id"`
			} `json:"ImpulseStatCounts"`
			ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
			MedalStatCounts   []struct {
				Count int    `json:"Count"`
				ID    string `json:"Id"`
			} `json:"MedalStatCounts"`
			MedalTimelapses []interface{} `json:"MedalTimelapses"`
		} `json:"FlexibleStats"`
		Impulses []struct {
			Count int `json:"Count"`
			ID    int `json:"Id"`
		} `json:"Impulses"`
		KilledByOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledByOpponentDetails"`
		KilledOpponentDetails []struct {
			GamerTag   string `json:"GamerTag"`
			TotalKills int    `json:"TotalKills"`
		} `json:"KilledOpponentDetails"`
		MedalAwards []struct {
			Count   int `json:"Count"`
			MedalID int `json:"MedalId"`
		} `json:"MedalAwards"`
		MetaCommendationDeltas []interface{} `json:"MetaCommendationDeltas"`
		Player                 struct {
			Gamertag string      `json:"Gamertag"`
			Xuid     interface{} `json:"Xuid"`
		} `json:"Player"`
		PlayerScore                   int         `json:"PlayerScore"`
		PostMatchRatings              interface{} `json:"PostMatchRatings"`
		PreMatchRatings               interface{} `json:"PreMatchRatings"`
		ProgressiveCommendationDeltas []struct {
			ID               string `json:"Id"`
			PreviousProgress int    `json:"PreviousProgress"`
			Progress         int    `json:"Progress"`
		} `json:"ProgressiveCommendationDeltas"`
		Rank       int `json:"Rank"`
		RewardSets []struct {
			CommendationLevelID string      `json:"CommendationLevelId"`
			CommendationSource  string      `json:"CommendationSource"`
			RewardSet           string      `json:"RewardSet"`
			RewardSourceType    int         `json:"RewardSourceType"`
			SpartanRankSource   interface{} `json:"SpartanRankSource"`
		} `json:"RewardSets"`
		TeamID                         int     `json:"TeamId"`
		TotalAssassinations            int     `json:"TotalAssassinations"`
		TotalAssists                   int     `json:"TotalAssists"`
		TotalDeaths                    int     `json:"TotalDeaths"`
		TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
		TotalGamesLost                 int     `json:"TotalGamesLost"`
		TotalGamesTied                 int     `json:"TotalGamesTied"`
		TotalGamesWon                  int     `json:"TotalGamesWon"`
		TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
		TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
		TotalGroundPoundDamage         int     `json:"TotalGroundPoundDamage"`
		TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
		TotalHeadshots                 int     `json:"TotalHeadshots"`
		TotalKills                     int     `json:"TotalKills"`
		TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
		TotalMeleeKills                int     `json:"TotalMeleeKills"`
		TotalPiesEarned                int     `json:"TotalPiesEarned"`
		TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
		TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
		TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
		TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
		TotalShotsFired                int     `json:"TotalShotsFired"`
		TotalShotsLanded               int     `json:"TotalShotsLanded"`
		TotalShoulderBashDamage        int     `json:"TotalShoulderBashDamage"`
		TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
		TotalSpartanKills              int     `json:"TotalSpartanKills"`
		TotalTimePlayed                string  `json:"TotalTimePlayed"`
		TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
		WarzoneLevel                   int     `json:"WarzoneLevel"`
		WeaponStats                    []struct {
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponStats"`
		WeaponWithMostKills struct {
			TotalDamageDealt    float64 `json:"TotalDamageDealt"`
			TotalHeadshots      int     `json:"TotalHeadshots"`
			TotalKills          int     `json:"TotalKills"`
			TotalPossessionTime string  `json:"TotalPossessionTime"`
			TotalShotsFired     int     `json:"TotalShotsFired"`
			TotalShotsLanded    int     `json:"TotalShotsLanded"`
			WeaponID            struct {
				Attachments []interface{} `json:"Attachments"`
				StockID     int           `json:"StockId"`
			} `json:"WeaponId"`
		} `json:"WeaponWithMostKills"`
		XpInfo struct {
			BoostAmount                  int     `json:"BoostAmount"`
			MatchSpeedWinAmount          int     `json:"MatchSpeedWinAmount"`
			ObjectivesCompletedAmount    int     `json:"ObjectivesCompletedAmount"`
			PerformanceXP                int     `json:"PerformanceXP"`
			PlayerRankXPAward            int     `json:"PlayerRankXPAward"`
			PlayerTimePerformanceXPAward int     `json:"PlayerTimePerformanceXPAward"`
			PrevSpartanRank              int     `json:"PrevSpartanRank"`
			PrevTotalXP                  int     `json:"PrevTotalXP"`
			SpartanRank                  int     `json:"SpartanRank"`
			SpartanRankMatchXPScalar     float64 `json:"SpartanRankMatchXPScalar"`
			TotalXP                      int     `json:"TotalXP"`
		} `json:"XpInfo"`
	} `json:"PlayerStats"`
	PlaylistID string `json:"PlaylistId"`
	SeasonID   string `json:"SeasonId"`
	TeamStats  []struct {
		Rank       int `json:"Rank"`
		RoundStats []struct {
			Rank        int `json:"Rank"`
			RoundNumber int `json:"RoundNumber"`
			Score       int `json:"Score"`
		} `json:"RoundStats"`
		Score  int `json:"Score"`
		TeamID int `json:"TeamId"`
	} `json:"TeamStats"`
	TotalDuration string `json:"TotalDuration"`
}

type EnemiesStruct []interface{}

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

type FlexibleStatsStruct []interface{}

type GameBaseVariantsStruct []interface{}

type GameVariantsStruct struct {
	ContentID         string `json:"contentId"`
	Description       string `json:"description"`
	GameBaseVariantID string `json:"gameBaseVariantId"`
	IconURL           string `json:"iconUrl"`
	ID                string `json:"id"`
	Name              string `json:"name"`
}

type ImpulsesStruct []interface{}

type MapsStruct []interface{}

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

type MedalsStruct []interface{}

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

type PlaylistsStruct []interface{}

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

type SeasonsStruct []interface{}

type ServiceRecordArenaStruct struct {
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
	Results []struct {
		ID     string `json:"Id"`
		Result struct {
			ArenaStats struct {
				ArenaGameBaseVariantStats []struct {
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []interface{} `json:"EnemyKills"`
					FlexibleStats          struct {
						ImpulseStatCounts []struct {
							Count int    `json:"Count"`
							ID    string `json:"Id"`
						} `json:"ImpulseStatCounts"`
						ImpulseTimelapses []struct {
							ID        string `json:"Id"`
							Timelapse string `json:"Timelapse"`
						} `json:"ImpulseTimelapses"`
						MedalStatCounts []struct {
							Count int    `json:"Count"`
							ID    string `json:"Id"`
						} `json:"MedalStatCounts"`
						MedalTimelapses []interface{} `json:"MedalTimelapses"`
					} `json:"FlexibleStats"`
					GameBaseVariantID string `json:"GameBaseVariantId"`
					Impulses          []struct {
						Count int `json:"Count"`
						ID    int `json:"Id"`
					} `json:"Impulses"`
					MedalAwards []struct {
						Count   int `json:"Count"`
						MedalID int `json:"MedalId"`
					} `json:"MedalAwards"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					TotalGroundPoundDamage         int     `json:"TotalGroundPoundDamage"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalHeadshots                 int     `json:"TotalHeadshots"`
					TotalKills                     int     `json:"TotalKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalShotsFired                int     `json:"TotalShotsFired"`
					TotalShotsLanded               int     `json:"TotalShotsLanded"`
					TotalShoulderBashDamage        int     `json:"TotalShoulderBashDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalSpartanKills              int     `json:"TotalSpartanKills"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
					WeaponStats                    []struct {
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponStats"`
					WeaponWithMostKills struct {
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponWithMostKills"`
				} `json:"ArenaGameBaseVariantStats"`
				ArenaPlaylistStats []struct {
					Csr                    interface{}   `json:"Csr"`
					CsrPercentile          interface{}   `json:"CsrPercentile"`
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []interface{} `json:"EnemyKills"`
					HighestCsr             interface{}   `json:"HighestCsr"`
					Impulses               []interface{} `json:"Impulses"`
					MeasurementMatchesLeft int           `json:"MeasurementMatchesLeft"`
					MedalAwards            []struct {
						Count   int `json:"Count"`
						MedalID int `json:"MedalId"`
					} `json:"MedalAwards"`
					PlaylistID                     string `json:"PlaylistId"`
					TotalAssassinations            int    `json:"TotalAssassinations"`
					TotalAssists                   int    `json:"TotalAssists"`
					TotalDeaths                    int    `json:"TotalDeaths"`
					TotalGamesCompleted            int    `json:"TotalGamesCompleted"`
					TotalGamesLost                 int    `json:"TotalGamesLost"`
					TotalGamesTied                 int    `json:"TotalGamesTied"`
					TotalGamesWon                  int    `json:"TotalGamesWon"`
					TotalGrenadeDamage             int    `json:"TotalGrenadeDamage"`
					TotalGrenadeKills              int    `json:"TotalGrenadeKills"`
					TotalGroundPoundDamage         int    `json:"TotalGroundPoundDamage"`
					TotalGroundPoundKills          int    `json:"TotalGroundPoundKills"`
					TotalHeadshots                 int    `json:"TotalHeadshots"`
					TotalKills                     int    `json:"TotalKills"`
					TotalMeleeDamage               int    `json:"TotalMeleeDamage"`
					TotalMeleeKills                int    `json:"TotalMeleeKills"`
					TotalPowerWeaponDamage         int    `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int    `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponKills          int    `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponPossessionTime string `json:"TotalPowerWeaponPossessionTime"`
					TotalShotsFired                int    `json:"TotalShotsFired"`
					TotalShotsLanded               int    `json:"TotalShotsLanded"`
					TotalShoulderBashDamage        int    `json:"TotalShoulderBashDamage"`
					TotalShoulderBashKills         int    `json:"TotalShoulderBashKills"`
					TotalSpartanKills              int    `json:"TotalSpartanKills"`
					TotalTimePlayed                string `json:"TotalTimePlayed"`
					TotalWeaponDamage              int    `json:"TotalWeaponDamage"`
					WeaponStats                    []struct {
						TotalDamageDealt    int    `json:"TotalDamageDealt"`
						TotalHeadshots      int    `json:"TotalHeadshots"`
						TotalKills          int    `json:"TotalKills"`
						TotalPossessionTime string `json:"TotalPossessionTime"`
						TotalShotsFired     int    `json:"TotalShotsFired"`
						TotalShotsLanded    int    `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponStats"`
					WeaponWithMostKills struct {
						TotalDamageDealt    int    `json:"TotalDamageDealt"`
						TotalHeadshots      int    `json:"TotalHeadshots"`
						TotalKills          int    `json:"TotalKills"`
						TotalPossessionTime string `json:"TotalPossessionTime"`
						TotalShotsFired     int    `json:"TotalShotsFired"`
						TotalShotsLanded    int    `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponWithMostKills"`
				} `json:"ArenaPlaylistStats"`
				ArenaPlaylistStatsSeasonID string `json:"ArenaPlaylistStatsSeasonId"`
				DestroyedEnemyVehicles     []struct {
					Enemy struct {
						Attachments []interface{} `json:"Attachments"`
						BaseID      int           `json:"BaseId"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"DestroyedEnemyVehicles"`
				EnemyKills         []interface{} `json:"EnemyKills"`
				HighestCsrAttained struct {
					Csr               int         `json:"Csr"`
					DesignationID     int         `json:"DesignationId"`
					PercentToNextTier int         `json:"PercentToNextTier"`
					Rank              interface{} `json:"Rank"`
					Tier              int         `json:"Tier"`
				} `json:"HighestCsrAttained"`
				HighestCsrPlaylistID string      `json:"HighestCsrPlaylistId"`
				HighestCsrSeasonID   interface{} `json:"HighestCsrSeasonId"`
				Impulses             []struct {
					Count int `json:"Count"`
					ID    int `json:"Id"`
				} `json:"Impulses"`
				MedalAwards []struct {
					Count   int `json:"Count"`
					MedalID int `json:"MedalId"`
				} `json:"MedalAwards"`
				TopGameBaseVariants []struct {
					GameBaseVariantID        string `json:"GameBaseVariantId"`
					GameBaseVariantRank      int    `json:"GameBaseVariantRank"`
					NumberOfMatchesCompleted int    `json:"NumberOfMatchesCompleted"`
					NumberOfMatchesWon       int    `json:"NumberOfMatchesWon"`
				} `json:"TopGameBaseVariants"`
				TotalAssassinations            int     `json:"TotalAssassinations"`
				TotalAssists                   int     `json:"TotalAssists"`
				TotalDeaths                    int     `json:"TotalDeaths"`
				TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
				TotalGamesLost                 int     `json:"TotalGamesLost"`
				TotalGamesTied                 int     `json:"TotalGamesTied"`
				TotalGamesWon                  int     `json:"TotalGamesWon"`
				TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
				TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
				TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
				TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
				TotalHeadshots                 int     `json:"TotalHeadshots"`
				TotalKills                     int     `json:"TotalKills"`
				TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
				TotalMeleeKills                int     `json:"TotalMeleeKills"`
				TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
				TotalShotsFired                int     `json:"TotalShotsFired"`
				TotalShotsLanded               int     `json:"TotalShotsLanded"`
				TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
				TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
				TotalSpartanKills              int     `json:"TotalSpartanKills"`
				TotalTimePlayed                string  `json:"TotalTimePlayed"`
				TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
				WeaponStats                    []struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponStats"`
				WeaponWithMostKills struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponWithMostKills"`
			} `json:"ArenaStats"`
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			Xp          int `json:"Xp"`
		} `json:"Result"`
		ResultCode int `json:"ResultCode"`
	} `json:"Results"`
}

type ServiceRecordCampaignStruct struct {
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
	Results []struct {
		ID     string `json:"Id"`
		Result struct {
			CampaignStat struct {
				CampaignMissionStats []struct {
					CoopStats struct {
						Three struct {
							AllSkullsOn           bool          `json:"AllSkullsOn"`
							FastestCompletionTime string        `json:"FastestCompletionTime"`
							HighestScore          int           `json:"HighestScore"`
							Skulls                []interface{} `json:"Skulls"`
							TotalTimesCompleted   int           `json:"TotalTimesCompleted"`
						} `json:"3"`
					} `json:"CoopStats"`
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []struct {
						Enemy struct {
							Attachments []interface{} `json:"Attachments"`
							BaseID      int           `json:"BaseId"`
						} `json:"Enemy"`
						TotalKills int `json:"TotalKills"`
					} `json:"EnemyKills"`
					FlexibleStats struct {
						ImpulseStatCounts []struct {
							Count int    `json:"Count"`
							ID    string `json:"Id"`
						} `json:"ImpulseStatCounts"`
						ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
						MedalStatCounts   []interface{} `json:"MedalStatCounts"`
						MedalTimelapses   []interface{} `json:"MedalTimelapses"`
					} `json:"FlexibleStats"`
					Impulses []struct {
						Count int `json:"Count"`
						ID    int `json:"Id"`
					} `json:"Impulses"`
					MedalAwards []interface{} `json:"MedalAwards"`
					MissionID   string        `json:"MissionId"`
					SoloStats   struct {
						Two struct {
							AllSkullsOn           bool          `json:"AllSkullsOn"`
							FastestCompletionTime string        `json:"FastestCompletionTime"`
							HighestScore          int           `json:"HighestScore"`
							Skulls                []interface{} `json:"Skulls"`
							TotalTimesCompleted   int           `json:"TotalTimesCompleted"`
						} `json:"2"`
					} `json:"SoloStats"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					TotalGroundPoundDamage         int     `json:"TotalGroundPoundDamage"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalHeadshots                 int     `json:"TotalHeadshots"`
					TotalKills                     int     `json:"TotalKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalShotsFired                int     `json:"TotalShotsFired"`
					TotalShotsLanded               int     `json:"TotalShotsLanded"`
					TotalShoulderBashDamage        int     `json:"TotalShoulderBashDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalSpartanKills              int     `json:"TotalSpartanKills"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
					WeaponStats                    []struct {
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponStats"`
					WeaponWithMostKills struct {
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponWithMostKills"`
				} `json:"CampaignMissionStats"`
				DestroyedEnemyVehicles []struct {
					Enemy struct {
						Attachments []interface{} `json:"Attachments"`
						BaseID      int           `json:"BaseId"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"DestroyedEnemyVehicles"`
				EnemyKills []struct {
					Enemy struct {
						Attachments []interface{} `json:"Attachments"`
						BaseID      int           `json:"BaseId"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"EnemyKills"`
				Impulses []struct {
					Count int `json:"Count"`
					ID    int `json:"Id"`
				} `json:"Impulses"`
				MedalAwards                    []interface{} `json:"MedalAwards"`
				TotalAssassinations            int           `json:"TotalAssassinations"`
				TotalAssists                   int           `json:"TotalAssists"`
				TotalDeaths                    int           `json:"TotalDeaths"`
				TotalGamesCompleted            int           `json:"TotalGamesCompleted"`
				TotalGamesLost                 int           `json:"TotalGamesLost"`
				TotalGamesTied                 int           `json:"TotalGamesTied"`
				TotalGamesWon                  int           `json:"TotalGamesWon"`
				TotalGrenadeDamage             float64       `json:"TotalGrenadeDamage"`
				TotalGrenadeKills              int           `json:"TotalGrenadeKills"`
				TotalGroundPoundDamage         float64       `json:"TotalGroundPoundDamage"`
				TotalGroundPoundKills          int           `json:"TotalGroundPoundKills"`
				TotalHeadshots                 int           `json:"TotalHeadshots"`
				TotalKills                     int           `json:"TotalKills"`
				TotalMeleeDamage               float64       `json:"TotalMeleeDamage"`
				TotalMeleeKills                int           `json:"TotalMeleeKills"`
				TotalPowerWeaponDamage         float64       `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int           `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponKills          int           `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponPossessionTime string        `json:"TotalPowerWeaponPossessionTime"`
				TotalShotsFired                int           `json:"TotalShotsFired"`
				TotalShotsLanded               int           `json:"TotalShotsLanded"`
				TotalShoulderBashDamage        float64       `json:"TotalShoulderBashDamage"`
				TotalShoulderBashKills         int           `json:"TotalShoulderBashKills"`
				TotalSpartanKills              int           `json:"TotalSpartanKills"`
				TotalTimePlayed                string        `json:"TotalTimePlayed"`
				TotalWeaponDamage              float64       `json:"TotalWeaponDamage"`
				WeaponStats                    []struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponStats"`
				WeaponWithMostKills struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponWithMostKills"`
			} `json:"CampaignStat"`
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			Xp          int `json:"Xp"`
		} `json:"Result"`
		ResultCode int `json:"ResultCode"`
	} `json:"Results"`
}

type ServiceRecordCustomStruct struct {
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
	Results []struct {
		ID     string `json:"Id"`
		Result struct {
			CustomStats struct {
				CustomGameBaseVariantStats []struct {
					DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
					EnemyKills             []interface{} `json:"EnemyKills"`
					FlexibleStats          struct {
						ImpulseStatCounts []struct {
							Count int    `json:"Count"`
							ID    string `json:"Id"`
						} `json:"ImpulseStatCounts"`
						ImpulseTimelapses []struct {
							ID        string `json:"Id"`
							Timelapse string `json:"Timelapse"`
						} `json:"ImpulseTimelapses"`
						MedalStatCounts []interface{} `json:"MedalStatCounts"`
						MedalTimelapses []interface{} `json:"MedalTimelapses"`
					} `json:"FlexibleStats"`
					GameBaseVariantID string `json:"GameBaseVariantId"`
					Impulses          []struct {
						Count int `json:"Count"`
						ID    int `json:"Id"`
					} `json:"Impulses"`
					MedalAwards []struct {
						Count   int `json:"Count"`
						MedalID int `json:"MedalId"`
					} `json:"MedalAwards"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGrenadeDamage             int     `json:"TotalGrenadeDamage"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					TotalGroundPoundDamage         int     `json:"TotalGroundPoundDamage"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalHeadshots                 int     `json:"TotalHeadshots"`
					TotalKills                     int     `json:"TotalKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalShotsFired                int     `json:"TotalShotsFired"`
					TotalShotsLanded               int     `json:"TotalShotsLanded"`
					TotalShoulderBashDamage        int     `json:"TotalShoulderBashDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalSpartanKills              int     `json:"TotalSpartanKills"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
					WeaponStats                    []struct {
						TotalDamageDealt    int    `json:"TotalDamageDealt"`
						TotalHeadshots      int    `json:"TotalHeadshots"`
						TotalKills          int    `json:"TotalKills"`
						TotalPossessionTime string `json:"TotalPossessionTime"`
						TotalShotsFired     int    `json:"TotalShotsFired"`
						TotalShotsLanded    int    `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponStats"`
					WeaponWithMostKills struct {
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponWithMostKills"`
				} `json:"CustomGameBaseVariantStats"`
				DestroyedEnemyVehicles []interface{} `json:"DestroyedEnemyVehicles"`
				EnemyKills             []interface{} `json:"EnemyKills"`
				Impulses               []struct {
					Count int `json:"Count"`
					ID    int `json:"Id"`
				} `json:"Impulses"`
				MedalAwards []struct {
					Count   int `json:"Count"`
					MedalID int `json:"MedalId"`
				} `json:"MedalAwards"`
				TopGameBaseVariants []struct {
					GameBaseVariantID        string `json:"GameBaseVariantId"`
					GameBaseVariantRank      int    `json:"GameBaseVariantRank"`
					NumberOfMatchesCompleted int    `json:"NumberOfMatchesCompleted"`
					NumberOfMatchesWon       int    `json:"NumberOfMatchesWon"`
				} `json:"TopGameBaseVariants"`
				TotalAssassinations            int     `json:"TotalAssassinations"`
				TotalAssists                   int     `json:"TotalAssists"`
				TotalDeaths                    int     `json:"TotalDeaths"`
				TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
				TotalGamesLost                 int     `json:"TotalGamesLost"`
				TotalGamesTied                 int     `json:"TotalGamesTied"`
				TotalGamesWon                  int     `json:"TotalGamesWon"`
				TotalGrenadeDamage             int     `json:"TotalGrenadeDamage"`
				TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
				TotalGroundPoundDamage         int     `json:"TotalGroundPoundDamage"`
				TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
				TotalHeadshots                 int     `json:"TotalHeadshots"`
				TotalKills                     int     `json:"TotalKills"`
				TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
				TotalMeleeKills                int     `json:"TotalMeleeKills"`
				TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
				TotalShotsFired                int     `json:"TotalShotsFired"`
				TotalShotsLanded               int     `json:"TotalShotsLanded"`
				TotalShoulderBashDamage        int     `json:"TotalShoulderBashDamage"`
				TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
				TotalSpartanKills              int     `json:"TotalSpartanKills"`
				TotalTimePlayed                string  `json:"TotalTimePlayed"`
				TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
				WeaponStats                    []struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponStats"`
				WeaponWithMostKills struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponWithMostKills"`
			} `json:"CustomStats"`
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			Xp          int `json:"Xp"`
		} `json:"Result"`
		ResultCode int `json:"ResultCode"`
	} `json:"Results"`
}

type ServiceRecordWarzoneStruct struct {
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
	Results []struct {
		ID     string `json:"Id"`
		Result struct {
			PlayerID struct {
				Gamertag string      `json:"Gamertag"`
				Xuid     interface{} `json:"Xuid"`
			} `json:"PlayerId"`
			SpartanRank int `json:"SpartanRank"`
			WarzoneStat struct {
				DestroyedEnemyVehicles []struct {
					Enemy struct {
						Attachments []interface{} `json:"Attachments"`
						BaseID      int           `json:"BaseId"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"DestroyedEnemyVehicles"`
				EnemyKills []struct {
					Enemy struct {
						Attachments []interface{} `json:"Attachments"`
						BaseID      int           `json:"BaseId"`
					} `json:"Enemy"`
					TotalKills int `json:"TotalKills"`
				} `json:"EnemyKills"`
				Impulses []struct {
					Count int `json:"Count"`
					ID    int `json:"Id"`
				} `json:"Impulses"`
				MedalAwards []struct {
					Count   int `json:"Count"`
					MedalID int `json:"MedalId"`
				} `json:"MedalAwards"`
				ScenarioStats []struct {
					DestroyedEnemyVehicles []struct {
						Enemy struct {
							Attachments []interface{} `json:"Attachments"`
							BaseID      int           `json:"BaseId"`
						} `json:"Enemy"`
						TotalKills int `json:"TotalKills"`
					} `json:"DestroyedEnemyVehicles"`
					EnemyKills []struct {
						Enemy struct {
							Attachments []interface{} `json:"Attachments"`
							BaseID      int           `json:"BaseId"`
						} `json:"Enemy"`
						TotalKills int `json:"TotalKills"`
					} `json:"EnemyKills"`
					FlexibleStats struct {
						ImpulseStatCounts []struct {
							Count int    `json:"Count"`
							ID    string `json:"Id"`
						} `json:"ImpulseStatCounts"`
						ImpulseTimelapses []interface{} `json:"ImpulseTimelapses"`
						MedalStatCounts   []struct {
							Count int    `json:"Count"`
							ID    string `json:"Id"`
						} `json:"MedalStatCounts"`
						MedalTimelapses []interface{} `json:"MedalTimelapses"`
					} `json:"FlexibleStats"`
					GameBaseVariantID string `json:"GameBaseVariantId"`
					Impulses          []struct {
						Count int `json:"Count"`
						ID    int `json:"Id"`
					} `json:"Impulses"`
					MapID       string `json:"MapId"`
					MedalAwards []struct {
						Count   int `json:"Count"`
						MedalID int `json:"MedalId"`
					} `json:"MedalAwards"`
					TotalAssassinations            int     `json:"TotalAssassinations"`
					TotalAssists                   int     `json:"TotalAssists"`
					TotalDeaths                    int     `json:"TotalDeaths"`
					TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
					TotalGamesLost                 int     `json:"TotalGamesLost"`
					TotalGamesTied                 int     `json:"TotalGamesTied"`
					TotalGamesWon                  int     `json:"TotalGamesWon"`
					TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
					TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
					TotalGroundPoundDamage         int     `json:"TotalGroundPoundDamage"`
					TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
					TotalHeadshots                 int     `json:"TotalHeadshots"`
					TotalKills                     int     `json:"TotalKills"`
					TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
					TotalMeleeKills                int     `json:"TotalMeleeKills"`
					TotalPiesEarned                int     `json:"TotalPiesEarned"`
					TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
					TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
					TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
					TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
					TotalShotsFired                int     `json:"TotalShotsFired"`
					TotalShotsLanded               int     `json:"TotalShotsLanded"`
					TotalShoulderBashDamage        int     `json:"TotalShoulderBashDamage"`
					TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
					TotalSpartanKills              int     `json:"TotalSpartanKills"`
					TotalTimePlayed                string  `json:"TotalTimePlayed"`
					TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
					WeaponStats                    []struct {
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponStats"`
					WeaponWithMostKills struct {
						TotalDamageDealt    float64 `json:"TotalDamageDealt"`
						TotalHeadshots      int     `json:"TotalHeadshots"`
						TotalKills          int     `json:"TotalKills"`
						TotalPossessionTime string  `json:"TotalPossessionTime"`
						TotalShotsFired     int     `json:"TotalShotsFired"`
						TotalShotsLanded    int     `json:"TotalShotsLanded"`
						WeaponID            struct {
							Attachments []interface{} `json:"Attachments"`
							StockID     int           `json:"StockId"`
						} `json:"WeaponId"`
					} `json:"WeaponWithMostKills"`
				} `json:"ScenarioStats"`
				TotalAssassinations            int     `json:"TotalAssassinations"`
				TotalAssists                   int     `json:"TotalAssists"`
				TotalDeaths                    int     `json:"TotalDeaths"`
				TotalGamesCompleted            int     `json:"TotalGamesCompleted"`
				TotalGamesLost                 int     `json:"TotalGamesLost"`
				TotalGamesTied                 int     `json:"TotalGamesTied"`
				TotalGamesWon                  int     `json:"TotalGamesWon"`
				TotalGrenadeDamage             float64 `json:"TotalGrenadeDamage"`
				TotalGrenadeKills              int     `json:"TotalGrenadeKills"`
				TotalGroundPoundDamage         float64 `json:"TotalGroundPoundDamage"`
				TotalGroundPoundKills          int     `json:"TotalGroundPoundKills"`
				TotalHeadshots                 int     `json:"TotalHeadshots"`
				TotalKills                     int     `json:"TotalKills"`
				TotalMeleeDamage               float64 `json:"TotalMeleeDamage"`
				TotalMeleeKills                int     `json:"TotalMeleeKills"`
				TotalPiesEarned                int     `json:"TotalPiesEarned"`
				TotalPowerWeaponDamage         float64 `json:"TotalPowerWeaponDamage"`
				TotalPowerWeaponGrabs          int     `json:"TotalPowerWeaponGrabs"`
				TotalPowerWeaponKills          int     `json:"TotalPowerWeaponKills"`
				TotalPowerWeaponPossessionTime string  `json:"TotalPowerWeaponPossessionTime"`
				TotalShotsFired                int     `json:"TotalShotsFired"`
				TotalShotsLanded               int     `json:"TotalShotsLanded"`
				TotalShoulderBashDamage        float64 `json:"TotalShoulderBashDamage"`
				TotalShoulderBashKills         int     `json:"TotalShoulderBashKills"`
				TotalSpartanKills              int     `json:"TotalSpartanKills"`
				TotalTimePlayed                string  `json:"TotalTimePlayed"`
				TotalWeaponDamage              float64 `json:"TotalWeaponDamage"`
				WeaponStats                    []struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponStats"`
				WeaponWithMostKills struct {
					TotalDamageDealt    float64 `json:"TotalDamageDealt"`
					TotalHeadshots      int     `json:"TotalHeadshots"`
					TotalKills          int     `json:"TotalKills"`
					TotalPossessionTime string  `json:"TotalPossessionTime"`
					TotalShotsFired     int     `json:"TotalShotsFired"`
					TotalShotsLanded    int     `json:"TotalShotsLanded"`
					WeaponID            struct {
						Attachments []interface{} `json:"Attachments"`
						StockID     int           `json:"StockId"`
					} `json:"WeaponId"`
				} `json:"WeaponWithMostKills"`
			} `json:"WarzoneStat"`
			Xp int `json:"Xp"`
		} `json:"Result"`
		ResultCode int `json:"ResultCode"`
	} `json:"Results"`
}

type SkullsStruct []interface{}

type SpartanRanksStruct []interface{}

type TeamColorsStruct []interface{}

type VehiclesStruct []interface{}

type WeaponsStruct []interface{}


