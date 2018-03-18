package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Line represents a single coin toss for the IChing.
type Line struct {
	CoinFlips [3]string
	Total     int
}

func (l *Line) toString() string {
	switch l.Total {
	case 6:
		return "--X--"
	case 7:
		return "-----"
	case 8:
		return "-- --"
	case 9:
		return "--0--"
	}
	return "uninitialized"
}

func flipCoin() string {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(2)
	if num == 0 {
		return "Tails"
	}
	return "Heads"
}

func main() {
	for i := 0; i < 6; i++ {
		var line Line
		for j := 0; j < 3; j++ {
			res := flipCoin()
			line.CoinFlips[j] = res
			if res == "Tails" {
				line.Total += 2
			} else {
				line.Total += 3
			}
		}
		fmt.Println(line.toString())
	}
}
