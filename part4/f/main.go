package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	ExampleNewIntGenerator()
	ExampleNewIntGenerator_multiple()

	//값으로 취급되는 함수
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 4))

	//예시가 좋지는 못하지만 이런식으로 함수의 내용을 바꿀 수 있다.
	add = func(x, y int) int {
		return x + y + x + y
	}
	fmt.Println(2, 4)

	fmt.Println(Sqrt(2))
}

// NewIntGenerator 함수는 생성기(generator)의 예시이자 고계함수이다.
// 호출될때마다 증가된 값을 받을 수 있는 기능을 지닌다.
func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
}

func ExampleNewIntGenerator_multiple() {
	gen1 := NewIntGenerator()
	gen2 := NewIntGenerator()
	fmt.Println(gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen2(), gen1())
	// Output:
	// 1 2 3
	// 1 2 3 4 5
	// 4 6 5
}

func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

func ExampleAddOne() {
	n := []int{1, 2, 3, 4}
	AddOne(n)
	fmt.Println(n)
	//Output:
	//[2,3,4,5]
}

// ExampleReadFrom_append 함수는 클로저의 예시이다.
// 클로저(closer)는 외부에서 선언한 변수를
// 함수 리터럴 내에서 마음대로 접근할수있는 코드를 말한다.
func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")

	//밑에서처럼 변수의 선언은 ExampleReadFrom_append에서 선언되었지만
	//ReadFrom에서 슬라이스에 읽은 줄들을 첨가 시킬 수 있다.
	var lines []string
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(lines)
	//Output:
	// [bill tom jane]
}

// ReadFrom 함수는 고계함수(higher-order function)이다.
// 고계함수는 함수를 넘기고 받는 함수를 뜻하기에 조금 더 고차원적인 함수라는 의미이다.
func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return nil
}

//아래는 명명되지 않은 자료형
type VertexID int
type EdgeId int

func NewVertexIDGenerator() func() VertexID {
	var next VertexID
	return func() VertexID {
		next++
		return VertexID(next) //VertexID형으로 형변환하여 반환한다.
	}
}

type runes []rune

//명명된 자료형(rune)과 명명되지않은 자료형(runes)의 사이에는
//표현이 같으면 호환이 가능하다.
func testRune() {
	var a []rune = runes{65, 66}
	fmt.Println(a)
}

type BinOp func(int, int) int
type BinSub func(int, int) int

func OpThreeAndFour(f BinOp) {
	fmt.Println(f(3, 4))
}

func BinOpToBinSub(f BinOp) BinSub {
	var count int
	return func(i1, i2 int) int {
		fmt.Println(f(i1, i2))
		count++
		return count
	}
}

func ExampleBinOpToBinSub() {
	sub := BinOpToBinSub(func(a, b int) int {
		return a + b
	})
	sub(5, 6)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	//Output:
	//12
	//12
	//12
	//count: 3
}
