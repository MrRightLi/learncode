package main

import (
	"fmt"
	"regexp"
)

const text string = `
My email is mr.rightli2750@gmail.com
My email2 is mr.xxx@abc.com.cn
My email3 is    mr.yyy@qq.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9)]+)(\.[a-zA-Z0-9]+)`)
	//mat := re.FindAllString(text, -1)
	//fmt.Println(mat)

	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		fmt.Println(match)
	}
}


/*
正则基本知识:
- .+ 匹配任意一个多个
- .* 匹配任意0个或多个


 */