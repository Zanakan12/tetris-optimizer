# Tetromino Placement Program

This README outlines the process flow of the Tetromino placement program. It details each step from the beginning of the program to the final output.
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

## Flowchart Description

Below is a flowchart that visually represents the steps involved in the Tetromino placement process:

```plaintext
+-------------------+     +------------------------+     +------------------------+
| Start Program     |     | Read Input File        |     | Validate Tetrominos    |
| (Check arguments) | --> | (Load Tetrominos)      | --> | (Check shapes and size)|
+-------------------+     +------------------------+     +------------------------+
                                      |
                                      V
                            +-------------------------+
                            | Compress Tetrominos     |
                            | (Remove empty rows and  |
                            |  columns)               |
                            +-------------------------+
                                      |
                                      V
                            +-------------------------+
                            | Calculate Initial Board |
                            | Size                    |
                            +-------------------------+
                                      |
                                      V
                            +-------------------------+
                            | Create Board            |
                            | (Fill with '.')         |
                            +-------------------------+
                                      |
                                      V
                            +-------------------------+
                            | Attempt to Place        |
                            | Tetrominos              |
                            | (Resolve function)      |
                            +-------------------------+
                                      |
                                      +----------------+
                                      |                |
                            Yes<----+ Can Place?      +---->No
                                      |                |
                                      +----------------+
                                      |                |
                                      V                |
                            +-----------------+  +------------------+
                            | Place Tetromino |  | Backtrack to     |
                            | on Board        |  | previous state   |
                            +-----------------+  | or Resize Board  |
                                                 +------------------+
                                      |                |
                                      +----------------+
                                      |
                                      V
                            +-------------------------+
                            | All Tetrominos Placed?  |
                            +-------------------------+
                                      |
                                      +----------------+
                                      |                |
                             Yes<-----+                +---->No
                                      |                |
                                      +----------------+
                                      |
                                      V
                            +-------------------------+
                            | Output Final Board      |
                            | Configuration           |
                            +-------------------------+
                                      |
                                      V
                            +-------------------------+
                            | End Program             |
                            +-------------------------+
```

## Explanation

The flowchart illustrates the logical steps involved in the Tetromino placement program. Here is what each step entails:

1. **Start**: The program begins by checking for input arguments to ensure a file path is provided.
2. **Reading and Validation**: Reads Tetrominos from the file and validates their structure.
3. **Compression**: Compresses Tetrominos by removing unnecessary rows and columns.
4. **Board Calculation and Creation**: Calculates the necessary board size and creates the board.
5. **Placement Attempts**: Tries to place each Tetromino on the board using the `Resolve` function.
6. **Backtracking and Resizing**: If placement fails, the program backtracks or enlarges the board, then retries.
7. **Output**: Once all Tetrominos are placed, the final board configuration is displayed.
8. **End**: The program concludes after the final board configuration is shown.


## Usage
To run the program, use the following commands:

```bash
go mod tidy
go run main.go
```

Ensure that the input file (tetrominos.txt) is present in the root directory of the project.

## Authors
DJIHADI RAFTANDJANI
