package main

import "math"

type AlmacRangeMapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type AlmanacMap struct {
	Source        string
	Destination   string
	RangeMappings []AlmacRangeMapping
}

func (am *AlmanacMap) mapValue(source int) int {
	for _, r := range am.RangeMappings {
		if source >= r.SourceRangeStart && source < r.SourceRangeStart+r.RangeLength {
			return r.DestinationRangeStart + (source - r.SourceRangeStart)
		}
	}
	return source
}

type Almanac struct {
	Seeds    []int
	Mappings []AlmanacMap
}

func (a *Almanac) findMapBySource(source string) *AlmanacMap {
	for _, m := range a.Mappings {
		if m.Source == source {
			return &m
		}
	}
	return nil
}

func (a *Almanac) findLowestLocationNumber() int {
	lowestLocationValue := math.MaxInt
	for _, seed := range a.Seeds {
		locationValue := a.findSeedMapping(seed, "location")
		if locationValue < lowestLocationValue {
			lowestLocationValue = locationValue
		}
	}
	return lowestLocationValue
}

func (a *Almanac) findSeedMapping(seed int, destination string) int {
	return a.mapFromSourceToDestination(seed, "seed", destination)
}

func (a *Almanac) mapFromSourceToDestination(value int, source string, destination string) int {
	am := a.findMapBySource(source)
	if am == nil {
		panic("Couldn't find Almanac Map by source")
	}
	mappedValue := am.mapValue(value)
	if am.Destination != destination {
		return a.mapFromSourceToDestination(mappedValue, am.Destination, destination)
	}
	return mappedValue
}
