# kmp-news-producer
Golang &amp; Kafka

## add lib
go get -v github.com/go-sql-driver/mysql \n
go get -v gopkg.in/olivere/elastic.v7
go get -v github.com/segmentio/kafka-go
go get -v github.com/go-yaml/yaml

# running kafka
## running zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
## running kafka server
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
