package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type configuration struct {
	red   int
	green int
	blue  int
}

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
		configurationPowers = append(configurationPowers, getPower(getLowestPossibleConfiguration(gameInput)))
		game++
	}

	sumOfPossibleGames := sum(possibleGames)
	fmt.Printf("sum of possible games with given configuration: %v\n", sumOfPossibleGames)

	sumOfConfigurationPowers := sum(configurationPowers)
	fmt.Printf("sum of power of minimal configurations: %v\n", sumOfConfigurationPowers)
}

func isGamePossible(gameInput string, config configuration) bool {
	splitted := strings.Split(gameInput, ":")
	game := splitted[1]
	sets := strings.Split(game, ";")

	for _, set := range sets {
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			switch {
			case strings.HasSuffix(cube, "red"):
				reds := strings.Split(cube, " ")
				if v, _ := strconv.Atoi(reds[1]); v > config.red {
					return false
				}
			case strings.HasSuffix(cube, "green"):
				greens := strings.Split(cube, " ")
				if v, _ := strconv.Atoi(greens[1]); v > config.green {
					return false
				}
			case strings.HasSuffix(cube, "blue"):
				blues := strings.Split(cube, " ")
				if v, _ := strconv.Atoi(blues[1]); v > config.blue {
					return false
				}
			}
		}
	}
	return true
}

func getLowestPossibleConfiguration(gameInput string) configuration {
	result := configuration{0, 0, 0}

	splitted := strings.Split(gameInput, ":")
	game := splitted[1]
	sets := strings.Split(game, ";")

	for _, set := range sets {
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			switch {
			case strings.HasSuffix(cube, "red"):
				reds := strings.Split(cube, " ")
				if v, _ := strconv.Atoi(reds[1]); v > result.red {
					result.red = v
				}
			case strings.HasSuffix(cube, "green"):
				greens := strings.Split(cube, " ")
				if v, _ := strconv.Atoi(greens[1]); v > result.green {
					result.green = v
				}
			case strings.HasSuffix(cube, "blue"):
				blues := strings.Split(cube, " ")
				if v, _ := strconv.Atoi(blues[1]); v > result.blue {
					result.blue = v
				}
			}
		}
	}

	return result
}

func getPower(config configuration) int {
	return config.red * config.green * config.blue
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
