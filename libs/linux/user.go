package linux

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	uid   int
	uname string
}

type Group struct {
	gid   int
	gname string
}

var users []User
var group []Group

func init() {
	if len(users) == 0 {
		//打开文件，并进行相关处理
		f, err := os.Open("/etc/passwd")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		//文件关闭
		defer f.Close()

		//将文件作为一个io.Reader对象进行buffered I/O操作
		bufferedReader := bufio.NewReader(f)
		for {
			//每次读取一行
			line, err := bufferedReader.ReadString('\n')
			if err != nil {
				break
			} else {
				array := strings.Split(line, ":")
				uid, _ := strconv.Atoi(array[2])
				user := User{uid: uid, uname: array[0]}
				users = append(users, user)
			}
		}
	}
	if len(group) == 0 {
		//打开文件，并进行相关处理
		f, err := os.Open("/etc/group")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		//文件关闭
		defer f.Close()

		//将文件作为一个io.Reader对象进行buffered I/O操作
		br := bufio.NewReader(f)
		for {
			//每次读取一行
			line, err := br.ReadString('\n')
			if err != nil {
				break
			} else {
				array := strings.Split(line, ":")
				gid, _ := strconv.Atoi(array[2])
				group = append(group, Group{gid: gid, gname: array[0]})
			}
		}
	}
}

func GetUname(uid int) string {

	for _, val := range users {
		if val.uid == uid {
			return val.uname
		}
	}
	return ""
}

func GetGname(gid int) string {

	for _, val := range group {
		if val.gid == gid {
			return val.gname
		}
	}
	return ""
}
