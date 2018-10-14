package parser

import (
	"project/learncode/crawler/engine"
	"project/learncode/crawler/model"
	"project/learncode/crawler/tools"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	// Age
	age, err := strconv.Atoi(extracrString(contents, ageRe))
	tools.CheeckErr(err)
	profile.Age = age

	//
	height, err := strconv.Atoi(extracrString(contents, heightRe))
	tools.CheeckErr(err)
	profile.Height = height

	profile.Income = extracrString(contents, incomeRe)
	profile.Weight = extracrString(contents, weightRe)
	profile.Gender = extracrString(contents, genderRe)
	profile.Xinzuo = extracrString(contents, xinzuoRe)
	// Marriage
	profile.Marriage = extracrString(contents, marriageRe)
	profile.Education = extracrString(contents, educationRe)
	profile.Occupation = extracrString(contents, occupationRe)
	profile.House = extracrString(contents, houseRe)
	profile.Car = extracrString(contents, carRe

	//profile.Name

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extracrString(contens []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contens)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
