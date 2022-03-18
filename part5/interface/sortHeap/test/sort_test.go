package main

import (
	"reflect"
	"sort"
	"testing"
)

// 문제 2.
// 정렬 인터페이스에 예제로 나와있는 ExampleCaseInsensitiveSort는 한가지 경우만 테스트한다.
// 테이블 기반 테스트를 활용한 TestCaseinsensitiveSort 함수를 구현하여 여러가지 경우의 수에 대하여 정렬이 제대로 동작하는지 확인.
// 대문자,소문자를 섞어서
// --> main_test.go
func TestCaseinsensitiveSort(t *testing.T) {
	s := [][]string{
		{"iPhone", "iPad", "MacBook", "AppStore"},
		{"One", "Two", "Three", "Four", "fIve"},
		{"one", "One", "ONe", "ONE"},
	}
	oks := [][]string{
		{"AppStore", "iPad", "iPhone", "MacBook"},
		{"fIve", "Four", "One", "Three", "Two"},
		{"ONE", "ONe", "One", "one"},
	}

	for i, ss := range s {
		a := CaseInsensivie(ss)
		sort.Sort(a)

		ok := CaseInsensivie(oks[i])

		if !reflect.DeepEqual(a, ok) {
			t.Errorf("TestCaseinsensivieSort is error. a is %s, ok is %s", a, ok)
		}
	}
}
