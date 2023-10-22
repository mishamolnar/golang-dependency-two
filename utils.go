package golang_dependency_two

import "strings"

func CapitalizeString(str string, another string, delimiter string) string {
	return strings.ToUpper(str) + delimiter + another
}
