package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type day8 struct {
	part1ConnectionCount int
}

func init() {
	registerDay(8, day8{1000})
}

func (d day8) part1(input string) string {
	junctions := d.parseInput(input)
	connections := d.shortestConnections(junctions)

	circuits := make(map[threeCoord]int)
	for i, j := range junctions {
		circuits[j] = i
	}

	for i := 0; i < d.part1ConnectionCount; i++ {
		circuitA := circuits[*connections[i].a]
		circuitB := circuits[*connections[i].b]

		for junction, circuit := range circuits {
			if circuit == circuitB {
				circuits[junction] = circuitA
			}
		}
	}

	circuitSizes := make(map[int]int)
	for _, circuit := range circuits {
		circuitSizes[circuit]++
	}

	a, b, c := 1, 1, 1
	for _, size := range circuitSizes {
		if size > a {
			a, b, c = size, a, b
		} else if size > b {
			b, c = size, b
		} else if size > c {
			c = size
		}
	}

	return fmt.Sprintf("%d", a*b*c)
}

func (d day8) part2(input string) string {
	junctions := d.parseInput(input)
	connections := d.shortestConnections(junctions)

	circuits := make(map[threeCoord]int)
	for i, j := range junctions {
		circuits[j] = i
	}

	var lastConnection junctionConnection
	for i, unified := 0, false; !unified && i < len(connections); i++ {
		lastConnection = connections[i]
		circuitA := circuits[*connections[i].a]
		circuitB := circuits[*connections[i].b]

		for junction, circuit := range circuits {
			if circuit == circuitB {
				circuits[junction] = circuitA
			}
		}

		unified = true
		for _, circuit := range circuits {
			if circuit != circuitA {
				unified = false
				break
			}
		}
	}

	return fmt.Sprintf("%d", lastConnection.a[0]*lastConnection.b[0])
}

type threeCoord [3]int

func (day8) parseInput(input string) []threeCoord {
	var coords []threeCoord
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		split := strings.Split(line, ",")
		if len(split) != 3 {
			panic("bad line: " + line)
		}

		a, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(split[2])
		if err != nil {
			panic(err)
		}

		coords = append(coords, threeCoord{a, b, c})
	}
	return coords
}

type junctionConnection struct {
	a *threeCoord
	b *threeCoord
}

func (j junctionConnection) distance() float64 {
	return math.Sqrt(math.Pow(float64(j.a[0]-j.b[0]), 2) +
		math.Pow(float64(j.a[1]-j.b[1]), 2) +
		math.Pow(float64(j.a[2]-j.b[2]), 2))
}

func cmpJunctionConnection(a, b junctionConnection) int {
	ad := a.distance()
	bd := b.distance()

	if ad < bd {
		return -1
	} else if ad == bd {
		return 0
	} else {
		return 1
	}
}

func (day8) shortestConnections(junctions []threeCoord) []junctionConnection {
	var connections []junctionConnection
	for i := 0; i < len(junctions); i++ {
		for j := i + 1; j < len(junctions); j++ {
			connections = append(connections, junctionConnection{&junctions[i], &junctions[j]})
		}
	}
	slices.SortFunc(connections, cmpJunctionConnection)
	return connections
}
