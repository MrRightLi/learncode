package container

import "fmt"

func lengthOfNonReatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonReatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonReatingSubStr("bbbbbb"))
	fmt.Println(lengthOfNonReatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonReatingSubStr("q"))
	fmt.Println(lengthOfNonReatingSubStr("asdfghjklzxcvbnm,."))
	fmt.Println(lengthOfNonReatingSubStr("这里是上海"))
}
