package convcase

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	wordRegExp  = regexp.MustCompile(`[a-z0-9]+|[A-Z0-9]+|[A-Z0-9][a-z0-9]*`)
	kebabRegExp = regexp.MustCompile(fmt.Sprintf(`^(%s)(-(%s))*$`, wordRegExp.String(), wordRegExp.String()))
	snakeRegExp = regexp.MustCompile(fmt.Sprintf(`^(%s)(_(%s))*$`, wordRegExp.String(), wordRegExp.String()))
	pathRegExp  = regexp.MustCompile(fmt.Sprintf(`^(%s)(/(%s))*$`, wordRegExp.String(), wordRegExp.String()))
	dotRegExp   = regexp.MustCompile(fmt.Sprintf(`^(%s)(.(%s))*$`, wordRegExp.String(), wordRegExp.String()))
	textRegExp  = regexp.MustCompile(fmt.Sprintf(`^(%s)( (%s))*$`, wordRegExp.String(), wordRegExp.String()))
)

func SplitWords(text string) (words []string, err error) {

	if textRegExp.MatchString(text) {
		words = strings.Split(text, " ")
		return words, nil
	}

	if kebabRegExp.MatchString(text) {
		words = strings.Split(text, "-")
		return words, nil
	}

	if snakeRegExp.MatchString(text) {
		words = strings.Split(text, "_")
		return words, nil
	}

	if pathRegExp.MatchString(text) {
		words = strings.Split(text, "/")
		return words, nil
	}

	if dotRegExp.MatchString(text) {
		words = strings.Split(text, ".")
		return words, nil
	}

	return nil, fmt.Errorf("'%s' unknown style, supports '-', '_', '/', '.' splited words", text)
}
