package main

import (
	"fmt"
)

type boardGame struct {
	Size  int
	Board []int
}

func (b boardGame) hasCollision() bool {
	// fmt.Println("Testando ->", b)
	n := b.Size
	for i := 0; i < n; i++ { //0 a 6
		for j := i + 1; j < n; j++ { //i a 7
			if b.Board[i*2] == b.Board[j*2] { //line
				return true
			}
			if b.Board[i*2+1] == b.Board[j*2+1] { //column
				return true
			}
			if b.Board[i*2]+b.Board[i*2+1] == b.Board[j*2]+b.Board[j*2+1] { //diagonal 1 x1+y1 == x2+y2
				return true
			}
			if b.Board[i*2]-b.Board[i*2+1] == b.Board[j*2]-b.Board[j*2+1] { //diagonal 2 x1-y1 == x2-y2
				return true
			}
		}
	}
	return false
}

func (b boardGame) next() bool {
	b.Board[b.Size*2-1]++
	for i := b.Size*2 - 1; i > 0; i-- {
		if b.Board[i] == b.Size {
			b.Board[i] = 0
			b.Board[i-1]++
		}
	}
	if b.Board[0] == b.Size {
		return false
	}
	return true
}

func findResult(s int) {
	board := boardGame{Size: s, Board: make([]int, s*2)}
	for board.next() {
		if !board.hasCollision() {
			fmt.Println("Solution found")
			fmt.Println(board)
			return
		}
	}
	fmt.Println("404 Solution not found")
}

func createBoard(begin, size int) boardGame {
	s := []int{begin}
	for i := 0; i < size*2-1; i++ {
		s = append(s, 0)
	}
	return boardGame{Size: size, Board: s}
}

func routineParallel(begin, end int, size int, ch chan int) {
	board := createBoard(begin, size)
	fmt.Println(begin, " | ", end)
	qtd := 0
	for board.next() && board.Board[0] != end {
		if !board.hasCollision() {
			//fmt.Println("Solution found:", board)
			qtd += 1
			// ch <- qtd //first result
			// break     //first result
		}
	}
	ch <- qtd //all results
}

func findResultParallel(boardSize, qtdGoRoutines int) {
	ch := make(chan int, 1)
	solutions := 0
	linesPerGoRoutine := boardSize / qtdGoRoutines
	i := 0
	for i = 0; i < qtdGoRoutines-1; i++ {
		go routineParallel(linesPerGoRoutine*i, linesPerGoRoutine*(i+1), boardSize, ch)
	}
	go routineParallel(linesPerGoRoutine*i, boardSize, boardSize, ch)
	for i := 0; i < qtdGoRoutines; i++ { //all results
		solutions += <-ch
	}
	// solutions += <-ch //first result
	if solutions > 0 {
		fmt.Println(solutions, " solutions found")
	} else {
		fmt.Println("404 Solution not found")
	}
}

func main() {
	boardSize := 6
	qtdGoRoutines := 6
	findResultParallel(boardSize, qtdGoRoutines)
}

//boardSize = 6
//qtdGoRoutines
//	1-> 74.257 seconds
//	2-> 49.165 seconds
//	3-> 41.946 seconds
//	6-> 27.173 seconds
