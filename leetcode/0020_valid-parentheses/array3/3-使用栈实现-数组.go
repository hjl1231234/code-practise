package array3

// leetcode 20 有效的括号

func IsValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]int, len(s))
	length := 0
	var match = map[rune]int{
		')': 1,
		'(': -1,
		']': 2,
		'[': -2,
		'}': 3,
		'{': -3,
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack[length] = match[char]
			length++
		case ')', ']', '}':
			if length == 0 {
				return false
			}
			if stack[length-1]+match[char] != 0 {
				return false
			} else {
				length--
			}
		}
	}
	return length == 0
}
