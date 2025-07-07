package main

import (
	"fmt"
	"strings"
	"regexp"
)

func main(){
	sampleStr := "hELlo$ _wORlD_"
	camelCaseResult := transformString(sampleStr, camelCase)
	pascalCaseResult := transformString(sampleStr, pascalCase)
	fmt.Printf("string: %v\ncamel case: %v\npascal case: %v\n", sampleStr, camelCaseResult, pascalCaseResult)
}

func transformString(str string, fn func(str string) string) string{
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	santisizedWords := strings.Fields(regex.ReplaceAllString(str, " "))
	for i, val := range santisizedWords{
		santisizedWords[i] = strings.ToUpper(val[:1]) + strings.ToLower(val[1:])
	}
	finalString := strings.Join(santisizedWords, "")
	return fn(finalString)
}

func pascalCase(str string) string{
	if len(str) <= 0{
		return strings.ToUpper(str[:1])
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

func camelCase(str string) string{
	if(len(str) <= 0){
		return strings.ToLower(str[:1])
	}
	return strings.ToLower(str[:1]) + str[1:]
}