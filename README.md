# shitwoch

Is it shitwoch yet?

## Run

```
go run cmd/server/*.go
```

## Develop

```
templ genereate
go run cmd/server/*.go
```

## Build

```
templ genereate
CGO_ENABLED=0 go build -o shitwoch cmd/server/*.go
```

## Setup

Install [templ](https://github.com/a-h/templ)
Use [maxmind](https://dev.maxmind.com/geoip/updating-databases) to gather geoip.
