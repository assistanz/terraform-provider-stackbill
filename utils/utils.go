package utils

import (
	"regexp"
	"strings"
	"terraform-provider-stackbill/config"
)

func GetProviderName() string {
	return config.Provider
}

func CamelCaseStringToSnakeCase(camelCase string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(camelCase, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
