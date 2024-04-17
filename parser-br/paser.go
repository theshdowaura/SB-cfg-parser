package parser_br

import (
	"encoding/json"
	"fmt"
	"github.com/theshdowaura/SB-cfg-parser/config"
	"os"
)

// LoadConfig 从指定的文件中加载配置
func LoadConfig(filename string) (*config.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer file.Close()

	var config config.Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// SaveConfig 将配置保存到指定的文件
func SaveConfig(config *config.Config, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建配置文件失败: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		return fmt.Errorf("写入新配置文件失败: %v", err)
	}

	return nil
}
