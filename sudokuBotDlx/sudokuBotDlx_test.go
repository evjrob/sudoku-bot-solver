package sudokuBotDlx

import (
  "bufio"
  "fmt"
  "os"
  "testing"
)

func solvePuzzles(puzzleFileName, solutionFileName string, blockXDim, blockYDim int, delimiter string) bool {

  passed := true

  puzzleFile, _ := os.Open(puzzleFileName)
  solutionFile, _ :=os.Open(solutionFileName)

  puzzleScanner := bufio.NewScanner(puzzleFile)
	puzzleScanner.Split(bufio.ScanLines)
  solutionScanner := bufio.NewScanner(solutionFile)
  solutionScanner.Split(bufio.ScanLines)

	// For each line
	for puzzleScanner.Scan() {

    solutionScanner.Scan()

		// Read the puzzle text in and split it into it's components
		providedPuzzle := puzzleScanner.Text()
    providedSolution := solutionScanner.Text()

    solver := NewSolver(providedPuzzle, blockXDim, blockYDim, delimiter)

    solution, err := solver.Solve()
    if err != nil {
      passed = false
    }

    success := solution == providedSolution

    if !(success) {
      passed = false
    }
	}

  return passed
}

func TestSolveNormalSudoku(t *testing.T) {

  passed := solvePuzzles("puzzles.txt", "solutions.txt", 3, 3, "")

  if passed {
    fmt.Println("Passed on all the provided 3x3 puzzles.")
  } else {
    t.Errorf("Failed on at least one of the 3x3 puzzles.")
  }
}

func TestSolve4x4Sudoku(t *testing.T) {

  passed := solvePuzzles("4x4-puzzles.txt", "4x4-solutions.txt", 4, 4, ",")

  if passed {
    fmt.Println("Passed on all the provided 4x4 puzzles.")
  } else {
    t.Errorf("Failed on at least one of the 4x4 puzzles.")
  }
}

func TestSolve7x2Sudoku(t *testing.T) {

  passed := solvePuzzles("7x2-puzzles.txt", "7x2-solutions.txt", 7, 2, ",")

  if passed {
    fmt.Println("Passed on all the provided 7x2 puzzles.")
  } else {
    t.Errorf("Failed on at least one of the 7x2 puzzles.")
  }
}

func TestSolveImpossibleSudoku(t *testing.T) {

  passed := solvePuzzles("impossible.txt", "impossible.txt", 3, 3, "")

  if !passed {
    fmt.Println("Did not falsely indicate a success for the impossible sudoku.")
  } else {
    t.Errorf("Falsely indicated a success for the impossible sudoku.")
  }
}

func TestInvalidPuzzleDimension(t *testing.T) {

  passed := solvePuzzles("invalid.txt", "invalid.txt", 3, 3, "")

  if !passed {
    fmt.Println("Did not falsely indicate a success for the sudoku of invaid dimensions.")
  } else {
    t.Errorf("Falsely indicated a success for a sudoku with invalid dimesnisons.")
  }
}
