package main

import "fmt"

type Config struct {
}

func buildInput(inputSet string) *Config {
	lines := splitNewLines(inputSet)
	configLine := splitSpaces(lines[0])
	fmt.Printf("Config line: %v\n", configLine)

	config := &Config{}

	return config
}

func buildOutput(result int) string {
	return "42"
}
