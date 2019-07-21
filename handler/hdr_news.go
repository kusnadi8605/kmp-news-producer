package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	dts "kmp-news-producer/datastruct"
	log "kmp-news-producer/logging"
	"net/http"

	"github.com/segmentio/kafka-go"
)

//NewsHandler  ..
func NewsHandler(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var rNews dts.News
		var rNewsResponse dts.NewsResponse

		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Logf("Error Read Body request: %v ", err)
		}

		err = json.Unmarshal(body, &rNews)

		log.Logf("Read Body request: %v ", string(body))

		if err != nil {
			rNewsResponse.ResponseCode = "500"
			rNewsResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(rNewsResponse)

			log.Logf("Response News : %v", rNewsResponse)

			return
		}
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
			Value: body,
		}

		err = kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			rNewsResponse.ResponseCode = "300"
			rNewsResponse.ResponseDesc = "Kafka error: " + err.Error()
			json.NewEncoder(w).Encode(rNewsResponse)

			log.Logf("Response News : %v", rNewsResponse)

			return
		}

		rNewsResponse.ResponseCode = "000"
		rNewsResponse.ResponseDesc = "success"
		json.NewEncoder(w).Encode(rNewsResponse)

		log.Logf("Response News : %v", rNewsResponse)

		return

	})
}
