package main

import (
	"strconv"
	"strings"
)

type configuration struct {
	red   int
	green int
	blue  int
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

func getConfigurationPower(config configuration) int {
	return config.red * config.green * config.blue
}
