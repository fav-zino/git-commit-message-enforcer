package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	var messageFile string = os.Args[1]
	data, err := os.ReadFile(messageFile)
	if err != nil {
		log.Fatalf("\033[31mFailed to read file: %v\033[0m", err)
	}
	re := regexp.MustCompile(`^(?i)(fix|feat|breaking\schange|docs|test|chore|refactor|build|style):`)

	if !re.MatchString(string(data)) {
		
		fmt.Println("\033[31m[POLICY] Your message is not formatted correctly\033[0m")
		os.Exit(1)
	}
}
