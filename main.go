package main

import (
	"fmt"
	"strings"
)

func getInitals(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitals("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitals("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitals("barret")
	fmt.Println(fn3, sn3)
}
