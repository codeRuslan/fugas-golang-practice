package configs

import (
	"flag"
	"fmt"
)

type DataFlag struct {
	ListenPort string
	TimeFormat string
}

func ReadFlags() *DataFlag {
	ListenPortFlag := flag.String("listenport", ":8000", "config string")
	TimeFormatFlag := flag.String("timeformat", "02.01.2006", "config string")

	FlagConfig := DataFlag{
		ListenPort: *ListenPortFlag,
		TimeFormat: *TimeFormatFlag,
	}

	return &FlagConfig
}

func (text *DataFlag) PrintConfig() {
	fmt.Printf("Listen Port is: %s\n", text.ListenPort)
	fmt.Printf("Time Format is: %s\n", text.TimeFormat)
}
