package main

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var returnTypePatterns = []*regexp.Regexp{
	regexp.MustCompile(`On success, an (?P<type>array of [A-Za-z]+) of the sent messages is returned`),
	regexp.MustCompile(`On success, the stopped (?P<type>[A-Za-z]+) is`),
	regexp.MustCompile(`On success, returns a (?P<type>[A-Za-z]+) object`),
	regexp.MustCompile(`On success, (?P<type>[A-Za-z]+) is returned`),
	regexp.MustCompile(`On success, a (?P<type>[A-Za-z]+) object is returned`),
	regexp.MustCompile(`On success, an (?P<type>array of [A-Za-z]+)s that were sent is returned`),
	regexp.MustCompile(`On success, the sent (?P<type>[A-Za-z]+) is returned`),
	regexp.MustCompile(`Returns a (?P<type>[A-Za-z]+) object`),
	regexp.MustCompile(`Returns the [a-z ]+ as ?a? (?P<type>[A-Za-z]+) `),
	regexp.MustCompile(`Returns the uploaded (?P<type>[A-Za-z]+)`),
	regexp.MustCompile(`Returns the (?P<type>[A-Za-z]+)`),
	regexp.MustCompile(`an (?P<type>Array of [A-Za-z]+) objects`),
	regexp.MustCompile(`a (?P<type>[A-Za-z]+) object`),
	regexp.MustCompile(`(?P<type>[A-Za-z]+) is returned`),
	regexp.MustCompile(`(?P<type>[A-Za-z]+) is returned, otherwise (?P<other>[a-zA-Z]+) is returned`),
	regexp.MustCompile(`(?P<type>[A-Za-z]+) on success`),
}

func isMethod(s string) bool {
	return strings.ToLower(s)[0] == s[0]
}

func isArray(t string) bool {
	return strings.HasPrefix(t, "[]")
}

func getType(key, s string, isOptional bool, sections []Section) string {
	if exactType := strings.TrimPrefix(s, "Array of "); exactType != s {
		return "[]" + getType(key, exactType, true, sections)
	} else if exactType := strings.TrimPrefix(s, "array of "); exactType != s {
		return "[]" + getType(key, exactType, true, sections)
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
			if isInterface(s, sections) {
				return s
			}

			return "*" + s
		}
		return s
	}
}

func isInterface(typeName string, sections []Section) bool {
	for _, section := range sections {
		if section.Name == typeName && section.IsInterface {
			return true
		}
	}

	return false
}

func snakeToPascal(s string) string {
	return strings.ReplaceAll(cases.Title(language.English).String(strings.ReplaceAll(s, "_", "  ")), " ", "")
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

func defaultValueOfType(t string) string {
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

func getNestedMediaFields(all []Section, s Section) (fields []Field) {
	for _, f := range s.Fields {
		ft := getType(f.Name, f.Type, f.IsOptional, all)
		ft = strings.TrimPrefix(ft, "[]")
		ft = strings.TrimPrefix(ft, "*")

		if ft == "InputMedia" || ft == "InputPaidMedia" {
			fields = append(fields, f)
			continue
		}

		for _, xs := range all {
			if xs.Name == ft {
				for _, xf := range xs.Fields {
					if getType(xf.Name, xf.Type, xf.IsOptional, all) == "*InputFile" {
						fields = append(fields, f)
						break
					}
				}
				break
			}
		}
	}

	return fields
}

// function naming level: INFINITY
func stringifyTypeAndSetInMap(name, fieldName, pType string) string {
	switch pType {
	case "*InputFile":
		return ""
	case "bool":
		return fmt.Sprintf(`payload["%s"] = strconv.FormatBool(x.%s)`, fieldName, name)
	case "int64":
		return fmt.Sprintf(`payload["%s"] = strconv.FormatInt(x.%s, 10)`, fieldName, name)
	case "float64":
		return fmt.Sprintf(`payload["%s"] = fmt.Sprintf("%%f", x.%s)`, fieldName, name)
	case "string":
		return fmt.Sprintf(`payload["%s"] = x.%s`, fieldName, name)
	case "ParseMode":
		return fmt.Sprintf(`payload["%s"] = string(x.%s)`, fieldName, name)
	default:
		return fmt.Sprintf(`if bb, err := json.Marshal(x.%s); err != nil {
			return nil, err
		} else {
			payload["%s"] = string(bb)
		}`, name, fieldName)
	}
}
