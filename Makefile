.PHONY: all clean format vet build

all: clean fmt vet build

##### Golang specific

clean:
	go clean
	rm -rf ./bin

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	mkdir -p bin
	go build -o bin/producer cmd/kafka-tryout/producer/main.go
	go build -o bin/consumer cmd/kafka-tryout/consumer/main.go

###### Docker Specific 

topic:
	docker exec -it broker /opt/kafka/bin/kafka-topics.sh --create \
	--topic my-topic \
	--partitions 3 \
	--replication-factor 1 \
	--bootstrap-server localhost:9092

cleand:
	docker compose down --volumes --remove-orphans

rund: cleand
	docker compose up --build -d
	$(MAKE) topic

logs:
	docker logs -f broker