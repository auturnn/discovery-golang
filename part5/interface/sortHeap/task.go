package sortHeap

import (
	"fmt"
	"time"
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
	fmt.Println(InCludeSubTasks(Task{
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
	}))
	// Output:
	// [ ] Laundry <nil>
	//   [ ] Wash <nil>
	//     [v] Push <nil>
	//	   [ ] Detergent <nil>
	// 	 [ ] Dry <nil>
	//   [ ] Fold <nil>
}
