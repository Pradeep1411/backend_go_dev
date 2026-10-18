package main

import "fmt"

func main() {
	fmt.Println("Jai Ma Saraswati, Jai BajarangBali, Jai Hedakhan Baba, Jai Guru GorkahNath !")
	s := "apradeep"
	fmt.Println(s[0])
	fmt.Println(string(s[0]))

	occ := make(map[rune]int)

	for i, v := range s {
		fmt.Printf("i : %d, v : %s \n", i, string(v))
		occ[v]++
	}
	fmt.Println(occ)

	for ch, count := range occ {
		if ch == 'p' {
			fmt.Println(string(ch), count)
		}
	}
}
