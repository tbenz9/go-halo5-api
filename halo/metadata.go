package halo

import "encoding/json"

// This one works!
type CampaignMissionsStruct []struct {
	MissionNumber json.Number `json:"missionNumber"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	ImageURL      string      `json:"imageUrl"`
	Type          string      `json:"type"`
	ID            string      `json:"id"`
	ContentID     string      `json:"contentId"`
}

// This one works!
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

// This one works!
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
	Faction           string      `json:"faction"`
	Name              string      `json:"name"`
	Description       interface{} `json:"description"`
	LargeIconImageURL string      `json:"largeIconImageUrl"`
	SmallIconImageURL string      `json:"smallIconImageUrl"`
	ID                string      `json:"id"`
	ContentID         string      `json:"contentId"`
}

type VehiclesStruct []struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	LargeIconImageURL string `json:"largeIconImageUrl"`
	SmallIconImageURL string `json:"smallIconImageUrl"`
	IsUsableByPlayer  bool   `json:"isUsableByPlayer"`
	ID                string `json:"id"`
	ContentID         string `json:"contentId"`
}

func CampaignMissions(baseurl, title string) CampaignMissionsStruct {
	var j CampaignMissionsStruct
	err := json.Unmarshal(metadataRequest(baseurl, title, "campaign-missions", ""), &j)
	checkErr(err)
	return j
}

func Commendations(baseurl, title string) CommendationsStruct {
	var j CommendationsStruct
	err := json.Unmarshal(metadataRequest(baseurl, title, "commendations", ""), &j)
	checkErr(err)
	return j
}

func CsrDesignations(baseurl, title string) CsrDesignationsStruct {
	var j CsrDesignationsStruct
	err := json.Unmarshal(metadataRequest(baseurl, title, "csr-designations", ""), &j)
	checkErr(err)
	return j
}

func Enemies(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "enemies", "")
}

func FlexibleStats(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "flexible-stats", "")
}

func GameBaseVariants(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "game-base-variants", "")
}

func GameVariants(baseurl, title, id string) []byte {
	return metadataRequest(baseurl, title, "game-variants", id)
}

func Impulses(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "impulses", "")
}

func MapVariants(baseurl, title, id string) []byte {
	return metadataRequest(baseurl, title, "map-variants", id)
}

func Maps(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "maps", "")
}

func Medals(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "medals", "")
}

func Playlists(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "playlists", "")
}

func RequisitionPacks(baseurl, title, id string) []byte {
	return metadataRequest(baseurl, title, "requisition-packs", id)
}

func Requisitions(baseurl, title, id string) []byte {
	return metadataRequest(baseurl, title, "requisitions", id)
}

func Seasons(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "seasons", "")
}

func Skulls(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "skulls", "")
}

func SpartanRanks(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "spartan-ranks", "")
}

func TeamColors(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "team-colors", "")
}

func Vehicles(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "vehicles", "")
}

func Weapons(baseurl, title string) []byte {
	return metadataRequest(baseurl, title, "weapons", "")
}
