package main

import (
	_ "github.com/go-sql-driver/mysql"

	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

func post(data string) {
	url := "http://localhost:8010/relate/n"
	payload := strings.NewReader(data)
	req, _ := http.NewRequest("POST", url, payload)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
}
func processLine(line []byte) {
	os.Stdout.Write(line)
}
func ReadLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func main() {
	ReadLine("test.txt", processLine)
}
