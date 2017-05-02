package util

import "github.com/koding/multiconfig"

// Our struct which is used for configuration
type ServerConfig struct {
	Index int `json:"index"`
	Host []string `json:"host"`
	Port string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func InitConfig() *ServerConfig{
	// Create a new DefaultLoader without or with an initial config file
	//m := multiconfig.New()
	m := multiconfig.NewWithPath("config.json") // supports TOML, JSON and YAML

	// Get an empty struct for your configuration
	serverConf := new(ServerConfig)

	// Populated the serverConf struct
	m.Load(serverConf) // Check for error
	//if err != nil {
	//	panic(err)
	//}
	//m.MustLoad(serverConf)    // Panic's if there is any error

	return serverConf
}

