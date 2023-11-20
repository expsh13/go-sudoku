package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Boardは数独のboard
// 0:未入力
// 1-9:入力

type Board [9][9]int

func pretty(b Board) string{
	var buf bytes.Buffer

	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

// 重複チェック
func duplicated(c [10]int) bool {
	// k:index,v:value
	for k,v := range c {
		if k == 0 {
			continue
		}
		// 出現回数二回以上
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(b Board) bool {
	// 行チェック
	for i := 0; i < 9; i++ {
		// 出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			return false
		}
	}
	// 列チェック
	for i := 0; i < 9; i++ {
		// 出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
		}
		if duplicated(c) {
			return false
		}
	}
	// 3*3チェック
	for i := 0; i < 9; i+=3 {
		for j := 0; j < 9; j+=3 {
			// 出現回数
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}
	return true
}

// 終了条件
func solved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func backTrack(b *Board) bool {
	// time.Sleep(time.Second * 1)
	// fmt.Printf("%+v\n", pretty(*b))

	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 探索する
			if b[i][j] == 0 {
				for c := 9; c >=1; c-- {
					b[i][j] = c
					// ルールに一致するかチェック
					if verify(*b) {
						if backTrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

// .5..83.17...1..4..3.4..56.8....3...9.9.8245....6....7...9....5...729..861.36.72.4
func short(input string) (*Board, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanRunes)
	var b Board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !s.Scan() {
				break
			}
			token := s.Text()
			if token == "." {
				b[i][j] = 0
				continue
			}
			n, err := strconv.Atoi(token)
			if err != nil {
				return nil ,err
			}
			b[i][j] = n
		}
	}
	return &b, nil
}

func main() {
	b := Board{
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
	}
	fmt.Printf("b = %+v\n", pretty(b))
}