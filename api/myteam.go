package api

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/jonnoking/vidukavindaloo-fpl/models"
)

func (api *API) GetMyTeamFromCache(entryID int, players *models.Players, teams *models.Teams, playerTypes *models.PlayerTypes) (*models.MyTeam, error) {

	// add teamID to file name
	f, err := ioutil.ReadFile(api.Config.API.GetMyTeamAPI(entryID))
	if err != nil {
		return nil, err
	}

	myteam, _ := models.NewMyTeam(f, players, teams, playerTypes)

	return myteam, nil
}

//GetMyTeam retrive my team from FPL
func (api *API) GetMyTeam(entryID int) (*models.MyTeam, error) {

	var myteam models.MyTeam

	byteValue, readErr := api.ExecuteFPLGet(api.Config.Files.GetMyTeamFilename(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	myteam = models.MyTeam{}
	json.Unmarshal([]byte(byteValue), &myteam)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetMyTeamFilename(entryID))

	return &myteam, nil
}
