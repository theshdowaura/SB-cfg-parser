package config

type Experimental struct {
	CacheFile CacheFile `json:"cache_file,omitempty"`
	ClashApi  ClashApi  `json:"clash_api,omitempty"`
	V2rayApi  V2rayApi  `json:"v2Ray_api,omitempty"`
}

/*
Since sing-box 1.8.0
*/
type CacheFile struct {
	Enabled     bool   `json:"enabled,omitempty"`
	Path        string `json:"path,omitempty"`
	CacheId     string `json:"cache_id,omitempty"`
	StoreFakeip bool   `json:"store_fakeip,omitempty"`
	StoreRdrc   bool   `json:"store_rdrc,omitempty"`
	RdrcTimeout string `json:"rdrc_timeout,omitempty"`
}

type ClashApi struct {
	// RESTful web API 监听地址。如果为空，则禁用 Clash API。
	ExperimentalController string `json:"external_controller,omitempty"`

	// 到静态网页资源目录的相对路径或绝对路径。sing-box 会在 http://{{external-controller}}/ui 下提供它。
	ExternalUI string `json:"external_ui"`

	// 静态网页资源的 ZIP 下载 URL，如果指定的 external_ui 目录为空，将使用。
	// 默认使用 https://github.com/MetaCubeX/Yacd-meta/archive/gh-pages.zip。
	ExternalUIDownloadURL string `json:"external_ui_download_url"`

	// 用于下载静态网页资源的出站的标签。
	// 如果为空，将使用默认出站。
	ExternalUIDownloadDetour string `json:"external_ui_download_detour"`

	// RESTful API 的密钥（可选） 通过指定 HTTP 标头 Authorization: Bearer ${secret} 进行身份验证
	// 如果 RESTful API 正在监听 0.0.0.0，请始终设置一个密钥。
	Secret string `json:"secret"`

	// Clash 中的默认模式，默认使用 Rule。
	// 此设置没有直接影响，但可以通过 clash_mode 规则项在路由和 DNS 规则中使用。
	DefaultMode string `json:"default_mode"`

	// 已在 sing-box 1.8.0 废弃
	// store_mode 已在 Clash API 中废弃，且默认启用当 cache_file.enabled。
	StoreMode bool `json:"store_mode"`

	// 已在 sing-box 1.8.0 废弃
	// store_selected 已在 Clash API 中废弃，且默认启用当 cache_file.enabled。
	// 必须为目标出站设置标签。
	StoreSelected bool `json:"store_selected"`

	// 已在 sing-box 1.8.0 废弃
	// store_selected 已在 Clash API 中废弃，且已迁移到 cache_file.store_fakeip。
	StoreFakeip bool `json:"store_fakeip"`

	// 已在 sing-box 1.8.0 废弃
	// cache_file 已在 Clash API 中废弃，且已迁移到 cache_file.enabled 和 cache_file.path。
	CacheFile string `json:"cache_file"`

	// 已在 sing-box 1.8.0 废弃
	// cache_id 已在 Clash API 中废弃，且已迁移到 cache_file.cache_id。
	CacheID string `json:"cache_id"`
}

type V2rayApi struct {
	// gRPC API 监听地址。如果为空，则禁用 V2Ray API。
	Listen string `json:"listen"`

	// 流量统计服务设置。
	Stats struct {
		// 启用统计服务。
		Enabled bool `json:"enabled"`

		// 统计流量的入站列表。
		Inbounds []string `json:"inbounds"`

		// 统计流量的出站列表。
		Outbounds []string `json:"outbounds"`

		// 统计流量的用户列表。
		Users []string `json:"users"`
	} `json:"stats"`
}
