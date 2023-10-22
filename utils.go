package golang_dependency_two

import "strings"

func CapitalizeString(str string, another string) string {
	return strings.ToUpper(str) + another
}
