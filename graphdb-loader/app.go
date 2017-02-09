package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type (
	User_Fans struct {
		Vid1 uint `json:"vid1"`
		Vid2 uint `json:"vid2"`
	}
)

var (
	folderpath                = *(flag.String("folder", "需要导入的数据文件夹所在路径", "."))
	lines      chan User_Fans = make(chan User_Fans, 1000000)
)

func post(data string) {
	url := "http://localhost:8010/relate/n"
	payload := strings.NewReader(data)
	req, _ := http.NewRequest("POST", url, payload)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
}
func posting() {
	i := 0
	datalist := make([]string, 0)
	for true {
		l := <-lines
		datalist = append(datalist, l)
		i++
		if i >= 10000 {
			data := strings.Join(datalist, "\n")
			post(data)
			i = 0
			datalist = make([]string, 0)
		}
	}

}
func processLine(line []byte) {
	u_f := strings.Split(string(line), ",")
	vid1, _ := strconv.Atoi(u_f[0])
	vid2, _ := strconv.Atoi(u_f[1])
	u := User_Fans{
		Vid1: vid1,
		Vid2: vid2,
	}
	lines <- u
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
	go posting()

	err := filepath.Walk(folderpath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ReadLine(path, processLine)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

}
