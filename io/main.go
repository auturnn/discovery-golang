package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if err := fileCreate("./test.txt", 111); err != nil {
		log.Fatalln(err)
	}
	if err := fileOpen("./test.txt"); err != nil {
		log.Fatalln(err)
	}

	ExampleWriteTo()
	ExampleReadFrom()

	TestGraphWriteTo()
	TestGraphReadFrom()
}

func fileOpen(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// 읽은 결과(하나의 파일을 읽었기에 num: 1)
	var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
		fmt.Println("파일에 쓰여진 숫자:", num)
		return nil
	}

	return errors.New("파일을 여는 것에 실패하였습니다")
}

func fileCreate(filename string, num int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, "%d\n", num); err != nil {
		return err
	}

	return nil
}

//문자열 슬라이스를 라인별로 파일에 출력.
func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}

	return nil
}

func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func ExampleWriteTo() {
	lines := []string{
		"auturnn21@gmail.com",
		"holiow21@naver.com",
		"acki@kakao.com",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
}

func ExampleReadFrom() {
	r := strings.NewReader("auturnn21\nholiow21\nacki\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [auturnn21 holiow21 acki]
}

func graphWriteTo(w io.Writer, adjList [][]int) error {
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

func graphReadFrom(r io.Reader, adjList *[][]int) error {
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

func TestGraphWriteTo() {
	adjList := [][]int{
		{3, 4},
		{0, 2},
		{3},
		{2, 4},
		{0},
	}
	w := bytes.NewBuffer(nil)
	if err := graphWriteTo(w, adjList); err != nil {
		fmt.Println(err)
	}
	expected := "5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n"
	if expected != w.String() {
		fmt.Printf("expected: %s\n", expected)
		fmt.Errorf("found: %s\n", w.String())
	}
}

func TestGraphReadFrom() {
	r := strings.NewReader("5\n2 3 4\n2 0 2\n1 3\n2 2 4\n1 0\n")
	var adjList [][]int
	if err := graphReadFrom(r, &adjList); err != nil {
		fmt.Println(err)
	}
	fmt.Println(adjList)
	//Output:
	// [[3 4] [0 2] [3] [2 4] [0]]
}
