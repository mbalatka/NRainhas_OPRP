package main

import (
	"fmt"
)

type boardGame struct {
	Size  int
	Board []int
}

func (b boardGame) hasCollision() bool {
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

func findResultParallel(s int, qtdRoutines int) {
	r1 := routineParallel([]int{0, 0, 0, 0, 0, 0, 0, 0}, []int{1, 1, 1, 1, 1, 1, 1, 1}, 4)
	r2 := routineParallel([]int{2, 0, 0, 0, 0, 0, 0, 0}, []int{3, 3, 3, 3, 3, 3, 3, 3}, 4)
	if r1 == false && r2 == false {
		fmt.Println("404 Solution not found")
	}
}

func routineParallel(begin []int, end []int, s int) bool {
	board := boardGame{Size: s, Board: begin}
	for board.next() { // || board.Board == end
		if !board.hasCollision() {
			fmt.Println("Solution found")
			fmt.Println(board)
			return true
		}
	}
	return false
}

func main() {
	//bCollision := boardGame{Size:4, Board:[]int{0,0,0,1,0,2,0,3}}
	//bNoCollision := boardGame{Size:4, Board:[]int{0,2,1,0,2,3,3,1}}
	//fmt.Println("1 -> ", bCollision.hasCollision())
	//fmt.Println("2 -> ", bNoCollision.hasCollision())
	//findResult(8)
	findResultParallel(4, 2)
}
