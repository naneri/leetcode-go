// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
    
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
        // if we get an opening bracket we append it to the stack
		if _, ok := pairs[r]; ok {
            stack = append(stack, r)

        // in case we get a closing bracket
		} else {
            // if the stack is empty, we return false
            if len(stack) == 0 {
                return false
            } 

            // if the last open bracket in the stack is not corresponding to the closing bracket, we return false
            if pairs[stack[len(stack)-1]] != r {
                return false
            }

            // we reduce the stack by one
            stack = stack[:len(stack)-1]
        }
	}

	return len(stack) == 0
}
