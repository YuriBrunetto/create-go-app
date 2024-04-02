package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: create-go-app <app-name>\n")
	}

	appName := os.Args[1]

	if err := os.Mkdir(appName, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
	}

	makefileContents := []byte(fmt.Sprintf("build:\n\t@go build -o bin/%s\n\nrun: build\n\t@./bin/%s", appName, appName))
	if err := os.WriteFile(fmt.Sprintf("%s/Makefile", appName), makefileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}

	mainFileContents := []byte(fmt.Sprintf("package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello from %s!\")\n}\n", appName))
	if err := os.WriteFile(fmt.Sprintf("%s/main.go", appName), mainFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}

	fmt.Printf("Directory `%s` created!", appName)
}
