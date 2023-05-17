package entity

import (
	"awesomeProject1/config"
	"fmt"
	"strings"
	"time"
)

type CivilTime time.Time

func (c *CivilTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	ConfigFile, err := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	t, err := time.Parse(ConfigFile.TimeFormat, value)
	if err != nil {
		return err
	}
	*c = CivilTime(t)
	return nil
}

func (c CivilTime) MarshalJSON() ([]byte, error) {
	ConfigFile, err := config.ReadJsonConfigFile("/Users/ruslanpilipyuk/GolandProjects/awesomeProject1/config/config.json")
	if err != nil {
		fmt.Println("MarhalJSON operation failed")
	}
	return []byte(`"` + time.Time(c).Format(ConfigFile.TimeFormat) + `"`), nil
}
