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
