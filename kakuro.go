package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var field Field

const size int = 5

// Constraint of a cell with horizontal and vertical sum
type Constraint struct {
	X int `json:"x"`
	Y int `json:"y"`
	H int `json:"h"`
	V int `json:"v"`
}

// Dimensions of the field
type Dimensions struct {
	W int `json:"w"`
	H int `json:"h"`
}

// Constraints is an array of Constraint objects and the Dimensions of the field
type Constraints struct {
	Dimensions  Dimensions   `json:"dimensions"`
	Constraints []Constraint `json:"constraints"`
}

var c Constraints

// Field is the w*h field
type Field struct {
	HSum       [size][size]int // -1 = no sum
	VSum       [size][size]int // -1 = no sum
	Max        [size][size]int // -1 = no number field
	Min        [size][size]int // -1 = no number field
	Candidates [size][size][]Candidate
	Solution   [size][size]int // -1 = no number field
}

// Candidate is a array of ints
type Candidate struct {
	values []int
}

func initField() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			field.Min[i][j] = 1
			field.Max[i][j] = 9
			var cand Candidate

			field.Candidates[i][j] = append(field.Candidates[i][j], cand)
			field.Solution[i][j] = 0
			field.HSum[i][j] = -1
			field.VSum[i][j] = -1

		}
	}
}
func applyConstraints() {
	for i := 0; i < len(c.Constraints); i++ {
		x := c.Constraints[i].X
		y := c.Constraints[i].Y
		field.HSum[x][y] = c.Constraints[i].H
		field.VSum[x][y] = c.Constraints[i].V
		field.Candidates[x][y] = nil
		field.Max[x][y] = -1
		field.Min[x][y] = -1
		field.Solution[x][y] = -1
	}
}

func printSolution() {
	fmt.Println("Current field state")
	for j := 0; j < size; j++ {
		var line = field.Solution[j]
		fmt.Println(line)
	}
}

func printCandidates() {
	fmt.Println("Current candidate")
	for j := 0; j < size; j++ {
		var line = field.Candidates[j]
		fmt.Println(line)
	}
}
func printConstraints() {
	fmt.Println("Constraints loaded")
	for j := 0; j < len(c.Constraints); j++ {
		var line = c.Constraints[j]
		fmt.Println(line)
	}
}

func readConstraintsFile() {
	jsonFile, err := os.Open("field.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened field.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	//	var users Con

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &c)

}

func main() {
	fmt.Printf("Starting Kakuro Solver\n")
	initField()
	readConstraintsFile()
	applyConstraints()

	printConstraints()
	printCandidates()
	printSolution()
}
