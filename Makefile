all: serve

.PHONY: deploy

run:
	@ grep -v "#" .env | sed 's/.*/export &/' > /tmp/.env
	@ . /tmp/.env && go run main.go
deploy:
	docker-compose -f ./deploy/local/docker-compose.yml up -d


teardown:
	docker-compose -f ./deploy/local/docker-compose.yml stop

