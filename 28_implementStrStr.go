package main

func main() {

}

func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}

	var strIndex = -1
	for i := 0;i < len(haystack)-len(needle)+1;i++ {
		var equal = true
		var iterI = i
		for j := 0;j < len(needle); {
			if haystack[iterI] != needle[j] {
				equal = false
				break
			}
			iterI++
			j++
		}
		if equal {
			strIndex = i
			break
		}
	}

	return strIndex
}
