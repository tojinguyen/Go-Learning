package main

import "fmt"

func main() {
	s := "Xin chào"
	fmt.Println(s[6])
	fmt.Println(string(s[6]))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // có thể bị lỗi với Unicode
	}

	fmt.Println()

	for _, r := range s {
		fmt.Printf("%c ", r) // in đúng ký tự
	}

	fmt.Println()

	s1 := "Tiếng Việt"
	fmt.Println([]byte(s1)) // mảng byte
	fmt.Println([]rune(s1)) // mảng rune (ký tự Unicode)
}
