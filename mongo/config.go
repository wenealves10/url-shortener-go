package mongo

import (
	"encoding/json"
	"os"
)

type Config struct {
	Uri        string `json:"uri"`
	Db         string `json:"db"`
	Collection string `json:"collection"`
}

func (c *Config) ToJson() string {
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ParseConfig(configFile string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
