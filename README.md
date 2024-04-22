## AI-Walk
The File Processor is a command-line tool that allows you to process multiple files and directories, extracting their contents and formatting them for use with AI systems. It supports a wide range of programming languages and file types, making it easy to prepare your codebase for AI analysis and processing.
Features

CLI Tool allowing you to format multiple files for AI's 

### Output Example 
```
ai-walk output.txt main.go src
```

```
main.go
\`\`\`go
// main.go contents 
\`\`\`
src/pkg/handlers.go
\`\`\`go
// src/pkg/handlers.go contents
\`\`\`
```

### Installation

Open a terminal or command prompt.
Run the following command to install the File Processor:
```
go install github.com/seal/ai-walk@latest
```
Wait for the installation to complete. The executable will be installed in your $GOPATH/bin directory.

### Usage


Open a terminal or command prompt.

```
ai-walk output.txt file1.go /path/to/directory
```

### Note:
If the output file already exists the program will not over-write it to avoid code los s


