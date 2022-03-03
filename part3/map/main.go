package main

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

func main() {
	/*
		Golang에서 map은 헤시테이블로 구현된다.
		해시맵은 key & value로 구분되어있고, 키를 이용해 값을 상수시간( O(n) )에 가져올 수 있지만,
		저장되는 데이터들에 대한 순서를 가지지는 않는다.
		아래는 map의 선언방법 3가지이다.
	*/

	//기초
	exampleOne()
	exampleTwo()
	exampleThree()
	exampleFour()

	//map에서 key를 삭제할때
	m := map[string]string{
		"key": "value",
		"hi":  "hello",
	}
	delete(m, "key")
	fmt.Println(m) // "key": "value"는 출력되지않는다.

	//활용
	badExampleCount()
	ExampleCount()

	fmt.Println(hasDupeRune("aabc"))
	fmt.Println(hasDupeRune("abc"))

	fmt.Println(betterHasDupeRune("aabc"))
	fmt.Println(betterHasDupeRune("abc"))

}

// exampleOne :
// map을 변수에 담을 수는 있지만, map자체가 생성은 되어있지 않다
// 따라서 해당 방법을 사용한 직후의 map은 nil 상태이기에 읽을 수는 있지만, 변경은 불가능한 상태
// 아래는 오류가 panic이 일어나는 코드(선언이후 초기화를 진행하지 않아 값 변경이 불가한 상황을 재현)
func exampleOne() {
	var m1 map[string]string
	m1["key"] = "value!"

	// key를 이용하여 value와 value의 존재여부를 bool형으로 반환가능
	val, ok := m1["key"]
	if !ok {
		fmt.Println("저장된 값이 없습니다.")
	} else {
		fmt.Println("Key에 저장된 값:", val)
	}
}

// exampleTwo :
// 해당 방법은 make()를 이용하여 초기화를 같이 진행
// 참고로 초기화가 진행된 뒤에 삽입되지 않은 key를 불러올 경우 panic을 일으키지 않고
// value 자료형의 기본값을 가져온다.(string="", int=0...)
func exampleTwo() {
	m2 := make(map[string]string)
	m2["key"] = "value!"
	val, ok := m2["key"]
	if !ok {
		fmt.Println("저장된 값이 없습니다.")
	} else {
		fmt.Println("Key에 저장된 값:", val)
	}
}

// exampleThree :
// exampleTwo 방법과 마찬가지로 빈맵으로 초기화할 수 있지만 선언형식의 차이존재
func exampleThree() {
	m3 := map[string]string{}
	m3["key"] = "value!"
	val, ok := m3["key"]
	if !ok {
		fmt.Println("저장된 값이 없습니다.")
	} else {
		fmt.Println("Key에 저장된 값:", val)
	}
}

// 4번 방법:
// 2,3번 방식에서 사용한 문법의 축약법
func exampleFour() {
	m4 := map[string]string{"key": "value!"}
	fmt.Println(m4["key"])
}

// count 함수는 이미 생성된 맵과 문자를 받고 문자열 안에 각 문자수를 세는 기능을 한다
// 특징으로는 반환값을 가지지도, 포인터를 가지지도 않는 다는 것이다
// map은 슬라이스와 다르게 맵은 이용할 때 맵 변수자체에 다시 할당하지 않기에
// 포인터를 사용하지 않아도 변경이 가능하다
func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

// badExample 함수는 count함수가 작동하여 map에 올바른 결과가 생성되었는지를 확인한다
func badExampleCount() {
	m5 := map[rune]int{}
	count("aabbccc", m5)

	if !reflect.DeepEqual(
		map[rune]int{'a': 2, 'b': 2, 'c': 3},
		m5,
	) {
		log.Panic("결과값이 같지않습니다. m5:", m5)
	}
	log.Println("성공: ", m5)
}

// ExampleCount 함수는 지정한 키 이외에 다른키가 있는 경우를 포함한 map의 활용법 예제
func ExampleCount() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)

	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)

	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}

	// Output:
	// 가 1
	// 나 2
	// 다 1
}

// hasDupeRune 함수는 문자열을 받아 중복되는 글자가 있는지 검사하는 기능을 지닌다.
// 다만 해당 함수는 map[rune]bool{}을 쓰고있지만, 더 나은 방법으로는 betterHasDupeRune 함수를 참조.
func hasDupeRune(s string) bool {
	runeSet := map[rune]bool{}
	for _, r := range s {
		if runeSet[r] {
			return true
		}
		runeSet[r] = true
	}

	return false
}

// betterHasDupeRune 함수는 HasDupeRune함수에서 사용되었던 자료형 map[rune]bool에서 value를 struct{}로 바꾼것
// 허나 이렇게 빈 구조체를 value로 사용할 경우 메모리를 따로 차지하지 않는다는 장점이 있다.(오버헤드x)
// 또다른 의문의 밑에서 주석으로 정리
func betterHasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		// 밑에서 보여지듯 runeSet[r]의 데이터를 받는데에 두가지 변수가 사용된다.
		// 하나는 '_' (ignore), 'ok' 두가지 인데 앞의 '_'에는 원래 struct{}가 담겨야하지만 사용하지 않으니 '_'처리.
		// 'ok'의 자리는 map[key]를 통해 저장된 value를 받을때, 해당 키값에 해당하는 데이터가 있는지, 없는지를 bool형태로 반환한다.
		// 이것은 golang에서 지원하는 map의 기능이기 때문에 map에서 데이터자료형으로 bool을 쓰기보단 빈 구조체를 사용하는 것이 좋다.
		if _, ok := runeSet[r]; ok {
			return true
		}
		// 위에서도 느꼈겠지만 구조체뒤에 괄호집합이 하나더 있는 것을 확인할 수 있는데,
		// 이것은 빈 구조체 자료형 struct{}로 초기화하는 문법이다.
		runeSet[r] = struct{}{}
	}
	return false
}
