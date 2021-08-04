package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var groupMap = make(map[string][]string)

	for _, str := range strs {
		sortstr := sortStr(str)
		if groupA, ok := groupMap[sortstr]; ok {
			groupMap[sortstr] = append(groupA, str)
		} else {
			groupMap[sortstr] = []string{str}
		}
	}

	var result = make([][]string, 0)
	for _, strs := range groupMap {
		result = append(result, strs)
	}

	return result
}

func sortStr(str string) string {
	var strs = make([]string, 0)
	for i := 0; i < len(str); i++ {
		strs = append(strs, str[i:i+1])
	}
	sort.Strings(strs)
	return strings.Join(strs, "")
}
