package config

import (
	"os"

	logger "kmp-news-producer/logging"
	"kmp-news-producer/parser"
)

// Configuration stores global configuration loaded from json file
type Configuration struct {
	ListenPort string `yaml:"listenPort"`
	Log        struct {
		FileName    string `yaml:"filename"`
		Level       string `yaml:"level"`
		KafkaOffset string `yaml:"kafkaOffset"`
	} `yaml:"log"`

	KafkaURL string `yaml:"kafkaURL"`
	//KafkaVerbose string `yaml:"kafkaVerbose"`
	//KafkaClient  string `yaml:"kafkaClient"`
	KafkaClientID string `yaml:"kafkaClientID"`
	KafkaTopic    string `yaml:"kafkaTopic"`
	KafkaGroup    string `yaml:"kafkaGroup"`
}

// Param use as global variable for configuration
var Param Configuration

// LoadConfigFromFile use to load global configuration
func LoadConfigFromFile(fn *string) {
	if err := parser.LoadYAML(fn, &Param); err != nil {
		logger.Errorf("LoadConfigFromFile() - Failed opening config file %s\n%s", &fn, err)
		os.Exit(1)
	}
	//logger.Logf("Loaded configs: %v", Param)
	logger.Logf("Config %s", "Loaded")
}
