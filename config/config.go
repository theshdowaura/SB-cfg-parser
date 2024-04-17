/*
Generate the package of sing-box configuration template, the current adaptation version is v.1.8.11
*/
package config

/*
sing-box global template
*/
type Config struct {
	Log []Log `json:"log"`
	DNS []DNS `json:"dns"`
	//sing-box inner NTP server for tls/shadowsocks/vmess protocol
	NTP       *NTP       `json:"ntp,omitempty"`
	Inbounds  []Inbound  `json:"inbounds"`
	Outbounds []Outbound `json:"outbounds"`
	Route     Route      `json:"route"`

	Experimental *Experimental `json:"experimental,omitempty"`
}
