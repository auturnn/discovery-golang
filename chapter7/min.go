package main

import (
	"fmt"
	"sync"
)

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	//최초 비교대상이 될 값 설정
	//반복문을 통해 값은 유동적으로 바뀔 것.
	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

func ExampleMin() {
	fmt.Println(ParalleMin([]int{
		2, 83, 22, 33, 45, 11,
		13, 56, 15, 64, 23,
	}, 4))
}

//ParalleMin 고루틴을 통해 최소값을 찾는 함수.
// 정렬되지않은 숫자나열을 가진 a와 몇개의 고루틴을 사용할 것인지를 지정하는 매개변수 n을 통해 최소값을 반환
func ParalleMin(a []int, n int) int {
	//작업에 할당될 사람(n)보다 작업량(len(a))이 적은 경우
	if len(a) < n {
		return Min(a)
	}
	mins := make([]int, n)

	//a:10, n:4 14-1  = 13/4 = 3(.25)
	size := (len(a) + n - 1) / n
	fmt.Println("size:", size)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// i:0 begin = 0 end = 3
			// i:1 begin = 3 end
			// i:4 begin = 12 end = 15
			begin, end := i*size, (i+1)*size
			// end가 len(a)를 넘어가는 경우 end를 len(a)의 마지막으로 수정
			if end > len(a) {
				end = len(a)
			}

			mins[i] = Min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return Min(mins)
}
