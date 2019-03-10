package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var app *AppStore

//AppStore stores all configuration
type AppStore struct {
	ProcessName string                                 `json:"process_name,omitempty"`
	LogLevel    string                                 `json:"log_severity,omitempty"`
	Port        string                                 `json:"service_port,omitempty"`
	DB          map[string]string                      `json:"db,omitempty"`
	Cache       map[string]string                      `json:"cache,omitempty"`
	Throttle    map[string]map[string]*ResourceActions `json:"throttle,omitempty"`
}

// ResourceActions list of allowed actions
type ResourceActions struct {
	Get    string `json:"get,omitempty"`
	Post   string `json:"post,omitempty"`
	Put    string `json:"put,omitempty"`
	Delete string `json:"delete,omitempty"`
}

//GetConfig returns init-ed application config instance
func GetConfig() AppStore {
	return *app
}

//InitConfig will initialize app config with config file name
func InitConfig(configFilepath string) error {
	if configFilepath == "" {
		configFilepath = path.Join(".", "config.json")
	}

	configFile := &configFilepath
	config := new(AppStore)
	fmt.Println(config)
	file, err := os.Open(*configFile)

	if err != nil {
		return err
	}

	if err := json.NewDecoder(file).Decode(config); err != nil {
		return err
	}

	app = config
	return nil
}
