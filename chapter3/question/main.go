package main

import (
	"fmt"
	"strings"
)

func main() {
	Answer_array()
	AnswerASC([]int{11, 1, 4, 6, 12, 5, 2, 5, 12, 123, 51, 0})

	//TODO 함수이름 수정
	AnswerThree([]string{"한개", "두개", "세개"}, "두개")
	AnswerQueue()
	AnswerFive()
}

// 마지막글자에 받침이 있는 경우에도 어색하지않은 조사가 붙어서 출력되도록 코드를 수정하라.
//1번 예제 답
func Answer_array() {
	fruits := [6]string{"사과", "바나나", "토마토", "감", "귤", "asdasd"}

	start := []rune("가") //한글유니코드의 시작점 (44032)
	end := []rune("힣")   //한글유니코드의 끝점   (55203)

	for _, fruit := range fruits {
		r := []rune(fruit)
		last := r[len(r)-1]
		if start[0] <= last && last <= end[0] { // 한글유니코드에 포함되는가?

			if ((last - start[0]) % 28) == 0 {
				// 받침 존재하지 않을 시 국어문법에 따라 '는'을 조사로 붙인다.
				fmt.Printf("%s는 맛있다.\n", string(r))
			} else {
				// 받침 존재하지 않을 시 국어문법에 따라 '은'을 조사로 붙인다.
				fmt.Printf("%s은 맛있다.\n", string(r))
			}
		} else { // 에러처리 한글 유니코드상 해당 문자가 없을 경우
			fmt.Printf("%s은(는) 판별에 유효한 문자가 아닙니다.\n", string(r))
		}
	}
	// Output
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 감은 맛있다.
	// 귤은 맛있다.
}

// []int 슬라이스를 넘겨받아 오름차순으로 정렬하는 함수를 작성하라.
// 슬라이스 a의 i,j값을 맞바꿀때는 a[i], a[j] = a[j], a[i]를 참고하라.
func AnswerASC(a []int) {
	fmt.Println(len(a))
	for i := range a {
		fmt.Println(i)
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}

//정렬된 문자열 슬라이스가 있을 때, 특정 문자열이 슬라이스에 있는지를 조사하는 함수를 작성하라.
func AnswerThree(ss []string, a string) {
	for _, s := range ss {
		if s == a {
			fmt.Printf("'%s'는 해당 슬라이스 안에 존재합니다.\n", a)
			break
		}
	}
}

type Queue struct {
	data interface{}
}

// 슬라이스를 이용하여 큐(queue)를 구현하라. 큐에 자료를 넣는 것은 append를 이용하고,
// 자료를 꺼낼때는 q = q[1:]를 이용하라
// 또한 위와같은 방법을 지속적으로 사용하였을때 문제가 일어날지, 일어나지 않을지를 생각하라.
func AnswerQueue() {
	// 어떠한 자료형의 queue인지 명시되어있지 않기에
	// 간단히 int형 슬라이스를 사용할 것이다.
	q := []int{}
	fmt.Printf("NIL >> append >> len(q): %d, cap(q): %d, q:%v, mem: %p\n", len(q), cap(q), q, &q)

	for i := range [5]int{} {
		q = append(q, i)
		fmt.Printf("append >> len(q): %d, cap(q): %d, q:%v, mem: %p\n", len(q), cap(q), q, &q)
	}

	pop := func() {
		q = q[1:]
	}

	for range [5]int{} {
		pop()
		fmt.Printf("pop >> len(q): %d, cap(q): %d, q:%v, mem: %p\n", len(q), cap(q), q, &q)
	}
	// 질문에서 "말하는 문제가 일어날지, 일어나지 않을 지"를 생각하라는 것이
	// q = q[1:]에 대한 문제를 뜻한다면 문제가 일어나지 않는다가 정답인듯하다.
	// Golang에서 슬라이스는 앞의 자료를 잘라낼 경우(q[1:])에는 용량도 함께 줄어들지만,
	// 뒤의 자료를 잘라낼 경우(q[:len(q)-1])는 슬라이스의 용량과 값은 그대로이기 때문이다.
	// 하지만 잘라내어졌다고 인식되는 것은 q라는 자료형이 비추는 범위가 바뀌었기 때문에
	// 그렇게 보이는 것일 뿐, 실제로 삭제는 행해지지 않는 것이다.
	// 결론적으로 위 질문에 대한 나의 답은
	// "Queue의 특성(FIFO)상 q = q[1:]의 경우 길이,용량이 모두 같이 비워지기 때문에
	// 문제가 일어나지 않는다." 가 된다.

	// 또 다른 관점으로 용량과 길이를 정하지 않은 Queue의 경우
	// Append를 계속 반복하였을 때 문제가 발생할 수 도 있다고 생각한다.
	// Append의 출력결과를 확인하면 길이는 1-2-3-4-5의 순으로 증가하지만,
	// 용량은 1-2-4-4-8로 증가하기 때문이다.
	// 이것은 용량이 넘칠 경우 늘어난 분량(+1,+1...)만큼 슬라이스 메모리를 확보하는 것이 아니라
	// 슬라이스(cap:4) + 슬라이스(cap:4)의 형식으로 추가한다.
	// 이처럼 자료가 없더라도 메모리에 해당 용량만큼의 자리를 미리 예약해두는 것은
	// 성능상으로 좋지않은 결과를 초래한다고 생각된다.

	// 두가지 관점 모두를 통합하여 결론을 내린다면 Queue에 들어갈 자료의 수를 알고, 길이와 용량을 사전 설정해두는 것이라면
	// 삽입과 삭제 모두 문제가 되지 않지만, 미정된 길이와 용량을 가진 슬라이스의 경우 Append가 문제될수 있다고 생각된다.
	// 허나 슬라이스가 배열보다 자주쓰이는 이유인 '유연한 구조' 라는 이점을 생각한다면
	// 틀린 답일 것이라고도 보인다.
	fmt.Println(q)
}

//같은 원소가 여러 번 들어갈 수 있는 집합인 MultiSet을 기본제공하는 Map을 이용하자
func AnswerFive() {
	//다음과 같은 예제가 동작하도록 함수를 작성한다.
	m := NewMultiSet()
	fmt.Println(m.String())
	fmt.Println(m.Count("3"))
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	fmt.Println(m.String())
	fmt.Println(m.Count("3"))
	m.Insert("1")
	m.Insert("2")
	m.Insert("5")
	m.Insert("7")
	m.Erase("3")
	m.Erase("5")
	fmt.Println(m.Count("3"))
	fmt.Println(m.Count("1"))
	fmt.Println(m.Count("2"))
	fmt.Println(m.Count("5"))
	//Output:
	// {  }
	// 0
	// 3 3 3 3
	// 4
	// 3
	// 1
	// 1
	// 0
}

type MultiSet map[string]int

//새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() MultiSet {
	return make(map[string]int)
}

//Insert 함수는 집합에 val을 추가한다.
func (m MultiSet) Insert(val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := []string{}
	for key, val := range m {
		s = append(s, strings.Repeat(string(key)+" ", val))
	}
	return fmt.Sprintf("{ %s}", strings.Join(s, ""))
}
