package main

import (
	"fmt"
	"math/rand"
	"time"
)

func flipCoin() string {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(2)
	if num == 0 {
		return "Tails"
	}
	return "Heads"
}

func main() {
	i := 0
	for i < 6 {
		j := 0
		total := 0
		line := make([]string, 0)
		for j < 3 {
			res := flipCoin()
			line = append(line, res)
			if res == "Tails" {
				total += 2
			} else {
				total += 3
			}
			j++
		}
		if total == 6 {
			fmt.Println("--X--")
		} else if total == 7 {
			fmt.Println("-----")
		} else if total == 8 {
			fmt.Println("-- --")
		} else {
			fmt.Println("--0--")
		}

		// fmt.Println(line)
		// fmt.Println(total)
		i++
	}
}
