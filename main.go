package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var fileTypeMap = map[string]string{
	".go":    "go",
	".rs":    "rust",
	".py":    "python",
	".js":    "javascript",
	".java":  "java",
	".c":     "c",
	".cpp":   "cpp",
	".cs":    "csharp",
	".rb":    "ruby",
	".php":   "php",
	".swift": "swift",
	".kt":    "kotlin",
	".scala": "scala",
	".hs":    "haskell",
	".ml":    "ocaml",
	".fs":    "fsharp",
	".clj":   "clojure",
	".erl":   "erlang",
	".ex":    "elixir",
	".dart":  "dart",
	".ts":    "typescript",
	".sh":    "bash",
	".html":  "html",
	".css":   "css",
	".sql":   "sql",
	".md":    "markdown",
	".txt":   "text",
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: ./%s <output_file> <file_or_directory1> [<file_or_directory2> ...]", os.Args[0])
		os.Exit(1)
	}

	outputFile := os.Args[1]
	if outputFile == "" {
		fmt.Println("Output file name cannot be empty.")
		os.Exit(1)
	}

	if fileExists(outputFile) {
		fmt.Printf("Output file '%s' already exists or contains text.\n", outputFile)
		os.Exit(1)
	}

	var output strings.Builder

	for _, path := range os.Args[2:] {
		err := processPath(path, &output)
		if err != nil {
			fmt.Printf("Error processing path %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	err := os.WriteFile(outputFile, []byte(output.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing to %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("Files processed successfully. Output saved to %s.\n", outputFile)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir() && info.Size() > 0
}

func processPath(path string, output *strings.Builder) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return processDirectory(path, output)
	} else {
		return processFile(path, output)
	}
}

func processDirectory(dir string, output *strings.Builder) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return processFile(path, output)
		}

		return nil
	})
}

func processFile(file string, output *strings.Builder) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	ext := strings.ToLower(filepath.Ext(file))
	languageType, ok := fileTypeMap[ext]
	if !ok {
		languageType = "text"
	}

	output.WriteString(file + "\n\n")
	output.WriteString("```" + languageType + "\n")
	output.WriteString(string(content) + "\n")
	output.WriteString("```\n\n")

	return nil
}

