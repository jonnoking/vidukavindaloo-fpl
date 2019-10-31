package api

import (
	"encoding/json"

	"github.com/jonnoking/vidukavindaloo-fpl/models"
)

//GetElementSummary retrive my team from FPL
func (api *API) GetPlayerSummary(elementID int) (*models.PlayerSummary, error) {

	var summary models.PlayerSummary

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetElementSummaryAPI(elementID))
	if readErr != nil {
		return nil, readErr
	}

	summary = models.PlayerSummary{}
	json.Unmarshal([]byte(byteValue), &summary)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetElementSummaryFilename(elementID))

	return &summary, nil
}
