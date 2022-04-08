package main

import (
	"context"
	"fmt"
	"sync"
)

// bufChannel에서 만든 PlusOne을 Context를 활용하여 수정한것.
func PlusOneWithContext(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			fmt.Println(num)
			select {
			case out <- num + 1:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func ExamplePlusOneWithContext() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()

	// 밖에서 취소신호를 받을 수 있는 경우에 이런 방식을 이용하면 좋다.
	// context.Context는 계층구조로 되어있다. 그중 Background()가 가장 상위에 위치하며
	// 프로그램이 종료될 때까지 계속 살아남는다.
	// 이것을 WithCancel을 통해 취소기능을 붙이고 ctx를 새로받아 변수형태로 저장하여 활용한다.
	ctx, cancel := context.WithCancel(context.Background())
	nums := PlusOneWithContext(ctx, PlusOneWithContext(ctx, PlusOneWithContext(ctx, PlusOneWithContext(ctx, PlusOneWithContext(ctx, c)))))
	for num := range nums {
		fmt.Println("num:", num)
		if num == 18 {
			//context.WithCancel을 통해 붙인 취소기능의 활용
			cancel()
			break
		}
	}
}

// Context의 이모저모
// Context는 관례상 다른 구조체안에 넣지 않고 함수의 첫인자로 주고받는다.
// WithDeadline, WithTimeout을 이용하여 만든 ctx를 이용하여 호출한뒤 지정된 시간후 취소하는 기능을 추가할 수도있다.
// 또한 WithValue를 이용한다면 인증 토큰 같이 요청범위 내에 있는 값들을 보낼 수 있어 편리하다고 한다.
type Request struct {
	Num int

	//구조체 내부에 chan내장
	Resp chan Response
}

type Response struct {
	Num      int
	WorkerID int
}

func PlusOneService(reqs <-chan Request, workerID int) {
	fmt.Println(workerID, "내부 대기")
	for req := range reqs {
		fmt.Println(workerID, "실행")
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}

func ExamplePlusOneService() {
	reqs := make(chan Request)
	defer close(reqs)

	//3개의 workerID 0,1,2
	for i := 0; i < 3; i++ {
		go PlusOneService(reqs, i)
	}
	fmt.Println("PlusOneService 대기")
	var wg sync.WaitGroup
	//5개의 요청 3,13,23,33,44
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resps := make(chan Response)

			//위에서 만들어진 reqs를 통해 대기중인 PlusOneService에 보낸다.
			reqs <- Request{i, resps}

			//PlusOneService의 결과는 Response 구조체의 채널을 통해 반환되어 받을 수 있다.
			fmt.Println(i, "=>", <-resps)

			//받는 부분을 다른방식으로 처리:
			// for resp := range resps {
			// 	fmt.Println(i, "=>", resp)
			// }
		}(i)
	}
	wg.Wait()
	//Output: 언제든 결과는 달라질 수 있다.(goroutine의 특성)
	/*
		0 내부 대기
		1 내부 대기
		PlusOneService 대기
		2 내부 대기
		2 실행
		1 실행
		0 실행
		2 실행
		33 => {34 2}
		1 실행
		43 => {44 1}
		13 => {14 0}
		3 => {4 2}
		23 => {24 1}
	*/

}
