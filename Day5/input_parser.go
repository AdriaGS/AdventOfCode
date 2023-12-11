package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type InputParser interface {
	ParseFile() Almanac
}

type inputParser struct {
	file *os.File
}

func NewInputParser(file *os.File) InputParser {
	return &inputParser{file: file}
}

func (i *inputParser) ParseFile() Almanac {
	almanac := Almanac{}
	var source string
	var destination string
	var rangeMappings []AlmacRangeMapping

	scanner := bufio.NewScanner(i.file)
	for scanner.Scan() {
		line := scanner.Text()

		// Seeds
		if strings.HasPrefix(line, "seeds") {
			parseSeedsLine(line, &almanac)
			continue
		}

		// Mapping title
		if match, _ := regexp.MatchString("^[a-z]+-to-[a-z]+ map:", line); match {
			almanacMap := parseAlmanacMapLine(line)
			source = almanacMap[0]
			destination = almanacMap[1]
			rangeMappings = []AlmacRangeMapping{}
			continue
		}

		// Mappings
		if match, _ := regexp.MatchString("^[\\d]+ [\\d]+ [\\d]+", line); match {
			rangeMapping := parseAlmanacMappingLine(line)
			rangeMappings = append(rangeMappings, rangeMapping)
			continue
		}

		// Empty line
		if source == "" || destination == "" {
			continue
		}
		almanacMap := AlmanacMap{
			Source:        source,
			Destination:   destination,
			RangeMappings: rangeMappings,
		}
		almanac.Mappings = append(almanac.Mappings, almanacMap)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return almanac
}


func convertStringSliceToIntSlice(stringSlice []string) []int {
	intSlice := make([]int, len(stringSlice))
	var err error
	for i, s := range stringSlice {
		intSlice[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return intSlice
}

func parseSeedsLine(line string, almanac *Almanac) {
	lineParts := strings.Split(line, ":")
	seeds := strings.Split(strings.TrimSpace(lineParts[1]), " ")

	// For Part Two
	fullSeedList := []int{}
	for i := 0; i < len(seeds); i+=2 {
		start, _ := strconv.Atoi(seeds[i])
		length, _ := strconv.Atoi(seeds[i+1])
		seedList := generateList(start, length)
		fullSeedList = append(fullSeedList, seedList...)
	}

	almanac.Seeds = fullSeedList
}

func generateList(start int, length int) []int {
	list := make([]int, length)
  
	current := start
	for i := 0; i < length; i++ {
	  list[i] = current
	  current++ 
	}
  
	return list
  }
  

func parseAlmanacMapLine(line string) []string {
	regex := regexp.MustCompile(`^([a-z]+)-to-([a-z]+) map:`)
	matches := regex.FindStringSubmatch(line)
	if len(matches) != 3 {
		panic("regex failed getting the amanac map values")
	}
	return []string(matches[1:])
}

func parseAlmanacMappingLine(line string) AlmacRangeMapping {
	mappings := strings.Split(strings.TrimSpace(line), " ")
	parsedMapping := convertStringSliceToIntSlice(mappings)
	return AlmacRangeMapping{
		DestinationRangeStart: parsedMapping[0],
		SourceRangeStart:      parsedMapping[1],
		RangeLength:           parsedMapping[2],
	}
}