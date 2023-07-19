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

//! use obly standart tools
// func ScanString() string {
// }

func ReadFile() ([]string, error) {
	file, err := os.Open("urls")
	if err != nil {
		fmt.Println("Error while opening file")
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
		log.Fatal(err)
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

			fmt.Println("BODY")
			fmt.Println(bodyStr)
			fmt.Println("END")

			bodyStrArr := strings.Split(bodyStr, " ")

			wordCnt := 0

			for _, word := range bodyStrArr {
				if strings.Compare(substr, word) == 0 {
					wordCnt++
				}
			}

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
