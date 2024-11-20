- On kafka, microservies can publish events
- microservices can subscribe to vents
- kafka can store events for later retrieval
- kafka can replay events for later retrival
- kafka can transform or process events


What we are building
1. Producer
Fiber(3000) -> ap/V1 -> /comments -> createComment() -> InitializeCommentStruct -> parse body request to comment -> convert comment into bytes and end i to kafka -> Push commentToBytes -> Define the brokersURL

2. Consumer
setTheTopic -> connectConsumer usning sarama.NewConfig
    |> consumePartition


`sarama` is a go clint library
kaka will be running in docker containers 

```bash
go mod init github.com/gaurav-225/golang-projects/EDA_withGOAndKafka
mkdir producer worker
touch producer/producer.go worker/worker.go
```

```bash
go get github.com/gofiber/fiber/v2
go get github.com/IBM/sarama
```


# Docker command
```bash
docker run -d --name zookeeper -p 22181:2181 \
-e ZOOKEEPER_CLIENT_PORT=2181 \
-e ZOOKEEPER_TICK_TIME=2000 \
confluentinc/cp-zookeeper:latest
```

```bash
docker run -d --name kafka \
--link zookeeper:zookeeper \
-e KAFKA_BROKER_ID=1 \
-e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:29092 \
-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT \
-e KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT \
-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
-p 29092:29092 \
confluentinc/cp-kafka:latest
```