package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
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
	copyToClipboard := flag.Bool("c", false, "Copy output to clipboard")
	fileList := flag.String("f", "", "Comma-separated list of files or directories to process")
	outputFileP := flag.String("o", "", "Output file list")
	flag.Parse()
	outputFile := *outputFileP
	if *fileList == "" {
		fmt.Println("Please provide a list of files or directories using the --f flag.")
		os.Exit(1)
	} else if outputFile == "" && !*copyToClipboard {
		fmt.Println("Please provide an output file via --o flag if clipboard is disabled")
		os.Exit(1)
	} else if fileExists(outputFile) && *copyToClipboard {
		fmt.Printf("Output file '%s' already exists or contains text.\n", outputFile)
		os.Exit(1)
	} else if outputFile != "" && *copyToClipboard {
		fmt.Println("Output file and clipboard=true passed, cannot write to file when copying to clipboard")
		os.Exit(1)
	}

	var output strings.Builder
	for _, path := range strings.Split(*fileList, ",") {
		if _, err := os.Stat(strings.TrimSpace(path)); os.IsNotExist(err) {
			fmt.Printf("File or directory '%s' does not exist.\n", path)
			os.Exit(1)
		}
		err := processPath(strings.TrimSpace(path), &output)
		if err != nil {
			fmt.Printf("Error processing path %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	if *copyToClipboard {
		err := clipboard.WriteAll(output.String())
		if err != nil {
			fmt.Println("Error copying output to clipboard:", err)
			os.Exit(1)
		}
		fmt.Println("Output copied to clipboard.")
	} else {
		err := os.WriteFile(outputFile, []byte(output.String()), 0644)
		if err != nil {
			fmt.Printf("Error writing to %s: %v\n", outputFile, err)
			os.Exit(1)
		}

		fmt.Printf("Files processed successfully. Output saved to %s.\n", outputFile)
	}
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
