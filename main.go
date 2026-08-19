package main

import (
	"fmt"
	"os"
)

// data_pipeline_engine - Build data pipelines
func data_pipeline_engine(path string) {
	fmt.Println("========================================")
	fmt.Println("  Data-Pipeline-Engine")
	fmt.Println("  Build data pipelines")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	data_pipeline_engine(path)
}
