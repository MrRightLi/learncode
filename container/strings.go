package container

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	fmt.Println(bytes)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	// 1. 使用range遍历pos,rune对
	// 2. 使用uft8.RuneCountInString 获得字符数量
	// 3. 使用len获得字节的长度
	// 4. 使用[]byte 获得字节
}