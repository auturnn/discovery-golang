package main

import (
	"fmt"
	"time"
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
type IntPipe func(<-chan int) <-chan int

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
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
