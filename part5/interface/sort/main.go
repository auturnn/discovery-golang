package main

import (
	"fmt"
	"sort"
	"strings"
)

//sort.Interface 참조 http://golang.org/pkg/sort/
type CaseInsensivie []string

func (c CaseInsensivie) Len() int {
	return len(c)
}
func (c CaseInsensivie) Less(i, j int) bool {
	fmt.Println("ㅗㅜㅑ")
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensivie) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func ExampleCaseInsensivie_sort() {
	apple := CaseInsensivie([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	//Output:
	//[AppStore iPad iPhone MacBook]
}
func main() {
	ExampleCaseInsensivie_sort()
}
