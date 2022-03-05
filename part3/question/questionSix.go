package main

import (
	"fmt"
	"strings"
)

//같은 원소가 여러 번 들어갈 수 있는 집합인 MultiSet을 기본제공하는 Map을 이용하자
func questionFive() {
	//다음과 같은 예제가 동작하도록 함수를 작성한다.
	m := NewMultiSet()
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "1")
	Insert(m, "2")
	Insert(m, "5")
	Insert(m, "7")
	Erase(m, "3")
	Erase(m, "5")
	fmt.Println(Count(m, "3"))
	fmt.Println(Count(m, "1"))
	fmt.Println(Count(m, "2"))
	fmt.Println(Count(m, "5"))
	//Output:
	// {  }
	// 0
	// 3 3 3 3
	// 4
	// 3
	// 1
	// 1
	// 0
}

//새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() map[string]int {
	return make(map[string]int)
}

//Insert 함수는 집합에 val을 추가한다.
func Insert(m map[string]int, val string) {
	m[val]++
}

func Erase(m map[string]int, val string) {
	m[val]--
}

func Count(m map[string]int, val string) int {
	return m[val]
}

func String(m map[string]int) string {
	s := []string{}
	for key, val := range m {
		for i := 0; i < val; i++ {
			s = append(s, key)
		}
	}
	return fmt.Sprintf("{ %s }", strings.Join(s, " "))
}
