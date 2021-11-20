package main

import "fmt"

func main() {
	a := buildBoard([][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}})
	fmt.Println(exist(a, "ABFD"))
}

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	a := buildBool(len(board), len(board[0]))

	for i, v := range board {
		for j, vv := range v {
			if vv == word[0] {
				if r(i, j, 0, a, board, word) {
					return true
				}
			}
		}
	}
	return false
}

func r(i, j, k int, dir [][]bool, board [][]byte, word string) bool {
	if k == len(word) {
		return true
	}
	if i == len(board) || j == len(board[0]) || i < 0 || j < 0 || dir[i][j] == true {
		return false
	}
	if word[k] == board[i][j] {
		dir[i][j] = true
	} else {
		return false
	}
	if r(i, j+1, k+1, dir, board, word) {
		return true
	}

	if r(i+1, j, k+1, dir, board, word) {
		return true
	}

	if r(i, j-1, k+1, dir, board, word) {
		return true
	}

	if r(i-1, j, k+1, dir, board, word) {
		return true
	}

	dir[i][j] = false
	return false
}

func buildBool(a, b int) [][]bool {
	var dir = make([][]bool, a)
	for i := 0; i < a; i++ {
		dir[i] = make([]bool, b)
	}
	return dir
}

func buildBoard(b [][]string) [][]byte {
	var board [][]byte
	for _, v := range b {
		cb := make([]byte, 0)
		for i := 0; i < len(v); i++ {
			cb = append(cb, byte(v[i][0]))
		}
		board = append(board, cb)
	}
	return board
}
