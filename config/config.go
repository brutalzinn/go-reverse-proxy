package config

import (
	"github.com/spf13/viper"
)

// //CONTEXT:
// I have a service that is running on my local machine, and i want to expose it to the internet
// but i dont have money to pay for a service like ngrok, so i will create my own redirect.
// The idea is to create a service that will receive a request and redirect to the local service
// the service will have a list of tunnels that will be used to redirect the request to the local service
// this is most economic to get 5 tunnels running at localtonet proxy server. I am paying five dollars foreach month
// /convert five dollars to BRL and you understand why i am doing this

type ProxyPath struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Path     string `json:"path,omitempty"`
}

type Routes struct {
	IN  ProxyPath `json:"in"`  //this is the path LOCAL that will be used to access the service
	OUT ProxyPath `json:"out"` //this is the path EXPOSED that will be used to access the service

}

type GeralConfig struct {
	Port string `json:"port"`
}

type Config struct {
	Geral  GeralConfig `json:"geral"`
	Routes []Routes    `json:"routes"`
}

var cfg *Config

func Init() {
	///localto net configs
	viper.SetDefault("geral.port", "8080")
	viper.SetDefault("routes", make([]Routes, 0))

}

func Load() (*Config, error) {
	Init()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}
	cfg = new(Config)

	cfg.Geral = GeralConfig{
		Port: viper.GetString("geral.port"),
	}
	var routes *[]Routes
	err = viper.UnmarshalKey("routes", &routes)
	if err != nil {
		return nil, err
	}
	cfg.Routes = *routes

	// for idx, route := range routes {
	// 	logrus.Info("Route %v: %v", idx, route)
	// IN: RedirectPath{
	// 	Protocol: route[idx],
	// 	Host:     route["in"].(map[string]any)["host"].(string),
	// 	Port:     route["in"].(map[string]any)["port"].(string),
	// }, OUT: RedirectPath{
	// 	Protocol: route["out"].(map[string]any)["protocol"].(string),
	// 	Host:     route["out"].(map[string]any)["host"].(string),
	// 	Port:     route["out"].(map[string]any)["port"].(string),
	// }}

	// }
	// cfg.Routes =

	viper.SafeWriteConfig()

	return cfg, nil
}

func Get() *Config {
	return cfg
}
