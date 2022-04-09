package main

import (
	"fmt"
	"sync"
)

func ExampleClosedChannel() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//Output:
	// 0
	// 0
	// 0
}

//동시성 패턴

//파이프라인 패턴
/*
	파이프라인은 한 단계의 출력이 다음 단계의 입력으로 이어지는 구조.
	즉 컨베이어벨트와 비슷하다. 상품을 만들기 위해 분업하여 일을 처리하는 것과 같음.
	파이프라인 패턴은 생성기 패턴의 일종으로 생성기 패턴과 동일하게 받기 전용채널을 반환한다.
	그러나 받기 전용 채널을 넘겨받아서 입력으로 활용한다는 점에서 차이가 있다.
*/

//PlusOne 받기 전용채널을 받아서 다른 받기 전용 채널을 돌려주는 함수. 받은 채널에서 숫자를 하나 증가시켜서 보내준다.
func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			// 위에서 말한 파이프라인과 생성기 패턴의 차이점.
			// in을 통해서 넘겨받은 데이터를 out에 입력하여 반환
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()

	//PlusOne을 두번 쓰기때문에 +2가 증가된 값이 나온다.
	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

//아래는 일직선 파이프 라인 구성이다.
type IntPipe func(in <-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func ExampleChain() {
	fmt.Println(Chain(PlusOne, PlusOne))
}

//채널공유로 팬 아웃하기
//팬 아웃(Fan-Out) : 논리 회로에서 주로 쓰이는 용어. 게이트 하나의 출력이 게이트 여러 입력으로 들어가는 경우를 팬아웃이라고 한다.
//파이프라인의 앞 과정은 시간이 적게 걸리나, 뒷 과정은 시간이 오래 걸려서 지연되는 경우에
//뒷 사람 역할의 인원을 보충하여 놀고있는 사람에게 넘겨주는 것.
//방법은 간단하다. 채널 하나를 여럿에게 공유하면 된다.

func FanOut() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				// time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

// 위와 반대로 팬인하기
// 팬 인(Fan-In) : 논리회로에서 주로 쓰이는 용어. 하나의 게이트에 여러개의 입력선이 들어가는 경우를 팬인이라고 한다.
// 우표를 만든다고 가정하면 여러 사람들이 만든 결과물을 모아 상자에 넣는 사람은 혼자서도 빠르게 일처리가 가능할 것이다.
// 이것을 코드로 옮겼을 때, 첫 방법은 채널을 공유하는 것.
// 다만 해당 방법은 채널을 닫는 것에 주의하여야 한다.
// 채널을 닫을 때는 채널을 닫기위해 고루틴을 하나 더 만들어 사용한다.

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		//아래의 고루틴을 보면 in을 넘겨주는 것을 볼 수 있는데,
		// 이렇게 하지않으면 in값이 변경된 이후에 고루틴이 돌아갈 수도 있다.
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}

	// 백그라운드에서 대기하다가 모든 고루틴이 끝나면(wg.Done()) 채널을 닫는다.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Distribute 함수는 IntPipe형태의 함수를 받고 n개로 분산처리하는 함수로 돌려주는 함수.
// 팬아웃해서 파이프라인을 통과시키고 다시 팬인시키면 분산처리가 된다.
// 그러므로 이 함수는 팬아웃과 팬인을 모두 수행하는 함수이다.
func Distribute(p IntPipe, n int) IntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		// IntPipe를 통해 FanOut
		for i := 0; i < 0; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

//고루틴의 갯수가 많은 것은 크게 걱정할 필요가 없다.
//Go에서는 고루틴마다 스레드를 모두 할당하지 않으며, 동시에 수행될 필요가 없는 고루틴들은
//모두 하나의 스레드에서 순차적으로 수행되며, 이것이 컴파일 시간에 예측 가능한 경우가 많기에
//스레드를 많이 만드는 경우에 생길 수 잇는 비용이 발생하지 않는다.

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3

	closeChan := func(c *<-chan int) bool {
		*c = nil
		//openCnt를 줄여 열린채널갯수를 조정
		openCnt--
		return openCnt == 0
	}
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			}
		}
	}()
	return out
}

func ExampleFanIn3() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)
	for n := range FanIn3(c1, c2, c3) {
		fmt.Print(n, ",")
	}
}
