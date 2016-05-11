package halo

import (
	"encoding/json"
	"fmt"
)

func (h *Halo) CampaignMissions() (CampaignMissionsStruct, error) {
	var j CampaignMissionsStruct
	jsonObject, err := h.metadataRequest("campaign-missions", "")
	if err != nil {
		return CampaignMissionsStruct{}, fmt.Errorf(
			"MetadataRequest campaign-missions Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return CampaignMissionsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Commendations() (CommendationsStruct, error) {
	var j CommendationsStruct
	jsonObject, err := h.metadataRequest("commendations", "")
	if err != nil {
		return CommendationsStruct{}, fmt.Errorf(
			"MetadataRequest commendations Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return CommendationsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) CsrDesignations() (CsrDesignationsStruct, error) {
	var j CsrDesignationsStruct
	jsonObject, err := h.metadataRequest("csr-designations", "")
	if err != nil {
		return CsrDesignationsStruct{}, fmt.Errorf(
			"MetadataRequest csr-designations Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return CsrDesignationsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Enemies() (EnemiesStruct, error) {
	var j EnemiesStruct
	jsonObject, err := h.metadataRequest("enemies", "")
	if err != nil {
		return EnemiesStruct{}, fmt.Errorf(
			"MetadataRequest enemies Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return EnemiesStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) FlexibleStats() (FlexibleStatsStruct, error) {
	var j FlexibleStatsStruct
	jsonObject, err := h.metadataRequest("flexible-stats", "")
	if err != nil {
		return FlexibleStatsStruct{}, fmt.Errorf(
			"MetadataRequest flexible-stats Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return FlexibleStatsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) GameBaseVariants() (GameBaseVariantsStruct, error) {
	var j GameBaseVariantsStruct
	jsonObject, err := h.metadataRequest("game-base-variants", "")
	if err != nil {
		return GameBaseVariantsStruct{}, fmt.Errorf(
			"MetadataRequest game-base-variants Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return GameBaseVariantsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) GameVariants(id string) (GameVariantsStruct, error) {
	var j GameVariantsStruct
	jsonObject, err := h.metadataRequest("game-variants", id)
	if err != nil {
		return GameVariantsStruct{}, fmt.Errorf(
			"MetadataRequest game-variants Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return GameVariantsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Impulses() (ImpulsesStruct, error) {
	var j ImpulsesStruct
	jsonObject, err := h.metadataRequest("impulses", "")
	if err != nil {
		return ImpulsesStruct{}, fmt.Errorf(
			"MetadataRequest impulses Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return ImpulsesStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) MapVariants(id string) (MapVariantsStruct, error) {
	var j MapVariantsStruct
	jsonObject, err := h.metadataRequest("map-variants", id)
	if err != nil {
		return MapVariantsStruct{}, fmt.Errorf(
			"MetadataRequest map-variants Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return MapVariantsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Maps() (MapsStruct, error) {
	var j MapsStruct
	jsonObject, err := h.metadataRequest("maps", "")
	if err != nil {
		return MapsStruct{}, fmt.Errorf(
			"MetadataRequest maps Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return MapsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Medals() (MedalsStruct, error) {
	var j MedalsStruct
	jsonObject, err := h.metadataRequest("medals", "")
	if err != nil {
		return MedalsStruct{}, fmt.Errorf(
			"MetadataRequest medals Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return MedalsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Playlists() (PlaylistsStruct, error) {
	var j PlaylistsStruct
	jsonObject, err := h.metadataRequest("playlists", "")
	if err != nil {
		return PlaylistsStruct{}, fmt.Errorf(
			"MetadataRequest playlists Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return PlaylistsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) RequisitionPacks(id string) (RequisitionPacksStruct, error) {
	var j RequisitionPacksStruct
	jsonObject, err := h.metadataRequest("requisition-packs", id)
	if err != nil {
		return RequisitionPacksStruct{}, fmt.Errorf(
			"MetadataRequest requisition-packs Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return RequisitionPacksStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Requisitions(id string) (RequisitionsStruct, error) {
	var j RequisitionsStruct
	jsonObject, err := h.metadataRequest("requisitions", id)
	if err != nil {
		return RequisitionsStruct{}, fmt.Errorf(
			"MetadataRequest requisitions Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return RequisitionsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Seasons() (SeasonsStruct, error) {
	var j SeasonsStruct
	jsonObject, err := h.metadataRequest("seasons", "")
	if err != nil {
		return SeasonsStruct{}, fmt.Errorf(
			"MetadataRequest seasons Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return SeasonsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Skulls() (SkullsStruct, error) {
	var j SkullsStruct
	jsonObject, err := h.metadataRequest("skulls", "")
	if err != nil {
		return SkullsStruct{}, fmt.Errorf(
			"MetadataRequest skulls Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return SkullsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) SpartanRanks() (SpartanRanksStruct, error) {
	var j SpartanRanksStruct
	jsonObject, err := h.metadataRequest("spartan-ranks", "")
	if err != nil {
		return SpartanRanksStruct{}, fmt.Errorf(
			"MetadataRequest spartan-ranks Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return SpartanRanksStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) TeamColors() (TeamColorsStruct, error) {
	var j TeamColorsStruct
	jsonObject, err := h.metadataRequest("team-colors", "")
	if err != nil {
		return TeamColorsStruct{}, fmt.Errorf(
			"MetadataRequest team-colors Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return TeamColorsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Vehicles() (VehiclesStruct, error) {
	var j VehiclesStruct
	jsonObject, err := h.metadataRequest("vehicles", "")
	if err != nil {
		return VehiclesStruct{}, fmt.Errorf(
			"MetadataRequest vehicles Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return VehiclesStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}

func (h *Halo) Weapons() (WeaponsStruct, error) {
	var j WeaponsStruct
	jsonObject, err := h.metadataRequest("weapons", "")
	if err != nil {
		return WeaponsStruct{}, fmt.Errorf(
			"MetadataRequest weapons Failed: %v",
			err,
		)
	}
	err = json.Unmarshal(jsonObject, &j)
	if err != nil {
		return WeaponsStruct{}, fmt.Errorf("Failure to unmarshal json: %v", err)
	}
	return j, nil
}
