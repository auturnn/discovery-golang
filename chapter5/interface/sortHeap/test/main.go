package main

// 문제 1.
// 이번장에서 만든 Task구조체를 MarkDone 메서드를 구현해보자.
// 이 메서드가 호출되면 해당 작업 및 SubTask 모두 상태가 DONE으로 바뀐다.

// 문제 2.
// 정렬 인터페이스에 예제로 나와있는 ExampleCaseInsensitiveSort는 한가지 경우만 테스트한다.
// 테이블 기반 테스트를 활용한 TestCaseinsensitiveSort 함수를 구현하여 여러가지 경우의 수에 대하여 정렬이 제대로 동작하는지 확인.
// 대문자,소문자를 섞어서
// --> sort_test.go

// 문제 3.
// JSON으로 구조체를 직렬화하여 보관하고, 이것을 역직렬화하여 이용하는 코드가 있다.
// 원래 필드에 int64형엥 이름을 붙인 ID라는 자료형을 이용하였는데, 자바스크립트에서 53byte이상의 정수를 제대로 읽을수 없는 문제때문에
// string형으로 변경하려고 한다. JSON 태그 `json:",string`을 이용하면 되는 문제지만 기존에 저장된 자료를 읽을 때 발생하는 문제가 있어서,
// 예전 형식이 정수로 되어있는 경우와 문자열로 되어 있는 경우 모두 읽을 수 있는 역직렬화코드를 작성하고 싶다. UnmarshalJSON과 형 단언을 이용하여 이것을 가능하게 해보자.
// --> json.go

// func main() {
// 	ExampleIncludeSubTasks_String()
// 	ExampleJSON()
// }
