package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func ReadFile() ([]string, error) {
	file, err := os.Open("urls")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scann := bufio.NewScanner(file)

	for scann.Scan() {
		lines = append(lines, scann.Text())
	}

	return lines, scann.Err()
}

func MakeRequest(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	} else {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
			return 0, err
		} else {
			substr := "Go"

			bodyStr := string(body)

			wordCnt := 0

			wordCnt = strings.Count(bodyStr, substr)
			return wordCnt, nil
		}
	}
}

func main() {
	lineArr, err := ReadFile()

	if err != nil {
		fmt.Println("Error", err)
	} else {
		for _, url := range lineArr {
			fmt.Println(url)
		}
	}

	cnt, err := MakeRequest(lineArr[0])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cnt)
}
