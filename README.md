# Anrop Squad API

Returns squad members on Anrop for Arma games.

## Requirements

Code is written in [Go](https://golang.org/) and uses [gb](https://getgb.io/) to compile.

## How To Use

Compile the sources with `gb build`.

Start the API with the `squad` binary in the `bin` folder.
Server will be available at `$PORT`.

## Environment Variables

Environment variables can be specified in `.env` file and will be autoloaded.

| Key | Required | Description |
| --- | -------- | ----------- |
| DATABASE_URL | Yes | MySQL url to your Anrop database |
| PORT | No | Port that HTTP Server is bound to. 8080 by default |

## Endpoints

### /arma2.json

Returns Arma 2 squad members as JSON

### /arma3.json

Returns Arma 3 squad members as JSON
