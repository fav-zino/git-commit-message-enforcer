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
	re := regexp.MustCompile(`^(?i)(fix|feat|breaking\schange|docs|test|perf|chore|revert|hotfix|refactor|build|style|release):`)

	if !re.MatchString(string(data)) {
		fmt.Println("\033[31m[MESSAGE FORMAT ERROR] Your commit message is not formatted correctly.\033[0m")
		fmt.Println("\033[33mPlease start your commit message with one of the following tags:\033[0m")
		fmt.Println("\033[32m- fix:      A bug fix")
		fmt.Println("- feat:     A new feature")
		fmt.Println("- BREAKING CHANGE: Introduces breaking API changes")
		fmt.Println("- docs:     Documentation changes")
		fmt.Println("- test:     Adding or modifying tests")
		fmt.Println("- perf:     Improving performance")
		fmt.Println("- chore:    Routine tasks, build process, or auxiliary tools")
		fmt.Println("- revert:    Revert a previous commit")
		fmt.Println("- hotfix:    An urgent fix")
		fmt.Println("- refactor: Code refactoring without changing functionality")
		fmt.Println("- build:    Changes that affect the build system or external dependencies")
		fmt.Println("- style:    Code style updates (formatting, etc.)")
		fmt.Println("- release:    For a release version\033[0m")
		os.Exit(1)
	}
}
