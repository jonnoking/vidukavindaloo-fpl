package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func (api *API) LoadBoostrapFromFile(filename string) ([]byte, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (api *API) LoadBootsrapFromCache() ([]byte, error) {
	return api.LoadBoostrapFromFile(api.Config.Files.GetBootstrapFilename())
}

func (api *API) RefreshBootstrap() map[string]interface{} {

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetBoostrapAPI())
	if readErr != nil {
		log.Fatal(readErr)
	}

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetBootstrapFilename())

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

// SaveBootstrap Gets bootstrap json and saves to supplied filename under config folder
func (api *API) SaveBootstrap(filename string) error {

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetBoostrapAPI())
	if readErr != nil {
		return readErr
	}

	se := api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.Folder+"/"+filename)
	if se != nil {
		return se
	}

	return nil
}
