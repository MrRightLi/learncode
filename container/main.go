package container

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const fileName = "project/demo/a.txt"
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	if contents, err := ioutil.ReadFile(fileName); err == nil {
		fmt.Printf("%s\n", contents)
	} else {
		fmt.Println(err)
	}

	fmt.Println(
		//grade(0),
		grade(59),
		grade(70),
		grade(80),
		grade(90),
		grade(100),
		//grade(200),
	)

}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		return "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g
}

func consts() {

	const a, c int = 3, 4

	const (
		cpp = iota
		_
		python
		golang
		js
	)

	// b, kb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		gb
		tb
		pb
	)

	fmt.Println(cpp, js, python, golang)
	fmt.Println(b, kb, gb, tb, pb)
}
