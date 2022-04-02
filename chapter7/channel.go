package main

import "fmt"

// 채널은 데이터를 전송할 수 있는 파이프와 같은 형태라고 생각할 수 있는 구조이다.
// 채널은 항상 동일한 채널을 가리킨다. 즉 매개변수로 쓸 때 일반적인 변수들은 복사되어 사용하지만,
// 채널의 경우엔 포인터 처럼 해당 채널을 그대로 가리킨다.
// 채널의 종류에는 단방향(받기, 보내기), 양방향(받기/보내기겸용) 가 있다.
// 양방향채널을 사용하는 것도 좋지만, 단방향 채널은 역할을 확실히 구분하기에
// 자칫 복잡해질 수 있는 채널 프로그래밍을 단순하게 만들어 준다.

//채널을 이용하는 방법에는 몇가지 장점이있다.
// 1.생성하는 쪽에서 상태 저장방법을 복잡하게 고민할 필요가 없다.
// 2.받는 쪽에서는 for의 range를 이용할 수 있다.
// 3.채널 버퍼를 이용하면 멀티코어를 활용하거나 입출력 성능상의 장점을 이용할 수 있다.

//ExampleSimpleChannel 자주쓰이는 패턴이라고 한다. 익혀둘 것
// 채널 하나를 만들어서 넘겨주고 넘겹받는 것이 깔끔해 보이지 않기 때문에
// 주로 함수가 채널을 반환하게 만드는 패턴을 쓰게 된다.
func ExampleSimpleChannel() {
	// 보내는 부분과 받는 부분이 서로 데이터의 개수를 알지 못하더라도 동작
	c := func() <-chan int {
		c := make(chan int)
		go func() {
			//채널닫기
			defer close(c)
			c <- 1
			c <- 2
			c <- 3
		}()

		return c
	}()

	for num := range c {
		fmt.Println(num)
	}
}

//생성기 패턴
//피보나치 수열을 max까지 생성한다.
func Fibonacci(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for a <= max {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ",")
	}

	// Output:
	// 0,1,1,2,3,5,8,13,
}

//클로저를 이용한 피보나치
func FibonacciGenerator(max int) func() int {
	next, a, b := 0, 0, 1
	return func() int {
		next, a, b = a, b, a+b
		if next > max {
			return -1
		}
		return next
	}
}

func ExampleFibonacciGenerator() {
	fib := FibonacciGenerator(15)
	for n := fib(); n >= 0; n = fib() {
		fmt.Print(n, ",")
	}
	// Output:
	// 0,1,1,2,3,5,8,13,
}

//받는 쪽에서는 for의 Range를 이용할 수 있다.에 해당하는 경우
//아이의 이름을 정한다고 가정하고 첫,두번째 글자의 후보들을 주고 모든 경우의 수를 생성한다.
func BabyName(first, second string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func ExampleBabyNames() {
	for n := range BabyName("성정명재경", "준호우훈진") {
		fmt.Print(n, ",")
	}
}
