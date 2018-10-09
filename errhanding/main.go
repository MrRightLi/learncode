package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	writeFile("abc.txt")
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); !ok{
			panic(err)
		} else {
			fmt.Println(pathErr.Op,
				pathErr.Path,
				pathErr.Err,
				)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, i)
	}
}
