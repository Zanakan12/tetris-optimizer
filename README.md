
# Tetromino Solver

## Overview

Tetromino Solver is a command-line tool designed to solve tetromino tiling puzzles on a defined grid size. The software utilizes an advanced backtracking algorithm to efficiently place tetromino pieces on the grid without overlaps and ensuring all pieces fit within the grid boundaries.

### Algorithm Description

The core algorithm behind Tetromino Solver uses a backtracking approach, enhanced with several optimizations to improve efficiency:
- **Bounding Box Check**: Ensures that a tetromino does not exceed the grid boundaries before attempting placement.
- **Early Exit**: Stops further checks as soon as an overlap is detected, enhancing performance by reducing unnecessary computations.
- **Heuristics**: Implements strategies like least-filling heuristic and most constraining piece placement to guide the search process, reducing the search space and focusing on more promising paths early in the computation.
```sh
+---------------------+
|  ParseTetrominos    |
+----------+----------+
           |
           |  [Entrée: lines[]]
           v
+----------+----------+
|  Initialisation de  |
|  tetrominos[]       |
+----------+----------+
           |
           |  [Boucle principale sur lines]
           v
+----------+----------+
|  Lecture des lignes |
|  pour chaque bloc   |
+----------+----------+
           |
           |  [Interprétation des caractères]
           v
+----------+----------+
|  Construction des   |
|  positions[]        |
+----------+----------+
           |
           |  [Validation]
           v
+-----------+-----------+
|  Ajout au tetrominos[]|
|  si valide            |
+-----------+-----------+
           |
           v
+----------+-----------+
|  Retour de tetrominos|
|  et gestion d'erreurs|
+----------------------+
```
## Installation

### Prerequisites

- Go (Golang) environment set up. The software is developed in Go and requires Go version 1.16 or higher.

### Setup

1. Clone the repository:
   ```bash
   git clone [URL-to-your-repository]
   cd tetris-optimizer
   ```

2. Build the project (optional):
   ```bash
   go mod init tetris-optimizer
   go build .
   ```

## Usage

To run Tetromino Solver, you need to provide an input file containing the grid size and the tetromino shapes. The input file format is expected to be specified as follows:

- First line: Grid size (N), where the grid will be N x N.
- Subsequent lines: Definitions of tetrominos, one tetromino per four lines, followed by an empty line.

### Example Input File

```
4
####
   #
   
####

```

**Note**: Each tetromino block is represented as a 4x4 matrix where `#` marks the blocks of the tetromino.

### Running the Solver

After setting up the input file, run the solver using:

```bash
./tetris-optimizer your_input_file.txt
```

or if you haven't built the project:

```bash
go run . your_input_file.txt
```

### Output

The program outputs the solved grid to the console with each tetromino placed on the grid. Each tetromino type is represented by a unique alphabet character starting from 'A'.

## Contributing

Contributions to Tetromino Solver are welcome. Please ensure to update tests as appropriate and follow the coding conventions and format outlined in existing files.

## License

Specify the license under which your project is made available.
