package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// LineOldYin - A broken line changing into a solid line.
	LineOldYin = 6

	// LineYoungYang - A solid line.
	LineYoungYang = 7

	// LineYoungYin - A broken line.
	LineYoungYin = 8

	// LineOldYang - A solid line changing into a broken line.
	LineOldYang = 9
)

const (
	// Yin ..
	Yin = false
	// Yang ..
	Yang = true
)

// Line represents a single coin toss for the IChing.
type Line struct {
	CoinFlips [3]string
	Total     int
}

func (l *Line) toString() string {
	switch l.Total {
	case LineOldYin:
		return "--X--"
	case LineYoungYang:
		return "-----"
	case LineYoungYin:
		return "-- --"
	case LineOldYang:
		return "--0--"
	}
	return "uninitialized"
}

func (l *Line) toBool() bool {
	if l.Total == LineOldYin || l.Total == LineYoungYin {
		return Yin
	} else if l.Total == LineOldYang || l.Total == LineYoungYang {
		return Yang
	} else {
		panic("Not Yin or Yang.")
	}
}

type Hexagram struct {
	// Number given to this hexagram.
	Number int
	// The series of lines that compose this hexagram.
	Lines [6]Line
	// Name of the hexagram
	Name string
	// The unicode character of the hexagram.
	Character string
	// Description
	Description string
}

// Reading ..
type Reading struct {
	Lines [6]Line
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
