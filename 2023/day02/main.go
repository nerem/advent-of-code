package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	utils "github.com/nerem/advent-of-code"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := configuration{
		red:   12,
		green: 13,
		blue:  14,
	}

	possibleGames := make([]int, 100)
	configurationPowers := make([]int, 100)
	game := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameInput := scanner.Text()
		if isGamePossible(gameInput, config) {
			possibleGames = append(possibleGames, game)
		}
		configurationPowers = append(configurationPowers, getConfigurationPower(getLowestPossibleConfiguration(gameInput)))
		game++
	}

	sumOfPossibleGames := utils.Sum(possibleGames)
	fmt.Printf("sum of possible games with given configuration: %v\n", sumOfPossibleGames)

	sumOfConfigurationPowers := utils.Sum(configurationPowers)
	fmt.Printf("sum of power of minimal configurations: %v\n", sumOfConfigurationPowers)
}
