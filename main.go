package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	conf "kmp-news-producer/config"
	hdr "kmp-news-producer/handler"
	log "kmp-news-producer/logging"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Logf("OS: %s", runtime.GOOS)
	log.Logf("architecture: %s", runtime.GOARCH)
	configFile := flag.String("conf", "config/conf.yml", "main configuration file")
	flag.Parse()

	log.Logf("reads configuration from %s", *configFile)
	conf.LoadConfigFromFile(configFile)

	log.Init(conf.Param.Log.Level, conf.Param.Log.FileName)

	// connect to kafka
	kafkaWriter, err := conf.Configure(strings.Split(conf.Param.KafkaURL, ","), conf.Param.KafkaClientID, conf.Param.KafkaTopic)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	http.HandleFunc("/api/save_news", hdr.Chain(hdr.NewsHandler(kafkaWriter), hdr.Method("POST"), hdr.ContentType("application/json")))

	var errors error
	errors = http.ListenAndServe(conf.Param.ListenPort, nil)

	if errors != nil {
		fmt.Println("error", errors)
		log.Logf("Unable to start the server: %s ", conf.Param.ListenPort)
		os.Exit(1)
	}
}
