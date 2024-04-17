package parser_br

import (
	"encoding/json"
	"fmt"
	"github.com/theshdowaura/SB-cfg-parser/config"
	"os"
)

func Configinit() {
	conf := config.Config{}
	Log := []config.Log{}
	DNS := []config.DNS{} // Create a slice of type []config.DNS
	Inbounds := []config.Inbound{}
	Outbounds := []config.Outbound{}
	RouteRule := []config.RouteRules{}
	conf.Log = Log
	conf.DNS = DNS
	conf.Inbounds = Inbounds
	conf.Outbounds = Outbounds
	conf.Route.Rules = RouteRule
	b, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}

	err = os.WriteFile("config.json", b, 0644)
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}

	fmt.Println("config.json generated successfully")
}
