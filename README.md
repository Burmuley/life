# Convey's "Game of Life" implementation in Go

## Description
 - `lifeform` - instance of a life form
 - `world` - instance of a world (grid) where lifeforms can live and survive
 - `ui` - how world depicted to user

## How to run
```bash
$ go build main.go -o life
$ ./life
```

By default `console` UI will be run. To watch life with GUI use option `-ui`:
```bash
$ ./life -ui gui
```
