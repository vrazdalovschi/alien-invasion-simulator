![Go Tests](https://github.com/vrazdalovschi/alien-invasion-simulator/workflows/Go/badge.svg)
# alien-invasion-simulator

Mad aliens are about to invade the earth and you are tasked with simulating the invasion.

## Overview

A CLI application to simulate aliens invasion described in the [task](./resources/TASK.md).

The simulation is based on the 3 steps:

1. Read the world map from a file
2. Place the aliens on the map
3. Move the aliens to the city direction and validate the STOP conditions:
    1. Simulation achieved the maximum number of steps.
    2. All aliens are killed.
    3. Remained only 1 alien.
4. Complete simulation, print the final result.

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

Output example of the simulation:

```sh
Running simulation with 10 aliens and 1000 maximum iterations to run on map ./resources/example-map.txt.
Placing aliens:
  * Alien 0 placed in Bar
  * Alien 1 placed in Baz
  * Alien 2 placed in Baz
  * Alien 3 placed in Bar
  * Alien 4 placed in Bar
  * Alien 5 placed in Bar
  * Alien 6 placed in Qu-ux
  * Alien 7 placed in Baz
  * Alien 8 placed in Baz
  * Alien 9 placed in Baz
Step: 0
  * Alien 8 moved from Baz to Foo
  * Alien 9 moved from Baz to Foo
  * Alien 0 moved from Bar to Foo
  * Alien 2 moved from Baz to Foo
  * Alien 3 moved from Bar to Foo
  * Alien 5 moved from Bar to Foo
  * Alien 7 moved from Baz to Foo
  * Alien 1 moved from Baz to Foo
  * Alien 4 moved from Bar to Bee
  * Alien 6 moved from Qu-ux to Foo
Step: 0, City: Foo has been destroyed by cityAliens: 8, 9, 6, 1, 0, 2, 3, 5, 7
Step: 0 Last one alien is alive, stop simulation
Simulation finished.
Statistics:
  * Aliens: 4
  * World:
     * Foo north=Bar west=Baz south=Qu-ux Aliens: 
     * Bar west=Bee Aliens: 
     * Baz Aliens: 
     * Qu-ux Aliens: 
     * Bee east=Bar Aliens: 4
```

## Notes:

### Expectations:

* City name doesn't include spaces.
* All directions follow the format `direction=city`
* All cities are linked bidirectionally by direction road.
* All cities can have only unique directions, the same direction cannot be used for 2 or more cities.

### Improvement points

The current version of the project follows the task description, but it can be extended with the points below:

* Random map generator.
* Map visualizer (console/graphical) before and after simulation.
* Aliens placement and movement algorithm with ability to seed predefined "random" number generator.
* City can reuse the same direction multiple times, example: `Foo north=[Bar, Baz]`
* Current tests cover only the most basic cases, but it can be extended with more tests.
