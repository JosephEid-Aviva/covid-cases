# Covid cases tracker

A basic HTTP server used to hit the UK Government's public COVID API for case date.

## Prerequisites

- [Go v1.16](https://golang.org/dl/)

## Getting Started

To install dependencies and start the server:

```sh
go run main.go
```

The HTTP server has a single `/cases` endpoint. Here is a `curl` example.

```sh
curl 'http://localhost:3000/cases?pageStart=1&pageEnd=4'
```
