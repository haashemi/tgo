package main

import (
	"fmt"
	"regexp"
	"strings"
)

var returnTypePatterns = []*regexp.Regexp{
	regexp.MustCompile(`Returns the [a-z ]+ as ?a? (?P<type>[A-Za-z]+) `),
	regexp.MustCompile(`Returns the uploaded (?P<type>[A-Za-z]+)`),
	regexp.MustCompile(`Returns the (?P<type>[A-Za-z]+)`),
	regexp.MustCompile(`On success, the stopped (?P<type>[A-Za-z]+) is`),
	regexp.MustCompile(`On success, returns a (?P<type>[A-Za-z]+) object`),
	regexp.MustCompile(`On success, (?P<type>[A-Za-z]+) is returned`),
	regexp.MustCompile(`On success, a (?P<type>[A-Za-z]+) object is returned`),
	regexp.MustCompile(`On success, an (?P<type>array of [A-Za-z]+)s that were sent is returned`),
	regexp.MustCompile(`On success, the sent (?P<type>[A-Za-z]+) is returned`),
	regexp.MustCompile(`an (?P<type>Array of [A-Za-z]+) objects`),
	regexp.MustCompile(`a (?P<type>[A-Za-z]+) object`),
	regexp.MustCompile(`(?P<type>[A-Za-z]+) is returned`),
	regexp.MustCompile(`(?P<type>[A-Za-z]+) is returned, otherwise (?P<other>[a-zA-Z]+) is returned`),
	regexp.MustCompile(`(?P<type>[A-Za-z]+) on success`),
}

func extractReturnType(rawDesc []string) string {
	var parts []string
	for _, part := range strings.Split(strings.Join(rawDesc, ". "), ".") {
		tP := strings.ToLower(part)
		if strings.Contains(tP, "returns") || strings.Contains(tP, "returned") {
			parts = append(parts, strings.TrimSpace(part))
		}
	}

	if parts == nil {
		return ""
	}
	desc := strings.Join(parts, ". ")

	for _, pattern := range returnTypePatterns {
		matches := pattern.FindStringSubmatch(desc)
		if len(matches) == 2 {
			return matches[1]
		} else if len(matches) > 2 {
			return "json.RawMessage"
		}
	}

	return ""
}

func getType(key, s string, isOptional bool) string {
	if exactType := strings.TrimPrefix(s, "Array of "); exactType != s {
		return "[]" + getType(key, exactType, true)
	} else if exactType := strings.TrimPrefix(s, "array of "); exactType != s {
		return "[]" + getType(key, exactType, true)
	}

	switch key {
	case "parse_mode":
		return "ParseMode"
	case "media":
		switch s {
		case "InputMedia", "InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo":
			return "InputMedia"
		case "InputFile", "String", "InputFile or String":
			return "*InputFile"
		}
	}

	switch s {
	// Basic types
	case "Int", "Integer":
		return "int64"
	case "String":
		return "string"
	case "True", "Boolean":
		return "bool"
	case "Float", "Float number":
		return "float64"

	// Special types
	case "Integer or String":
		return "ChatID"
	case "InputFile", "InputFile or String":
		return "*InputFile"
	case "InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo":
		return "InputMedia"
	case "InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply":
		return "ReplyMarkup"

	// Other types, it should be a struct.
	default:
		if isOptional {
			return "*" + s
		}
		return s
	}
}

func getDefaultValue(t string) string {
	if strings.HasPrefix(t, "*") || strings.HasPrefix(t, "[]") {
		return "nil"
	}

	switch t {
	case "int64", "float64":
		return "0"
	case "string":
		return `""`
	case "bool":
		return "false"
	case "ReplyMarkup", "InputFile", "ChatID":
		return "nil"
	case "ParseMode":
		return `ParseModeNone`
	default:
		return "UNKNOWN_FIX_ME_NOW"
	}
}

func getStringerMethod(prefix, param, paramType string, isOptional bool) string {
	pascalCasedParam := snakeToPascal(param)

	switch getType(param, paramType, isOptional) {
	case "*InputFile":
		return ""
	case "bool":
		return fmt.Sprintf(`payload["%s"] = strconv.FormatBool(%s.%s)`, param, prefix, pascalCasedParam)
	case "int64":
		return fmt.Sprintf(`payload["%s"] = strconv.FormatInt(%s.%s, 10)`, param, prefix, pascalCasedParam)
	case "float64":
		return fmt.Sprintf(`payload["%s"] = fmt.Sprintf("%%f", %s.%s)`, param, prefix, pascalCasedParam)
	case "string":
		return fmt.Sprintf(`payload["%s"] = %s.%s`, param, prefix, pascalCasedParam)
	case "ParseMode":
		return fmt.Sprintf(`payload["%s"] = string(%s.%s)`, param, prefix, pascalCasedParam)
	default:
		return fmt.Sprintf(`if bb, err := json.Marshal(%s.%s); err != nil {
			return nil, err
		} else {
			payload["%s"] = string(bb)
		}`, prefix, pascalCasedParam, param)
	}
}
