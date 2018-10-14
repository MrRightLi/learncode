package parser

import (
	"project/learncode/crawler/engine"
	"regexp"
)

const cityRe string = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">(.+)</a>`

func ParseCity(contens []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contens, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, string(m[2]))
			},
		})
	}

	return result
}
