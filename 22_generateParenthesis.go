package main

func main() {

}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	var result = []string{}

	generateParenthesisDFS("", n, n, &result)
	return result
}

func generateParenthesisDFS(str string, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, str)
		return
	}
	if left > right {
		return
	}

	if left > 0 {
		generateParenthesisDFS(str+"(", left-1, right, result)
	}
	if right > 0 {
		generateParenthesisDFS(str+")", left, right-1, result)
	}
}
