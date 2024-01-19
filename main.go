package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Table struct {
	rows    int
	columns int
	data    []string
	// cells  []Cell
}

type Position struct {
	x int
	y int
}

type CellType string

type Cell[T any] struct {
	value    T
	cellType CellType
	position Position
}

const (
	CELL_KIND_TEXT       CellType = "CELL_KIND_TEXT"       // if no number text
	CELL_KIND_NUMBER     CellType = "CELL_KIND_NUMBER"     //if number number
	CELL_KIND_EXPRESSION CellType = "CELL_KIND_EXPRESSION" //if starts wit = expression
)

func generateTable(table Table, s []string) {
	fmt.Print("Printing table\n")
	fmt.Printf("%#v\n", table)
	fmt.Printf("%#v\n", s)

	parsed_string_array := [][]string{}

	for i := range s {
		parsed_string_array = append(parsed_string_array, strings.Split(s[i], "|"))
	}

	for i := range parsed_string_array {
		for j := range parsed_string_array[i] {
			if parsed_string_array[i][j] == "" {
				parsed_string_array[i][j] = " "
			}
		}
	}

	for i := range parsed_string_array {
		for len(parsed_string_array[i]) < table.columns {
			parsed_string_array[i] = append(parsed_string_array[i], " ")
		}
	}

	for len(parsed_string_array) < table.rows {
		parsed_string_array = append(parsed_string_array, []string{})
		for i := 0; i < table.columns ; i++ {
			parsed_string_array[len(parsed_string_array)-1] = append(parsed_string_array[len(parsed_string_array)-1], " ")
		}
	}



	printCells(parsed_string_array, table.rows, table.columns)

}

func printCells(cells [][]string, rows int, columns int) {
	fmt.Println("Printing cells")
    fmt.Printf("cells: %v\n", cells)

    fmt.Printf("rows: %d\n", rows)
    fmt.Printf("columns: %d\n", columns)

    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            fmt.Printf("|%s|", cells[i][j] )
        }
        fmt.Printf("\n")
    }
}

func importFile() {

	file, err := os.Open("./input.csv")
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	line_number := 0
	max_col_length := 0
	col_length := 0

	var s []string
	// var cells[] int

	for scanner.Scan() {
		line_number += 1
		s = append(s, strings.ReplaceAll(scanner.Text(), " ", ""))
	}

	for i := range s {

		for _, char := range s[i] {

			if string(char) == "|" {

				col_length = col_length + 1
			}

			if col_length > max_col_length {
				max_col_length = col_length
			}
		}
		col_length = 0

	}

	rows := line_number
	columns := max_col_length + 1

	generateTable(Table{
		rows, columns, s,
	}, s)

}

func main() {
	importFile()
}
