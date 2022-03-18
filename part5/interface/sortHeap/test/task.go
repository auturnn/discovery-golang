package main

import (
	"fmt"
	"sort"
	"time"
)

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
	SubTasks Tasks     `json:"subTask,omitempty"`
}

type Tasks []Task

func (t Tasks) init() {
	sort.Sort(t)
	for _, st := range t {
		st.MarkDone()
		fmt.Println(InCludeSubTasks(st))
	}
}

func (st Tasks) Len() int {
	return len(st)
}

func (st Tasks) Less(i, j int) bool {
	return st[i].Priority < st[j].Priority ||
		(st[i].Priority == st[j].Priority && st[i].Priority < st[j].Priority)
}

func (st Tasks) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
}

func (t *Task) MarkDone() {
	t.Status = DONE
	if len(t.SubTasks) > 0 {
		for i := range t.SubTasks {
			//재귀호출 실행. 하지만 t.SubTasks의 길이만큼 실행되기 때문에 따로 break를 걸 필요는 없어보인다.
			t.SubTasks[i].MarkDone()

		}
		sort.Sort(t.SubTasks)
	}
}

type status int

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %d %s", check, t.Title, t.Priority, t.Deadline)
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
	ts := []Task{{
		Title:    "Laundry2",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 3,
			SubTasks: []Task{
				{"PUT", DONE, nil, 3, []Task{
					{"PUT", DONE, nil, 3, nil},
					{"Detergent", TODO, nil, 2, nil},
				}},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 5,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}, {
		Title:    "Laundry1",
		Status:   TODO,
		Deadline: nil,
		Priority: 1,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 3,
			SubTasks: []Task{
				{"PUT", DONE, nil, 3, []Task{
					{"PUT", DONE, nil, 3, nil},
					{"Detergent", TODO, nil, 2, nil},
				}},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 5,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}}
	Tasks(ts).init()
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
