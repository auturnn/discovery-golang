package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
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

	var urls = []string{
		"https://image.com/img01.jpg",
		"https://image.com/img02.jpg",
		"https://image.com/img03.jpg",
	}
	// waitGroup을 통해 동시성 제어
	var wg sync.WaitGroup
	// wg.Add(len(urls)) // 기다리는 타이머
	for _, url := range urls {
		// for문이 실행될 때마다 Add를 +1씩 증가
		// 다만 아래의 방법은 for문이 반복되기 전에 이전에 진행되던 고루틴들이 모두 끝날 경우
		// 고루틴이 wg.Wait()을 통과해버릴수도 있다. 이것을 rage Condition이라고 한다.
		wg.Add(1)
		go func(url string) {
			defer wg.Done() // 해당 함수가 종료되었음을 알림
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait() // 위의 Add숫자만큼 Done이 실행되어야 다음으로 넘길 수 있도록 기다림

	filenames, err := filepath.Glob("*.jpg")
	if err != nil {
		log.Fatal(err)
	}

	err = writeZip("images.zip", filenames)
	if err != nil {
		log.Fatal(err)
	}
}

func download(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	return filename, err
}

func urlToFilename(rawUrl string) (string, error) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	return filepath.Base(url.Path), nil
}

func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}

	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}

		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}

	return zw.Close()
}
