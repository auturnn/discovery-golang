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
			//요청 당 0개,1개 혹은 여러개의 응답을 받을 수 있다.
			//이것은 검색요청의 결과처럼 여러 개의 결과를 받아야하는 경우에 유용.
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

//동적으로 고루틴 이어붙이기

/*
	필요에 따라 동적으로 채널을 통하여 고루틴을 이어붙일 수 있다.
	(https://golang.org/doc/play/sieve.go) 이곳의 예제를 보면 prime의 배수들을 걸러내는 고루틴을 이어붙이는 형태로 구성되어있다.
	예제에서는 2부터 숫자를 하나씩 증가시켜 가면서 채널에 숫자를 보낸다.다른 고루틴에서는 이 채널에서 숫자를 받을 때마다 출력하고,
	출력된 숫자의 배수가 되는 숫자들을 걸러내는 필터 고루틴을 이어붙인다. 그러면 이미 출력된 숫자의 배수들은 다시는 출력되지 않고,
	이로인해 출력된 숫자들은 모두 소수가 된다.

	이 예제에서 나온 아이디어를 차용하여 파이프라인을 일직선으로 이어붙일 때 이전에 사용했던 생성기 패턴으로 필터를 만들고 컨텍스트 등을 이용하여
	아래에 소수생성기를 만들어본다.
*/

//Range 정수 생성기로 start부터 시작해서 step만큼 무한정 생성한다.
func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

//IntPipe2 IntPipe에서 context를
type IntPipe2 func(context.Context, <-chan int) <-chan int

//FilterMultiple n의 배술르 걸러내는 파이프라인을 반환한다.
//클로저를 이용하여 함수를 반환한 것은 파이프라인 함수형을 맞춰서 다른 파이프 라인에 연결하여 쓸 수 있게 하기위함.
func FilterMultiple(n int) IntPipe2 {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

//Primes 무한 소수 생성기. 여기서 조금 어려운 부분은 <-c와 같이 받는 부분에서 막혀 있을 때
//ctx가 취소가 될 수 있고, 여기서 받은 뒤에 그 값을 out <- i와 같이 보낼 때 막혀있다가 ctx가 취소될수 도있기 때문에
//아래의 형태처럼 select를 다중으로 만들어야 한다는 것.
//이처럼 작성해야 프로그램 종료까지 안죽고 살아있는 좀비 고루틴이 없어진다.
func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

//PrintPrimes max 숫자가 나올 때까지 Primes에서 소수를 순서대로 꺼내어 이용하다가
//범위를 넘어가버리면 반복문을 빠져 나간다.
//defer cancel()을 통해 최상위 Context를 종료하여 Primes로 넘어간 ctx가 취소되고 생성된 고루틴을 모두 소멸할 수 있게 되었다.
func PrintPrimes(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for prime := range Primes(ctx) {
		if prime > max {
			break
		}
		fmt.Print(prime, " ")
	}
	fmt.Println()
}
