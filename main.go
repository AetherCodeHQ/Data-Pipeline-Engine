package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: data-pipeline <input.csv>")
		fmt.Println("Pipeline: Extract -> Transform -> Validate -> Load")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var rows [][]string
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		rows = append(rows, parts)
	}
	fmt.Printf("Pipeline Report\n")
	fmt.Printf("==============\n")
	fmt.Printf("Extract:  %d rows loaded\n", len(rows))
	if len(rows) > 0 {
		fmt.Printf("Headers:  %v\n", rows[0])
	}
	transformed := 0
	validated := 0
	for i, row := range rows[1:] {
		_ = i
		for j, cell := range row {
			row[j] = strings.ToUpper(cell)
			transformed++
		}
		if len(row) == len(rows[0]) {
			validated++
		}
	}
	fmt.Printf("Transform: %d cells uppercased\n", transformed)
	fmt.Printf("Validate:  %d/%d rows pass schema check\n", validated, len(rows)-1)
	fmt.Printf("Load:      Ready to write %d valid rows\n", validated)
}