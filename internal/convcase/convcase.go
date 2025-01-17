package convcase

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ConvCase struct {
	converter func(words []string) []string
	joiner    func(words []string) string
}

var (
	CamelCase = ConvCase{
		converter: camel,
		joiner:    joinFactory(""),
	}
	PascalCase = ConvCase{
		converter: title,
		joiner:    joinFactory(""),
	}
	SnakeCase = ConvCase{
		converter: lower,
		joiner:    joinFactory("_"),
	}
	ConstantCase = ConvCase{
		converter: upper,
		joiner:    joinFactory("_"),
	}
	KebabCase = ConvCase{
		converter: lower,
		joiner:    joinFactory("-"),
	}
	TrainCase = ConvCase{
		converter: title,
		joiner:    joinFactory("-"),
	}
	PathStyle = ConvCase{
		converter: nop,
		joiner:    joinFactory("/"),
	}
	LowerPathStyle = ConvCase{
		converter: lower,
		joiner:    joinFactory("/"),
	}
	DotStyle = ConvCase{
		converter: nop,
		joiner:    joinFactory("."),
	}
	LowerDotStyle = ConvCase{
		converter: lower,
		joiner:    joinFactory("."),
	}
)

func (c *ConvCase) Convert(words []string) string {
	if len(words) == 0 {
		return ""
	}

	return c.joiner(c.converter(words))
}

var (
	lowerCase = cases.Lower(language.English)
	upperCase = cases.Upper(language.English)
	titleCase = cases.Title(language.English)
)

func nop(words []string) []string {
	return words
}

func lower(words []string) []string {
	for i := range words {
		words[i] = lowerCase.String(words[i])
	}
	return words
}

func upper(words []string) []string {
	for i := range words {
		words[i] = upperCase.String(words[i])
	}
	return words
}

func title(words []string) []string {
	for i := range words {
		words[i] = titleCase.String(words[i])
	}
	return words
}

func camel(words []string) []string {
	for i := range words {
		if i == 0 {
			words[i] = lowerCase.String(words[i])
		} else {
			words[i] = titleCase.String(words[i])
		}
	}
	return words
}

func joinFactory(sep string) func([]string) string {
	return func(s []string) string {
		return strings.Join(s, sep)
	}
}
