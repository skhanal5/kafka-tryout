## kafka-tryout

### About
Toy example using Apache Kafka in Golang.

## Local Development

### Setting up Kafka

Run `make rund` to spin up a Docker container with an individual Kafka broker running with KRaft. 

### Interacting with Kafka

1. Set up the following environment variables
```bash
KAFKA_TOPIC="my_topic"
KAFKA_BROKER="localhost:29092"
```
2. Run `make build` to build the executable for the consumer and producer. 
3. Execute the binary for both

There are re-usable scripts in `scripts` that can be referred to. 

### System Observability

#### Prometheus
Prometheus is consuming metrics from kafka-exporter and the jmx-exporter defined in the Docker compose file. The configuration for Prometheus is defined in prometheus/prometheus.yml`.

In the UI, you can view the status of its targets by viewing http://localhost:9090/targets. 

#### Using Grafana
1. Visit the Grafana UI at http://localhost:3000 
2. Sign in using username: z `admin` and password: `password`. 
3. Create a new Dashboard
4. Setup a Prometheus datasource using `http://localhost:9090`
5. Continue the dashboard creation with this [Kafka Overview](https://grafana.com/grafana/dashboards/721-kafka/) template
6. View your metrics

### Container Observability

In case there are issues with the Docker setup, you can make use of Dozzle to easily view each individual container's logs. Visit `http://localhost:8080` to do so.