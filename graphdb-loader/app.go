package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type (
	User_Fans struct {
		Vid1 uint `json:"vid1"`
		Vid2 uint `json:"vid2"`
	}
)

func (this *User_Fans) String() string {
	return strconv.Itoa(int(this.Vid1)) + "," + strconv.Itoa(int(this.Vid2))
}

var (
	folderpath                = flag.String("folder", "mydumperdata", "需要导入的数据文件夹所在路径")
	chancount                 = flag.Int("c", 10000, "单次上传数据量")
	lines      chan User_Fans = make(chan User_Fans, 1000000)
)

func post(data string) {
	url := "http://localhost:8010/fans/n?type=row"
	payload := strings.NewReader(data)
	req, _ := http.NewRequest("POST", url, payload)
	res, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	bd, err := ioutil.ReadAll(res.Body)
	fmt.Println(err)
	fmt.Println(string(bd))
	defer res.Body.Close()
}
func posting() {
	i := 0
	datalist := make([]string, 0)
	for true {
		l := <-lines
		fmt.Println("+")
		datalist = append(datalist, l.String())
		i++
		if i >= *chancount {
			data := strings.Join(datalist, "\n")
			post(data)
			i = 0
			datalist = make([]string, 0)
		}
	}
}
func processLine(line []byte) {
	l := string(line)
	l = strings.Replace(l, "(", "", -1)
	l = strings.Replace(l, ")", "", -1)
	l = strings.Replace(l, ";", "", -1)
	l = strings.Replace(l, " ", "", -1)
	l = strings.Replace(l, "\n", "", -1)
	u_f := strings.Split(l, ",")
	if len(u_f) >= 2 {
		vid1, err := strconv.Atoi(u_f[0])
		if err != nil {
			fmt.Println(string(line), "u_f[0]", u_f[0])
			return
		}

		vid2, err := strconv.Atoi(u_f[1])
		if err != nil {
			fmt.Println(string(line), "u_f[1]", u_f[1])
			return
		}
		u := User_Fans{
			Vid1: uint(vid1),
			Vid2: uint(vid2),
		}
		lines <- u
	}

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
				fmt.Println("read ok", filePth)
				return nil
			}
			return err
		}
	}
	return nil
}
func main() {
	flag.Parse()
	fmt.Println("folderpath", *folderpath)
	go posting()

	err := filepath.Walk(*folderpath, func(path string, f os.FileInfo, err error) error {
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
	time.Sleep(time.Second * 10)
}
