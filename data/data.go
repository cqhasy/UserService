package data

import (
	"gopkg.in/yaml.v3"
	"os"
)

type MysqlConfig struct {
	Mysql struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"mysql"`
}

// NewMysqlConfig 读取并解析 YAML 配置文件
func NewMysqlConfig(configFilePath string) (*MysqlConfig, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return &MysqlConfig{}, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var configs MysqlConfig
	// 解析 YAML 文件
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&configs)
	if err != nil {
		return &MysqlConfig{}, err
	}

	return &configs, nil
}
