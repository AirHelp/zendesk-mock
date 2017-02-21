## Download

```bash
git clone ...
or
go get github.com/Airhelp/zendesk-mock (requires go installed and GOPATH exported)
cd $GOPATH/src/github.com/AirHelp/zendesk-mock
```

## Build
```bash
make release
```

## Run locally on docker
```bash
docker run -p 8080:8080 airhelp/zendesk-mock
```

## Push to docker hub
```bash
make push
```

## Execute tests

```
make test
```
