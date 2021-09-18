package main

func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return []string{}
	}
	var result = make([]string, 0)
	backtraceRestoreIpAddresses(s, 0, "", &result)
	return result
}

func backtraceRestoreIpAddresses(s string, start int, solution string, result *[]string) {
	var length = len(s)
	var sollen = len(solution)
	if sollen - 4 == length {
		if start == length {
			*result = append(*result, solution[:sollen-1])
		}
		return
	}

	if start == length {
		return
	}

	backtraceRestoreIpAddresses(s, start+1, solution+s[start:start+1]+".", result)
	if s[start] == '0' {
		return
	}

	if length-start >= 2 {
		backtraceRestoreIpAddresses(s, start+2, solution+s[start:start+2]+".", result)
	}

	if length-start >= 3 && s[start:start+3] <= "255" {
		backtraceRestoreIpAddresses(s, start+3, solution+s[start:start+3]+".", result)
	}
}
