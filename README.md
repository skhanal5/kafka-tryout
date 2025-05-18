## kafka-tryout

### About
Toy example using Apache Kafka in Golang.

### Local Development

#### 1. Setting up Kafka 
Run `make rund` to spin up a Docker container with an individual Kafka broker running with KRaft. 


#### 2. Producer & Consumer
Run `make build` to build the executable for the consumer and producer. The following environment variables need to be passed in for both programs:

```bash
KAFKA_TOPIC = "my_topic"
KAFKA_BROKER = "localhost:9082"
```

Once this is set, you can run the producer in one terminal and the consumer in another terminal.