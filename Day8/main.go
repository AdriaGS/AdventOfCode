package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
  "regexp"
  "strings"
)

var COMMAND_MAPPING = map[string]int{"L": 0, "R": 1}
var KNOWN_PATHS = map[string][]string{}

func parseMapLine(line string) (string, []string) {
  splitLine := strings.Split(line, "=")
  origin := strings.TrimSpace(splitLine[0])
  r, _ := regexp.Compile(`\(([^,]+), ([^,]+)\)`)
  destination := r.FindStringSubmatch(splitLine[1])
  return origin, destination[1:]
}

func getStepsToDestination(desertMap map[string][]string, commands []string, source string, destination string) int {
  steps := 0
  commandIndex := 0
  currentPosition := source
  for {
    steps++
    currentPosition = desertMap[currentPosition][COMMAND_MAPPING[commands[commandIndex]]]
    if currentPosition == destination {
      break
    }
    commandIndex = (commandIndex + 1) % len(commands)
  }
  return steps
}

func getStepsToFirstDestination(desertMap map[string][]string, commands []string, source string) int {
  steps := 0
  commandIndex := 0
  coordinate := source
  for {
    steps++
    coordinate = desertMap[coordinate][COMMAND_MAPPING[commands[commandIndex]]]
    if coordinate[len(coordinate)-1] == 'Z' {
      break
    }

    commandIndex = (commandIndex + 1) % len(commands)
  }
  return steps
}

func greatestCommonDenominator(a, b int) int {
  for b != 0 {
          t := b
          b = a % b
          a = t
  }
  return a
}

func leastCommonMultiple(a, b int, integers ...int) int {
  result := a * b / greatestCommonDenominator(a, b)

  for i := 0; i < len(integers); i++ {
          result = leastCommonMultiple(result, integers[i])
  }

  return result
}

func stepsToDestinationGhostMode(desertMap map[string][]string, commands []string, sourceCoordinates []string) int {
  loopValues := []int{}
  for _, coordinate := range sourceCoordinates {
    loop := getStepsToFirstDestination(desertMap, commands, coordinate)
    fmt.Printf("Loop for %s is %d\n", coordinate, loop)
    loopValues = append(loopValues, loop)
  }

  return leastCommonMultiple(loopValues[0], loopValues[1], loopValues[2:]...)
}

func main() {

	filename := flag.String("file", "", "filename to read")

	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		return
	}

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

  commands := []string{}
  desertMap := map[string][]string{}
  startingCoordinates := []string{}

  r := regexp.MustCompile(`^[LR]+$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

    if r.MatchString(line) {
      commands = strings.Split(line, "")
    } else if len(line) != 0 {
      origin, destination := parseMapLine(line)
      desertMap[origin] = destination

      if (origin[len(origin)-1] == 'A') {
        startingCoordinates = append(startingCoordinates, origin)
      }
    }
	}

  // Part One
  // stepsToDestination := getStepsToDestination(desertMap, commands, "AAA", "ZZZ")

  stepsToDestination := stepsToDestinationGhostMode(desertMap, commands, startingCoordinates)
  fmt.Println(stepsToDestination)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
