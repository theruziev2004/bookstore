package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Settings struct {
	Dsn    string `json:"dsn"`
	JWTKey string `json:"jwt_key"`
}

func ReadSettings(path string) (*Settings, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var settings Settings
	if err = json.Unmarshal(byteValue, &settings); err != nil {
		return nil, err
	}

	return &settings, nil
}
