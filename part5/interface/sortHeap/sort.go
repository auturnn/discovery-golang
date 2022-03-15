package sortHeap

import (
	"fmt"
	"sort"
	"strings"
)

// 일반적으로 싱글 스레드에서 비교 정렬 알고리즘은 빠른정렬(Quicksort)가 좋다.
// 빠른정렬은 O(nlog n)의 평균 시간 복잡도를 가지지만 최악의 경우 O(n^2)이 될 수 있다.
// 7개 이하의 값들에 대해서는 삽입정렬(insertion sort)가 가장 효율적.
//삽입정렬은 O(n^2)의 비효율적인 알고리즘이지만, 작은 크기의 자료에 대해서는 빠른정렬보다 더 빠르다.
// sort.Sort에서는 기본적으로 빠른 정렬을 이용한다. 빠른정렬의 최악의 경우를 피하기 위해 피벗3개를 골라서 가운데 값을 고르는 중위법을 이용한다.
// 그렇게했지만 너무 깊이 빠른 정렬에 빠지게 되면 힙 정렬을 이용한다.
// 빠른 정렬을 이용하다가도 7개 이하의 자료에 대하여 정렬하는 상황을 만나면 삽입 정렬을 이용한다.
// 책의 저자는 이것이 꽤효율적으로 잘 구현되어있다고 한다.

// 밑의 Interface는 sort.Interface 참조: http://golang.org/pkg/sort/
// type Interface interface {
// 	//Len is the number of elements in the collection
// 	Len() int
// 	//Less reports whether the element with
// 	//index i should sort before the element with index j.
// 	Less(i, j int) bool
// 	//Swap swaps the elements with indexes i and j
// 	Swap(i, j int)
// }

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
