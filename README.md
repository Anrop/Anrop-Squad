# Anrop Squad API

Returns squad members on Anrop for Arma games.

## Requirements

Code is written in [Go](https://golang.org/) and uses [gb](https://getgb.io/) to compile.

## How To Use

Compile the sources with `gb build`.

Start the API with the `squad` binary in the `bin` folder.
Server will be available at `$PORT`.

### Static Data

Static data is embedded using [go-bindata](https://github.com/jteeuwen/go-bindata)

Install `go-bindata` with `go get -u github.com/jteeuwen/go-bindata/...`

Regenerate static assets using `go-bindata -o src/static/static.go -pkg static static/`

## Environment Variables

Environment variables can be specified in `.env` file and will be autoloaded.

| Key | Required | Description |
| --- | -------- | ----------- |
| DATABASE_URL | Yes | MySQL url to your Anrop database |
| PORT | No | Port that HTTP Server is bound to. 8080 by default |

## Endpoints

### /arma1.json

Returns Arma 1 squad members as JSON

### /arma2.json

Returns Arma 2 squad members as JSON

### /arma3.json

Returns Arma 3 squad members as JSON

### /ofp.json

Returns OFP squad members as JSON
