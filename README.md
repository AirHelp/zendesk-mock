## Download

```bash
git clone ...
or
go get github.com/Airhelp/zendesk-mock (requires go installed and GOPATH exported)
cd zendesk-mock
```

## Build
```bash
env GOOS=linux GOARCH=386 go build
docker build -t airhelp/zendesk-mock .
```

## Run locally on docker
```bash
docker run -p 8080:8080 airhelp/zendesk-mock
```

## Push to docker hub
```bash
docker push airhelp/zendesk-mock
```
