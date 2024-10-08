package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toggleLights(instructions map[[2]int]string, currentGrid map[[2]int]string) map[[2]int]string {
	for coordinates, instruction := range instructions {
		switch instruction {
		case "on":

			currentGrid[coordinates] = "on"

		case "off":

			currentGrid[coordinates] = "off"

		case "toggle":

			if currentGrid[coordinates] == "off" {
				currentGrid[coordinates] = "on"
			} else if currentGrid[coordinates] == "on" {
				currentGrid[coordinates] = "off"
			}

		}
	}
	return currentGrid
}

func originalGrid() map[[2]int]string {
	origin := make(map[[2]int]string)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			origin[[2]int{i, j}] = "off"
		}
	}

	return origin
}

func getInstructions(s string) map[[2]int]string {
	commands := make(map[[2]int]string)
	// fmt.Println(s)

	inputSlice := strings.Split(s, " ")

	if inputSlice[0] == "toggle" {
		command := "toggle"
		startx, _ := strconv.Atoi(inputSlice[1])
		starty, _ := strconv.Atoi(inputSlice[2])
		endx, _ := strconv.Atoi(inputSlice[4])
		endy, _ := strconv.Atoi(inputSlice[5])

		for i := startx; i <= endx; i++ {
			for j := starty; j <= endy; j++ {
				commands[[2]int{i, j}] = command
			}
		}

	}
	if inputSlice[0] == "turn" && inputSlice[1] == "off" {
		command := "off"
		startx, _ := strconv.Atoi(inputSlice[2])
		starty, _ := strconv.Atoi(inputSlice[3])
		endx, _ := strconv.Atoi(inputSlice[5])
		endy, _ := strconv.Atoi(inputSlice[6])

		for i := startx; i <= endx; i++ {
			for j := starty; j <= endy; j++ {
				commands[[2]int{i, j}] = command
			}
		}

	}
	if inputSlice[0] == "turn" && inputSlice[1] == "on" {
		command := "on"
		startx, _ := strconv.Atoi(inputSlice[2])
		starty, _ := strconv.Atoi(inputSlice[3])
		endx, _ := strconv.Atoi(inputSlice[5])

		endy, _ := strconv.Atoi(inputSlice[6])

		for i := startx; i <= endx; i++ {
			for j := starty; j <= endy; j++ {
				commands[[2]int{i, j}] = command
			}
		}
		return commands
	}

	return commands
}

func lightsOn(currentGrid map[[2]int]string) int {
	var count int
	for _, v := range currentGrid {
		if v == "on" {
			count++
		}
	}
	return count
}

func countLights(s []string) int {
	lineCount := len(s)
	var latestGrid map[[2]int]string

	var instructions map[[2]int]string
	var count int
	latestGrid = originalGrid()

	for i := 0; i < lineCount; i++ {
		inputSlice := strings.ReplaceAll(s[i], ",", " ")

		instructions = getInstructions(inputSlice)

		latestGrid = toggleLights(instructions, latestGrid)
		// fmt.Println(latestGrid)
		count = lightsOn(latestGrid)

	}
	return count

	// fmt.Println(instructions)
}

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	input := bufio.NewReader(file)

	scan := bufio.NewScanner(input)

	var out int
	var Lines []string

	for scan.Scan() {
		line := scan.Text()
		Lines = append(Lines, line)
		// fmt.Println(lines)
		// count += count
	}
	// fmt.Println(lineCount)
	out = countLights(Lines)
	fmt.Println(out)
}
