build:
	docker build -t airhelp/zendesk-mock .

push: build
	docker push airhelp/zendesk-mock

tests:
	docker-compose run --rm tests

dev:
	go build
