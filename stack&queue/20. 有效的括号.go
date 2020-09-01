package main

import "fmt"

func main() {
	fmt.Println(isValid("{{{}}}}"))
}

// 1,如果为左括号，则直接放入栈中
// 2,若为右括号，
//	如果栈中无元素，则绝对不匹配，返回false
//	若有元素，则比较栈中是否有对应右括号，有则弹出，无则返回false
// 3,检查stack是否为空
func isValid(s string) bool {
	var m = map[string]string{")": "(", "]": "[", "}": "{"}

	stack := make([]string, 0)

	for _, str := range s {
		stackLen := len(stack)

		// 如果为左括号，则直接放入栈中
		if _, ok := m[string(str)]; !ok {
			stack = append(stack, string(str))
			continue
		}

		// 若为右括号，
		// 如果栈中无元素，则绝对不匹配，返回false
		// 若有元素，则比较栈中是否有对应右括号，有则弹出，无则返回false
		if stackLen > 0 {
			if stack[stackLen-1] == m[string(str)] {
				stack = stack[:stackLen-1]
			} else {
				return false
			}
		} else {
			return false
		}

	}

	//检查stack是否为空
	return len(stack) == 0
}
