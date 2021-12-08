# Infra Processor 
This is a microservice which will read data from a kafka topic and write it to a file.

## How to run
Run the service by passing the topic name and the bootstrap server address.
```
infraprocessor --topic CPU_Topic --bootstrap-server localhost:9092
```