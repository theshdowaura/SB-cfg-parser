package config

type DNS struct {
	Servers          []DNSServer `json:"servers"` //DNS
	Rules            []DNSRules  `json:"rules"`
	Final            string      `json:"final"`
	Strategy         string      `json:"strategy"`
	DisableCache     bool        `json:"disable_cache"`
	DisableExpire    bool        `json:"disable_expire"`
	IndependentCache bool        `json:"independent_cache"`
	ReverseMapping   bool        `json:"reverse_mapping"`
	Fakeip           DNSFakeip   `json:"fakeip"`
}
type DNSServer struct {
	Tag     string `json:"tag,omitempty"`
	Address string `json:"address"` /*
		example filed:
			System	local
			TCP	tcp://1.0.0.1
			UDP	8.8.8.8 udp://8.8.4.4
			TLS	tls://dns.google
			HTTPS	https://1.1.1.1/dns-query
			QUIC	quic://dns.adguard.com
			HTTP3	h3://8.8.8.8/dns-query
			RCode	rcode://refused
			DHCP	dhcp://auto 或 dhcp://en0
			FakeIP	fakeip
	*/
	AddressResolver string `json:"address_resolver,omitempty"` //if serverAddress include host filed else it must filed
	AddressStrategy string `json:"address_strategy"`
	Strategy        string `json:"strategy"` //Default parsing strategy.
	Detour          string `json:"detour"`
	ClientSubnet    string `json:"client_subnet"`
}

type DNSRules struct {
	Domain  string `json:"domain"`
	Geosite string `json:"geosite"` //Deprecated in sing-box 1.8.0
	Server  string `json:"server"`
}
type DNSFakeip struct {
	Enable     bool   `json:"enable"`
	Inet4Range string `json:"inet4_range"`
	Inet6Range string `json:"inet6_range"`
}
