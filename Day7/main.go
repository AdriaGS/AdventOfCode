package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var CARD_SUITE map[string]int = map[string]int{"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1}
var CARD_SUITE map[string]int = map[string]int{"A": 13, "K": 12, "Q": 11, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1, "J": 0}

type OrderedHands struct {
  Hands []string
  Values []int
}

func getLowerHand(hand1 string, hand2 string) string {
  for i, card := range hand1 {
    if string(card) == string(hand2[i]) {
      continue
    }
    if CARD_SUITE[string(card)] < CARD_SUITE[string(hand2[i])] {
      return hand1
    }
    return hand2
  }
  return hand1
}

func (oh *OrderedHands) Append(hand string, value int) {
  for i, v := range oh.Values {
    if value < v || (value == v && getLowerHand(oh.Hands[i], hand) == hand) {
      oh.Hands = append(oh.Hands, "")
      oh.Values = append(oh.Values, 0)
      copy(oh.Hands[i+1:], oh.Hands[i:])
      copy(oh.Values[i+1:], oh.Values[i:])
      oh.Hands[i] = hand
      oh.Values[i] = value
      return
    }
  }
  oh.Hands = append(oh.Hands, hand)
  oh.Values = append(oh.Values, value)
}

func parseLine(line string) (string, int) {
	hand := strings.Split(line, " ")
	cards := hand[0]
	bid, _ := strconv.Atoi(hand[1])
	return cards, bid
}

// Value: 
// 6 --> Five of a kind
// 5 --> Four of a kind
// 4 --> Full house
// 3 --> Three of a kind
// 2 --> Two pair
// 1 --> Pair
// 0 --> High card

func getHandValue(cards string) int {
	cardMap := map[string]int{}
  maxValue := 0
  maxCard := ""
	for _, card := range cards {
		cardMap[string(card)] += 1
    if string(card) != "J" && cardMap[string(card)] > maxValue {
      maxValue = cardMap[string(card)]
      maxCard = string(card)
    }
	}

  if _, ok := cardMap["J"]; ok {
    maxValue += cardMap["J"]
    if maxCard == "" {
      maxCard = "A"
    }
    cardMap[maxCard] = maxValue
    delete(cardMap, "J")
  }

  if len(cardMap) == 1 {
    return 6
  }
  if len(cardMap) == 2 {
    if maxValue == 4 {
      return 5
    }
    return 4
  }
  if len(cardMap) == 3 {
    if maxValue == 3 {
      return 3
    }
    return 2
  }
  if len(cardMap) == 4 {
    return 1
  }
  return 0
}

func getOrderedHands(cardHands []string) []string {
  orderedHands := OrderedHands{}
  for _, cardHand := range cardHands {
    value := getHandValue(cardHand)
    orderedHands.Append(cardHand, value)
  }
  return orderedHands.Hands
}

func getTotalWinnings(cardRanking []string, handsMap map[string]int) int {
  totalWinnings := 0
  
  orderedHands := getOrderedHands(cardRanking)

  for i, hand := range orderedHands {
    totalWinnings += handsMap[hand] * (i + 1)
  }

  return totalWinnings
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

	handsMap := map[string]int{}
  cardHands := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cards, bid := parseLine(line)
		handsMap[cards] = bid
    cardHands = append(cardHands, cards)
	}

  totalWinnings := getTotalWinnings(cardHands, handsMap)
  fmt.Printf("Total winnings: %d\n", totalWinnings)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
