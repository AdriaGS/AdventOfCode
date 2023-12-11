package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func parseLine(line string) []int {
	splitLine := strings.Split(line, ":")
	// Part One
	// values := strings.Split(strings.TrimSpace(splitLine[1]), " ")
	// Part Two
	values := []string{strings.ReplaceAll(splitLine[1], " ", "")}
	return convertStringSliceToIntSlice(values)
}

func findWaysToWin(time int, record int) int {
	waysToWin := 0

	// Checking half of the time and multiplying for two the result
	for i := 0; i <= time/2; i++ {
		speed := 1 * i
		distance := (time - i) * speed
		if distance > record {
			waysToWin++
		}
	}

	waysToWin *= 2
	if (time % 2) == 0 {
		waysToWin -= 1
	}

	return waysToWin
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

	times := []int{}
	records := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		space := regexp.MustCompile(`\s+`)
		line := space.ReplaceAllString(scanner.Text(), " ")

		if strings.HasPrefix(line, "Time") {
			times = parseLine(line)
			continue
		}

		if strings.HasPrefix(line, "Distance") {
			records = parseLine(line)
			continue
		}
	}

	waysToWin := 1
	for i := 0; i < len(times); i++ {
		waysToWin *= findWaysToWin(times[i], records[i])
	}
	fmt.Printf("Found %d ways to win", waysToWin)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
