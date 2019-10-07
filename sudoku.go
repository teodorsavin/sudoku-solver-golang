package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	sudoku = [][]int{}
)

const matrixSize = 9

func main() {
	// Initialize Sudoku Matrix
	initializeSudokuMatrix()

	matrix := readFile()

	// Parse the numbers in the line
	ints, err := parseInts(matrix)
	if err != nil {
		log.Fatal(err)
	}

	finished := true

	for {
		finished = true
		for i := 0; i < matrixSize; i++ {
			for j := 0; j < matrixSize; j++ {

				// Reset
				posibilitiesHorizontal := []int{}
				posibilities := []int{}
				posibilitiesSquare := []int{}

				if ints[i][j] == -1 {
					finished = false

					// Check Horizontal Posibilities
					posibilitiesHorizontal = checkHorizontal(i, j, ints)

					if len(posibilitiesHorizontal) == 1 {
						ints[i][j] = posibilitiesHorizontal[0]
						continue
					} 

					// Check Vertical Posibilities
					if len(posibilitiesHorizontal) > 1 {
						posibilities = checkVertical(i, j, ints, posibilitiesHorizontal)

						if len(posibilities) == 1 {
							ints[i][j] = posibilities[0]
							continue
						}
					}

					// Check Square Posibilities
					if len(posibilities) > 1 {
						posibilitiesSquare = checkSquare(i, j, ints, posibilities)

						if len(posibilitiesSquare) == 1 {
							ints[i][j] = posibilitiesSquare[0]
						}
					}

					posibilitiesHorizontal = nil
					posibilities = nil
					posibilitiesSquare = nil
				}
			}
		}

		if finished == true {
			break
		}
	}

	// Print the numbers
	for i := 0; i < matrixSize; i++ {
		fmt.Println(ints[i])
	}
}

func checkHorizontal(i int, j int, matrix [][]int) []int {
	alreadyTaken := []int{}
	posibilities := []int{}

	for x := 0; x < matrixSize; x++ {
		if matrix[i][x] != -1 {
			alreadyTaken = append(alreadyTaken, matrix[i][x])
		}
	}

	for y := 1; y <= matrixSize; y++ {
		taken := false
		for x := 0; x < len(alreadyTaken); x++ {
			if alreadyTaken[x] == y {
				taken = true
				continue
			}
		}

		if taken == false {
			posibilities = append(posibilities, y)
		}
	}

	return posibilities
}

func checkVertical(i int, j int, matrix [][]int, posibilitiesBefore []int) []int {
	alreadyTaken := []int{}
	posibilitiesAfter := []int{}
	posibilities := []int{}

	for x := 0; x < matrixSize; x++ {
		if matrix[x][j] != -1 {
			alreadyTaken = append(alreadyTaken, matrix[x][j])
		}
	}

	for y := 1; y <= matrixSize; y++ {
		taken := false
		for x := 0; x < len(alreadyTaken); x++ {
			if alreadyTaken[x] == y {
				taken = true
				continue
			}
		}

		if taken == false {
			posibilitiesAfter = append(posibilitiesAfter, y)
		}
	}

	for y := 0; y < len(posibilitiesBefore); y++ {
		for x := 0; x < len(posibilitiesAfter); x++ {
			if posibilitiesAfter[x] == posibilitiesBefore[y] {
				posibilities = append(posibilities, posibilitiesAfter[x])
				continue
			}
		}
	}

	return posibilities
}

func checkSquare(i int, j int, matrix [][]int, posibilitiesBefore []int) []int {
	alreadyTaken := []int{}
	posibilitiesAfter := []int{}
	posibilities := []int{}

	squareI := i / 3
	i = i % 3
	squareJ := j / 3
	j = j % 3

	for x := (squareI * 3); x < (squareI * 3 + 3) ; x++ {
		for y := (squareJ * 3); y < (squareJ * 3 + 3) ; y++ {
			if matrix[x][y] != -1 {
				alreadyTaken = append(alreadyTaken, matrix[x][y])
			}
		}
	}

	for z := 1; z <= matrixSize; z++ {
		taken := false
		for x := 0; x < len(alreadyTaken); x++ {
			if alreadyTaken[x] == z {
				taken = true
				continue
			}
		}

		if taken == false {
			posibilitiesAfter = append(posibilitiesAfter, z)
		}
	}

	for y := 0; y < len(posibilitiesBefore); y++ {
		for x := 0; x < len(posibilitiesAfter); x++ {
			if posibilitiesAfter[x] == posibilitiesBefore[y] {
				posibilities = append(posibilities, posibilitiesAfter[x])
				continue
			}
		}
	}

	return posibilities
}

func parseInts(s []string) ([][]int, error) {
	var (
		fields = make([][]string, matrixSize)
		ints   = make([][]int, matrixSize)
		err    error
	)

	for x := range ints {
		ints[x] = make([]int, matrixSize)
	}

	for y := range fields {
		fields[y] = make([]string, matrixSize)
	}

	// fields = strings.Fields(s[0])
	for x := 0; x < matrixSize; x++ {
		fields[x] = strings.Fields(s[x])
	}

	for j := 0; j < matrixSize; j++ {
		for i, f := range fields[j] {
			if f == "*" {
				ints[j][i] = -1
				continue
			}

			ints[j][i], err = strconv.Atoi(f)

			if err != nil {
				return nil, err
			}
		}
	}

	return ints, nil
}

func readStdinMatrix() []string {

	line := make([]string, matrixSize)
	for i := 0; i < matrixSize; i++ {
		fmt.Print("Enter Numbers for line", i, ":")

		// Read line until enter
		r := bufio.NewReader(os.Stdin)
		redLine, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line[i] = redLine
	}

	return line
}

func initializeSudokuMatrix() {

	matrix := make([][]int, matrixSize)
	for x := range matrix {
		matrix[x] = make([]int, matrixSize)
	}

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			matrix[i][j] = -1
		}
	}

	sudoku = matrix
}

func readFile() []string {
	file, err := os.Open("sudoku.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	line := make([]string, matrixSize)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line[i] = scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}
