package main

import (
	"fmt"
	"os"
	"strconv"
)

func resolvSudoku(sudoku *[9][9]rune) bool {
	i, j := findEmptyCel(sudoku)
	if i == -1 {
		return true
	}
	for nb := 1; nb <= 9; nb++ {
		if isValid(sudoku, i, j, nb) {
			r := '0' + rune(nb%10)
			sudoku[i][j] = r
			if resolvSudoku(sudoku) {
				return true
			}
			sudoku[i][j] = '.'
		}
	}
	return false
}
func isValid(tab *[9][9]rune, i, j, nb int) bool {
	a := strconv.Itoa(nb)
	// balayage ligne
	for i := range tab {
		if string(tab[i][j]) == a {
			return false
		}
	}
	//balayage colone
	for j := range tab {
		if string(tab[i][j]) == a {
			return false
		}
	}
	//balayage grid 3x3
	sectX := 3 * (i / 3)
	sectY := 3 * (j / 3)
	for l := sectX; l < sectX+3; l++ {
		for k := sectY; k < sectY+3; k++ {
			if l != i && k != j {
				if string(tab[l][k]) == a {
					return false
				}
			}
		}
	}
	return true
}
func findEmptyCel(tab *[9][9]rune) (x, y int) {
	for i, v1 := range tab {
		for j := range v1 {
			if string(tab[i][j]) == "." {
				return i, j
			}
		}
	}
	return -1, -1
}
func printSudoku(tab *[9][9]rune) {
	for i, v := range tab {
		for j := range v {
			fmt.Print(string(tab[i][j]), " ")
		}
		fmt.Print("\n")
	}
}
func testParam(s []string) bool {
	for i := range s {
		R := []rune(s[i])
		if len(R) != 9 {
			return false
		}
	}
	return true
}
func IsAlpha(s []string) bool {
	for _, v := range s {
		r := []rune(v)
		for i := range r {
			if r[i] >= 65 && r[i] <= 90 || r[i] >= 97 && r[i] <= 122 {
				return true
			}
		}
	}
	return false
}

func IsDuplicate(s []string) bool {
	for _, v := range s {
		//fmt.Println(v)
		r := []rune(v)
		for i := '1'; i <= '9'; i++ {
			cpt := 0
			for j := 0; j < len(r); j++ {
				if r[j] == i {
					cpt++
				}
			}
			if cpt >= 2 {
				return true
			}
		}

	}
	return false
}

func main() {
	l := os.Args[1:]
	if len(l) != 9 {
		fmt.Println("Error")
	} else if testParam(l) && !IsAlpha(l) && !IsDuplicate(l) {
		var tab [9][9]rune
		for i, v1 := range l {
			for j, v2 := range v1 {
				tab[i][j] = v2
			}
		}
		if resolvSudoku(&tab) {
			printSudoku(&tab)
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}
