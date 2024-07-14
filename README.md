
# Tetris Optimizer
## Description
The `tetris-optimizer` project is a tool developed in Go to validate and optimize Tetris game configurations.

## Project Architecture
The project is structured as follows:

```
tetris-optimizer/
│
├── main.go                   # Program entry point
├── go.mod                    # Go module configuration file
├── README.md                 # Project description and instructions
├── pkg/
│   ├── cutline.go            # Management of cut lines
│   ├── Verify.go             # Validation of tetrominos
│   ├── Reader_file.go        # Reading input files
│   └── tetromino_resolver.go # Resolution of tetromino configurations
└── veryhard.txt              # Example input file
```

## General Operation
- **File Reading**: The program reads the specified input file (`tetrominos.txt`).
- **Validation**: The tetrominos in the file are validated to ensure they are correct.
- **Resolution**: The program attempts to solve the configuration using optimization algorithms.
- **Displaying Results**: Results are displayed to the user in the form of color blocks.

## Workflow Diagram
(Describe the workflow diagram here if necessary)

## Usage
To run the program, use the following commands:

```bash
go mod tidy
go run main.go
```

Ensure that the input file (`tetrominos.txt`) is present in the root directory of the project.

## Authors
DJIHADI RAFTANDJANI
