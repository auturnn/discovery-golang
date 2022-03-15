package sortHeap

import (
	"container/heap"
	"fmt"
)

// Heap은 자료중에 가장 작은 값을 O(log N)의 시간 복잡도로 꺼낼 수 있는 자료구조이다.
// 여기서 가장 작은값이란 정렬의 순서상 가장 먼저 나오는 자료를 말한다.

//heap.Interface는 sort.Interface를 내장하고 있다.
// type Interface interface {
// 	sort.Interface
// 	Push(x interface{}) // add x as element Len()
// 	Pop() interface{}   //remove and return element Len()-1
// }

func (c *CaseInsensivie) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensivie) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	return last
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensivie([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		fmt.Println(heap.Pop(&apple))
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}
