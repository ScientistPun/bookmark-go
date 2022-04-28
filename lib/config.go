package lib

import (
	"log"
)

type ConfigForSystem struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Session   string `yaml:"session"`
	SignInUrl string `yaml:"sign-in-url"`
}

type ConfigForService struct {
	Skin    string `yaml:"skin"`
	Port    string `yaml:"port"`
	SSLMode bool   `yaml:"ssl-mode"`
	Csr     string `yaml:"csr"`
	Key     string `yaml:"key"`
}

type ConfigForAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Configuration struct {
	System  ConfigForSystem
	Service ConfigForService
	Auth    ConfigForAuth
}

var GlobalConfig Configuration

func InitConfig() {
	system := ConfigForSystem{
		Name:      "bookmark",
		Version:   "1.0.2",
		Session:   "aid",
		SignInUrl: "/sign-in",
	}
	service := ConfigForService{
		Skin:    "skin-1",
		Port:    "8080",
		SSLMode: false,
		Csr:     "",
		Key:     "",
	}
	auth := ConfigForAuth{
		Username: "root",
		Password: "888888",
	}
	GlobalConfig = Configuration{
		System:  system,
		Service: service,
		Auth:    auth,
	}

	WriteYamlFile(ConfigPath, GlobalConfig)
}

func LoadConfiguration() {
	if !IsFileExists(ConfigPath) {
		InitConfig()
	} else {
		err := LoadYamlFile(ConfigPath, &GlobalConfig)
		if err != nil {
			log.Printf("file get err: %v ", err.Error())
		}
	}
	if len(GlobalConfig.Service.Skin) < 1 {
		GlobalConfig.Service.Skin = "default"
	}
	if len(GlobalConfig.Service.Port) < 1 {
		GlobalConfig.Service.Port = "8080"
	}
	if len(GlobalConfig.Service.Csr) < 0 || len(GlobalConfig.Service.Key) < 0 {
		GlobalConfig.Service.SSLMode = false
	}
}
