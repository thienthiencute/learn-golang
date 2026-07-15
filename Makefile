include .env
export 


.PHONY: server help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

server:
	templ generate && go run main.go server

pubsub:
	go run main.go pubsub

schedule:
	go run main.go schedule

dev:
	air