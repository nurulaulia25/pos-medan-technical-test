package main

import (
	"fmt"
	"sort"
	"strings"
)

func countLetters(s string) string {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	var letters []string
	for letter, count := range counts {
		letters = append(letters, fmt.Sprintf("%c=%d", letter, count))
	}
	sort.Strings(letters)
	return strings.Join(letters, ", ")
}

func groupAndSortLetters(strings []string) string {
	counts := make(map[rune]int)
	for _, s := range strings {
		for _, r := range s {
			counts[r]++
		}
	}
	letters := make([]rune, 0, len(counts))
	for letter := range counts {
		letters = append(letters, letter)
	}
	sort.Slice(letters, func(i, j int) bool {
		countI := counts[letters[i]]
		countJ := counts[letters[j]]
		if countI != countJ {
			return countI > countJ
		}
		return letters[i] < letters[j]
	})
	var result []rune
	for _, letter := range letters {
		result = append(result, letter)
	}
	return string(result)
}

func main() {
	fmt.Println("----------Output untul soal nomer 1----------")
	fmt.Println(countLetters("We Always Mekar"))
	fmt.Println(countLetters("coding is fun"))
	fmt.Println("----------Output untuk soal nomer 2----------")
	fmt.Println(groupAndSortLetters([]string{"Abc", "bCd"}))
	fmt.Println(groupAndSortLetters([]string{"Oke", "One"}))
	strings := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	fmt.Println(groupAndSortLetters(strings))
}
