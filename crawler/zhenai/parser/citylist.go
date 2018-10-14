package parser

import (
	"project/learncode/crawler/engine"
	"regexp"
)

const cityListRe string = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
const cityLimit int = 10
func ParseCityList(contens []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contens, -1)

	result := engine.ParseResult{}
	limit := cityLimit
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if  limit <= 0{
			break
		}
	}

	return result
}
