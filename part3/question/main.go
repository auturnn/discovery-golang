package main

import (
	"fmt"
	"strings"
)

func main() {
	questionOne([]string{"사과", "바나나", "토마토", "약", "ㅁㄴ"})
	questionTwo([]int{11, 1, 4, 6, 12, 5, 2, 5, 12, 123, 51, 0})
	questionThree([]string{"한개", "두개", "세개"}, "두개")
	questionFour("pop", "")
	questionFive()
}

// 마지막글자에 받침이 있는 경우에도 어색하지않은 조사가 붙어서 출력되도록 코드를 수정하라.
func questionOne(ss []string) {
	//일반적으로 유니코드를 외우지는 않는다고 생각하여
	//한글유니코드의 시작점, 끝점을 알아내고 실행하는 것으로 시작한다.
	start := []rune("가")
	end := []rune("힣")
	for _, s := range ss {
		// 이전에는 for를 통해 unicode를 추출하여 사용했다.
		// 하지만 이번 문제에서는 조사를 어느것으로 붙일지, 즉 마지막 글자의 유니코드만 확인하면 된다.
		// 그렇기때문에 밑의 방법처럼 글자를 []rune에 담은뒤
		// 마지막 글자만을 확인하기 위한 코드를 작성하였다.
		r := []rune(s)
		l := r[len(r)-1]
		if start[0] <= l && l <= end[0] {
			if ((l - start[0]) % 28) == 0 {
				fmt.Printf("%s는 받침이 없다.\n", string(r))
			} else {
				fmt.Printf("%s은 받침이 있다.\n", string(r))
			}
		} else {
			fmt.Printf("%s은(는) 판별에 유효한 문자가 아닙니다.\n", string(r))
		}
	}
}

// []int 슬라이스를 넘겨받아 오름차순으로 정렬하는 함수를 작성하라.
// 슬라이스 a의 i,j값을 맞바꿀때는 a[i], a[j] = a[j], a[i]를 참고하라.
func questionTwo(a []int) {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}

//정렬된 문자열 슬라이스가 있을 때, 특정 문자열이 슬라이스에 있는지를 조사하는 함수를 작성하라.
func questionThree(ss []string, a string) {
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
// 자료를 꺼낼때는 q= q[1:]를 이용하라
// 또한 위와같은 방법을 지속적으로 사용하였을때 문제가 일어날지, 일어나지 않을지를 생각하라.
func questionFour(action string, data ...interface{}) {
	q := []interface{}{"firstIn"}
	if action == "add" {
		// 용량을 정해두지 않은 슬라이스가 append를 할때 초과되면,
		// 원래의 용량만큼을 다시 만들어 붙이는 방법이기 때문에 해당 방법은 계속되면 메모리 누수가 심해질 것이다.
		// ex) cap:2인 슬라이스 q 에 append를 3번 진행할 경우 용량이 초과되기 때문에
		// q+q의 용량을 가진 새로운 슬라이스를 생성하여 append를 진행하게 된다.
		q = append(q, data...)
	} else if action == "pop" {
		// 이렇게 삭제를 진행할 경우 슬라이스의 용량은 변함이 없는데,
		// 슬라이스가 보고있는 범위만 바뀔뿐이다.
		q = q[1:]
	} else {
		fmt.Println("해당 명령은 유효하지않습니다. [info]: add, pop")
	}

	fmt.Println(q)
}

//같은 원소가 여러 번 들어갈 수 있는 집합인 MultiSet을 기본제공하는 Map을 이용하자
func questionFive() {
	//다음과 같은 예제가 동작하도록 함수를 작성한다.
	m := NewMultiSet()
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "1")
	Insert(m, "2")
	Insert(m, "5")
	Insert(m, "7")
	Erase(m, "3")
	Erase(m, "5")
	fmt.Println(Count(m, "3"))
	fmt.Println(Count(m, "1"))
	fmt.Println(Count(m, "2"))
	fmt.Println(Count(m, "5"))
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

//새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() map[string]int {
	return make(map[string]int)
}

//Insert 함수는 집합에 val을 추가한다.
func Insert(m map[string]int, val string) {
	m[val]++
}

func Erase(m map[string]int, val string) {
	m[val]--
}

func Count(m map[string]int, val string) int {
	return m[val]
}

func String(m map[string]int) string {
	s := []string{}
	for key, val := range m {
		for i := 0; i < val; i++ {
			s = append(s, key)
		}
	}
	return fmt.Sprintf("{ %s }", strings.Join(s, " "))
}
