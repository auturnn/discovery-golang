package main

import (
	"fmt"
	"strconv"
	"strings"
)

type StrSet map[string]struct{}

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

//Map keyed by operator to set of higher percedence operators
type PrecMap map[string]StrSet

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	ops := []string{"("} // 초기 여는 괄호
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				//더 낮은 순위 연산자이므로 여기서 계산 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				//괄호를 제거하였으므로 종료
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			//닫는 괄호 여는 괄호까지 계산하고 제거
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}

	reduce(")") //초기의 여는 괄호까지 모두 계산
	return nums[0]
}

func ExampleNewEvaluator() {
	eval := NewEvaluator(map[string]BinOp{
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}
			if b < 0 {
				return 0
			}
			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}
			return r
		},
		"*":   func(a, b int) int { return a * b },
		"/":   func(a, b int) int { return a / b },
		"mod": func(a, b int) int { return a % b },
		"+":   func(a, b int) int { return a + b },
		"-":   func(a, b int) int { return a - b },
	}, PrecMap{
		"**":  NewStrSet(),
		"*":   NewStrSet("**", "*", "/", "mod"),
		"/":   NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+":   NewStrSet("**", "*", "/", "mod"),
		"-":   NewStrSet("**", "*", "/", "mod"),
	})

	fmt.Println("---Start Func ExampleNewEvaluator---")
	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1 - 2 - 4"))
	fmt.Println(eval("( 3 - 2 ** 3) * ( -2 )"))
	fmt.Println(eval("3 * ( 3 + 1 * 3 ) / ( -2 )"))
	fmt.Println(eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	fmt.Println(eval("1 + 2 ** 10 * 2"))
	fmt.Println(eval("2 ** 3 mod 3"))
	fmt.Println(eval("2 ** 2 ** 3"))
	fmt.Println("---End Func ExampleNewEvaluator---")
	// Output:
	// 5
	// 3
	// -5
	// 10
	// -9
	// 18
	// 2049
	// 2
	// 256
}
