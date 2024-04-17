package config

type Log struct {
	Disabled  bool   `json:"disabled"`  //日志状态是否启用
	Level     string `json:"info"`      //信息输出等级:trace,debug,info,warn,error,fatal,panic
	Output    string `json:"output"`    //文件输出路径，启动将不输出到控制台
	Timestamp bool   `json:"timestamp"` //为每行输出添加时间线
}
