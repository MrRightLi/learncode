package container

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(
		convertToBin(0),
		convertToBin(5),
		convertToBin(13),
		convertToBin(341234234),
	)
}

func convertToBin(n int) string {

	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}
