package main

import (
	"fmt"
	"sort"
	"strings"
)

// 문제 2.
// 정렬 인터페이스에 예제로 나와있는 ExampleCaseInsensitiveSort는 한가지 경우만 테스트한다.
// 테이블 기반 테스트를 활용한 TestCaseinsensitiveSort 함수를 구현하여 여러가지 경우의 수에 대하여 정렬이 제대로 동작하는지 확인.
// --> sort_test.go
type CaseInsensivie []string

func (c CaseInsensivie) Len() int {
	return len(c)
}

func (c CaseInsensivie) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensivie) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensivie([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	//Output:
	//[AppStore iPad iPhone MacBook]
}
