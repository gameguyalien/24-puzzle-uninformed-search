# Project 24-puzzle-uninformed-search

A project To Solve versions of the [8 puzzle](https://sliding.toys/mystic-square/8-puzzle/) using uninformed search techniques such as depth-first search and breadth-first search.

## Supported Puzzle size
|Puzzle Size | BFS | DFS |
|---|---|---|
|3|x|x|
|4|x|x|
|5|x|x|

## System Requirements 
## Memory in Gigabytes
|Puzzle Size | BFS | DFS |
|---|---|---|
|3|8|8|
|4|32|?|
|5|300+|?|


# Usage
us-slover in your local path
./us-slover 
## Optional Flags
|Flag|Description|Vaild Options|Default|
|---|---|---|---|
|-s | Size of Board, 3 is the default 8 Puzzle, 4 15 Puzzle, and 5 24 Puzzle.| 3,4,5| 3|
|-m | Search Method | bfs,dfs|bfs|
|-d |BFS Only, sets a deepth limit to BFS to increase lucklyhood of giving up if impossible| any int | 30|
|-i | User json Inputed puzzle |  Relative or full file path| null|

## Json Input
When using Json input the -s flag is ignored and instead the tableSize json value is used instead.
0 Is Interpreted as the blank space.
Only one board can be processed at a time.

### 8 Puzzle Custom Board
``` json
{
    "tableSize": "3",
    "initialBoard": [
        [1, 2, 3],
        [6, 7, 8],
        [0, 9, 4],
    ]
}
```
### 15 Puzzle Custom Board
``` json
{
    "tableSize": "4",
    "initialBoard": [
        [1, 2, 3, 4],
        [6, 7, 8, 9],
        [11, 12, 13, 14],
        [5, 10, 15, 0],
    ]
}
```
### 24 Puzzle Custom Board
``` json
{
    "tableSize": "5",
    "initialBoard": [
        [1, 2, 3, 4, 5],
        [6, 7, 8, 9, 10],
        [11, 12, 13, 14, 15],
        [16, 17, 18, 0, 19],
        [21, 22, 23, 24, 20]
    ]
}
```
## Getting Started

To Compile this program you'll require [Golang](https://go.dev/doc/install) 1.25^ 
To use the prebuilt Binarys just select the build for your system architecture and fire away.

## MakeFile

### General running and compling instructions.
Run the application
```bash
make run
```

Build the application for the local system 
```bash
make build
```

### Compile for other platforms

Build for all systems
```
make build_all
```

Build the application for windows AMD64
```bash
make build_windows_AMD64
```

Build the application for windows ARM64
```bash
make build_windows_ARM64
```

Build the application for Mac OSX ARM64 (M series processors)
```bash
make build_mac_ARM64
```

Build the application for Mac OSX AMD64 (Intel processors)
```bash
make build_mac_AMD64
```

Build the application for linux AMD64 
```bash
make build_linux_AMD64
```

Build the application for linux ARM64 
```bash
make build_linux_ARM64
```

Clean up binary from the last builds:
```bash
make clean
```
