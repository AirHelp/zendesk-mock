build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o zendesk-mock .

release: build
	docker build -t airhelp/zendesk-mock -f Dockerfile .

push: release
	docker push airhelp/zendesk-mock
