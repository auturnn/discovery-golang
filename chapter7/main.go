package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

//동시성

// 동시성이 있는 루틴은 서로 의존관계가 없다.
// 동시성과 병렬성은 다르지만 동시성이 있어야 병렬성이 생기게 된다.
// 이것은 두가지 행동에 대해 의존 관계가 성립할 경우
// 동시성은 존재 할 수 없다는 것을 의미한다.
// 유닉스에서의 명령어 &을 붙이는 것과 마찬가지.
// 백그라운드 프로세스는 현재의 흐름과는 연관이 없어진다.

//입출력이 오래걸리는 작업들을 동시에 수행시켜 더욱 빠른 프로그램을 작성할 수있다.
// 또한, 멀티코어도 활용할 수 있는 코드를 작성할 수 있게된다.

// 고루틴은 일반적인 코드처럼 기능하지 않는데, 이것을 컨트롤 하기 위해서 sync라이브러리가 제공된다.
func main() {

	// var urls = []string{
	// 	"https://image.com/img01.jpg",
	// 	"https://image.com/img02.jpg",
	// 	"https://image.com/img03.jpg",
	// }
	// // waitGroup을 통해 동시성 제어
	// var wg sync.WaitGroup
	// wg.Add(len(urls)) // 기다리는 타이머
	// for _, url := range urls {
	// 	// for문이 실행될 때마다 Add를 +1씩 증가
	// 	// 다만 아래의 방법은 for문이 반복되기 전에 이전에 진행되던 고루틴들이 모두 끝날 경우
	// 	// 고루틴이 wg.Wait()을 통과해버릴수도 있다. 이것을 rage Condition이라고 한다.
	// 	// wg.Add(1)
	// 	go func(url string) {
	// 		defer wg.Done() // 해당 함수가 종료되었음을 알림
	// 		if _, err := download(url); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}(url)
	// }
	// wg.Wait() // 위의 Add숫자만큼 Done이 실행되어야 다음으로 넘길 수 있도록 기다림

	// //파일경로중 .jpg로 형성된 파일들의 이름을 가져온다.
	// filenames, err := filepath.Glob("*.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //완성될 압축파일명과 압축대상 파일명 목록을 보낸다.
	// err = writeZip("images.zip", filenames)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ExampleMin()
	ExampleSimpleChannel()
	fmt.Println("----")
	ExampleFibonacci()
	fmt.Println("----")
	ExampleBabyNames()
	fmt.Println("----")
	ExampleFanIn3()
	fmt.Println("----")
	ExamplePlusOneWithContext()
	fmt.Println("----")
	ExamplePlusOneService()
	fmt.Println("----")
}

//download는 url에 담겨있는 파일을 다운받아 파일이름과 에러를 반환한다.
func download(url string) (string, error) {
	//http에 GET요청 후 반환된 값을 res에 저장
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}

	//위에서 받은 filename을 이름으로 파일을 생성
	//이 시점에서 파일 내용은 없고, 수정할 수 있도록 열린 상태로 대기중
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	//함수가 끝나면 다운이 끝났다는 것이므로 대기파일(다운파일)을 닫는다.
	defer f.Close()

	//os.Create로 생성된 파일내용에 res.Body를 복사한다.
	_, err = io.Copy(f, res.Body)
	return filename, err
}

//urlToFilename은 rawurl을 받아 이미지 파일 이름 부분을 반환한다.
func urlToFilename(rawUrl string) (string, error) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	//filepath.Base는 url의 가장 마지막부분을 반환한다. 따라서 img01-03.jpg를 결과로 반환
	return filepath.Base(url.Path), nil
}

//writeZip은 다운된 파일들의 이름과, 이를 압축하여 생성될 파일의 이름을 받아 결과를 error로 반환한다.
func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}

	// 압축된 데이터가 담길 파일(outf)를 넣는다.
	zw := zip.NewWriter(outf)

	for _, filename := range filenames {
		//파일이름을 사용하여 압축리스트에 추가한다.
		//반환된 w는 압축될 데이터(파일)를 기록하는 역할을 수행한다.
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}

		// 파일에 대한 데이터를 얻기위해 실행
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		//함수가 종료될때 마찬가지로 대기중은 f(파일)을 닫는다.
		defer f.Close()

		//압축 대상의 데이터를 zip에 복사한다.
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}

	return zw.Close()
}
