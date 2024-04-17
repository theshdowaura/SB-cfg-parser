package config

type Route struct {
	Rules   []RouteRules `json:"rules"`
	RuleSet []RuleSet    `json:"rule_set"`

	// 默认出站标签。如果为空，将使用第一个可用于对应协议的出站。
	Final string `json:"final"`

	// 仅支持 Linux、Windows 和 macOS。
	//
	// 默认将出站连接绑定到默认网卡，以防止在 tun 下出现路由环路。
	//
	// 如果设置了 outbound.bind_interface 设置，则不生效。
	AutoDetectInterface bool `json:"auto_detect_interface"`

	// 仅支持 Android。
	//
	// 启用 auto_detect_interface 时接受 Android VPN 作为上游网卡。
	OverrideAndroidVpn bool `json:"override_android_vpn"`

	// 仅支持 Linux、Windows 和 macOS。
	//
	// 默认将出站连接绑定到指定网卡，以防止在 tun 下出现路由环路。
	//
	// 如果设置了 auto_detect_interface 设置，则不生效。
	DefaultInterface string `json:"default_interface"`

	// 仅支持 Linux。
	//
	// 默认为出站连接设置路由标记。
	//
	// 如果设置了 outbound.routing_mark 设置，则不生效。
	DefaultMark int `json:"default_mark"`
}

type RouteRules struct {
	// 入站 标签。
	Inbound []string `json:"inbound"`
	// 4 或 6。
	//
	// 默认不限制。
	IpVersion int `json:"ip_version"`
	// tcp 或 udp。
	Network []string `json:"network"`
	// 认证用户名，参阅入站设置。
	AuthUser []string `json:"auth_user"`
	// 探测到的协议, 参阅 协议探测。
	Protocol []string `json:"protocol"`
	// 匹配完整域名。
	Domain []string `json:"domain"`
	// 匹配域名后缀。
	DomainSuffix []string `json:"domain_suffix"`
	// 匹配域名关键字。
	DomainKeyword []string `json:"domain_keyword"`
	// 匹配域名正则表达式。
	DomainRegex []string `json:"domain_regex"`
	// 已在 sing-box 1.8.0 废弃
	//
	// Geosite 已废弃且可能在不久的将来移除，参阅 迁移指南。
	//
	// 匹配 Geosite。
	Geosite []string `json:"geosite"`
	// 已在 sing-box 1.8.0 废弃
	//
	// GeoIP 已废弃且可能在不久的将来移除，参阅 迁移指南。
	//
	// 匹配源 GeoIP。
	SourceGeoip []string `json:"source_geoip"`
	// 已在 sing-box 1.8.0 废弃
	//
	// GeoIP 已废弃且可能在不久的将来移除，参阅 迁移指南。
	//
	// 匹配 GeoIP。
	Geoip []string `json:"geoip"`
	// 匹配源 IP CIDR。
	SourceIpCidr []string `json:"source_ip_cidr"`
	// 自 sing-box 1.8.0 起
	//
	// 匹配非公开源 IP。
	SourceIpIsPrivate bool `json:"source_ip_is_private"`
	// 匹配 IP CIDR。
	IpCidr []string `json:"ip_cidr"`
	// 自 sing-box 1.8.0 起
	//
	// 匹配非公开 IP。
	IpIsPrivate bool `json:"ip_is_private"`
	// 匹配源端口。
	SourcePort []int `json:"source_port"`
	// 匹配源端口范围。
	SourcePortRange []string `json:"source_port_range"`
	// 匹配端口。
	Port []int `json:"port"`
	// 匹配端口范围。
	PortRange []string `json:"port_range"`
	// 仅支持 Linux、Windows 和 macOS。
	//
	// 匹配进程名称。
	ProcessName []string `json:"process_name"`
	// 仅支持 Linux、Windows 和 macOS.
	//
	// 匹配进程路径。
	ProcessPath []string `json:"process_path"`
	// 匹配 Android 应用包名。
	PackageName []string `json:"package_name"`
	// 仅支持 Linux.
	//
	// 匹配用户名。
	User []string `json:"user"`
	// 仅支持 Linux.
	//
	// 匹配用户 ID。
	UserId []int `json:"user_id"`
	// 匹配 Clash 模式。
	ClashMode string `json:"clash_mode"`
	// 仅在 Android 与 Apple 平台图形客户端中支持。
	//
	// 匹配 WiFi SSID。
	WifiSsid []string `json:"wifi_ssid"`
	// 仅在 Android 与 Apple 平台图形客户端中支持。
	//
	// 匹配 WiFi BSSID。
	WifiBssid []string `json:"wifi_bssid"`
	// 自 sing-box 1.8.0 起
	//
	// 匹配规则集。
	RuleSet []string `json:"rule_set"`
	// 自 sing-box 1.8.0 起
	//
	// 使规则集中的 ipcidr 规则匹配源 IP。
	RuleSetIpcidrMatchSource bool `json:"rule_set_ipcidr_match_source"`
	// 反选匹配结果。
	Invert bool `json:"invert"`
	// 必填
	//
	// 目标出站的标签。
	Outbound string `json:"outbound"`
}
type RuleSet struct {
	// 必填
	//
	// 规则集类型，本地或远程。
	Type string `json:"type"`
	// 必填
	//
	// 规则集标签。
	Tag string `json:"tag"`
	// 必填
	//
	// 规则集格式，源或二进制。
	Format string `json:"format"`

	// 本地字段
	// 必填
	//
	// 规则集文件路径。
	Path string `json:"path"`

	// 远程字段
	// 必填
	//
	// 规则集下载 URL。
	URL string `json:"url"`
	// 下载规则集的出站标签。
	//
	// 如果为空，将使用默认出站。
	DownloadDetour string `json:"download_detour"`
	// 规则集更新间隔。
	//
	// 如果为空，将使用 1 天。
	UpdateInterval string `json:"update_interval"`
}
