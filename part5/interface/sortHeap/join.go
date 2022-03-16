package sortHeap

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
	}
	return strings.Join(t, sep)
}

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// func ExampleJoin() {
// 	t := task.Task{
// 		Title:  "Laundry",
// 		Status: task.Done,
// 	}
// }

func ExampleJoin() {
	t := Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: nil,
	}
	fmt.Println(Join(",", 1, "two", 3, t))
	// OutPut: 1, two,3,[v] Laundry <nil>
}
