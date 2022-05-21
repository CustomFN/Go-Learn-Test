package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	fmt.Println("1")
	ReadFile("/Users/zhujiakun/Downloads/pic2")
}

func ReadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", fileName, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	num := 1000
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
		if len(line) < 5 {
			continue
		}

		resp, _ := http.Get(line)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		_ = ioutil.WriteFile("/Users/zhujiakun/Downloads/cc/"+strconv.FormatInt(int64(num), 10)+".jpg", body, 0755)
		num++

		time.Sleep(100 * time.Millisecond)
	}
}
