package main

import (
	"fmt"
	"time"
)

// 문제 1.
// 이번장에서 만든 Task구조체를 MarkDone 메서드를 구현해보자.
// 이 메서드가 호출되면 해당 작업 및 SubTask 모두 상태가 DONE으로 바뀐다.
func (t *Task) MarkDone() {
	t.Status = DONE
	if len(t.SubTasks) > 0 {
		for i := range t.SubTasks {
			//재귀호출 실행. 하지만 t.SubTasks의 길이만큼 실행되기 때문에 따로 break를 걸 필요는 없어보인다.
			t.SubTasks[i].MarkDone()
		}
	}
}

// 문제 2.
// 정렬 인터페이스에 예제로 나와있는 ExampleCaseInsensitiveSort는 한가지 경우만 테스트한다.
// 테이블 기반 테스트를 활용한 TestCaseinsensitiveSort 함수를 구현하여 여러가지 경우의 수에 대하여 정렬이 제대로 동작하는지 확인.
func TestCaseinsensitiveSort() {

}

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Deadline struct {
	time.Time
}

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTask,omitempty"`
}

type status int

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func PrintStringer(data fmt.Stringer) {
	fmt.Print(data.String())
}

type InCludeSubTasks Task

func (t InCludeSubTasks) String() string {
	return t.indentedString("")
}

func (t InCludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + InCludeSubTasks(st).indentedString(prefix+" ")
	}
	return str
}

func ExampleIncludeSubTasks_String() {
	t := Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"PUT", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}
	t.MarkDone()
	fmt.Println(InCludeSubTasks(t))
	// Output:
	// [ ] Laundry <nil>
	//   [ ] Wash <nil>
	//     [v] Push <nil>
	//	   [ ] Detergent <nil>
	// 	 [ ] Dry <nil>
	//   [ ] Fold <nil>
}

func main() {
	ExampleIncludeSubTasks_String()
}
