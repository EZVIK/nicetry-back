package main

import "fmt"

/**

给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。


*/

func exist(board [][]byte, word string) bool {
	height, length := len(board), len(board[0])

	i, j := -1, -1

	for x, y := 0, 0; x != height && y != length; y++ {
		fmt.Println(x, y)
		if board[x][y] == word[0] {
			i, j = x, y
			break
		}

		if y == length-1 {
			x += 1
			y = -1
		}
	}

	if i == -1 || j == -1 {
		return false
	}

	dp := make([][]int, len(word))

	dp[0] = []int{i, j}

	//flag := true

	for i := 1; i <= len(word); i++ {

	}

	return true
}

func moveXY(board [][]byte, x, y int) int32 {
	return int32(board[x][y])
}

func main() {

	//board := make([][]byte, 3)
	//
	//board[0] = []byte{'A','B','C','E'}
	//board[1] = []byte{'S','F','C','S'}
	//board[2] = []byte{'A','D','E','E'}
	//
	//fmt.Println(exist(board, "ABCCED"))
	aa := "ABCCED"
	fmt.Println(len(aa))
}

//[
//[1,  2, 3, 4, 5]
//[6,  7, 8, 9,10]
//[11,12,13,14,15]
//
//
//

// 1. find the header
// 2. compare other three direction
// 3. mark the right index
