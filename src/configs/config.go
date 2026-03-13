package configs

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config 配置结构体
type Config struct {
	Server    ServerConfig    `toml:"server"`
	Transport TransportConfig `toml:"transport"`
	Log       LogConfig       `toml:"log"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	IP   string `toml:"ip"`
	Port int    `toml:"port"`
}

// TransportConfig 传输配置
type TransportConfig struct {
	WebSocket WebSocketConfig `toml:"websocket"`
}

// WebSocketConfig WebSocket配置
type WebSocketConfig struct {
	Enabled bool   `toml:"enabled"`
	IP      string `toml:"ip"`
	Port    int    `toml:"port"`
}

// LogConfig 日志配置
type LogConfig struct {
	LogFormat string `toml:"log_format"`
	LogLevel  string `toml:"log_level"`
	LogDir    string `toml:"log_dir"`
	LogFile   string `toml:"log_file"`
}

// AppConfig 全局配置实例
var AppConfig *Config

// LoadConfig 加载配置文件
func LoadConfig(filePath string) error {
	// 读取配置文件
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析配置
	config := &Config{}
	if err := toml.Unmarshal(content, config); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	AppConfig = config
	return nil
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return AppConfig
}
