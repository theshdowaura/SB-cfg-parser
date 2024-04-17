package config

type NTP struct {
	Enabled    bool   `json:"enabled"`
	NTPServer  string `json:"server"`
	ServerPort int    `json:"server_port"`
	Interval   string `json:"interval"`
}
