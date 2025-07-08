package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	sampleStr := "hELlo$ _wORlD_"
	camelCaseResult := transformString(sampleStr, camelCase)
	pascalCaseResult := transformString(sampleStr, pascalCase)
	fmt.Printf("string: %v\ncamel case: %v\npascal case: %v\n", sampleStr, camelCaseResult, pascalCaseResult)
}

// transforms string with given function
func transformString(str string, fn func(str []string) string) string {
	// sanitising the string
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	santisizedWords := strings.Fields(regex.ReplaceAllString(str, " "))
	return fn(santisizedWords)
}

// formats the words slice as per pascal format
func pascalCase(words []string) string {
	for i, val := range words {
		words[i] = strings.ToUpper(val[:1]) + strings.ToLower(val[1:])
	}
	resultString := strings.Join(words, "")
	return resultString
}

// formats the words slice as per pascal format
func camelCase(words []string) string {
	resultString := pascalCase(words)
	resultString = strings.ToLower(resultString[:1]) + resultString[1:]
	return resultString
}
