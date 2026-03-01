# hands

Minimal but robust toolkit for your agents: Read, Write, Edit, Glob, Grep, Bash.

## Install

```sh
go install github.com/WinPooh32/hands
```

## Debug

You can use [f/mcptools](https://github.com/f/mcptools) for testing tools.

Example of the Glob tool execution from a terminal:

```sh
$ mcptools call Glob --params '{"pattern":"*.go"}' hands
Found 15 files:
- /home/winpooh/workspace/GO/hands/main.go
...
```

List of available tools:

```sh
$ mcptools tools hands
Bash(command:str, [workingDir:str])
     Execute a bash command

Edit(path:str, replace:str, search:str)
     Edit a file using search and replace

Glob(pattern:str, [dir:str])
     Find files by pattern

Grep(path:str, pattern:str, [ignoreCase:bool])
     Search for patterns in files

Read(path:str)
     Read file contents

Write(content:str, path:str)
     Write content to a file
```
