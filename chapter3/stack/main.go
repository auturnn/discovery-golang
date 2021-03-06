package main

import (
	"fmt"
	"strconv"
	"strings"
)

//stack은 LIFO의 구조를 갖는 자료구조.
//go에는 따로 스택이 없는데, 이것은 이미 슬라이스로도 충분히 구현이 가능하기 때문.
//책의 예제에서 스택을 활용한 함수예제가 있기때문에 그것을 밑에 적어놓도록함.
func main() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("1 + 2"))
	fmt.Println(Eval("1 - 2 + 3"))
	fmt.Println(Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))

	// Outputs :
	// 5
	// 3
	// 2
	// 9
	// 18
}

func Eval(expr string) int {
	// ops는 연산자를 담는 스택(슬라이스).
	var ops []string
	// nums는 정수들을 담는 스택(슬라이스).
	var nums []int

	//pop은 스택의 마지막 숫자를 출력하는 익명함수를 변수로써 지정
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	//reduce는 숫자,연산자들을 구분하여 처리하는 익명함수를 변수로써 지정
	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				//목록에 없는 연산자이므로 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				//괄호를 종료하였으므로 종료
				return
			}
			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			//덧셈과 뺄셈 이상의 우선순위를 가진 사칙연산 적용
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			reduce("*/")
			ops = append(ops, token)
		case ")":
			//닫는 괄호는 여는 괄호까지 계산하고 제거
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}

	reduce("+-*/")

	return nums[0]
}
