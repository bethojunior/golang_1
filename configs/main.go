package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName string `yaml:"app_name"`
	Port    int    `yaml:"port"`
	Debug   bool   `yaml:"debug"`
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug: %v\n", config.Debug)
}