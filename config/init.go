package config

import (
	"FurbotServer-Go/constants"
	"github.com/spf13/viper"
	"strings"
)

// Config 配置
type Config struct {
	Name string
}

// InitConfig 初始化配置
func InitConfig(path string) error {
	config := Config{
		Name: path,
	}
	if err := config.initConfig(); err != nil {
		return err
	}
	// 配置热加载
	viper.WatchConfig()

	return nil
}
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("config")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(constants.EnvPrefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
