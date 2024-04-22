## AI-Walk
CLI Tool allowing you to format multiple files for AI's 

### Output Example 
```
ai-walk -f=main.go,src -o=text.txt
```
#### Output:
~~~
main.go
```go
// main.go contents 
```
src/pkg/handlers.go
```go
// src/pkg/handlers.go contents
```
~~~
### Installation

Open a terminal or command prompt.
Run the following command to install the File Processor:
```
go install github.com/seal/ai-walk@latest
```
Wait for the installation to complete. The executable will be installed in your $GOPATH/bin directory.

### Usage


#### Output to file 
```
ai-walk -o=text.txt -f=main.go,directoryName
```

#### Output to clipboard
```
ai-walk -c=true -f=main.go,directoryName
```


### Note:
If the output file already exists the program will not over-write it to avoid code loss


