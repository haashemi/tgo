package tggen

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var tCase = cases.Title(language.English)

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

func upperFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func snakeToPascal(s string) string {
	return strings.ReplaceAll(tCase.String(strings.ReplaceAll(s, "_", "  ")), " ", "")
}

func typeOf(key, s string) string {
	if exactType := strings.TrimPrefix(s, "Array of "); exactType != s {
		return "[]" + typeOf(key, exactType)
	} else if exactType := strings.TrimPrefix(s, "array of "); exactType != s {
		return "[]" + typeOf(key, exactType)
	}

	switch key {
	case "parse_mode":
		return "ParseMode"
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
		return "InputFile"
	case "InputMediaAudio, InputMediaDocument, InputMediaPhoto and InputMediaVideo":
		return "InputMedia"
	case "InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply":
		return "InlineKeyboard"

	// Other types, it should be a struct.
	default:
		return "*" + s
	}
}

func extractReturnType(desc string) []string {
	var parts []string
	for _, part := range strings.Split(desc, ".") {
		tP := strings.ToLower(part)
		if strings.Contains(tP, "returns") || strings.Contains(tP, "returned") {
			parts = append(parts, strings.TrimSpace(part))
		}
	}

	if parts == nil {
		return nil
	}
	desc = strings.Join(parts, ". ")

	for _, pattern := range returnTypePatterns {
		if matches := pattern.FindStringSubmatch(desc); len(matches) != 0 {
			return matches[1:]
		}
	}

	return nil
}
