package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files := []string{
		"a_an_example",
		// "b_better_start_small",
		// "c_collaboration",
		// "d_dense_schedule",
		// "e_exceptional_skills",
		// "f_find_great_mentors",
	}

	for _, fileName := range files {
		fmt.Printf("Processing input: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.in.txt", fileName))

		config := buildInput(inputSet)
		printInputMetrics(config)

		result := algorithm(config)

		output := buildOutput(result)
		printResultMetrics(config)

		ioutil.WriteFile(fmt.Sprintf("./result/%s.out", fileName), []byte(output), 0644)
		fmt.Printf("Wrote output for: %s\n", fileName)
	}
}
