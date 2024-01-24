package config

import (
	"flag"
	"os"
)

const (
	serverAddressEnv = "SERVER_ADDRESS"
	baseURLEnv       = "BASE_URL"
)

type Option struct {
	Name     string
	FlagName string
	Value    string
}

type Config struct {
	ServerAddress Option
	BaseURL       Option
}

var URLConfig Config = Config{
	Option{
		"Server address",
		"a",
		"localhost:8080",
	},
	Option{
		"Base URL address",
		"b",
		"http://localhost:8080",
	},
}

func NewConfig() {

	configFlags := flag.NewFlagSet("Config flagset", flag.ContinueOnError)

	serverAddress := getServerAddr(configFlags)
	baseURL := getBaseURL(configFlags)

	configFlags.Parse(os.Args[1:])

	if configFlags.Parsed() {
		URLConfig.ServerAddress.Value = *serverAddress
		URLConfig.BaseURL.Value = *baseURL
	}
}

func getServerAddr(flags *flag.FlagSet) *string {
	var serverAddr *string
	if addr := os.Getenv(serverAddressEnv); addr != "" {
		serverAddr = &addr
	} else {
		serverAddr = flags.String(URLConfig.ServerAddress.FlagName, URLConfig.ServerAddress.Value, URLConfig.BaseURL.Name)
	}
	return serverAddr
}

func getBaseURL(flags *flag.FlagSet) *string {
	var baseURL *string
	if url := os.Getenv(baseURLEnv); url != "" {
		baseURL = &url
	} else {
		baseURL = flags.String(URLConfig.BaseURL.FlagName, URLConfig.BaseURL.Value, URLConfig.BaseURL.Name)
	}
	return baseURL
}
