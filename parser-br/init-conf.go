package parser_br

import (
	"encoding/json"
	"fmt"
	"github.com/theshdowaura/SB-cfg-parser/config"
	"os"
)

func Configinit() {
	conf := config.Config{}

	// Initialize config fields
	conf.Log = []config.Log{
		{
			Disabled:  false,
			Level:     "",
			Output:    "",
			Timestamp: false,
		},
	}
	conf.DNS = []config.DNS{
		{
			Servers:          nil,
			Rules:            nil,
			Final:            "",
			Strategy:         "",
			DisableCache:     false,
			DisableExpire:    false,
			IndependentCache: false,
			ReverseMapping:   false,
			Fakeip:           config.DNSFakeip{},
		},
	}
	conf.Inbounds = []config.Inbound{}
	conf.Outbounds = []config.Outbound{}
	conf.Route.Rules = []config.RouteRules{}

	var b []byte
	var err error

	// Check for nil values
	if conf.Log == nil {
		conf.Log = []config.Log{}
	}
	if conf.DNS == nil {
		conf.DNS = []config.DNS{}
	}
	if conf.Inbounds == nil {
		conf.Inbounds = []config.Inbound{}
	}
	if conf.Outbounds == nil {
		conf.Outbounds = []config.Outbound{}
	}
	if conf.Route.Rules == nil {
		conf.Route.Rules = []config.RouteRules{}
	}

	switch {
	case len(conf.Log) > 0:
		b, err = json.MarshalIndent(conf.Log, "", "    ")
		if err != nil {
			fmt.Println("json marshal error:", err)
			return
		}
		fallthrough
	case len(conf.DNS) > 0:
		b, err = json.MarshalIndent(conf.DNS, "", "    ")
		if err != nil {
			fmt.Println("json marshal error:", err)
			return
		}
		fallthrough
	case len(conf.Inbounds) > 0:
		b, err = json.MarshalIndent(conf.Inbounds, "", "    ")
		if err != nil {
			fmt.Println("json marshal error:", err)
			return
		}
		fallthrough
	case len(conf.Outbounds) > 0:
		b, err = json.MarshalIndent(conf.Outbounds, "", "    ")
		if err != nil {
			fmt.Println("json marshal error:", err)
			return
		}
		fallthrough
	case len(conf.Route.Rules) > 0:
		b, err = json.MarshalIndent(conf.Route.Rules, "", "    ")
		if err != nil {
			fmt.Println("json marshal error:", err)
			return
		}
		fallthrough
	default:
		b, err = json.MarshalIndent(conf, "", "    ")
		if err != nil {
			fmt.Println("json marshal error:", err)
			return
		}
	}

	// Open file for writing
	f, err := os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer f.Close() // Ensure file is closed when function returns

	// Write config to file
	if _, err := f.Write(b); err != nil {
		fmt.Println("error writing to file:", err)
		return
	}

	fmt.Println("config.json generated successfully")
}
