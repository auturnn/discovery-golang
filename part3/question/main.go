package main

import (
	"fmt"
)

func main() {
	questionOne([]string{"사과", "바나나", "토마토", "약", "ㅁㄴ"})
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
			fmt.Printf("%s은(는) 판별에 유효한 문자가 아닙니다.", string(r))
		}
	}
}
