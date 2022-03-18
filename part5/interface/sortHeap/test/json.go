package main

import (
	"encoding/json"
	"fmt"
)

// 문제 3.
// JSON으로 구조체를 직렬화하여 보관하고, 이것을 역직렬화하여 이용하는 코드가 있다.
// 원래 필드에 int64형에 이름을 붙인 ID라는 자료형을 이용하였는데, 자바스크립트에서 53byte이상의 정수를 제대로 읽을수 없는 문제때문에
// string형으로 변경하려고 한다. JSON 태그 `json:",string`을 이용하면 되는 문제지만 기존에 저장된 자료를 읽을 때 발생하는 문제가 있어서,
// 예전 형식이 정수로 되어있는 경우와 문자열로 되어 있는 경우 모두 읽을 수 있는 역직렬화코드를 작성하고 싶다. UnmarshalJSON과 형 단언을 이용하여 이것을 가능하게 해보자.
// --> json.go

// ...해당 문제 의문. 요구받은 역직렬화의 경우 json(byte)데이터를 구조체를 통해 string, int로 형 단언을 통해 출력할 수 있어야한단 것이 맞는지.

type JSON struct {
	ID ID `json:",string"`
}

type ID interface{}

func UnmarshalJSON(data []byte, j *JSON) {
	json.Unmarshal(data, &j)
	switch j.ID.(type) {
	case string:
		fmt.Println("string", j.ID)
	case int64:
		fmt.Println("int64", j.ID)
	case float64:
		fmt.Println("float64", j.ID.(float64))
	case int:
		fmt.Println("int", j.ID.(int))
	default:
		fmt.Println("이것은 잘못되었다.")
	}
}

func ExampleJSON() {
	j := JSON{}
	j.ID = ID(23).(int64)
	bytes, _ := json.Marshal(j)

	js := JSON{}
	UnmarshalJSON(bytes, &js)
}
