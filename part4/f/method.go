package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// type VertexID int

func ExampleVertexID_print() {
	i := VertexID(100)
	fmt.Println(i)
	// Output:
	// 100
}

//String 함수가 존재하지않을 경우 위 ExampleVertexID_print()의 Output대로 출력된다.
//참고로 리시버를 이용한 String() 함수는 id.String()을 호출하지않아도 자동으로 실행된다.
//이는 go에서의 reciver의 기본적 기능으로써 선언되면 VertexID를 사용할 경우 자동으로 사용된다.
func (id VertexID) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}

func ExampleVertexID_String() {
	i := VertexID(100)
	fmt.Println(i)
	// Output:
	// VertexID(100)
}

type MultiSet map[string]int

func (m MultiSet) Insert(val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "( "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + ")"
}

func ExampleWriteTo() {
	graph := Graph{
		{3, 4},
		{0, 2},
		{3},
		{2, 4},
		{0},
	}
	w := bytes.NewBuffer(nil)
	if err := graph.WriteTo(w); err != nil {
		fmt.Println(err)
	}
	expected := "5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n"
	if expected != w.String() {
		fmt.Printf("expected: %s\n", expected)
		fmt.Errorf("found: %s\n", w.String())
	}
}

type Graph [][]int

func (adjList Graph) WriteTo(w io.Writer) error {
	size := len(adjList)
	if _, err := fmt.Fprintf(w, "%d", size); err != nil {
		return err
	}
	for i := 0; i < size; i++ {
		lsize := len(adjList[i])
		if _, err := fmt.Fprintf(w, "\n%d", lsize); err != nil {
			return err
		}

		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fprintf(w, " %d", adjList[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}

func (adjList *Graph) ReadFrom(r io.Reader) error {
	var size int
	if _, err := fmt.Fscanf(r, "%d", &size); err != nil {
		return err
	}

	*adjList = make([][]int, size)
	for i := 0; i < size; i++ {
		var lsize int
		if _, err := fmt.Fscanf(r, "\n%d", &lsize); err != nil {
			return err
		}
		(*adjList)[i] = make([]int, lsize)
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fscanf(r, " %d", &(*adjList)[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fscanf(r, "\n"); err != nil {
		return err
	}

	return nil
}
