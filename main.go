package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	conf "kmp-news-producer/config"
	hdr "kmp-news-producer/handler"
	log "kmp-news-producer/logging"
	mdw "kmp-news-producer/middleware"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Logf("OS: %s", runtime.GOOS)
	log.Logf("architecture: %s", runtime.GOARCH)

	//load config file
	configFile := flag.String("conf", "config/conf.yml", "main configuration file")
	flag.Parse()

	log.Logf("reads configuration from %s", *configFile)
	conf.LoadConfigFromFile(configFile)

	log.Init(conf.Param.Log.Level, conf.Param.Log.FileName)

	// connect to kafka
	kafkaWriter, err := conf.ConfigKafka(strings.Split(conf.Param.KafkaURL, ","), conf.Param.KafkaClientID, conf.Param.KafkaTopic)
	if err != nil {
		log.Errorf("Unable to open kafka %v", err)
		return
	}

	http.HandleFunc("/api/save_news", mdw.Chain(
		hdr.NewsHandler(kafkaWriter),
		mdw.Method("POST"),
		mdw.ContentType("application/json"),
	))

	var errors error
	errors = http.ListenAndServe(conf.Param.ListenPort, nil)

	if errors != nil {
		fmt.Println("error", errors)
		log.Errorf("Unable to start the server: %s ", conf.Param.ListenPort)
		os.Exit(1)
	}
}
