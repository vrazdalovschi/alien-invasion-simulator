# alien-invasion-simulator
Mad aliens are about to invade the earth and you are tasked with simulating the invasion.

## Overview

A CLI application to simulate aliens invasion described in the [task](./resources/TASK.md).

## Usage
### Build

```bash
go build -o alien-invasion-simulator main.go
```

### Test

```bash
go test -v -cover ./...
```

### Run

```bash
./alien-invasion-simulator -h
```

```
Run the simulation of the invasion.

Usage:
   [flags]

Examples:
./alien-invasion-simulator ./resources/example-map.txt --aliens 100 --iterations 10000

Flags:
  -a, --aliens int       Number of aliens to invade. (default 10)
  -h, --help             help for this command
  -i, --iterations int   The maximum number of iterations to run. (default 1000)
  -v, --verbose          Verbosity level, print all aliens steps.
```

#### Run a test simulation with [example map](./resources/example-map.txt)


```bash
make simulate-example 
```
