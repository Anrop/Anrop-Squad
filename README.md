# Anrop Squad API

Returns squad members on Anrop for Arma games.

## Requirements

Code is written in [Go](https://golang.org/) and uses [Go modules](https://github.com/golang/go/wiki/Modules) for dependency management.

## How To Use

Use `go build` to download all dependencies and compile the sources.

Start the API with the `Anrop-Squad` binary.
Server will be available at `$PORT`.

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
