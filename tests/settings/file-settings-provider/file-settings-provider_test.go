package file_settings_provider_tests

import (
	file_settings_provider "simple-hosting/go.common/settings/file-settings-provider"
	"testing"
)

func TestYamlReading(t *testing.T) {
	type settings struct {
		A string `yaml:"A"`
		B int    `yaml:"B"`
		C struct {
			D float32 `yaml:"D"`
		} `yaml:"C"`
	}
	s := file_settings_provider.GetSetting[settings]("test-settings.yaml")
	etalon := settings{
		A: "some-string",
		B: 10,
		C: struct {
			D float32 `yaml:"D"`
		}{
			D: 20.01,
		},
	}
	if s != etalon {
		t.Errorf("Read struct is not equal to etalon")
	}
}
