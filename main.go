package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files := []string{
		// "a_an_example",
		"b_better_start_small",
		"c_collaboration",
		"d_dense_schedule",
		"e_exceptional_skills",
		"f_find_great_mentors",
	}

	for _, fileName := range files {
		fmt.Printf("Processing input: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.in.txt", fileName))

		maxDays, config, contributors, projects, rolesContributor := buildInput(inputSet)
		// fmt.Printf("Config %+v\n", config)
		// for _, contrib := range contributors {
		// 	fmt.Printf("Contributor %+v\n", contrib)
		// }
		// for _, project := range projects {
		// 	fmt.Printf("Project %+v\n", project)
		// 	for _, role := range project.rolesMap {
		// 		fmt.Printf(">> Role %+v\n", role)
		// 	}
		// }

		printInputMetrics(config)

		result := algorithm(maxDays, config, contributors, projects, rolesContributor)

		output := buildOutput(result)
		printResultMetrics(config)

		ioutil.WriteFile(fmt.Sprintf("./result/%s.out", fileName), []byte(output), 0644)
		fmt.Printf("Wrote output for: %s\n", fileName)
	}
}
