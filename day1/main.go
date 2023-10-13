package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	raw, err := os.ReadFile("./adventofcode.com_2022_day_1_input.txt")
	check(err)
	data := string(raw)

	elfInventories := strings.Split(data, "\n\n")
	elfInventories = elfInventories[0 : len(elfInventories)-1]

	var elfCalories []int

	for i := 0; i < len(elfInventories); i++ {
		inventory := strings.Split(elfInventories[i], "\n")
		totalCalories := 0
		for j := 0; j < len(inventory); j++ {
			calories, err := strconv.Atoi(inventory[j])
			check(err)
			totalCalories = totalCalories + calories
		}

		elfCalories = append(elfCalories, totalCalories)
	}

	slices.SortFunc(elfCalories, func(a, b int) int { return b - a })
	fmt.Println("Highest calories:", elfCalories[0])
	fmt.Println("Top's 3 sum:", elfCalories[0]+elfCalories[1]+elfCalories[2])

	// TODO: Refactor, make it efficient
}
