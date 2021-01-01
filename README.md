[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/onigra/wasataro/blob/main/LICENSE)
![Test](https://github.com/onigra/wasataro/workflows/UnitTest/badge.svg)

# wasataro

wasataro is one-liner http server cli.

## Install

```sh
go get -v github.com/onigra/wasataro/cmd/wasataro
```

## Binary packages

[Releases](https://github.com/onigra/wasataro/releases)

## Docker images

[GitHub Packages](https://github.com/users/onigra/packages/container/package/wasataro)

## Usage

### CLI

Start http server on 3000 port(default).  
Access to http://localhost:3000

```sh
$ wasataro
```

Change port.
Access to http://localhost:8000

```sh
$ wasataro -port=8000
```

Change resposne.

```sh
$ wasataro -response='{ "response": "ok" }'
```

### Docker

Basic usage.

```sh
$ docker run --rm -p 3000:3000 ghcr.io/onigra/wasataro:latest
```

Use options.

```sh
$ docker run --rm -p 8000:8000 \
  ghcr.io/onigra/wasataro:latest \
  -port=8000 \
  -response='use options'
```

## About name

ワンライナーHTTPサーバー太郎 -> One-liner http server taro -> wasataro

## LICENCE

MIT

## Author

onigra
