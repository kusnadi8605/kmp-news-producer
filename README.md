# kmp-news-producer
Golang &amp; Kafka

## Add Library
go get -v github.com/go-sql-driver/mysql   
go get -v gopkg.in/olivere/elastic.v7  
go get -v github.com/segmentio/kafka-go  
go get -v github.com/go-yaml/yaml  

# Running Kafka
## Running Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
## Running Kafka Server
bin/kafka-server-start.sh config/server.properties

# Running App
go run main.go

## Post News
url -X POST \
  http://localhost:8181/api/save_news \
  -H 'Content-Type: application/json' \
  -d '{
	
	"author":"kusnadi",
	"body":" beritaku hari ini"
}'
