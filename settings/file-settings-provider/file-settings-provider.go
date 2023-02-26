package file_settings_provider

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetSetting[settingsT any](filepath string) settingsT {
	var settings settingsT
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &settings)
	if err != nil {
		panic(err)
	}
	return settings
}
