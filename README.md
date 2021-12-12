# Infra Processor 
This is a microservice which will read cpu metrics from a kafka topic and write it to a file.

## How to run
In order to run the executable, export the following environment variables:
| Variable Name | Description | Defaults (If any) |
| ------ | ------ | ------|
| TOPIC | Name of the topic which is created on kafka |  |
| BROKER_ADDRESS | Address of broker service in format host:port | |
| GROUP_ID | Group ID for a group of consumers. Can be any string value | |
| OUT_FILE_PATH | Location including filename on which the cpu metrics will be written | |
| RUN_TIMEOUT | If you want to run for a limited time, you can do so by passing seconds as integer. For indefinite running, pass -1 | -1 |

## Docker image
Docker image has been published on docker hub. Pull it using below command:
```
docker pull tejaswikasat/infraprocessor:latest
```
While running it, set the above variables as environment variables