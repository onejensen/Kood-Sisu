package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	return sort_string(strings.ToLower(str1)) == sort_string(strings.ToLower(str2))
}
func sort_string(s string) string {
	sort_arr := []string{}
	sort_arr = strings.Split(s, "")
	sort.Strings(sort_arr)
	result := strings.Join(sort_arr, "")
	return strings.Trim(result, " ")
}
