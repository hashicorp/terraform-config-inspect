package tfconfig

import (
	"regexp"
	"strings"
)

// Variable represents a single variable from a Terraform module.
type Variable struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`

	// Default is an approximate representation of the default value in
	// the native Go type system. The conversion from the value given in
	// configuration may be slightly lossy. Only values that can be
	// serialized by json.Marshal will be included here.
	Default   interface{} `json:"default"`
	Required  bool        `json:"required"`
	Sensitive bool        `json:"sensitive,omitempty"`

	Validation *Validation `json:"validation,omitempty"`

	Pos SourcePos `json:"pos"`
}

// Validation represents a validation object from a single variable from a Terraform module.
type Validation struct {
	Condition    string            `json:"condition,omitempty"`
	ErrorMessage string            `json:"error_message,omitempty"`
	Fields       map[string]string `json:"fields,omitempty"`
}

type HclValidation struct {
	Condition    string `hcl:"condition"`
	ErrorMessage string `hcl:"error_message"`
}

func Between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func ReturnFields(str string) map[string]string {
	validations := map[string]string{}
	var levelUp []string
	strSplit := strings.Split(str, "(,)?")
	for y := 0; y < len(strSplit); y++ {
		fName := ""
		for {
			if -1 != strings.Index(strSplit[y], "\\{") {
				strSplitint := strings.Split(strSplit[y], "\\{")
				levels := strings.SplitAfter(strSplitint[0], "\"")
				if len(levels) > 1 {
					levelUp = append(levelUp, strings.Replace(levels[1], "\\\"", "", -1))
				}
				strSplit[y] = strSplitint[1]
				continue
			}
			break
		}
		for {
			if -1 != strings.Index(strSplit[y], "\\}") {
				strSplitint := strings.Split(strSplit[y], "\\}")
				strSplit[y] = strSplitint[0]
				if len(levelUp) > 0 {
					levelUp = levelUp[:len(levelUp)-1]
				}
				continue
			}
			break
		}

		//string
		lastSplit := strings.SplitAfter(strSplit[y], "\"")
		for u := 0; u < len(lastSplit); u++ {
			if len(lastSplit) < 4 {
				break
			}
			_, err := regexp.Compile("^" + strings.Replace(lastSplit[3], "\"", "", -1) + "$")
			if nil == err {
				if len(levelUp) > 0 {
					fName = strings.Join(levelUp, ":") + ":" + strings.Replace(lastSplit[1], "\\\"", "", -1)
				} else {
					fName = strings.Replace(lastSplit[1], "\\\"", "", -1)
				}
				validations[fName] = "^" + strings.Replace(lastSplit[3], "\\\"", "", -1) + "$"
				break
			}
		}
		if len(lastSplit) > 1 && validations[strings.Replace(lastSplit[1], "\\\"", "", -1)] != "" {
			continue
		}
		//fmt.Println(y)
		//number
		numbSplit := strings.SplitAfter(strSplit[y], ":")
		strSplit := strings.SplitAfter(numbSplit[0], "\"")
		if len(numbSplit) < 2 || len(strSplit) < 2 {
			continue
		}
		_, err := regexp.Compile("^" + strings.Replace(numbSplit[1], "\"", "", -1) + "$")
		if nil == err {
			if len(levelUp) > 0 {
				fName = strings.Join(levelUp, ":") + ":" + strings.Replace(lastSplit[1], "\\\"", "", -1)
			} else {
				fName = strings.Replace(strSplit[1], "\\\"", "", -1)
			}
			validations[fName] = "^" + strings.Replace(numbSplit[1], "\\\"", "", -1) + "$"
			continue
		}
	}
	return validations
}
