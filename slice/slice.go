package main

import (
	"log"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 2 //길이상 표기

	//맨뒤 숫자 추가
	s = append(s, 99)
	log.Println(s)

	// index 다음에 숫자 추가 ( O(n) )
	s = add(s, index, 999)
	log.Println(s) // O(n)의 시간복잡도를 갖지만, 순서가 바뀌지않는다.

	// index 다음 숫자 삭제
	s[index] = s[len(s)-1] // s[2] : 지우고싶은 숫자, s[len(s)-1] : 슬라이스 최후미의 숫자
	s = s[:len(s)-1]       // s[:len(s)-1] : 원래의 최후미숫자를 제외한 나머지를 다시 복사하여 넣기
	log.Println(s)         // O(1)의 시간복잡도를 가지지만, 순서가 바뀜

	// index 다음 숫자 삭제
	s = delete(s, index)
	log.Println(s)
}

// add 슬라이스 중간추가
// index : 삽입하려는 차례
// d : 삽입하려는 숫자
func add(s []int, index, d int) []int {
	s = append(s[:index+1], s[index:]...)
	s[index] = d
	return s
}

func delete(s []int, index int) []int {
	s = append(s[:index], s[index+1:]...)
	return s
}
