package halo

import (
	"encoding/json"
	"fmt"
	"log"
)

func (h *Halo) CampaignMissions() CampaignMissionsStruct {
	var j CampaignMissionsStruct
	jsonObject, err := h.metadataRequest("campaign-missions", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Commendations() CommendationsStruct {
	var j CommendationsStruct
	jsonObject, err := h.metadataRequest("commendations", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) CsrDesignations() CsrDesignationsStruct {
	var j CsrDesignationsStruct
	jsonObject, err := h.metadataRequest("csr-designations", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Enemies() EnemiesStruct {
	var j EnemiesStruct
	jsonObject, err := h.metadataRequest("enemies", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) FlexibleStats() FlexibleStatsStruct {
	var j FlexibleStatsStruct
	jsonObject, err := h.metadataRequest("flexible-stats", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		fmt.Println("hjere")
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) GameBaseVariants() GameBaseVariantsStruct {
	var j GameBaseVariantsStruct
	jsonObject, err := h.metadataRequest("game-base-variants", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) GameVariants(id string) GameVariantsStruct {
	var j GameVariantsStruct
	jsonObject, err := h.metadataRequest("game-variants", id)
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Impulses() ImpulsesStruct {
	var j ImpulsesStruct
	jsonObject, err := h.metadataRequest("impulses", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) MapVariants(id string) MapVariantsStruct {
	var j MapVariantsStruct
	jsonObject, err := h.metadataRequest("map-variants", id)
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Maps() MapsStruct {
	var j MapsStruct
	jsonObject, err := h.metadataRequest("maps", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Medals() MedalsStruct {
	var j MedalsStruct
	jsonObject, err := h.metadataRequest("medals", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Playlists() PlaylistsStruct {
	var j PlaylistsStruct
	jsonObject, err := h.metadataRequest("playlists", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) RequisitionPacks(id string) RequisitionPacksStruct {
	var j RequisitionPacksStruct
	jsonObject, err := h.metadataRequest("requisitions-packs", id)
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Requisitions(id string) RequisitionsStruct {
	var j RequisitionsStruct
	jsonObject, err := h.metadataRequest("requisitions", id)
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Seasons() SeasonsStruct {
	var j SeasonsStruct
	jsonObject, err := h.metadataRequest("seasons", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Skulls() SkullsStruct {
	var j SkullsStruct
	jsonObject, err := h.metadataRequest("skulls", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) SpartanRanks() SpartanRanksStruct {
	var j SpartanRanksStruct
	jsonObject, err := h.metadataRequest("spartan-ranks", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) TeamColors() TeamColorsStruct {
	var j TeamColorsStruct
	jsonObject, err := h.metadataRequest("team-colors", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Vehicles() VehiclesStruct {
	var j VehiclesStruct
	jsonObject, err := h.metadataRequest("vehicles", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}

func (h *Halo) Weapons() WeaponsStruct {
	var j WeaponsStruct
	jsonObject, err := h.metadataRequest("weapons", "")
	if err != nil {
		log.Fatal("MetadataRequest Failed: ", err)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		log.Fatal("Failure to unmarshal json: ", err)
	}
	return j
}
