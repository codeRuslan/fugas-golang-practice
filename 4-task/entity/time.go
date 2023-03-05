package entity

import (
	"awesomeProject1/config"
	"strings"
	"time"
)

type CivilTime time.Time

func (c *CivilTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	ConfigFile := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	t, err := time.Parse(ConfigFile.TimeFormat, value)
	if err != nil {
		return err
	}
	*c = CivilTime(t)
	return nil
}

func (c CivilTime) MarshalJSON() ([]byte, error) {
	ConfigFile := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	return []byte(`"` + time.Time(c).Format(ConfigFile.TimeFormat) + `"`), nil
}
